package main

import (
	"fmt"
	"io"
)

type OneWayList struct {
	value int
	next  *OneWayList
}

//Кладем в список элементы рекурсией
func (l *OneWayList) PutElement(dig int) {
	if l.next != nil {
		l.next.PutElement(dig)
	} else {
		next := OneWayList{value: dig}
		l.next = &next
	}
}

//Получаем все элементы по порядку
func (l OneWayList) GetAll() {
	for {
		fmt.Println(l)
		if l.next == nil {
			break
		} else {
			l = *l.next
		}
	}
}

//Разворот рекурсией
func (l OneWayList) ReverseByRecurse() {
	if l.next != nil {
		l.next.ReverseByRecurse()
	}
	fmt.Println(l)
}

//Разворот циклом
func (l OneWayList) ReverseByLoop() {
}

func main() {
	fmt.Println("Вводи элементы списка:")
	first := OneWayList{}
	firstElem := true
	for {
		var digit int
		_, err := fmt.Scanln(&digit)
		if err == io.EOF {
			break
		}
		if firstElem == true {
			first.value = digit
			firstElem = false
		} else {
			first.PutElement(digit)
		}

	}

	first.GetAll()
	first.ReverseByRecurse()
}
