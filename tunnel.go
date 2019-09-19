package main

import (
    "fmt"
    "flag"
    "net"
    "quick/rw"
    "strconv"
    "io/ioutil"
    "strings"
)

var mainIP = ""
var mainPort = ""
var fileLoc = ""


func main(){
	flagtype := flag.String("type","R","R for recieve, T for transmit")
	flagip := flag.String("ip","192.168.1.108","cominication device ip")
	flagport := flag.String("port","20","cominication port")
	flagfile := flag.String("file","/","file to transmit")
	flag.Parse()

    if len(*flagip) == 3 {
        mainIP = "192.168.1." + *flagip
    }else{
        mainIP = *flagip
    }
    addr, err := net.ResolveIPAddr("ip", mainIP)
    if err != nil {


        fmt.Println("IP Resolving Error:", err)
        return
    }

    mainIP = addr.String()
    tempPort, err := strconv.Atoi(*flagport)

    if err != nil {
        fmt.Println("Port  Error:", err)
        return
    }

    if  tempPort < 1 && tempPort > 65535 {
        fmt.Println("Port  Error:", err)
       return
    }

    mainPort = *flagport

    file := *flagfile
    if file[0] == ' ' {
        file = file[1:]
  }
    fileLoc = rw.PreProcess(file)

    if *flagtype == "T" {
        if !rw.IsFileExist(fileLoc) || rw.IsDir(fileLoc){
            fmt.Println("File Does Not Exist")
            return
        }
    }

    if *flagtype == "R" {
    recieveFile()
    }else if *flagtype == "T" {
        transmitfile()
    }

}


func recieveFile(){
    fmt.Printf("Listen Has Begun at port:%v\n", mainPort)
    listener, err := net.Listen("tcp", ":" + mainPort)
    if err != nil {
        fmt.Println("Listen Error")
        panic(err)
    }
    defer listener.Close()
    conn, err := listener.Accept()
    if err != nil {
        fmt.Println("Listen Accept Error")
        panic(err)
    }
    msg, err :=  ioutil.ReadAll(conn)
    if err != nil {
        fmt.Println("Message Read Error")
        panic(err)
    }
    namesize := msg[0]
    name := msg[1:namesize + 1]
    realmsg := msg[namesize + 1:]
    dir := rw.GetDesktop() + rw.Sep() + string(name)
    rw.OWrite(dir, realmsg)
    fmt.Println("File Receive Sequence Successful")
    fmt.Println("Received File Saved Here:", dir)

}

func transmitfile(){

    name := dirToName(fileLoc)
    namesize := len([]byte(name))
    file := rw.Read(fileLoc)
    msg := make([]byte,0, 1 + namesize + len(file))
    msg = append(msg, byte(namesize))
    msg = append(msg, []byte(name)...)
    msg = append(msg, file...)
    conn := connect("tcp",mainIP, mainPort)
    conn.Write(msg)
    conn.Close()
    fmt.Println("File Trasmission Sequence successful")
}

func connect(protocol, ip, port string) net.Conn {
    conn, err := net.Dial(protocol, ip + ":" + port)
    if err != nil {
        fmt.Println("Connection Fail")
        panic(err)
    }
    return conn
}

func dirToName(dir string) string{
    tokens := strings.Split(dir, rw.Sep())
    name := tokens[len(tokens) - 1]
    return name
}
