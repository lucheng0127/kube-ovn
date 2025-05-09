package netconf

import (
	"github.com/containernetworking/plugins/pkg/hns"

	"github.com/kubeovn/kube-ovn/pkg/request"
)

type NetConf struct {
	hns.NetConf
	ServerSocket string          `json:"server_socket"`
	Provider     string          `json:"provider"`
	Routes       []request.Route `json:"routes"`
	IPAM         *IPAMConf       `json:"ipam"`
	// PciAddrs in case of using sriov
	DeviceID string `json:"deviceID"`
	VfDriver string `json:"vf_driver"`
	// for dpdk
	VhostUserSocketVolumeName string `json:"vhost_user_socket_volume_name"`
	VhostUserSocketName       string `json:"vhost_user_socket_name"`
}

func (n *NetConf) PostLoad() {
	if len(n.DNS.Nameservers) == 0 {
		n.DNS.Nameservers = n.RuntimeConfig.DNS.Nameservers
	}
	if len(n.DNS.Search) == 0 {
		n.DNS.Search = n.RuntimeConfig.DNS.Search
	}
}
