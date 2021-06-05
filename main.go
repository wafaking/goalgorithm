package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	//tt()
	//ChanT()
	//time.Sleep(3 * time.Second)

}

type People struct {
	Name string
	Age  int
}

func ChanT() {
	conNum := 3
	chBuf := 10
	chSli := make([]chan *People, conNum)

	for i := 0; i < conNum; i++ {
		chSli[i] = make(chan *People, chBuf)
	}

	wg := &sync.WaitGroup{}
	total := 20

	wg.Add(1)
	go func() {
		for i := 0; i < total; i++ {
			chOrder := i % conNum
			peo := &People{Name: fmt.Sprintf("name-%d", i), Age: i}
			chSli[chOrder] <- peo
		}

		for i := 0; i < conNum; i++ {
			close(chSli[i])
		}
		wg.Done()
	}()

	wg.Add(conNum)
	for i := 0; i < conNum; i++ {
		go func(i int) {
			for {
				select {
				case v, ok := <-chSli[i]:
					if !ok {
						log.Println(".... .. . . ....ok: ", ok)
						wg.Done()
						return
					}
					fmt.Printf("ch%d: %#v\n", i, v)
				}
			}
		}(i)
	}

	wg.Wait()
	time.Sleep(3 * time.Second)
	log.Println("==========end... ...==========")
}

func tt() {
	// 值类型
	var a = 5
	b := a

	// 指针类型
	var m = []int{1, 2, 3}
	n := m
	log.Printf("%v, %v\n", &a, &b) // 0xc000098000, 0xc000098008
	log.Printf("%v, %v\n", &m, &n) // &[1 2 3], &[1 2 3]
	n[2] = 0
	log.Printf("%v, %v\n", &m, &n) // &[1 2 0], &[1 2 0]

}
