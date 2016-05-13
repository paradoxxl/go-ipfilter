package ipfilter

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
)


type ExtIpnet struct {
	ipnet net.IPNet
	permit bool
}
type ExtIp struct {
	ip net.IP
	permit bool
}
type IPFilter struct {
	ipnets []ExtIpnet
	ips    []ExtIp
}

func (f *IPFilter) FilterIP(ip net.IP) bool {
	for _, item := range f.ipnets {
		if item.ipnet.Contains(ip) {

			return item.permit
		}
	}
	for _, item := range f.ips {
		if item.ip.Equal(ip) {
			return item.permit
		}
	}
	return false
}

func (f *IPFilter) FilterIPString(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	return f.FilterIP(ip)
}

func (f *IPFilter) AddIPNet(item net.IPNet) {
	entry := ExtIpnet{item,true}
	f.ipnets = append(f.ipnets, entry)
}
func (f *IPFilter) AddIPNetExt(item net.IPNet,permit bool) {
	entry := ExtIpnet{item,permit}
	f.ipnets = append(f.ipnets, entry)
}

func (f *IPFilter) AddIPNetString(s string) error {
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		return err
	}
	f.AddIPNet(*ipnet)
	return nil
}
func (f *IPFilter) AddIPNetStringExt(s string,permit bool) error {
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		return err
	}
	f.AddIPNetExt(*ipnet,permit)
	return nil
}

func (f *IPFilter) AddIP(ip net.IP) {
	entry := ExtIp{ip,true}
	f.ips = append(f.ips, entry)
}

func (f *IPFilter) AddIPExt(ip net.IP,permit bool) {
	entry := ExtIp{ip,true}
	f.ips = append(f.ips, entry)
}

func (f *IPFilter) AddIPString(s string) error {
	ip := net.ParseIP(s)
	if ip == nil {
		return errors.New("Parse IP Error: " + s)
	}
	f.AddIP(ip)
	return nil
}
func (f *IPFilter) AddIPStringExt(s string ,permit bool) error {
	ip := net.ParseIP(s)
	if ip == nil {
		return errors.New("Parse IP Error: " + s)
	}
	f.AddIPExt(ip,permit)
	return nil
}

func (f *IPFilter) Load(data []byte) error {
	for _, item := range bytes.Fields(data) {
		if bytes.IndexByte(item, '/') < 0 {
			if err := f.AddIPString(string(item)); err != nil {
				return err
			}
		} else {
			if err := f.AddIPNetString(string(item)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *IPFilter) LoadFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}
	return f.Load(data)
}

func (f *IPFilter) LoadHttp(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return f.Load(data)
}
