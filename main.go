package main

import (
	"fmt"
	"sync"
)

//使用channel 来模拟生产者和消费者:
//1个producer, 3 个 consumer.
//生产者生产 n 个数据(int即可)后结束, 消费者消费完所有数据后也结束(打印即算消费).
//要求:
//1. 编译运行通过, 且程序正常退出.
//2. 消费者必须消费完所有数据
func main() {
	var conNum = 3
	var ch = make(chan int, 20)
	var wg sync.WaitGroup

	wg.Add(1)
	go producer(ch, &wg)

	wg.Add(conNum)
	for i := 0; i < conNum; i++ {
		go func() {
			defer wg.Done()
			consumer(ch)
		}()
	}
	wg.Wait()
}

func producer(ch chan int, wg *sync.WaitGroup) {
	count := 100
	for i := 0; i < count; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func consumer(ch <-chan int) {
	for {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		} else {
			return
		}
	}
}

func te(ch <-chan int) {
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(msg)
	}
}

//tt()
//ChanT()
//time.Sleep(3 * time.Second)

/*
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

*/
