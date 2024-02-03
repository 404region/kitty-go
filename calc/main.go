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

var romanArr = []string{
	"I",
	"II",
	"III",
	"IV",
	"V",
	"VI",
	"VII",
	"VIII",
	"IX",
	"X",
	"XI",
	"XII",
	"XIII",
	"XIV",
	"XV",
	"XVI",
	"XVII",
	"XVIII",
	"XIX",
	"XX",
	"XXI",
	"XXII",
	"XXIII",
	"XXIV",
	"XXV",
	"XXVI",
	"XXVII",
	"XXVIII",
	"XXIX",
	"XXX",
	"XXXI",
	"XXXII",
	"XXXIII",
	"XXXIV",
	"XXXV",
	"XXXVI",
	"XXXVII",
	"XXXVIII",
	"XXXIX",
	"XL",
	"XLI",
	"XLII",
	"XLIII",
	"XLIV",
	"XLV",
	"XLVI",
	"XLVII",
	"XLVIII",
	"XLIX",
	"L",
	"LI",
	"ЛИИ",
	"LIII",
	"LIV",
	"LV",
	"LVI",
	"LVII",
	"LVIII",
	"LIX",
	"LX",
	"LXI",
	"LXII",
	"LXIII",
	"LXIV",
	"LXV",
	"LXVI",
	"LXVII",
	"LXVIII",
	"LXIX",
	"LXX",
	"LXXI",
	"LXXII",
	"LXXXIII",
	"LXXIV",
	"LXXV",
	"LXXVI",
	"LXXVII",
	"LXXVIII",
	"LXXXIX",
	"LXXX",
	"LXXXI",
	"LXXXII",
	"LXXXIII",
	"LXXXIV",
	"LXXXV",
	"LXXXVI",
	"LXXXVII",
	"LXXVIII",
	"LXXXIX",
	"XC",
	"XCI",
	"XII",
	"XIII",
	"XCIV",
	"XCV",
	"XCVI",
	"XVII",
	"XCVIII",
	"ХХIХ",
	"C",
}

func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func translateRomanToArabic(romanNum string) (arabicNum int) {
	findEl := find(romanArr, romanNum)
	if findEl != -1 {
		fmt.Println("findEl + 1 ", findEl+1)
		return findEl + 1
	}
	return findEl
}

func translateArabicToRoman(arabicNum int) (romanNum string) {
	return romanArr[arabicNum-1]
}
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
	re, _ := regexp.Compile(`(\d+)([-+/*])(\d+)`) // 1+2
	res := re.FindAllStringSubmatch(text, -1)
	fmt.Println("res ", res) // [[2+2 2 + 2]]

	reRoman, _ := regexp.Compile(`([IVX]+)([-+/*])([IVX]+)`) // 1+2
	resRoman := reRoman.FindAllStringSubmatch(text, -1)
	// fmt.Println("resRoman ", resRoman) // [[I+I I + I]]

	// Проверяем входящие значения
	if len(res) > 0 && len(res[0]) >= 4 {
		fmt.Println("res ", res) // [[2+2 2 + 2]]

		// Арабские цифры
		fmt.Println("res[0] ", res[0][1], res[0][2], res[0][3])
		fmt.Println("len(res[0]) ", len(res[0]))

		//  конвертируем
		num1, _ := strconv.Atoi(res[0][1])
		num2, _ := strconv.Atoi(res[0][3])

		// результат
		fmt.Println("Результат: ")
		fmt.Println(calc(num1, num2, res[0][2]))
	} else if len(resRoman) > 0 && len(resRoman[0]) >= 4 {
		fmt.Println("resRoman ", resRoman) // [[I+I I + I]]
		// конвертируем в римские
		fmt.Println("res[0] ", resRoman[0][1], resRoman[0][2], resRoman[0][3])
		num1 := translateRomanToArabic(resRoman[0][1])
		num2 := translateRomanToArabic(resRoman[0][3])
		fmt.Println("translateArabicToRoman result : ", translateArabicToRoman(calc(num1, num2, resRoman[0][2])))
	} else {
		// Ошибка - вывести пользователю ошибку
	}
}
