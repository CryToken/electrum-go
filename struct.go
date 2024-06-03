package electrum

import "net"

type ElectrumServer struct {
	Conn    net.Conn
	Network string
}
