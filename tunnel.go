package main

import (
    "fmt"
    "flag"
    "net"
    "strconv"
    "io/ioutil"
    "os/user"
    "strings"
    "os"
    "time"
    "unsafe"
)

var mainIP = ""
var mainPort = ""
var fileLoc = ""


func main(){
	flagstate := flag.String("state","R","R for recieve, T for transmit")
	flagip := flag.String("ip","192.168.1.108","cominication device ip")
	flagport := flag.String("port","8080","cominication port")
	flagfile := flag.String("file","/","file to transmit")
	flag.Parse()

    if len(*flagip) > 0 &&  len(*flagip) < 4 {
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
    fileLoc = PreProcess(file)

    if *flagstate == "T" || *flagstate == "t" {
        if !IsFileExist(fileLoc) || IsDir(fileLoc){
            fmt.Println("File Does Not Exist")
            return
        }
    }

    if *flagstate == "R" || *flagstate == "r"{
        recieveFile()
    }else if *flagstate == "T" || *flagstate == "t" {
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
    }else{
        fmt.Println("Connected.")
    }
    start := time.Now()
    msg, err :=  ioutil.ReadAll(conn)
    if err != nil {
        fmt.Println("Message Read Error")
        panic(err)
    }
    end := time.Now()
    fmt.Printf("File Trasfer Ended in %v\n", end.Sub(start))
    namesize := msg[0]
    checkSum := ByteArrayToInt(msg[1:4 + 1])
    name := msg[4 + 1:namesize + 4 + 1]
    realmsg := msg[namesize + 4 + 1:]
    msgSize := len(msg)
    if msgSize == checkSum {
        fmt.Println("File Creating.")
        start = time.Now()
        dir := GetDesktop() + Sep() + string(name)
        OWrite(dir, realmsg)
        end = time.Now()
        fmt.Printf("File Creation Done in %v\n", end.Sub(start))
        fmt.Println("File Receive Sequence Successful")
        fmt.Println("File Saved Here:", dir)
    }else{
        fmt.Printf("%v bytes expected but %v bytes has receive.\n", checkSum, msgSize)
        fmt.Println("File Receive Sequence Faild!")
    }

}

func transmitfile(){
    name := dirToName(fileLoc)
    namesize := len([]byte(name))
    file := Read(fileLoc)
    checkSum := 1 + 4 + namesize + len(file)
    msg := make([]byte,0, checkSum)
    msg = append(msg, byte(namesize))
    msg = append(msg, IntToByteArray(checkSum, 4)...)
    msg = append(msg, []byte(name)...)
    msg = append(msg, file...)
    conn := connect("tcp",mainIP, mainPort)
    fmt.Println("Connected")
    start := time.Now()
    conn.Write(msg)
    conn.Close()
    end := time.Now()
    fmt.Printf("File Trasmission Sequence ended in %v\n", end.Sub(start))
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
    tokens := strings.Split(dir, Sep())
    name := tokens[len(tokens) - 1]
    return name
}

func PreProcess(dir string) string {
    if strings.HasPrefix(strings.ToLower(dir), "desk") {
        dir = strings.Replace(dir , "desk", GetDesktop(), -1)
    }else if strings.HasPrefix(strings.ToLower(dir), "curr"){
        dir = strings.Replace(dir , "curr", GetCurrentDir(), -1)
    }
    return dir
}

func IsDir(dir string) bool{
    fi, err := os.Stat(dir)
    if err != nil {
        return false
    }
    if fi.Mode().IsDir() {
        return true
    }
    return false
}

func IsFileExist(file string) bool {
    if _, err := os.Stat(file); os.IsNotExist(err){
        return false
    }
    return true
}

func GetDesktop() string{
    myself, _ := user.Current()
    var deskdir string = myself.HomeDir
    deskdir = deskdir + Sep() + "Desktop"
    return deskdir
}

func GetCurrentDir() string {
    wd, _ := os.Getwd()
    return wd
}

func Sep() string{
    return string(os.PathSeparator)
}

func Read(dir string) []byte{
    _, filedir := SplitDir(dir)
    buff, err := ioutil.ReadFile(filedir)
    if err != nil {
        fmt.Printf("Read File Error:%v\n", err)
    }else{
        return buff
    }
    return []byte{}
}

func SplitDir(dir string) (string, string){
    dir = PreProcess(dir)
    sp := Sep()
    tokens := strings.Split(dir, sp)
    dirPart := strings.Join(tokens[:len(tokens) - 1], sp)
    return dirPart, dir
}

func OWrite(dir string, buff []byte){
    newdir, newfile := SplitDir(dir)
    err := os.MkdirAll(newdir, os.ModePerm)
    if err != nil {
        fmt.Println("Make Directory Error:", err)
    }else{
        writeFile(newfile, buff)
    }
}

// main write function
func writeFile(filedir string, buffer []byte) {
    err := ioutil.WriteFile(filedir, buffer, os.ModePerm)   
    if err != nil {
        fmt.Printf("File Write Error:%v\n", err)
    }
}

func IntToByteArray(num int, size int) []byte {
    // size := int(unsafe.Sizeof(num))
    arr := make([]byte, size)
    for i := 0 ; i < size ; i++ {
        byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
        arr[i] = byt
    }
    return arr
}

func ByteArrayToInt(arr []byte) int{
    val := 0
    size := len(arr)
    for i := 0 ; i < size ; i++ {
        *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
    }
    return val
}