package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"testing"
	"text/tabwriter"
	"time"
)

func TestFunc(t *testing.T) {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Println("id: ", id)
	}

	var wg sync.WaitGroup
	var num = 1000
	wg.Add(num)
	for i := 0; i < num; i++ {
		go hello(&wg, i)
	}
	wg.Wait()
	fmt.Println("end ... ...")
}

func TestFunc2(t *testing.T) {

	producer := func(wg *sync.WaitGroup, l sync.Locker) { //1 defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1) //2
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutext\tMutex\n")
	for i := 0; i <= 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(tw, "%d\t%v\t%v\n",
			count, test(count, &m, m.RLocker()), test(count, &m, &m))

	}

}