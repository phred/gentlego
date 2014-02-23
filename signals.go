package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, os.Kill)
	fmt.Println("Started")

	// do some work ...

	sig := <- quit  // now wait

	fmt.Println("Got signal:", sig)
}

func init () {
	after := time.After(2 * time.Second)
	go func () {
		<- after
		proc,_ := os.FindProcess(os.Getpid())
		proc.Signal(os.Interrupt)
	}()
}
