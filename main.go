package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(1000)
	// sli = []int{1, 3, 46, 5, 2, 10, 2, 31, 18, 24, 30, 12, 9, 4}
	fmt.Println("s: ", num)
}

var wg sync.WaitGroup

var ch1 = make(chan os.Signal, 0)

func main() {
	/*
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
		go G3(&wg, stop)
		go G4(&wg, stop)

		go func() {
			signal.Notify(ch1, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
			<-ch1
			close(stop)
		}()

		wg.Wait()
	*/
}

    //求数组内所有数的异或运算值
func findoutMissingDigits(sli []int) int{
        var result = 0;
        for  i, _ := range  sli{
            result ^= sli[i];
        }
        return result;
}

     // 这个方法的返回值用来区分两个数，参数num是两个数的异或运算值
func getShedNum(num int){
	var numShift = 1;

	/**
	 * 这个while用来判断num从右往左数哪一位出现了1：
	 * 假如num的值是 10 = parseInt('1010',2),numShift的值是 1 = parseInt('0001'，2)
	 * 执行numShift = numShift<<1; 将numShift左位移一位得到 2 = numShift = parseInt('0010',2)
	 * 直到& 结果不为零得到使两个数异或运算结果不为1的那个数，两个数分别&numShift 得到的结果是0或1
	*/
	for res:=num & numShift; ; res==0 {
		numShift = numShift<<1
	}
	return numShift;
}

func  findoutMissingDigits2(arr []int)[]int{
	var A = 0;
	var B = 0;
	var result = findoutMissingDigits(arr);
	//从两个数的异或结果，获得数组分治的“分水岭”
	var shedNum = getShedNum(result);//tips:到这以后如果想用右位移吧1和0放到最后一位的话就计算shedNum转成字符串的字符串长度，右移长度-1位
	for i,_:= range arr{
		A = ((arr[i] & shedNum) == 0)? (A^arr[i]):A;
		B = ((arr[i] & shedNum) != 0)? (B^arr[i]):B;
	}
	return [A, B];
}

// <script>
 
 
    //  //这个方法的返回值用来区分两个数，参数num是两个数的异或运算值
    //  function getShedNum(num){
    //      var numShift = 1;
 
    //      /**
    //       * 这个while用来判断num从右往左数哪一位出现了1：
    //       * 假如num的值是 10 = parseInt('1010',2),numShift的值是 1 = parseInt('0001'，2)
    //       * 执行numShift = numShift<<1; 将numShift左位移一位得到 2 = numShift = parseInt('0010',2)
    //       * 直到& 结果不为零得到使两个数异或运算结果不为1的那个数，两个数分别&numShift 得到的结果是0或1
    //      */
    //      while((num & numShift) ==0){
    //          numShift = numShift<<1;
    //      }
 
    //      return numShift;
    //  }
 
//      function findoutMissingDigits2(arr){
//          var A = 0;
//          var B = 0;
//          var result = findoutMissingDigits(arr);
//          //从两个数的异或结果，获得数组分治的“分水岭”
//          var shedNum = getShedNum(result);//tips:到这以后如果想用右位移吧1和0放到最后一位的话就计算shedNum转成字符串的字符串长度，右移长度-1位
//          for (var i in arr){
//              A = ((arr[i] & shedNum) == 0)? (A^arr[i]):A;
//              B = ((arr[i] & shedNum) != 0)? (B^arr[i]):B;
//          }
//          return [A, B];
//      }
 
//      findoutMissingDigits2([1,1,2,2,3,3,4,5,6,6,7,7])//[4,5]
//  </script>