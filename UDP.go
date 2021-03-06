package main

import (
"net";
"fmt"
."time"
)

/*type UDPAddr() struct{
IP IP
Port int

}*/

func write(sock2 *net.UDPConn) {
	for {
		sock2.Write([]byte("Hello\n"))
		Sleep(100*Millisecond)
	}
}

func read(sock *net.UDPConn){
	for {
		buf := make([]byte, 1024)		
		rlen, raddr, _ := sock.ReadFromUDP(buf)
		//fmt.Printf("This is length: %s\n %s\n",rlen,raddr)
		if raddr.IP.String() != "129.241.187.157" {
			fmt.Println(string(buf))
		}
	_ = rlen
	_ = raddr
	}
}

func main() {
	address, err := net.ResolveUDPAddr("udp", "129.241.187.255:20019")//UDP is requred to broadcast and listen to the same port(?)
	if err != nil {
		fmt.Printf("code red")	
	}
	sock2, err := net.DialUDP("udp", nil, address)
	
	if err != nil {
			fmt.Printf("code cyan")
		}

	sock, err := net.ListenUDP("udp", address)
	if err != nil {
			fmt.Printf("code blue")	
		}
	
	go write(sock2)
	go read(sock)

	a := make(chan string)
	<-a

}
