package fib

import "fmt"

type Fib struct {
	FibMap map[int]int
	dig    int
	iter   int
}

func (f *Fib) Init() {
	f.FibMap = make(map[int]int)
	f.FibMap[0], f.FibMap[1] = 0, 1
	f.iter = 2
}

func (f *Fib) fibCalc() {
	if f.FibMap == nil {
		f.Init()
	}
	if f.iter <= f.dig {
		f.FibMap[f.iter] = f.FibMap[f.iter-2] + f.FibMap[f.iter-1]
		f.iter++
		f.fibCalc()
	}
}

func main() {
	fmt.Println("Введите количество итераций вычисления числа Фибоначи:")
	var fib Fib
	fmt.Scanln(&fib.dig)
	fib.fibCalc()
	fmt.Println(fib.FibMap)
}
