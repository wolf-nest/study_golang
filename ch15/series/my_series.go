package series

import "fmt"

func init() {
	fmt.Println("init...")
}
func init() {
	fmt.Println("init this package")
}
func Square(n int) int {
	return n * n
}

func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
