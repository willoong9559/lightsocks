package server

import (
	"fmt"
	"log"
	"net"

	"github.com/willoong9559/lightsocks/common"
	"github.com/willoong9559/lightsocks/conf"
)

type LsServer struct {
	*conf.Config
	*common.Forwarder
}

func NewLsServer() (*LsServer, error) {
	config := conf.NewConfig()
	forwarder, err := common.NewForwarderWithStr(config.Password)
	if err != nil {
		return nil, err
	}
	return &LsServer{
		Config:    config,
		Forwarder: forwarder,
	}, nil
}

func (l *LsServer) Listen() error {
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

func (l *LsServer) handler(rconn *net.TCPConn) {
	defer rconn.Close()

	rconn.SetLinger(0)
	buf := common.BufferPoolGet()
	defer common.BufferPoolPut(buf)

	_, err := l.DecodeRead(rconn, buf)
	if err != nil || buf[0] != 0x05 {
		log.Printf("不支持的socks版本: %v", buf[0])
		return
	}
	_, err = l.EncodeWrite(rconn, []byte{0x05, 0x00})
	if err != nil {
		log.Printf("写入socks5响应失败")
		return
	}

	n, _ := l.DecodeRead(rconn, buf)
	if buf[1] != 0x01 {
		// 目前只支持 CONNECT
		return
	}

	dstAddr, err := GerSocksDstAddr(buf, n)
	if err != nil {
		log.Print(err)
		return
	}

	// 连接真正的远程服务
	dconn, err := net.DialTCP("tcp", nil, dstAddr)
	if err != nil {
		log.Print(err)
		return
	} else {
		defer dconn.Close()
		// Conn被关闭时直接清除所有数据 不管没有发送的数据
		dconn.SetLinger(0)
		l.EncodeWrite(rconn, []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	}

	// 进行转发
	// 从 localUser 读取数据发送到 dstServer
	go func() {
		err := l.DecodeCopy(dconn, rconn)
		if err != nil {
			rconn.Close()
			dconn.Close()
		}
	}()
	l.EncodeCopy(rconn, dconn)
}

func (l *LsServer) PrintInfo() {
	info := fmt.Sprintf(`
服务端启动成功，配置如下：
本地监听地址：
%s`, conf.ListenAddr)
	log.Println(info)
}
