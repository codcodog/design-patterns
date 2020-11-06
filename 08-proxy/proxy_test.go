package proxy

import "testing"

func TestProxy(t *testing.T) {
	proxy := &Proxy{}
	if proxy.Do() != "pre:real:after" {
		t.Fatal("test proxy failed.")
	}
}
