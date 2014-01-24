package main
import("net"; "fmt"; "bufio";"time")

func main(){
     connect, err := net.Dial("tcp", "129.241.187.161:34932")  //Hva legges i connect?
     if err != nil {
     	    fmt.Printf("unable to connect to server, code red: %s\n", err.Error())
     }
	fmt.Print(connect)
	fmt.Fprintf(connect, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(connect).ReadString('\n')
	fmt.Print(status)
	
