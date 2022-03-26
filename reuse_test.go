package reuseport

import (
	"context"
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var device = "lo"

func testDialFromListeningPort(t *testing.T, network, host string) {
	lc := net.ListenConfig{
		Control: GetBindToDeviceControl(device),
	}
	ctx := context.Background()
	ll, err := lc.Listen(ctx, network, host+":0")
	if err != nil && strings.Contains(err.Error(), "cannot assign requested address") {
		t.Skip(err)
	}
	require.NoError(t, err)
	rl, err := lc.Listen(ctx, network, host+":0")
	require.NoError(t, err)
	d := net.Dialer{
		LocalAddr: ll.Addr(),
		Control:   GetBindToDeviceControl(""),
	}
	c, err := d.Dial(network, rl.Addr().String())
	require.NoError(t, err)
	c.Close()
}

func TestDialFromListeningPort(t *testing.T) {
	testDialFromListeningPort(t, "tcp", "localhost")
}

func TestDialFromListeningPortTcp6(t *testing.T) {
	testDialFromListeningPort(t, "tcp6", "[::1]")
}

func TestListenPacketWildcardAddress(t *testing.T) {
	pc, err := ListenPacket(device, "udp", ":0")
	require.NoError(t, err)
	pc.Close()
}

func TestErrorWhenDialUnresolvable(t *testing.T) {
	_, err := Dial(device, "asd", "127.0.0.1:1234", "127.0.0.1:1234")
	require.ErrorIs(t, err, net.UnknownNetworkError("asd"))
	_, err = Dial(device, "tcp", "a.b.c.d:1234", "a.b.c.d:1235")
	require.Error(t, err)
}
