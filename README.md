# tunnel
TCP file transfer program for devices that sharing same network

#### usage for receive:
./tunnel --status=r

with this commend device enters the listening state. with port 8080

#### usage for trasmit:
./tunnel --status=t --ip=local_ip_of_receiver --port=same_port_with_receiveing_device --file=filelocaation

with this command device transmits the file to the receiver receiving file save location is the Desktop/filename.file_extension
if your receiving devices local IP is starts with 192.168.1 you can just use last 3 digits of it.

sample use:

--ip=192. 198.1.108 

is equal to

--ip=108 

desk or curr prefix can be used to describe other directories

Sample use:

/home/you/Desktop/hello.txt           -> desk/hello.txt

/home/you/go/src/tunnel/otherfile.go  -> curr/otherfile.go

In the end of the transmission program shut down.
