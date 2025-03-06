/*
Copyright © 2025 Ernesto Celis <ernesto@patito.ninja>
*/
package route

import (
	"io/ioutil"

	"encoding/json"
)

type Route struct {
	Device      string
	Destination string
	Netmask     string
	Gateway     string
}

func ReadConfig(filename string) ([]Route, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Route{}, err
	}

	var routes []Route
	if err := json.Unmarshal(b, &routes); err != nil {
		return []Route{}, err
	}

	return routes, nil
}
