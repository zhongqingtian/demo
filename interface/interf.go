package main
import (
	"fmt"
)

//定义interface
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (t *MyString)Tmethod()  {
	fmt.Println("我是一个测试方法")
}
//实现接口
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Sam Anderson") // 类型转换
	var v VowelsFinder // 定义一个接口类型的变量
	//fmt.Printf("Vowels are %c\n", v.FindVowels()) //没有 把结构体地址复制给接口 接口类型没有指向会报错
        	v = name
	fmt.Printf("Vowels are %c\n", v.FindVowels())
    fmt.Printf("Vowels are %c\n",name.FindVowels())
    name.Tmethod() //接口v 只能调用接口的方法
}
