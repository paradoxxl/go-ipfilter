package ipfilter

import (
	"testing"
)

func TestIPFilter(t *testing.T) {
	var f IPFilter
	if err := f.AddIPNetString("192.168.1.0/24"); err != nil {
		t.Errorf("Add: 192.168.1.0/24, %s", err.Error())
	}
	if err := f.AddIPNetStringExt("2001:abcd::0/64",true); err != nil {
		t.Errorf("Add: 2001:abcd::0/64, %s", err.Error())
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
	if f.FilterIPString("2001:abcd::1/64") {
		t.Error("Filter Fail: 2001:abcd::1/64")
	}
}

func TestBehaviour(t *testing.T){
	var f2 IPFilter

	if err := f2.AddIPNetStringExt("172.16.1.30/32",true); err != nil {
		t.Errorf("Add: 172.16.1.0/24, %s", err.Error())
	}

	if err := f2.AddIPNetStringExt("172.16.1.0/24",false); err != nil {
		t.Errorf("Add: 192.168.1.0/24, %s", err.Error())
	}
	if err := f2.AddIPNetStringExt("172.16.1.0/24",true); err != nil {
		t.Errorf("Add: 172.16.1.0/24, %s", err.Error())
	}

	f2.defaultBehaviour = true
	if !f2.FilterIPString("10.10.10.10") {
		t.Error("Filter Fail: 10.10.10.10 should return true")
	}
	f2.defaultBehaviour = false
	if f2.FilterIPString("10.10.10.10") {
		t.Error("Filter Fail: 10.10.10.10 should return false")
	}

	if err := f2.AddIPNetStringExt("0.0.0.0/0",true); err != nil {
		t.Errorf("Add: 0.0.0.0/0, %s", err.Error())
	}
	if f2.FilterIPString("172.16.1.20") {
		t.Error("Filter Fail: 172.16.1.20 should return false")
	}
	if !f2.FilterIPString("10.10.10.10") {
		t.Error("Filter Fail: 10.10.10.10 should return true")
	}


	if !f2.FilterIPString("172.16.1.30") {
		t.Error("Filter Fail: 172.16.1.20 should return true")
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
