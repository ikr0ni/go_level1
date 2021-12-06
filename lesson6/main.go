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

//Разворот циклом
func (l *OneWayList) ReverseByLoop() (reversedList *OneWayList) {
	//Ставим первую метку, откуда начинаем разворот
	currentElement := l

	for {
		if currentElement == nil {
			break
		}
		// Помещаем следующий элемент во времянку
		temporaryElement := currentElement.next
		// Помещаем в следующий элемент текущеий элемент в reversed листе
		currentElement.next = reversedList
		// назначаем в reversed листе текущим элементом текущий элемент из начального списка
		reversedList = currentElement
		//переводим метку на следующий элемент списка
		currentElement = temporaryElement
	}

	return
}

//Ну и чисто поржатьразвернем вывод спомощью defer
func (l *OneWayList) ReversByDefer() {
	for {
		if l == nil {
			break
		}
		defer fmt.Println(l)
		l = l.next
	}
}

//Псевдоразворот рекурсией
func (l OneWayList) ReversByRecurse() {
	if l.next != nil {
		l.next.ReversByRecurse()
	}
	fmt.Println(l)
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
	fmt.Println("Выведем первоначальный список:")
	first.GetAll()
	fmt.Println("Выведем рекурсией список в обратном порядке:")
	first.ReversByRecurse()
	reversedList := first.ReverseByLoop()
	fmt.Println("Выведем список после разворота циклом:")
	reversedList.GetAll()
	fmt.Println("А сейчас еще и defer используем, но уже на reversed:")
	reversedList.ReversByDefer()

}
