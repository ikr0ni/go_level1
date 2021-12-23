package main

import (
	"fmt"
	"io"
)

//Заворачиваем стрингу в []byte
type MyString struct {
	str []byte
}

//Реализуем метод Write для интерфеса Writer
func (s *MyString) Write(str []byte) (n int, err error) {
	s.str = str
	return len(str), nil
}

//Реализуем метод Read для интерфеса Reader
func (s MyString) Read(str []byte) (n int, err error) {
	fmt.Println(string(s.str))
	return len(str), io.EOF
}

func main() {
	var str MyString
	n, err := io.WriteString(&str, "БлаблаБла")
	if err != nil {
		panic("Приляжем")
	}
	str2, err := io.ReadAll(&str)
	if err != nil {
		panic("А ввод все не кончался")
	}
	fmt.Println(n, string(str2))
	fmt.Println(n, string(str.str))
}
