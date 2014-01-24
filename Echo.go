package main
import ("net"; 
	"fmt"; 
	"bufio"; 
	."time")

func main(){
	
	serverAddr, _ := net.ResolveTCPAddr("tcp", "129.241.187.161:33546")
	connect, err := net.DialTCP("tcp", nil, serverAddr)
	
	if err != nil {
		fmt.Printf("Fail: %s\n",err.Error())	
	}

	for {
		Echo(connect)}

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
	defer c.Close()  //Closes the TCP connection so you dont abuse it, defer is called when the function returns.
}
