package main

import (
	"fmt"
	"io"
)

type ArrayToSort struct {
	unsorted []int
	sorted   []int
}

func (sort *ArrayToSort) Sort() {

	for i := 1; i < len(sort.unsorted); i++ {
		for j := i; j > 0 && sort.unsorted[j-1] > sort.unsorted[j]; j-- {
			sort.sorted = append(sort.unsorted[:j-1], sort.unsorted[j], sort.unsorted[j-1])
			sort.unsorted = append(sort.sorted, sort.unsorted[j+1:]...)
		}
	}
}

func (a *ArrayToSort) GetUnsorted() {
	fmt.Println("Не сортированный массив выглядел так:", a.unsorted)
}

func (a *ArrayToSort) GetSorted() {
	fmt.Println("Сортированный массив выглядит так:", a.unsorted)
}

func main() {
	fmt.Println("Вводи цифры пока не надоест, а потом CRTL-D:")
	array := new(ArrayToSort)
	for {
		var digit int
		_, err := fmt.Scanln(&digit)
		if err == io.EOF {
			break
		}
		array.unsorted = append(array.unsorted, digit)
	}
	fmt.Println("1")
	array.GetUnsorted()
	array.Sort()
	array.GetSorted()

}
