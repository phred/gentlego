package main

import (
	"log"
	"os"
	"os/signal"
	"time"
	"syscall"
)

func main() {
	tick := time.Tick(time.Second)
	quit := make(chan os.Signal)
	winch := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	signal.Notify(winch, syscall.SIGWINCH)

	for {
		select {
		case <- tick:
			log.Println("Tick!")
		case <- winch:
			log.Println("WINCH!")			
		case <- quit:
			log.Println("Signal received, quitting")
			return
		}
	}
}


