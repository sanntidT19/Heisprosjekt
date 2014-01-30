package main

import (
"net";
"bufio";
"fmt";
."time"
)

func main() {

	raddr, err := ResolveUDPAddr("udp", "129.241.187.255:20019")

	c ,err :=  net.DialUDP("udp","129.241.187.157:20019")
	if err != nil {
     	    fmt.Printf("unable to connect to server, code red: %s\n", err.Error())
     	}
	
	
	
	for {
		
		Echo(raddr)
		Broadcast(c)
		Sleep(100*Millisecond)
	}
	c.Close()
	raddr.Close()
}

func Echo(raddr net.UDPConn) {
	line, err := bufio.NewReader(raddr).ReadString('\x00')
	if err != nil {
		fmt.Printf("Failure to read: %s\n",err.Error())	
		return
	}
	fmt.Println("Received: ", line)
}

func Broadcast(c net.UDPConn) {
	_, err = c.WriteToUDP([]byte("Hello yo yo hello\x00"))
	if err != nil {
		fmt.Printf("Failure to write: %s\n", err.Error())
		return	
	}
}
	// defer c.Close() //Closes the TCP connection so you dont abuse it, defer is called when the function returns.
