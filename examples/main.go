package main

import "github.com/envoyproxy/dynamic-modules-examples/go/gosdk"

func main() {}

func init() {
	gosdk.NewHttpFilterConfig = newHttpFilterConfig
}

func newHttpFilterConfig(name string, config []byte) gosdk.HttpFilterConfig {
	switch name {
	case "passthrough":
		return &passThroughFilterConfig{}
	default:
		panic("unknown filter config: " + name)
	}
}
