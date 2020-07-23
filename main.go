package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

var ch1 = make(chan os.Signal, 0)

func main() {
	var stop = make(chan struct{}, 0)

	var num1 = 5
	var num2 = 5

	wg.Add(num1 + num2)
	for i := 0; i < num1; i++ {
		go func() {
			G1(&wg, stop)
		}()
	}

	for i := 0; i < num2; i++ {
		go func() {
			G2(&wg, stop)
		}()
	}
	// go G3(&wg, stop)
	// go G4(&wg, stop)

	go func() {
		signal.Notify(ch1, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
		<-ch1
		close(stop)
	}()

	wg.Wait()
}

func G1(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			return
		default:
			fmt.Println("g1")
			time.Sleep(5 * time.Second)
			fmt.Println("end g1.......")

			time.Sleep(3 * time.Second)
		}
	}
}

// func G2() {
func G2(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			// break
			return
		default:
			fmt.Println("g2")
			time.Sleep(5 * time.Second)
			fmt.Println("end g2.......")
			time.Sleep(3 * time.Second)
		}
	}
}

func G3(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			// break
			return
		default:
			fmt.Println("g3")
			time.Sleep(5 * time.Second)
			fmt.Println("end g3.......")
			time.Sleep(3 * time.Second)

		}
	}
}

func G4(wg *sync.WaitGroup, ch chan struct{}) {
	// func G4() {
	defer wg.Done()
	for {
		select {
		case <-ch:
			// break
			return
		default:
			fmt.Println("g4")
			time.Sleep(5 * time.Second)
			fmt.Println("end g4.......")
			time.Sleep(3 * time.Second)

		}

	}
}
