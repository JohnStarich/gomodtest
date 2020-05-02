package init

import (
	"net"

	"github.com/johnstarich/gomodtest/dns"
)

func init() {
	net.DefaultResolver = dns.New()
}
