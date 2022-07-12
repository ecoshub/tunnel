# Tunnel
Local network file transfer program (TCP/IP). Windows/Linux/Mac support.

## How to Build
    go build .

## How to Use
Receiver side
```bash
    # run tunnel without flags
    ./tunnel
```

Transmitter side
```bash
    ./tunnel --state <t|r> --ip <IP_of_receiver> --file <file_path> --dest <file_destination_dir>
```

## Flags

-   **--state:** Defines transfer state it can be `receiver` or `transmitter`. Default state is `r` for receive. if you want to send a file your state must be transmitter adn use state `t`.

-   **--port:** Communication port. Default port number is 8080 but you can set an arbitrary port number like this.
-    **--ip:** IP flag is only required for transmitter side. It is local IP of receiver side. if your receiving device local IP is starts with 192.168.1 you can just write last octet.
-   **--dest:** Destination Directory for received file.

example:
```bash
    # receiver ip: 192.168.1.104
    # file path: /home/eco/Desktop/hello.txt
    # dest. dir: /home/sbl/Downloads
    # port: 2020
    
    # receiver
    ./tunnel --port "2020"

    # transmitter
    ./tunnel --state "t" --port "2020" --ip "104" --file "desk/hello.txt" --dest "/home/sbl/Downloads"
```
