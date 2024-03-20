package peer_store

import (
	"encoding/binary"
	"net"
	"testing"

	"github.com/obitoquilt/dht/v2/krpc"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	var r krpc.NodeAddr
	b := krpc.NodeAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 42069,
	}
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, uint16(b.Port))
	var key []byte
	key = append(key, b.IP...)
	key = append(key, bs...)
	r.UnmarshalBinary(key)
	assert.Equal(t, r.IP, b.IP)
	assert.Equal(t, r.Port, b.Port)
}
