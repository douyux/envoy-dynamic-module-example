package main

import (
	"github.com/envoyproxy/dynamic-modules-examples/go/gosdk"
)

type passThroughFilterConfig struct{}

func (c *passThroughFilterConfig) NewFilter() gosdk.HttpFilter {
	return &passThroughFilter{}
}

func (c *passThroughFilterConfig) Destroy() {}

type passThroughFilter struct{}

func (f *passThroughFilter) RequestHeaders(e gosdk.EnvoyHttpFilter, endOfStream bool) gosdk.RequestHeadersStatus {
	return gosdk.RequestHeadersStatusContinue
}

func (f *passThroughFilter) RequestBody(e gosdk.EnvoyHttpFilter, endOfStream bool) gosdk.RequestBodyStatus {
	return gosdk.RequestBodyStatusContinue
}

func (f *passThroughFilter) ResponseHeaders(e gosdk.EnvoyHttpFilter, endOfStream bool) gosdk.ResponseHeadersStatus {
	return gosdk.ResponseHeadersStatusContinue
}

func (f *passThroughFilter) ResponseBody(e gosdk.EnvoyHttpFilter, endOfStream bool) gosdk.ResponseBodyStatus {
	return gosdk.ResponseBodyStatusContinue
}

func (f *passThroughFilter) Scheduled(e gosdk.EnvoyHttpFilter, eventID uint64) {}

func (f *passThroughFilter) Destroy() {}
