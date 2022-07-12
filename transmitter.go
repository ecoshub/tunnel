package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func transmitFile() error {

	ip, err := resolveIPFlag()
	if err != nil {
		fmt.Println(err)
		return err
	}

	port, err := resolvePort()
	if err != nil {
		fmt.Println(err)
		return err
	}

	filePath := strings.TrimSpace(*flagFile)
	destinationDir := strings.TrimSpace(*flagDest)

	p, err := NewPacket(filePath, destinationDir)
	if err != nil {
		return err
	}

	msg := p.encoder()

	addr := ip + ":" + port

	fmt.Printf("[info] connecting to %s\n", addr)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Printf("[info] connection success. addr: %s\n", addr)

	start := time.Now()
	n, err := conn.Write(msg)
	if err != nil {
		return err
	}

	if n != len(msg) {
		return fmt.Errorf("connection write failed. sent byte size and message size are not same. sent byte: %d, message size: %d", n, len(msg))
	}

	fmt.Printf("[info] file send success. Duration: %s\n", time.Since(start))
	return nil
}
