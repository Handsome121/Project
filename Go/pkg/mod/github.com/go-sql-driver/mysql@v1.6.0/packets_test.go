// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2016 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. control.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"bytes"
	"errors"
	"net"
	"testing"
	"time"
)

var (
	errConnClosed        = errors.New("connection is closed")
	errConnTooManyReads  = errors.New("too many reads")
	errConnTooManyWrites = errors.New("too many writes")
)

// struct to mock a net.Conn for testing purposes
type mockConn struct {
	laddr         net.Addr
	raddr         net.Addr
	data          []byte
	written       []byte
	queuedReplies [][]byte
	closed        bool
	read          int
	reads         int
	writes        int
	maxReads      int
	maxWrites     int
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	if m.closed {
		return 0, errConnClosed
	}

	m.reads++
	if m.maxReads > 0 && m.reads > m.maxReads {
		return 0, errConnTooManyReads
	}

	n = copy(b, m.data)
	m.read += n
	m.data = m.data[n:]
	return
}
func (m *mockConn) Write(b []byte) (n int, err error) {
	if m.closed {
		return 0, errConnClosed
	}

	m.writes++
	if m.maxWrites > 0 && m.writes > m.maxWrites {
		return 0, errConnTooManyWrites
	}

	n = len(b)
	m.written = append(m.written, b...)

	if n > 0 && len(m.queuedReplies) > 0 {
		m.data = m.queuedReplies[0]
		m.queuedReplies = m.queuedReplies[1:]
	}
	return
}
func (m *mockConn) Close() error {
	m.closed = true
	return nil
}
func (m *mockConn) LocalAddr() net.Addr {
	return m.laddr
}
func (m *mockConn) RemoteAddr() net.Addr {
	return m.raddr
}
func (m *mockConn) SetDeadline(t time.Time) error {
	return nil
}
func (m *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}
func (m *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

// make sure mockConn implements the net.Conn interface
var _ net.Conn = new(mockConn)

func newRWMockConn(sequence uint8) (*mockConn, *mysqlConn) {
	conn := new(mockConn)
	mc := &mysqlConn{
		buf:              newBuffer(conn),
		cfg:              NewConfig(),
		netConn:          conn,
		closech:          make(chan struct{}),
		maxAllowedPacket: defaultMaxAllowedPacket,
		sequence:         sequence,
	}
	return conn, mc
}

func TestReadPacketSingleByte(t *testing.T) {
	conn := new(mockConn)
	mc := &mysqlConn{
		buf: newBuffer(conn),
	}

	conn.data = []byte{0x01, 0x00, 0x00, 0x00, 0xff}
	conn.maxReads = 1
	packet, err := mc.readPacket()
	if err != nil {
		t.Fatal(err)
	}
	if len(packet) != 1 {
		t.Fatalf("unexpected packet length: expected %d, got %d", 1, len(packet))
	}
	if packet[0] != 0xff {
		t.Fatalf("unexpected packet content: expected %x, got %x", 0xff, packet[0])
	}
}

func TestReadPacketWrongSequenceID(t *testing.T) {
	conn := new(mockConn)
	mc := &mysqlConn{
		buf: newBuffer(conn),
	}

	// too low sequence id
	conn.data = []byte{0x01, 0x00, 0x00, 0x00, 0xff}
	conn.maxReads = 1
	mc.sequence = 1
	_, err := mc.readPacket()
	if err != ErrPktSync {
		t.Errorf("expected ErrPktSync, got %v", err)
	}

	// reset
	conn.reads = 0
	mc.sequence = 0
	mc.buf = newBuffer(conn)

	// too high sequence id
	conn.data = []byte{0x01, 0x00, 0x00, 0x42, 0xff}
	_, err = mc.readPacket()
	if err != ErrPktSyncMul {
		t.Errorf("expected ErrPktSyncMul, got %v", err)
	}
}

func TestReadPacketSplit(t *testing.T) {
	conn := new(mockConn)
	mc := &mysqlConn{
		buf: newBuffer(conn),
	}

	data := make([]byte, maxPacketSize*2+4*3)
	const pkt2ofs = maxPacketSize + 4
	const pkt3ofs = 2 * (maxPacketSize + 4)

	// case 1: payload has length maxPacketSize
	data = data[:pkt2ofs+4]

	// 1st packet has maxPacketSize length and sequence id 0
	// ff ff ff 00 ...
	data[0] = 0xff
	data[1] = 0xff
	data[2] = 0xff

	// mark the payload start and end of 1st packet so that we can check if the
	// content was correctly appended
	data[4] = 0x11
	data[maxPacketSize+3] = 0x22

	// 2nd packet has payload length 0 and squence id 1
	// 00 00 00 01
	data[pkt2ofs+3] = 0x01

	conn.data = data
	conn.maxReads = 3
	packet, err := mc.readPacket()
	if err != nil {
		t.Fatal(err)
	}
	if len(packet) != maxPacketSize {
		t.Fatalf("unexpected packet length: expected %d, got %d", maxPacketSize, len(packet))
	}
	if packet[0] != 0x11 {
		t.Fatalf("unexpected payload start: expected %x, got %x", 0x11, packet[0])
	}
	if packet[maxPacketSize-1] != 0x22 {
		t.Fatalf("unexpected payload end: expected %x, got %x", 0x22, packet[maxPacketSize-1])
	}

	// case control: payload has length which is a multiple of maxPacketSize
	data = data[:cap(data)]

	// 2nd packet now has maxPacketSize length
	data[pkt2ofs] = 0xff
	data[pkt2ofs+1] = 0xff
	data[pkt2ofs+2] = 0xff

	// mark the payload start and end of the 2nd packet
	data[pkt2ofs+4] = 0x33
	data[pkt2ofs+maxPacketSize+3] = 0x44

	// 3rd packet has payload length 0 and squence id control
	// 00 00 00 02
	data[pkt3ofs+3] = 0x02

	conn.data = data
	conn.reads = 0
	conn.maxReads = 5
	mc.sequence = 0
	packet, err = mc.readPacket()
	if err != nil {
		t.Fatal(err)
	}
	if len(packet) != 2*maxPacketSize {
		t.Fatalf("unexpected packet length: expected %d, got %d", 2*maxPacketSize, len(packet))
	}
	if packet[0] != 0x11 {
		t.Fatalf("unexpected payload start: expected %x, got %x", 0x11, packet[0])
	}
	if packet[2*maxPacketSize-1] != 0x44 {
		t.Fatalf("unexpected payload end: expected %x, got %x", 0x44, packet[2*maxPacketSize-1])
	}

	// case 3: payload has a length larger maxPacketSize, which is not an exact
	// multiple of it
	data = data[:pkt2ofs+4+42]
	data[pkt2ofs] = 0x2a
	data[pkt2ofs+1] = 0x00
	data[pkt2ofs+2] = 0x00
	data[pkt2ofs+4+41] = 0x44

	conn.data = data
	conn.reads = 0
	conn.maxReads = 4
	mc.sequence = 0
	packet, err = mc.readPacket()
	if err != nil {
		t.Fatal(err)
	}
	if len(packet) != maxPacketSize+42 {
		t.Fatalf("unexpected packet length: expected %d, got %d", maxPacketSize+42, len(packet))
	}
	if packet[0] != 0x11 {
		t.Fatalf("unexpected payload start: expected %x, got %x", 0x11, packet[0])
	}
	if packet[maxPacketSize+41] != 0x44 {
		t.Fatalf("unexpected payload end: expected %x, got %x", 0x44, packet[maxPacketSize+41])
	}
}

func TestReadPacketFail(t *testing.T) {
	conn := new(mockConn)
	mc := &mysqlConn{
		buf:     newBuffer(conn),
		closech: make(chan struct{}),
	}

	// illegal empty (stand-alone) packet
	conn.data = []byte{0x00, 0x00, 0x00, 0x00}
	conn.maxReads = 1
	_, err := mc.readPacket()
	if err != ErrInvalidConn {
		t.Errorf("expected ErrInvalidConn, got %v", err)
	}

	// reset
	conn.reads = 0
	mc.sequence = 0
	mc.buf = newBuffer(conn)

	// fail to read header
	conn.closed = true
	_, err = mc.readPacket()
	if err != ErrInvalidConn {
		t.Errorf("expected ErrInvalidConn, got %v", err)
	}

	// reset
	conn.closed = false
	conn.reads = 0
	mc.sequence = 0
	mc.buf = newBuffer(conn)

	// fail to read body
	conn.maxReads = 1
	_, err = mc.readPacket()
	if err != ErrInvalidConn {
		t.Errorf("expected ErrInvalidConn, got %v", err)
	}
}

// https://github.com/go-sql-driver/mysql/pull/801
// not-NUL terminated plugin_name in init packet
func TestRegression801(t *testing.T) {
	conn := new(mockConn)
	mc := &mysqlConn{
		buf:      newBuffer(conn),
		cfg:      new(Config),
		sequence: 42,
		closech:  make(chan struct{}),
	}

	conn.data = []byte{72, 0, 0, 42, 10, 53, 46, 53, 46, 56, 0, 165, 0, 0, 0,
		60, 70, 63, 58, 68, 104, 34, 97, 0, 223, 247, 33, 2, 0, 15, 128, 21, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 98, 120, 114, 47, 85, 75, 109, 99, 51, 77,
		50, 64, 0, 109, 121, 115, 113, 108, 95, 110, 97, 116, 105, 118, 101, 95,
		112, 97, 115, 115, 119, 111, 114, 100}
	conn.maxReads = 1

	authData, pluginName, err := mc.readHandshakePacket()
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if pluginName != "mysql_native_password" {
		t.Errorf("expected plugin name 'mysql_native_password', got '%s'", pluginName)
	}

	expectedAuthData := []byte{60, 70, 63, 58, 68, 104, 34, 97, 98, 120, 114,
		47, 85, 75, 109, 99, 51, 77, 50, 64}
	if !bytes.Equal(authData, expectedAuthData) {
		t.Errorf("expected authData '%v', got '%v'", expectedAuthData, authData)
	}
}
