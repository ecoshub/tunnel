# Tunnel
#### File transfer program for devices that sharing same network (TCP/IP)
#### Windows/Linux/Mac compatible.
#### Build with "go build tunnel.go" and you are good to go. 

## Simple usage:
### Receiver Side:
./tunnel --state=r
### Transmitter Side:
./tunnel --state=t --ip=local_IP_of_receiver --port=same_port_with_receiveing_device --file=filelocaation

*do not use './' prefix with windows. just a reminder.*

**There is another sample usage on the bottom of this page**

## Other functionalities

* **--state** : Defines computers comminication state. Default device state is "r" for receive. if you want to send a file your state must be "t" for trasmit.

* **--port** : Device comminication port. Default port number is 8080 but you can set an arbitrary port number like this,

**./tunnel --state=r --port=2020**

the port numbers that below 1024 can ask root permission.

* **--ip** : IP flag is only nessesary for transmitter side. It is local IP of receiver side. if your receiving device local IP is starts with 192.168.1 you can just write last byte.

**sample use:**

**--ip=108**

**is equal to**

**--ip=192.168.1.108**

* **--file** : File directoy of the file that you want to send.

desk or curr prefix can be used to describe other directories quickly.

Sample use:

*/home/you/Desktop/hello.txt*                  -> **desk/hello.txt**

*/home/you/where_tunnel_running/otherfile.go*  -> **curr/otherfile.go**

* **--dest** : Destination directory. Default is desktop. This flag is sets the destination directory of the file that you want to send.


## Sample Usage:

### computer 1 
* IP = 192.168.1.104
* receiver side

### computer 2
* IP = 192.168.1.101
* transmitter side

**objective: send a file from computer 2 to computer 1**
**file directory is '/home/comp2/Desktop/tunnel.go'**
**port is 2020** *arbitrary port number* *optional, just for example*
**destination directory is /home/comp1/Desktop/Downloads** *optional, just for example*

#### computer 1:
./tunnel --port=2020

#### computer 2
./tunnel --state=t --port=2020 --IP=104 --file=desk/tunnel.go --dest=desk/Downloads

**This procedure sends the 'tunnel.go' file from computer 2 to computer 1. Creates nessesary directories ( /home/comp1/Desktop/Downloads ) and save 'tunnel.go' filr in it.
