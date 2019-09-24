# tunnel
TCP file transfer program for devices that sharing same network

### build with "go build tunnel.go" and you are good to go. 

#### Receiver Side:
##### linux/mac:
./tunnel --state=r
##### windows:
tunnel --state=r

with this commend device enters the listening state. with port 8080

you can set an arbitrary port number just add --port=your_port_number to the command line. like this:

./tunnel --state=r --port=2020

the port numbers that below 1024 can ask root permission.

#### Transmitter Side:
##### linux/mac:
./tunnel --state=t --ip=local_ip_of_receiver --port=same_port_with_receiveing_device --file=filelocaation
##### windows:
tunnel --state=t --ip=local_ip_of_receiver --port=same_port_with_receiveing_device --file=filelocaation

with this command device transmits the file to the receiver, save location is the Desktop/filename.file_extension

if your receiving devices local IP is starts with 192.168.1 you can just use last 3 digits of it.

sample use:

--ip=192. 198.1.108 

is equal to

--ip=108 

desk or curr prefix can be used to describe other directories

Sample use:

/home/you/Desktop/hello.txt           -> desk/hello.txt

/home/you/go/src/tunnel/otherfile.go  -> curr/otherfile.go

sample usage for trasmitter side:
Reciever:
./tunnel --state=r --port=2020
Trasmitter:
./tunnel --state=t --port=2020 --ip=192.168.1.108 --file=/home/myname/Desktop/view.jpg

In the end of the transmission program shut down.
