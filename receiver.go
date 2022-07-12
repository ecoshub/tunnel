package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func receiveFile() error {

	port, err := resolvePort()
	if err != nil {
		return err
	}

	addr := "0.0.0.0:" + port

	fmt.Printf("[info] listen started. addr: %s\n", addr)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		return err
	}

	fmt.Printf("[info] connection accepted. addr: %s\n", conn.RemoteAddr().String())

	start := time.Now()
	msg, err := ioutil.ReadAll(conn)
	if err != nil {
		return err
	}

	fmt.Printf("[info] File transfer done. Duration: %v\n", time.Since(start))

	if len(msg) == 0 {
		return errors.New("message length is 0")
	}

	p, err := decoder(msg)
	if err != nil {
		return err
	}

	fmt.Println(" - File info:")
	fmt.Println(" - Name: " + p.filename)
	fmt.Println(" - Destination: " + p.destinationDir)

	if p.destinationDir == "" {
		p.destinationDir = pathDownload
	}

	path := p.destinationDir + separator() + p.filename

	start = time.Now()
	err = ioutil.WriteFile(path, p.content, 0655)
	if err != nil {
		return err
	}

	sizeString := fileSizeToString(len(p.content))
	fmt.Printf("[info] File write done. Size:%s, Duration: %v\n", sizeString, time.Since(start))
	fmt.Printf("[info] File receive sequence success\n")
	fmt.Printf("[info] New file path: %s\n", path)

	return nil
}
