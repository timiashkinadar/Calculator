package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanSym = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
}

func printRes(ans int, romanFlag bool) {
	if romanFlag && ans > 0 {
		fmt.Printf("Ответ: %s\n", arabToRoman(ans))
	} else if romanFlag && ans <= 0 {
		fmt.Printf("Ответ: %s\n", "Результатом работы с римскими числами могут быть только положительные числа.")
	} else {
		fmt.Printf("Ответ: %d\n", ans)
	}
}

func arabToRoman(number int) string {
	arabs := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	roms := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var roman string
	for i, arab := range arabs {
		for number >= arab {
			roman = roman + roms[i]
			number -= arab
		}
	}
	return roman
}

func checkParam(num1, num2 int, sign string) error {
	if (num1 <= 0 || num1 >= 11) || (num2 <= 0 || num2 >= 11) {
		return fmt.Errorf("Числа %d и %d должны быть в диапазоне от 1 до 10 включительно.", num1, num2)
	}
	if len(sign) != 1 || !strings.Contains("+-*/", sign) {
		return fmt.Errorf("Калькулятор умеет выполнять только операции сложения, вычитания, умножения и деления.")
	}
	return nil
}

func parseNums(strNum1, strNum2 string) (num1, num2 int, romanFlag bool, err error) {
	if isRoman(strNum1) && isRoman(strNum2) {
		return romanToAr(strNum1), romanToAr(strNum2), true, nil
	}

	if isArab(strNum1) && isArab(strNum2) {
		num1, _ = strconv.Atoi(strNum1)
		num2, _ = strconv.Atoi(strNum2)
		return num1, num2, false, nil
	}
	return 0, 0, false, fmt.Errorf("Оба числа должны быть только целыми и арабскими или римскими одновременно.")
}

func romanToAr(str string) (x int) {
	res := 0
	max := 0
	for i := len(str) - 1; i >= 0; i-- {
		s := str[i]
		num := romanSym[rune(s)]
		if num >= max {
			max = num
			res += num
			continue
		}
		res -= num
	}
	return res
}

func isRoman(numStr string) bool {
	for _, sym := range numStr {
		_, exist := romanSym[sym]
		if !exist {
			return false
		}
	}
	return true
}

func isArab(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}

func startCalc(sign string, num1 int, num2 int) (ans int) {
	switch sign {
	case "-":
		ans = num1 - num2
	case "+":
		ans = num1 + num2
	case "*":
		ans = num1 * num2
	case "/":
		ans = num1 / num2
	}
	return ans
}

func getString() (str string) {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	str = strings.Trim(str, "\n")
	return str
}

func main() {
	for {
		str := getString()
		parts := strings.Split(str, " ")
		if len(parts) != 3 {
			fmt.Println("Формат математической операции не удовлетворяет заданию.")
		}
		num1, num2, romanFlag, err := parseNums(parts[0], parts[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		err = checkParam(num1, num2, parts[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		ans := startCalc(parts[1], num1, num2)
		printRes(ans, romanFlag)
	}
}
