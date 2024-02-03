package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func checkNums(num1 int, num2 int) {
	if num1 < 1 || num2 < 1 || num1 > 10 || num2 > 10 {
		panic("Входные значения меньше 1 или больше 10")
	}
}

func main() {
	// Получаем значения
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите математическое выражение с использованием знаков -+/* и арабских или римских цифр:")
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ToUpper(text)

	// Регулярки
	matched, _ := regexp.MatchString(`[.,]`, text)
	if matched {
		panic("Калькулятор умеет работать только с целыми числами.")
	}

	re, _ := regexp.Compile(`(\d+)([-+/*])(\d+)`)
	res := re.FindAllStringSubmatch(text, -1)

	reRoman, _ := regexp.Compile(`([IVX]+)([-+/*])([IVX]+)`)
	resRoman := reRoman.FindAllStringSubmatch(text, -1)

	// Проверяем входящие значения
	if len(res) > 0 && len(res[0]) >= 4 {

		//  конвертируем
		num1, _ := strconv.Atoi(res[0][1])
		num2, _ := strconv.Atoi(res[0][3])

		checkNums(num1, num2)

		// результат
		fmt.Println(calc(num1, num2, res[0][2]))
	} else if len(resRoman) > 0 && len(resRoman[0]) >= 4 {
		// конвертируем в римские
		num1 := translateRomanToArabic(resRoman[0][1])
		num2 := translateRomanToArabic(resRoman[0][3])
		result := calc(num1, num2, resRoman[0][2])

		if result <= 0 || result > len(romanArr) {
			panic("Результат выходит за возможные границы для римских чисел.")
		}

		// результат
		fmt.Println(translateArabicToRoman(result))
	} else {
		panic("Введены неверные данные!")
	}
}
