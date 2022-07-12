package main

import (
	"fmt"
	"os"
	"os/user"
	"unsafe"
)

var (
	pathDownload string
	dataSizeList = []string{"Bytes", "KB", "MB", "GB", "TB"}
)

func fileSizeToString(size int) string {
	sf := float64(size)
	c := 0
	for sf > 1 {
		sf = sf / 1024
		c++
	}
	if c > 0 {
		sf = sf * 1024
		c--
	}
	if c > 5 {
		return "NaN"
	}
	return fmt.Sprintf("%0.2f %s", sf, dataSizeList[c])
}

func init() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("[error] error getting current user. err: " + err.Error())
		os.Exit(0)
	}
	pathDownload = u.HomeDir + separator() + "Desktop"
}

func separator() string {
	return string(os.PathSeparator)
}

func IntToByteArray(num int, size int) []byte {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr
}

func ByteArrayToInt(arr []byte) int {
	val := 0
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}
