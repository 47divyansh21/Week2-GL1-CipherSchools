package main

//channel.go
import (
	"fmt"
	"sync"
	"time"
)

func mainjs() {
	channel := make(chan string, 1)
	go func(ch chan<- string) {
		ch <- "2"
		// mess := <-ch
		fmt.Println(1)

	}(channel)
	message := <-channel
	time.Sleep(time.Second * 5)
	fmt.Println(message)
}
func mai2n() {
	channel := make(chan string, 1)
	go func(ch chan string) {
		mess := <-ch
		//ch <- "hey from anonymous"
		// "hey from anonymous" -> ch
		fmt.Println(mess)
		fmt.Println(1)

	}(channel)
	message := "Hello from MaAIN FUNCTION"
	channel <- message
	time.Sleep(time.Second * 5)
	message = <-channel
	fmt.Println(message)
}

func main() {
	channel := make(chan string, 1)
	go func(ch chan<- string) {
		ch <- "2"
		channel <- "3"
		time.Sleep(time.Second * 5)
		fmt.Println(1)

	}(channel)
	message := <-channel
	mess := <-channel
	time.Sleep(time.Second * 5)
	fmt.Println(message)
	fmt.Println(mess)
}

func main() {
	channel := make(chan string, 1)
	go func() {
		channel <- "2"
		channel <- "3"
		time.Sleep(time.Second * 5)
		fmt.Println(1)

	}()
	message := <-channel
	mess := <-channel
	time.Sleep(time.Second * 5)
	fmt.Println(message)
	fmt.Println(mess)
}

func main() {
	channel := make(chan string)
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		channel <- "Hello from Anonymous"
		time.Sleep(time.Second * 5)
		fmt.Println(1)
		wait.Done()
	}()
	go func() {
		message := <-channel
		time.Sleep(time.Second * 5)
		fmt.Println(message)
		wait.Done()
	}()
	time.Sleep(time.Second * 5)
	message := <-channel
	time.Sleep(time.Second * 5)
	fmt.Println(message)
	wait.Wait()
}

/*

//select.go


package main

import (
	"fmt"
	"time"
)

 func main() {
 	channel := make(chan int)
 	go func() {
 		channel <- 1
 		time.Sleep(time.Second)
 		channel <- 2
		close(channel)

 	}()
 	for ch := range channel {
 		fmt.Println(ch)
 	}
 }

func sendInt(ch chan int) {
	ch <- 1
}
func sendBool(ch chan bool) {
	ch <- true
}
func main321() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	go sendBool(ch2)
	go sendInt(ch1)
	select {
	case getInt := <-ch1:
		fmt.Println(getInt)
	case getBool := <-ch2:
		fmt.Println(getBool)
		 default:
		 	fmt.Println("bye")
	}
}

 func main() {
	helloCh := make(chan string, 1)
	goodByeCh := make(chan string, 1)
	quitCh := make(chan bool)
	go receiveMessage(helloCh, goodByeCh, quitCh)
	go sendMessage(helloCh, "Hello world")
	time.Sleep(time.Second)
	go sendMessage(goodByeCh, "Good Bye world")
	fmt.Println("message from quitCh=", <-quitCh)
 }
func sendMessage(ch chan<- string, message string) {
	ch <- message
}
func receiveMessage(helloCh, goodByeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case message := <-helloCh:
			fmt.Println("Message from helloCh:	", message)
		case message := <-goodByeCh:
			fmt.Println("Message from goodByeCh:	", message)
		case <-time.After(time.Second * 2):
			fmt.Println("Nothing receiving in 2 seconds:	Exiting the receiveMessage goroutine")
			quitCh <- true
			break
		}
	}
}

*/
