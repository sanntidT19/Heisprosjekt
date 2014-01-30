package main

import (
"net";
"fmt"
"time"
)

func main() {
	var buf [1024]byte
	
	addr, err := net.ResolveUDPAddr("udp", "129.241.187.255:20019")//UDP is requred to broadcast and listen to the same port(?)
	sock, err := net.ListenUDP("udp", "129.241.187.157:20019")
	for {
	    rlen, remote, err := sock.ReadFromUDP(buf)
		
		fmt.Fprint(rlen)
		sleep(100*Millieseconds)			    
}
