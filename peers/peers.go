package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type Peer struct{
	IP net.IP
	Port uint16
}

func Unmarshal(peerBin []byte) ([]Peer,error){
	const peersize = 6
	// 4 bytes for the IPv4 and 2 bytes for Ports
	numPeers := len(peerBin)/peersize
	if len(peerBin)%peersize!=0{
		err := fmt.Errorf("Received malformed peers")
        return nil, err
	}
	peers:=make([]Peer,numPeers)
	for i:=0;i<numPeers;i++{
		offset := i*peersize
		peers[i].IP = net.IP(peerBin[offset:offset+4])
		peers[i].Port = binary.BigEndian.Uint16([]byte(peerBin[offset+4:offset+6]))

	}
	return peers,nil

}

func (p Peer) String () string{
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}