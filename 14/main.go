package main

// detectType определяет тип переменной с использованием switch по типу
func detectType(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "chan"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	default:
		return "unknown"
	}
}

func main() {
	// Целое число
	var num int = 1
	println(detectType(num))

	// Строка
	var str string = "Hello, I'm string!"
	println(detectType(str))

	// Булево значение
	var boolType bool = true
	println(detectType(boolType))

	// Каналы разных типов
	var chInt chan int = make(chan int)
	println(detectType(chInt))

	var chString chan string = make(chan string)
	println(detectType(chString))

	var chBool chan bool = make(chan bool)
	println(detectType(chBool))

	var chInterface chan interface{} = make(chan interface{})
	println(detectType(chInterface))

	// Другие типы для демонстрации
	var sliceVal []int = []int{1}
	println(detectType(sliceVal))

	var floatVal float64 = 1.1
	println(detectType(floatVal))
}
