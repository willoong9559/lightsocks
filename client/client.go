package client

import (
	"log"
	"net"

	"github.com/willoong9559/lightsocks/common"
)

type LsClient struct {
	*common.Config
	*common.Forwarder
}

func NewLsClient(listenAddr, remoteAddr, password string) (*LsClient, error) {
	lsPassword, err := common.Atp(password)
	if err != nil {
		return nil, err
	}
	return &LsClient{
		Config:    common.NewClientConfig(listenAddr, remoteAddr, password),
		Forwarder: common.NewForWarder(lsPassword),
	}, nil
}

func (l *LsClient) Listen() error {
	listener, err := common.GetTcpListener(l.ListenAddr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Print(err)
			continue
		}
		go l.handler(conn)
	}
}

func (l *LsClient) handler(rconn *net.TCPConn) {
	defer rconn.Close()
	dconn, err := common.DialTCP(l.ListenAddr, l.RemoteAddr)
	if err != nil {
		return
	}
	// Conn被关闭时直接清除所有数据 不管没有发送的数据
	rconn.SetLinger(0)
	dconn.SetLinger(0)
	// 转发
	go func() {
		err := l.DecodeCopy(rconn, dconn)
		if err != nil {
			dconn.Close()
			return
		}
	}()
	l.EncodeCopy(dconn, rconn)
}
