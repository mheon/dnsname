package main

import (
	"errors"

	"github.com/containernetworking/cni/pkg/types"
)

const (
	//	dnsNameConfPath is where we store the conf, pid, and hosts files
	dnsNameConfPath = "/run/containers/cni/dnsname"
	// confFileName is the name of the dns masq conf file
	confFileName = "dnsmasq.conf"
	// hostsFileName is the name of the addnhosts file
	hostsFileName = "addnhosts"
	// pidFileName is the file where the dnsmasq file is stored
	pidFileName = "pidfile"
)

const dnsMasqTemplate = `## WARNING: THIS IS AN AUTOGENERATED FILE
## AND SHOULD NOT BE EDITED MANUALLY AS IT
## LIKELY TO AUTOMATICALLY BE REPLACED.
strict-order
local=/{{.Domain}}/
domain={{.Domain}}
expand-hosts
pid-file={{.PidFile}}
except-interface=lo
bind-dynamic
no-hosts
interface={{.NetworkInterface}}
addn-hosts={{.AddOnHostsFile}}`

var (
	// ErrBinaryNotFound means that the dnsmasq binary was not found
	ErrBinaryNotFound = errors.New("unable to locate dnsmasq in path")
	// ErrNoIPAddressFound means that CNI was unable to resolve an IP address in the CNI configuration
	ErrNoIPAddressFound = errors.New("no ip address was found in the network")
)

// DNSNameConf represents the cni config with the domain name attribute
type DNSNameConf struct {
	types.NetConf
	DomainName string `json:"domainName"`
}

// dnsNameFile describes the plugin's attributes
type dnsNameFile struct {
	AddOnHostsFile   string
	Binary           string
	ConfigFile       string
	Domain           string
	NetworkInterface string
	PidFile          string
}
