package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/satori/uuid"
)

type Person struct {
	Name string
	Age  int
	Attr *Student
}

type Student struct {
	Addr string
	Name string
}

func main() {
	GenUUid()
}

func GenUUid() {
	uid := uuid.NewV4()

	traceID := fmt.Sprintf("%s-%s", time.Now().Format("20060102"), strings.Replace(uid.String(), "-", "", -1))

	h := md5.New()
	md5 := h.Sum([]byte("wafa"))
	r := fmt.Sprintf("md5: %x", md5)
	fmt.Printf("md5: %s, %d\n", r, len(r))

	uuids := uuid.NewV3(uid, "wafa")
	log.Println("uuid: ", traceID, len(traceID), uid, uuids.String())
	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	u2s := strings.Replace(u2.String(), "-", "", -1)
	fmt.Println("Successfully parsed: ", u2, u2s, len(u2s))

}

func GenerateRandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())

	b := []byte(str)
	var res []byte
	for i := 0; i < length; i++ {
		val := b[rand.Intn(len(str)-1)]
		res = append(res, val)
	}
	return string(res)
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
