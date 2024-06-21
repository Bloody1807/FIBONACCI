package logic

import (
	"math"
)

type FibonacciService struct {
}

func (fs *FibonacciService) IsFibonacci(num int) bool {
	return fs.isPerfectSquare(5*num*num+4) || fs.isPerfectSquare(5*num*num-4)
}

func (fs *FibonacciService) GetNearestFibonacci(num int) int {
	for i := 1; ; i++ {
		if fs.IsFibonacci(num + i) {
			return num + i
		}
		if fs.IsFibonacci(num - i) {
			return num - i
		}
	}
}

func (fs *FibonacciService) GetAdjacentFibonacci(num int) (int, int) {
	if num == 0 {
		return 0, 1
	}
	a, b := 0, 1
	for b < num {
		a, b = b, a+b
	}
	return a, b + a
}

func (fs *FibonacciService) isPerfectSquare(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	return sqrt*sqrt == num
}
