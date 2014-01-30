package main

import (
"net";
"bufio";
"fmt";
."time"
)

func main() {
	c ,err :=  net.Dial("tcp","129.241.187.161:33546")
	
	if err != nil {
     	    fmt.Printf("unable to connect to server, code red: %s\n", err.Error())
     }
	line := "Connect to: 192.241.187.153\x00"
	_, err = c.Write([]byte(line))
	if err != nil {
     	    fmt.Printf("unable to write to server, code red: %s\n", err.Error())
	}
for {
Echo(c)
}
c.Close()
}

func Echo(c net.Conn) {
	Sleep(100*Millisecond)
	line, err := bufio.NewReader(c).ReadString('\x00')
	if err != nil {
		fmt.Printf("Failure to read: %s\n",err.Error())	
		return
	}
	fmt.Println("Received: ", line)
	_, err = c.Write([]byte("Hello yo yo hello\x00"))
	if err != nil {
		fmt.Printf("Failure to write: %s\n", err.Error())
		return	
	}
	// defer c.Close() //Closes the TCP connection so you dont abuse it, defer is called when the function returns.
}
