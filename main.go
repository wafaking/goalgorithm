package main

import "fmt"

func main() {
	var sli = []int{1, 1, 4, 2, 5, 3, 2, 5, 3, 5, 4, 4, 5}
	for i := 1; i < len(sli); i++ {
		sli[0] = sli[0] ^ sli[i]
		fmt.Println("sli[0]: ", sli[0])
	}
	fmt.Println("res: ", sli[0]) // 4
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