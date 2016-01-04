package main

import (
	//"errors"
	//"bufio"
	"fmt"
	"net"
	//"os"
	"container/list"
	"strconv"
)

var localaddr string

type StreamAddr struct {
	addr *net.UDPAddr
	conn *net.UDPConn
}

type Stream struct {
	conn net.Conn
	uuid string
	dest *list.List
}

type StreamManger struct {
	stream_counter int
	stream_map     map[string]*Stream
}

var g_sm StreamManger

func InitStreamServer(port int) {
	localaddr = ":" + strconv.Itoa(port)
	fmt.Println("Server Start At ", localaddr)
}

func StartStreamServer() {
	// net.ResolveTCPAddr is create a TCPAddr
	if tcpaddr, err := net.ResolveTCPAddr("tcp4", localaddr); err != nil {
		panic(err)
	} else {
		listener, err := net.ListenTCP("tcp", tcpaddr)
		CheckErrorExit(err)
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			uuid := Rand()
			ret, err := createTCPStreamChannel(conn, uuid.Hex())
			if ret {
				go handleClient(conn, uuid.Hex())
			}
		}
	}
}

func CloseStreamServer() {

}

func createTCPStreamChannel(c net.Conn, uuid string) (ret bool, err error) {
	fmt.Println("Remote ip addr ", c.RemoteAddr())
	stream := new(Stream)
	stream.Init(c, uuid)
	g_sm.AddStream(uuid, stream)
	ret = true
	return
}

func handleClient(c net.Conn, uuid string) {
	var buf [4096]byte
	n, err := c.Read(buf[0:])
	if err != nil {
		if _, ok := g_sm.stream_map[uuid]; !ok {
			g_sm.stream_map[uuid].SendData(buf[0:n])
		}
	}
}

func ListStreamFormatJson() (ret string, err error) {
	ret = ""
	return
}

func (s *Stream) Init(c net.Conn, u string) {
	s.conn = c
	s.uuid = u
	s.dest = list.New()
}

func (s *Stream) AddDestAddr(addr string) (ret bool, err error) {
	//udpaddr := net.UDPAddr
	udpaddr, err := net.ResolveUDPAddr("udp", addr)
	sa := new(StreamAddr)
	sa.addr = udpaddr
	sa.conn, err = net.DialUDP("udp", nil, udpaddr)
	if err == nil {
		s.dest.PushBack(sa)
		return true, err
	}
	return false, err
}

func (s *Stream) DelDestAddr(addr string) (ret bool, err error) {
	udpaddr, err := net.ResolveUDPAddr("udp", addr)
	fmt.Println(udpaddr.String())
	for e := s.dest.Front(); e != nil; e = e.Next() {
		if e.Value.(*StreamAddr).addr.IP.Equal(udpaddr.IP) {
			if e.Value.(*StreamAddr).addr.Port == udpaddr.Port {
				defer s.dest.Remove(e)
			}
		}
	}
	return true, err
}

func (s *Stream) SendData(buf []byte) {
	for e := s.dest.Front(); e != nil; e = e.Next() {
		e.Value.(*StreamAddr).conn.Write(buf)
	}
}

func (sm *StreamManger) AddStream(uuid string, s *Stream) (ret bool, err error) {
	if _, ok := sm.stream_map[uuid]; !ok {
		sm.stream_counter++
		sm.stream_map[uuid] = s
		ret = true
		return
	}
	ret = false
	return
}

func (sm *StreamManger) DelStream(uuid string) (ret bool, err error) {
	if _, ok := sm.stream_map[uuid]; !ok {
		sm.stream_counter--
		delete(sm.stream_map, uuid)
		ret = true
		return
	}
	err.Error()
	ret = false
	return
}

func (sm *StreamManger) AddStreamAddr(uuid string, addr string) (ret bool, err error) {
	if _, ok := sm.stream_map[uuid]; !ok {
		ret, err := sm.stream_map[uuid].AddDestAddr(addr)
		return ret, err
	}
	ret = false
	return
}

func (sm *StreamManger) DelStreamAddr(uuid string, addr string) (ret bool, err error) {
	if _, ok := sm.stream_map[uuid]; !ok {
		ret, err := sm.stream_map[uuid].DelDestAddr(addr)
		return ret, err
	}
	ret = false
	return
}
