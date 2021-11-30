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
	sort.sorted = make([]int, len(sort.unsorted))
	copy(sort.sorted, sort.unsorted)
	for i := 1; i < len(sort.sorted); i++ {
		for j := i; j > 0 && sort.sorted[j-1] > sort.sorted[j]; j-- {
			temp := append(sort.sorted[:j-1], sort.sorted[j], sort.sorted[j-1])
			sort.sorted = append(temp, sort.sorted[j+1:]...)
		}
	}
}

func (a *ArrayToSort) GetUnsorted() {
	fmt.Println("Не сортированный массив выглядел так:", a.unsorted)
}

func (a *ArrayToSort) GetSorted() {
	fmt.Println("Сортированный массив выглядит так:", a.sorted)
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
	array.Sort()
	array.GetUnsorted()
	array.GetSorted()

}
