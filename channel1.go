package main

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		println(val)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

// The producer function sends 10 integers to the channel ch, and then closes the channel. The consumer function reads the integers from the channel ch and prints them to the console. The main function creates a channel ch, starts the producer goroutine, and then calls the consumer function to read the integers from the channel ch. The consumer function reads the integers from the channel ch until the channel is closed. When the channel is closed, the consumer function breaks out of the loop and terminates. The main function waits for the consumer function to finish before it exits. The producer function sends 10 integers to the channel ch, and then closes the channel. The consumer function reads the integers from the channel ch and prints them to the console. The main function creates a channel ch, starts the producer goroutine, and then calls the consumer function to read the integers from the channel ch. The consumer function reads the integers from the channel ch until the channel is closed. When the channel is closed, the consumer function breaks out of the loop and terminates. The main function waits for the consumer function to finish before it exits. The producer function sends 10 integers to the channel ch, and then closes the channel. The consumer function reads the integers from the channel ch and prints them to the console. The main function creates a channel ch, starts the producer goroutine, and then calls the consumer function to read the integers from the channel ch. The consumer function reads the integers from the channel ch until the channel is closed. When the channel is closed, the consumer function breaks out of the loop and terminates. The main function waits for the consumer function to finish before it exits. The producer function sends 10 integers to the channel ch, and then closes the channel. The consumer function reads the integers from the channel ch and prints them to the console. The main function creates a channel ch, starts the producer goroutine, and then calls the consumer function to read the integers from the channel ch. The consumer function reads the integers from the channel ch until the channel is closed. When the channel is closed, the consumer function breaks out of the loop and terminates. The main function waits for the consumer function to finish before it exits. The producer function sends 10 integers to the channel ch, and then closes the channel. The consumer function reads the integers from the channel ch and prints them to the console. The main function creates a channel ch, starts the producer goroutine, and then calls the consumer
