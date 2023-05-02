package common

import "net"

func GetTcpListener(listenAddr string) (*net.TCPListener, error) {
	structListenAddr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenTCP("tcp", structListenAddr)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func DialTCP(listenAddr, remoteAddr string) (*net.TCPConn, error) {
	structListenAddr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		panic(err)
	}
	structRemoteAddr, err := net.ResolveTCPAddr("tcp", remoteAddr)
	if err != nil {
		panic(err)
	}
	dconn, err := net.DialTCP("tcp", structListenAddr, structRemoteAddr)
	if err != nil {
		return nil, err
	}
	return dconn, err
}
