package stack

import (
	"fmt"
	"testing"
)

func TestFit(t *testing.T) {

}

func TestPMMD(t *testing.T) {

	// ( 4 - 3 ) * 5 - 2 + ( 6 - 2 ) / 2 * 3 - 3 + 10 = 16
	// [4 3 - 5 * 2 - 6 2 - 2 / 3 * + 1 - 5 +]
	// var exp = []string{"(", "4", "-", "3", ")", "*", "5", "-", "2", "+", "(", "6", "-", "2", ")", "/", "2", "*", "3", "-", "1", "+", "5"}
	// 9+(3-1)*3+10/2 =20
	// [9 3 1 - 3 * + 10 2 / +]
	// var exp = []string{"9", "+", "(", "3", "-", "1", ")", "*", "3", "+", "10", "/", "2"}

	var exp = "( 4 - 3 ) * 15 - 2 + ( 6 - 2 ) / 2 * 3 - 10 + 10"

	rear := midConvertRear(exp)
	fmt.Println(rear)
	result := rearCalculation(rear)
	fmt.Println(result)
}
