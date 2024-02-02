package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
1) Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами:
a + b, a - b, a * b, a / b. Данные передаются в одну строку (смотри пример ниже).
Решения, в которых каждое число и арифметическая операция передаются с новой строки,
считаются неверными.
ВЫПОЛНЕНО

2) Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5…),
так и с римскими (I, II, III, IV, V…) числами.
*/
func calc(a int, b int, sign string) (result int) {
	switch {
	case sign == "+":
		return a + b
	case sign == "-":
		return a - b
	case sign == "/":
		return a / b
	case sign == "*":
		return a * b
	default:
		// обработка ошибки
		return 0
	}

}
func main() {
	// Получаем значения
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите значение")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ToUpper(text)
	fmt.Println("Текст: " + text)

	// Регулярки
	re, _ := regexp.Compile(`(\d+)(-|\+|//|\*)(\d+)`) // 1+2
	res := re.FindAllStringSubmatch(text, -1)
	//fmt.Println("res ", res) // [[2+2 2 + 2]]

	reRoman, _ := regexp.Compile(`(I|V|X+)(-|\+|//|\*)(I|V|X+)`) // 1+2
	resRoman := reRoman.FindAllStringSubmatch(text, -1)
	// fmt.Println("resRoman ", resRoman) // [[I+I I + I]]

	// Проверяем входящие значения
	if len(res) > 0 && len(res[0]) >= 4 {
		fmt.Println("res ", res) // [[2+2 2 + 2]]

		// Арабские цифры
		fmt.Println("res[0] ", res[0][1], res[0][2], res[0][3])
		fmt.Println("len(res[0]) ", len(res[0]))

		//  считаем
		num1, _ := strconv.Atoi(res[0][1])
		num2, _ := strconv.Atoi(res[0][3])

		// результат
		fmt.Println(calc(num1, num2, res[0][2]))
	} else if len(resRoman) > 0 && len(resRoman[0]) >= 4 {
		fmt.Println("resRoman ", resRoman) // [[I+I I + I]]
	} else {
		// Ошибка
	}
}
