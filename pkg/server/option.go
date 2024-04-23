package server

import "net"

type Option func(builder *ApiServer)

func Port(port string) Option {
	return func(builder *ApiServer) {
		builder.server.Addr = net.JoinHostPort("", port)
	}
}
