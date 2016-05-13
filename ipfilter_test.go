package ipfilter

import (
	"testing"
)

func TestIPFilter(t *testing.T) {
	var f IPFilter
	if err := f.AddIPNetString("192.168.1.0/24"); err != nil {
		t.Errorf("Add: 192.168.1.0/24, %s", err.Error())
	}
	if err := f.AddIPString("192.168.100.3"); err != nil {
		t.Errorf("Add: 192.168.100.3, %s", err.Error())
	}
	if err := f.AddIPStringExt("172.16.10.1",false); err != nil {
		t.Errorf("Add: 172.16.10.1, %s", err.Error())
	}
	if err := f.AddIPNetStringExt("172.16.1.0/24",false); err != nil {
		t.Errorf("Add: 172.16.1.0/24, %s", err.Error())
	}
	if !f.FilterIPString("192.168.1.3") {
		t.Error("Filter Fail: 192.168.1.3")
	}
	if !f.FilterIPString("192.168.100.3") {
		t.Error("Filter Fail: 192.168.100.3")
	}
	if !f.FilterIPString("172.16.10.1") {
		t.Error("Filter Fail: 172.16.10.1")
	}
	if f.FilterIPString("172.16.1.20") {
		t.Error("Filter Fail: 172.16.1.20")
	}
}

func TestLoad(t *testing.T) {
	var f IPFilter
	var data = []byte("192.168.1.0/24 192.168.1.3")
	if err := f.Load(data); err != nil {
		t.Error(err)
	}
	if !f.FilterIPString("192.168.1.3") {
		t.Error("Filter Fail: 192.168.1.3")
	}
}
