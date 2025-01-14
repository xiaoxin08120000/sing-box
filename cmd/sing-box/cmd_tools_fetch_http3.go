//go:build with_quic

package main

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/xiaoxin08120000/quic-go"
	"github.com/xiaoxin08120000/quic-go/http3"
	box "github.com/xiaoxin08120000/sing-box"
	"github.com/xiaoxin08120000/sing/common/bufio"
	M "github.com/xiaoxin08120000/sing/common/metadata"
	N "github.com/xiaoxin08120000/sing/common/network"
)

func initializeHTTP3Client(instance *box.Box) error {
	dialer, err := createDialer(instance, commandToolsFlagOutbound)
	if err != nil {
		return err
	}
	http3Client = &http.Client{
		Transport: &http3.RoundTripper{
			Dial: func(ctx context.Context, addr string, tlsCfg *tls.Config, cfg *quic.Config) (quic.EarlyConnection, error) {
				destination := M.ParseSocksaddr(addr)
				udpConn, dErr := dialer.DialContext(ctx, N.NetworkUDP, destination)
				if dErr != nil {
					return nil, dErr
				}
				return quic.DialEarly(ctx, bufio.NewUnbindPacketConn(udpConn), udpConn.RemoteAddr(), tlsCfg, cfg)
			},
		},
	}
	return nil
}
