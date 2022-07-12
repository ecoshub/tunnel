package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

// frame
// --------------------------------------------------------------------------------------
// | filename-length | filename | destination-path-length | destination-path | checksum |
// --------------------------------------------------------------------------------------
// |     1 byte     |   vary   |        1 byte           |       vary       |  4 bytes |
// --------------------------------------------------------------------------------------

type Packet struct {
	filename       string
	filePath       string
	destinationDir string
	content        []byte
}

func decoder(buff []byte) (*Packet, error) {

	filenameSize := buff[0]

	filename := buff[1 : 1+filenameSize]

	destinationDirSize := buff[1+filenameSize]

	destinationDir := buff[1+filenameSize+1 : 1+filenameSize+1+destinationDirSize]

	packetSize := ByteArrayToInt(buff[1+filenameSize+1+destinationDirSize : 1+filenameSize+1+destinationDirSize+4])

	fileContent := buff[1+filenameSize+1+destinationDirSize+4:]

	msgSize := len(buff)

	if msgSize != packetSize {
		return nil, fmt.Errorf("file transfer failed. %v bytes expected but %v bytes has receive", packetSize, msgSize)
	}

	return &Packet{
		filename:       string(filename),
		destinationDir: string(destinationDir),
		content:        fileContent,
	}, nil
}

func NewPacket(filePath, destinationDir string) (*Packet, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	_, filename := path.Split(filePath)
	return &Packet{
		filename:       filename,
		filePath:       filePath,
		destinationDir: destinationDir,
		content:        file,
	}, nil
}

func (p *Packet) encoder() []byte {

	filenameSize := len(p.filename)
	destinationDirSize := len(p.destinationDir)
	packetSize := 1 + 4 + filenameSize + len(p.content) + 1 + destinationDirSize

	// NOTE replace 'append' with 'copy' after
	msg := make([]byte, 0, packetSize)
	msg = append(msg, byte(filenameSize))
	msg = append(msg, []byte(p.filename)...)
	msg = append(msg, byte(destinationDirSize))
	msg = append(msg, []byte(p.destinationDir)...)
	msg = append(msg, IntToByteArray(packetSize, 4)...)
	msg = append(msg, p.content...)
	return msg
}
