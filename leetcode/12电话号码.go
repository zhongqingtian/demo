package leetcode

import "strconv"

var num2string map[byte][]string = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func LetterCombinations(digits string) []string {
	res := []string{}
	if digits == "" {
		return res
	}

	res = append(res, "")
	for i := 0; i < len(digits); i++ {
		nums := num2string[digits[i]]
		pre := res
		res = []string{}
		for _, num := range nums {
			for _, v := range pre {
				res = append(res, num+v)
			}
		}
	}
	return res
}

func EvalRPN(tokens []string) int {
	number := []int{}

	for _, val := range tokens {
		l := len(number)
		switch val {
		case "+":
			number = append(number[:l-2], number[l-2]+number[l-1])
		case "-":
			number = append(number[:l-2], number[l-2]-number[l-1])
		case "*":
			number = append(number[:l-2], number[l-2]*number[l-1])
		case "/":
			number = append(number[:l-2], number[l-2]/number[l-1])
		default:
			n, _ := strconv.Atoi(val)
			number = append(number, n)
		}
	}
	return number[0]
}
