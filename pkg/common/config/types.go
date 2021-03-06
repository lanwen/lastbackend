//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package config

type Config struct {
	LogLevel        *int
	Token           *string
	SystemDomain    *string
	Etcd            ETCD
	Registry        Registry
	APIServer       APIServer
	ProxyServer     ProxyServer
	AgentServer     AgentServer
	DiscoveryServer DiscoveryServer
	Runtime         Runtime
	Host            Host
}

type APIServer struct {
	Host *string
	Port *int
}

type ProxyServer struct {
	Host *string
	Port *int
}

type ETCD struct {
	Endpoints *[]string
	TLS       struct {
		Key  *string
		Cert *string
		CA   *string
	}
	Quorum *bool
}

type Host struct {
	Hostname *string
	IP       *string
}

type Runtime struct {
	CRI    *string
	Docker Docker
}

type Docker struct {
	Host, Certs, Version *string
	TLS                  *bool
}

type Registry struct {
	Server, Username, Password *string
}

type AgentServer struct {
	Host *string
	Port *int
}

type DiscoveryServer struct {
	Port *int
}
