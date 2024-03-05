package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	"github.com/oliverisaac/kasa-light-controller/pkg/messages"
	"github.com/oliverisaac/kasa-light-controller/pkg/xortransport"
	"github.com/sirupsen/logrus"
)

func main() {
	ip := "192.168.1.93"
	port := 9999

	mg := messages.MessageGenerator{
		Encrypter: xortransport.Encrypter{},
	}

	// Create a UDP address.
	dest := fmt.Sprintf("%s:%d", ip, port)
	logrus.Infof("Talking to %s", dest)
	addr, err := net.ResolveTCPAddr("tcp", dest)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a UDP connection.
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	sendMessage(conn, mg.On())
	for {
		sendMessage(conn, mg.HSV(40, 40+rand.Intn(30), 40+rand.Intn(30)))
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("Message sent and response received!")
}

func sendMessage(conn *net.TCPConn, msg []byte) error {
	dec, err := xortransport.DecryptBytes(msg)
	logrus.Infof("Sending string: %s", dec)
	logrus.Infof("Sending bytes: % x", msg)
	l, err := conn.Write(msg)
	if err != nil {
		return fmt.Errorf("Failed to write to conn: %w", err)
	}

	logrus.Infof("Wrote %d bytes", l)

	// Read the response.
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		return fmt.Errorf("Failed to read from conn: %w", err)
	}

	buf, _ = xortransport.DecryptBytes(buf[:n])
	logrus.Infof("Read %d bytes: %s", n, buf)
	return nil
}
