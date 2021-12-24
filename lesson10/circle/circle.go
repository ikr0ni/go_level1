package circle

import (
	"flag"
	"fmt"
	"math"
)

// Считаем площадь прямоугольника
func calcSquare(top int, left int) int {

	return top * left
}

// Считаем диаметр и длинну окружности по площади круга
func calcCircleDL(square int) (float64, float64) {
	length := math.Sqrt(float64(square) * 4 * math.Pi)
	dia := math.Sqrt(4 * float64(square) / math.Pi)
	return dia, length
}

//Подивайдим трехзначное на дигиты
func divideByDigits(number int) (int, int, int) {
	hundreds := number / 100
	dozens := (number % 100) / 10
	units := (number % 100) % 10
	return hundreds, dozens, units
}

func main() {
	fmt.Println("hello")
	oper := flag.String("operation", "", "Введите операцию, принимаются операции calcSquare - вычисление площади прямоугольника, calcCircleDL - определение длинны окружности и радиуса по плозади, divideByDigits - разбиение трехзначного чилсла на частные.")
	left := flag.Int("left", 0, "Длинна левой стороны. Данный параметр обязателен для calcSquare.")
	top := flag.Int("top", 0, "Длинна верхней стороны стороны. Данный параметр обязателен для calcSquare.")
	square := flag.Int("square", 0, "Площадь круга. Данный параметр обязателен для calcCircleDL.")
	number := flag.Int("number", 0, "Разбиваем число на сотни, десятки и единицы. Обязательный параметр для divideByDigits.")
	flag.Parse()
	if *oper == "" {
		flag.Usage()
	}
	switch *oper {
	case "calcSquare":
		if *left == 0 || *top == 0 {
			flag.Usage()
		} else {
			fmt.Println("Будем считать площадь")

			square := calcSquare(*left, *top)
			fmt.Println("Площадь прямоугольника с длинной левой стороны", *left, "правой", *top, "=", square)
		}
	case "calcCircleDL":
		if *square == 0 {
			flag.Usage()
		} else {
			fmt.Println("Будем считать радиус и длинну")
			dia, length := calcCircleDL(*square)
			fmt.Println("Диаметр круга", dia, "длинна окружности", length, "для круга площадью")
		}
	case "divideByDigits":
		if *number == 0 {
			flag.Usage()
		} else {
			fmt.Println("Бей, круши, ломай")
			hundreds, dozens, units := divideByDigits(*number)
			fmt.Println("Число", *number, "содержит сотен:", hundreds, "десятков:", dozens, "единиц:", units)
		}
	}

}
