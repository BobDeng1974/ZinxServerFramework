package zinxInterface

import "net"

/*
	定义抽象的Tcp链接层,,用于将链接请求和用户的业务进行绑定
*/
type InterfaceConnection interface {
	//启动链接
	Start()

	//停止链接
	Stop()

	//获取链接ID
	GetConnID() uint32

	//获取conn的原声socket套接字
	GetTCPConnection() *net.TCPConn

	//获取远程客户端的ip地址
	GetRemoteAddr() *net.Addr

	//发送数据给对方客户端
	Send(data[]byte)error
}

//定义抽象的业务处理方法
type HandleFunc func(*net.TCPConn,[]byte,int) error