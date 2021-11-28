package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
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

//Поиск простых чисел алгоритмом Эратосфена, Аткинс это жесть
func getPrime(number int) []int {
	a := []int{2, 3}
	for i := 4; i <= number; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				break
			}
			if j == i-1 {
				a = append(a, i)
			}
		}
	}
	return a
}

//Калькулятор на операции +,-,*,/,sqrt,^
func simpleCalc(first float64, second float64, oper string) (float64, error) {
	switch oper {
	case "+":
		return first + second, nil
	case "-":
		return first - second, nil
	case "*":
		return first * second, nil
	case "/":
		if second == 0 {
			return 0, errors.New("Devide by zero, in hell you will be in a good company")
		}
		return first / second, nil
	case "^":
		return math.Pow(first, second), nil
	case "sqrt":
		return math.Sqrt(first), nil
	default:
		return 0, errors.New("Something goes wrong, but I don't know what exectly")
	}

}

func main() {
	oper := flag.String("operation", "", "Введите операцию, принимаются операции calcSquare - вычисление площади прямоугольника, calcCircleDL - определение длинны окружности и радиуса по плозади, divideByDigits - разбиение трехзначного чилсла на частные. Операцию * вводите в кавычках")
	first := flag.Int("first", 0, "Длинна левой стороны.")
	second := flag.Int("second", 0, "Длинна верхней стороны стороны.")
	flag.Parse()
	switch *oper {
	case "calcSquare":
		if *first == 0 || *second == 0 {
			flag.Usage()
		} else {
			fmt.Println("Будем считать площадь")

			square := calcSquare(*first, *second)
			fmt.Println("Площадь прямоугольника с длинной левой стороны", *first, "правой", *second, "=", square)
		}
	case "calcCircleDL":
		if *first == 0 {
			flag.Usage()
		} else {
			fmt.Println("Будем считать радиус и длинну")
			dia, length := calcCircleDL(*first)
			fmt.Println("Диаметр круга", dia, "длинна окружности", length, "для круга площадью")
		}
	case "divideByDigits":
		if *first == 0 {
			flag.Usage()
		} else {
			fmt.Println("Бей, круши, ломай")
			hundreds, dozens, units := divideByDigits(*first)
			fmt.Println("Число", *first, "содержит сотен:", hundreds, "десятков:", dozens, "единиц:", units)
		}
	case "+", "-", "/", "*", "^", "sqrt":
		res, err := simpleCalc(float64(*first), float64(*second), *oper)
		if err == nil {
			fmt.Println(*first, *oper, *second, "=", res)
		} else {
			log.Fatal(err)
		}
	case "getPrime":
		if *first < 3 {
			log.Fatal(errors.New("Так не интересно, первое простое число 2, давай заново."))
		}
		fmt.Println(getPrime(*first))

	default:
		flag.Usage()
	}

}
