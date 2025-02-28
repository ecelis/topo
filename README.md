# Topo - simple network cross-platform network route manager

Test in Vagrant

```sh
vagrant up
```

## Linux Setup

```sh
go build -o topo cmd/topo/main.go
```

Create a systemd service file (e.g., `topo.service`) and place it in `/etc/systemd/system/`:

```ini
Ini, TOML

[Unit]
Description=My Route Service
After=network.target

[Service]
ExecStart=/path/to/your/route-service/route-service  // Path to your binary
Restart=always

[Install]
WantedBy=multi-user.target
```

```sh
sudo systemctl enable route-service.service
sudo systemctl start route-service.service
```

## Windows Setup

Run the r`topo.exe` from an administrator command prompt.

You can use a service manager like NSSM if you want to run it as a service.
