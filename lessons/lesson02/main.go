package main

import (
	"fmt"
)

func FibonacciIterative(n int) int {
	// TODO: імплементуйте ітераційно.
	// Підказка: тримайте два останні значення й оновлюйте їх у циклі.
	// Вхід вважаємо: n >= 0.
	// При отриманні негативного n повертаємо його без змін

	if n < 0 || n == 0 || n == 1 {
		return n
	}
	n1 := 0
	n2 := 1
	for i := 2; i <= n; i++ {
		temp := n1 + n2
		n1 = n2
		n2 = temp
	}

	return n2
}

func FibonacciRecursive(n int) int {
	// TODO: імплементуйте рекурсивно.
	// База: n==0 -> 0; n==1 -> 1.
	// Рекурсія: F(n-1)+F(n-2)
	// При отриманні негативного n повертаємо його без змін

	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func IsPrime(n int) bool {
	// TODO: імплементуйте перевірку на просте число.
	// Підказка: n<=1 -> false; 2 -> true; парні >2 -> false;
	// Далі перевіряйте дільники до sqrt(n).

	if n == 2 {
		return true
	}

	if n <= 1 || n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsBinaryPalindrome(n int) bool {
	// TODO: імплементуйте перевірку числа на те що це паліндром.
	// Підказка: перетворіть n у строку (strconv ефективніший за fmt)
	// потім перевірте паліндромність.

	if n < 0 {
		return false
	}
	bin := ""
	x := n
	if x == 0 {
		bin = "0"
	} else {
		for x > 0 {
			if x%2 == 0 {
				bin = "0" + bin
			} else {
				bin = "1" + bin
			}
			x = x / 2
		}
	}
	for i := 0; i < len(bin)/2; i++ {
		if bin[i] != bin[len(bin)-1-i] {
			return false
		}
	}
	return true
}

func ValidParentheses(s string) bool {
	// TODO: імплементуйте перевірку дужок.
	// Правила:
	// 1. Допустимі дужки (, [, {, ), ], }
	// 2. У кожної відкритої дужки є відповідна закриваюча дужка того ж типу
	// 3. Закриваючі дужки стоять у правильному порядку
	// "[{}]" - правильно
	// "[{]}" - не правильно
	// 4. Кожна закриваюча дужка має відповідну відкриваючу дужку
	// Підказка: використовуйте стек (можна зробити через масив рун []rune)

	stack := []rune{}
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			// Если стек пустой — нет пары
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (ch == ')' && top != '(') ||
				(ch == ']' && top != '[') ||
				(ch == '}' && top != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}

func Increment(num string) int {
	// TODO: Імплементуйте функцію яка на вхід отримує строку яка складається лише з символів 0 та 1
	// Тобто строка містить певне число у бінарному вигляді
	// Потрібно повернути число на один більше
	// Додайте валідацію вхідної строки, якщо вона містить недопустимі символи, повертайте 0

	if num == "" {
		return 0
	}
	value := 0
	for _, ch := range num {
		if ch != '0' && ch != '1' {
			return 0
		}
		value = value*2 + int(ch-'0')
	}
	return value + 1
}

func main() {
	// Невеликі демонстраційні виклики (для наочного запуску `go run .`)
	fmt.Println("FibonacciIterative(10):", FibonacciIterative(10)) // очікуємо 55
	fmt.Println("FibonacciRecursive(10):", FibonacciRecursive(10)) // очікуємо 55

	fmt.Println("IsPrime(2):", IsPrime(2))   // true
	fmt.Println("IsPrime(15):", IsPrime(15)) // false
	fmt.Println("IsPrime(29):", IsPrime(29)) // true

	fmt.Println("IsBinaryPalindrome(7):", IsBinaryPalindrome(7)) // true (111)
	fmt.Println("IsBinaryPalindrome(6):", IsBinaryPalindrome(6)) // false (110)

	fmt.Println(`ValidParentheses("[]{}()"):`, ValidParentheses("[]{}()")) // true
	fmt.Println(`ValidParentheses("[{]}"):`, ValidParentheses("[{]}"))     // false

	fmt.Println(`Increment("101") ->`, Increment("101")) // 6
}
