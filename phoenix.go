package main
import(
	"fmt"
	//"io"
	"os"
	"time"
	"bufio"
)

const(
	fileName = "phoenix.txt"
)

func process() {
	file, _  := os.Open(fileName)
	i := 0	
	//for {
		//t := time.Now()
		//f.Write("%v %i \n", t, i)
		file.Write([]byte("2014-02-20 10:34:43.179628499 +0100 CET 3 \n"))
		file.Write([]byte("2014-02-21 10:35:43.179628499 +0100 CET 4 \n"))
		file.Write([]byte("2014-02-22 10:34:46.179628499 +0100 CET 5 \n"))
		fmt.Println(i)
		i++
	//}

	read := bufio.NewReader(file)
	l,_ ,_ := read.ReadLine()
	fmt.Println(l)
	t,_ ,_ := read.ReadLine()
	fmt.Println(t)
	q,_ ,_ := read.ReadLine()
	fmt.Println(q)
	
	fmt.Println("Hello")
}

func createFile() {
	f, err := os.Create(fileName)
	if err != nil {panic(err)}
	defer f.Close()
}

func main() {
 	createFile()
	c1 make (chan 
	//t1 := time.Now()
	//fmt.Printf("%v\n", t1)
	c1 <- go process() //primary
	// go process() //backup
	time.Sleep(10000)
}



//compare with time.LocalTime()
