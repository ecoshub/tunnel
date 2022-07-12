package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	StateReceiver    string = "r"
	StateTransmitter string = "t"
)

var (
	flagState = flag.String("state", "R", "R for receive, T for transmit")
	flagIP    = flag.String("ip", "", "communication device ip")
	flagPort  = flag.String("port", "8080", "communication port")
	flagFile  = flag.String("file", "", "file to transmit")
	flagDest  = flag.String("dest", "", "destination directory")
)

func main() {

	flag.Parse()

	state, err := resolveState()
	if err != nil {
		fmt.Println(err)
		return
	}

	switch state {
	case StateReceiver:
		err := receiveFile()
		if err != nil {
			fmt.Println(err)
			return
		}
	case StateTransmitter:
		err := transmitFile()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func resolveIPFlag() (string, error) {
	if *flagIP == "" {
		return "", errors.New("ip flag is required")
	}
	addr, err := net.ResolveIPAddr("ip", *flagIP)
	if err != nil {
		fmt.Println("IP Resolving Error:", err)
		return "", errors.New("ip resolve error. check 'ip' flag. err: " + err.Error())
	}
	return addr.String(), nil
}

func resolvePort() (string, error) {
	tempPort, err := strconv.Atoi(*flagPort)
	if err != nil {
		return "", errors.New("port value must be a number between 0 ~ 65535. check 'port' flag")
	}
	if tempPort < 1 || tempPort > 65535 {
		return "", errors.New("given port not valid. check 'ip' flag")
	}
	return *flagPort, nil
}

func resolveState() (string, error) {
	state := strings.ToLower(*flagState)
	switch state {
	case StateReceiver:
		return StateReceiver, nil
	case StateTransmitter:
		return StateTransmitter, nil
	default:
		return "", errors.New("state is not valid. check 'state' flag")
	}
}
