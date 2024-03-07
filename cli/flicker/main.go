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
	"github.com/spf13/pflag"
)

var mg messages.MessageGenerator

func init() {
	mg = messages.MessageGenerator{
		Encrypter: xortransport.Encrypter{},
	}
}

func main() {
	hue := pflag.IntP("hue", "h", 45, "Hue of the flicker")
	saturation := pflag.IntP("saturation", "s", 50, "Saturation of the flicker")
	brightness := pflag.IntP("brightness", "b", 80, "brightness of the flicker")
	variance := pflag.IntP("variance", "v", 10, "Variance of the flicker")
	pflag.Parse()

	// Create a UDP socket
	addr, err := net.ResolveUDPAddr("udp", ":4444")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := mg.GetInfo()
	dec := xortransport.DecryptBytes(msg)
	logrus.Infof("Sending string: %s", dec)
	logrus.Infof("Sending bytes: % x", msg)
	// Send a broadcast message
	_, err = conn.WriteToUDP(msg, &net.UDPAddr{IP: net.IPv4bcast, Port: 9999})
	if err != nil {
		fmt.Println(err)
		return
	}

	clientCh := make(chan string)
	go func() {
		for {
			select {
			case addr := <-clientCh:
				logrus.Infof("Talking to %s", addr)
				go flickerLight(addr, 9999, *hue, *saturation, *brightness, *variance)
			}
		}
	}()

	// Receive a broadcast message
	buf := make([]byte, 4096)
	for {
		logrus.Info("Reading from udp...")
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Received message from", addr)
		fmt.Println(string(xortransport.DecryptBytes(append([]byte{0, 0, 0, 0}, buf[:n]...))))
		clientCh <- fmt.Sprintf("%s", addr.IP)
	}
}

func flickerLight(ip string, port int, hue int, saturation int, brightness int, variance int) {
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
		_, err := sendMessage(conn, mg.HSV(
			hue,
			saturation-(variance/2)+rand.Intn(variance),
			brightness-(variance/2)+rand.Intn(variance),
		))
		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func sendMessage(conn *net.TCPConn, msg []byte) ([]byte, error) {
	dec := xortransport.DecryptBytes(msg)
	logrus.Debugf("Sending string: %s", dec)
	logrus.Tracef("Sending bytes: % x", msg)
	l, err := conn.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Failed to write to conn: %w", err)
	}

	logrus.Tracef("Wrote %d bytes", l)

	// Read the response.
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("Failed to read from conn: %w", err)
	}
	return buf[:n], nil
}
