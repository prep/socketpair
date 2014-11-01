// Copyright (c) 2014 Maurice Nonnekes <maurice@codeninja.nl>
// All rights reserved.

package socketpair

import (
	"net"
	"testing"
)

var (
	testByteString = []byte("Hello World")
	testNetworks   = []string{"unix", "unixgram"}
)

func slicesAreEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestSocketpair(t *testing.T) {
	for _, network := range testNetworks {
		sock1, sock2, err := New(network)
		if err != nil {
			t.Fatalf("Error creating socket: %s: %s", network, err)
		}

		defer sock1.Close()
		defer sock2.Close()

		if _, ok := sock1.(*net.UnixConn); !ok {
			t.Fatalf("Expected to be able to typecast to a new.UnixConn pointer")
		}

		if _, ok := sock2.(*net.UnixConn); !ok {
			t.Fatalf("Expected to be able to typecast to a new.UnixConn pointer")
		}

		if _, err := sock1.Write(testByteString); err != nil {
			t.Fatalf("Error writing to socket: %s: %s", network, err)
		}

		byteString := make([]byte, len(testByteString))
		if _, err := sock2.Read(byteString); err != nil {
			t.Fatalf("Error reading from socket: %s: %s", network, err)
		}

		if !slicesAreEqual(byteString, testByteString) {
			t.Fatalf("Unexpected data read from unix socket: %s", byteString)
		}
	}
}

func TestIllegalNetwork(t *testing.T) {
	if _, _, err := New("foobar"); err == nil {
		t.Fatalf("Expected error when requesting a bogus network type")
	}
}
