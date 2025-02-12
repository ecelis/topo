//go:build linux
// +build linux

package platform

import (
	"fmt"
	"net"
	"syscall"
	"unsafe"

	"github.com/ecelis/topo/internal/route"
)

// Linux Implementation (using netlink)
func AddRoute(r route.Route) error {
	dst := net.ParseIP(r.Destination)
	if dst == nil {
		return fmt.Errorf("invalid destination IP: %s", r.Destination)
	}

	gateway := net.ParseIP(r.Gateway)
	if gateway == nil {
		return fmt.Errorf("invalid gateway IP: %s", r.Gateway)
	}

	// Convert IP addresses to 4-byte representations
	dst4 := dst.To4()
	gw4 := gateway.To4()

	if dst4 == nil || gw4 == nil {
		return fmt.Errorf("IP addresses must be ipv4")
	}

	mask, _ := net.ParseIPMask(r.Netmask)
	mask4 := net.IPv4Mask(mask[0], mask[1], mask[2], mask[3])

	maskLen, _ := mask4.Size()

	// Create a netlink socket
	s, err := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_DGRAM, syscall.NETLINK_ROUTE)
	if err != nil {
		return fmt.Errorf("error creating netlink socket: %w", err)
	}
	defer syscall.Close(s)

	// Build the netlink message
	nlmsg := &syscall.NlMsghdr{
		Len:   syscall.NLM_LENGTH(syscall.SizeofRtmsg),
		Type:  syscall.RTM_NEWROUTE,
		Flags: syscall.NLM_F_REQUEST | syscall.NLM_F_CREATE,
		Seq:   1, // You can increment this for each message
		Pid:   uint32(syscall.Getpid()),
	}

	rtmsg := &syscall.Rtmsg{
		Family:  syscall.AF_INET,
		Dst_len: uint8(maskLen),            // Prefix length
		Scope:   syscall.RT_SCOPE_UNIVERSE, // Or appropriate scope
		Type:    syscall.RT_TABLE_MAIN,
	}

	// Pack the netlink message data
	rtmsgBytes := unsafe.Slice((*byte)(unsafe.Pointer(rtmsg)), syscall.SizeofRtmsg)
	nlmsg.Data = append(rtmsgBytes)

	// Add destination attribute
	dstAttr := &syscall.NlAttr{
		Type: syscall.RTA_DST,
	}
	dstAttrBytes := unsafe.Slice((*byte)(unsafe.Pointer(dstAttr)), syscall.SizeofNlAttr)
	nlmsg.Data = append(nlmsg.Data, dstAttrBytes...)
	nlmsg.Data = append(nlmsg.Data, dst4...)

	// Add gateway attribute
	gwAttr := &syscall.NlAttr{
		Type: syscall.RTA_GATEWAY,
	}
	gwAttrBytes := unsafe.Slice((*byte)(unsafe.Pointer(gwAttr)), syscall.SizeofNlAttr)
	nlmsg.Data = append(nlmsg.Data, gwAttrBytes...)
	nlmsg.Data = append(nlmsg.Data, gw4...)

	// Send the netlink message
	if _, err := syscall.Sendto(s, nlmsg.Data, nil); err != nil {
		return fmt.Errorf("error sending netlink message: %w", err)
	}

	return nil
}
