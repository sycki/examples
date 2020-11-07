package imitate_tcp_base_ip

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

const (
	SrcPortOffset   = 0
	DstPortOffset   = 2
	SeqNumOffset    = 4
	AckNumOffset    = 8
	DataOffset      = 12
	FlagsOffset     = 13
	WinSizeOffset   = 14
	ChecksumOffset  = 16
	UrgentPtrOffset = 18
)

type TCPHeader struct {
	Source   uint16
	Target   uint16
	SeqNum   uint32
	AckNum   uint32
	HeadLen  uint8
	Flags    uint8
	WinSize  uint16
	CheckSum uint16
	UrgentP  uint16
}

func (p *TCPHeader) Marshal() []byte {
	b := make([]byte, 20)
	binary.BigEndian.PutUint16(b[SrcPortOffset:], p.Source)
	binary.BigEndian.PutUint16(b[DstPortOffset:], p.Target)
	binary.BigEndian.PutUint32(b[SeqNumOffset:], p.SeqNum)
	binary.BigEndian.PutUint32(b[AckNumOffset:], p.AckNum)
	b[DataOffset] = (p.HeadLen / 4) << 4
	b[FlagsOffset] = p.Flags
	binary.BigEndian.PutUint16(b[WinSizeOffset:], p.Source)
	binary.BigEndian.PutUint16(b[ChecksumOffset:], p.Source)
	binary.BigEndian.PutUint16(b[UrgentPtrOffset:], p.Source)

	return b
}

func Unmarshal(b []byte) *TCPHeader {
	if len(b) < 20 {
		return nil
	}

	tcpHeader := &TCPHeader{
		binary.BigEndian.Uint16(b[SrcPortOffset:]),
		binary.BigEndian.Uint16(b[DstPortOffset:]),
		binary.BigEndian.Uint32(b[SeqNumOffset:]),
		binary.BigEndian.Uint32(b[AckNumOffset:]),
		(b[DataOffset] >> 4) * 4,
		b[FlagsOffset],
		binary.BigEndian.Uint16(b[WinSizeOffset:]),
		binary.BigEndian.Uint16(b[ChecksumOffset:]),
		binary.BigEndian.Uint16(b[UrgentPtrOffset:]),
	}

	return tcpHeader
}

func ListenTCP(address, port string) error {
	addr, _ := net.ResolveIPAddr("ip4", address)
	conn, err := net.ListenIP("ip4:tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	buf := make([]byte, 1480)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			return err
		}

		tcpHeader := Unmarshal(buf[:20])
		realPort := fmt.Sprint(tcpHeader.Target)
		if realPort != port {
			fmt.Printf("the package is not to me, expectPort [%s], realPort [%s]\n", port, realPort)
			continue
		}
		fmt.Printf("receve data headLen [%d], remoteAddr [%s], msg [%s]\n", n, addr.String(), buf[tcpHeader.HeadLen:n])
	}
}

func DialTCP(address, port string) error {
	conn, err := net.Dial("ip4:tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	portNum, _ := strconv.Atoi(port)
	tcpHeader := &TCPHeader{
		65530,
		uint16(portNum),
		0,
		0,
		(20 / 4) << 4,
		0,
		0,
		0,
		0,
	}
	b := tcpHeader.Marshal()
	opts := make([]byte, ((tcpHeader.HeadLen>>4)*4)-20)
	b = append(b, opts...)
	b = append(b, []byte("hello!")...)

	n, err := conn.Write(b)
	fmt.Println("client write len: ", n)

	return nil
}
