package main

import (
	"fmt"
	"strconv"
)

const Max_Size = 100

type Stack struct {
	data [Max_Size]string
	top  int //下标
}

func (s *Stack) init() {
	s.top = -1
}

//压入栈中
func (s *Stack) push(elem string) {
	if s.top == Max_Size-1 {
		fmt.Println("the stack is full")
	} else {
		s.top++
		s.data[s.top] = elem
	}
}

//弹弹弹
func (s *Stack) pop() string {
	var elem string
	if s.top < 0 {
		fmt.Println("the stack is empty")
	} else {
		elem = s.data[s.top]
		s.top--
	}
	return elem
}

func main() {
	fmt.Println("Input an expression")
	var expr string
	fmt.Scanln(&expr)
	//中缀到后缀
	RPN := Get_RPN(expr)
	//计算后缀
	Calculate(RPN)

}

func Get_RPN(src string) [Max_Size]string {
	var dst [Max_Size]string
	var op Stack
	op.init()
	var j int
	var t string
	for i := 0; i < len(src); {
		//数字直接到dst
		if i < len(src) && (string(src[i]) >= "0" && string(src[i]) <= "9") {
			//for t = " "; i < len(src) && (string(src[i]) >= "0" && string(src[i]) <= "9"); i++ {
			for t = ""; i < len(src) && (string(src[i]) >= "0" && string(src[i]) <= "9"); i++ {
				t = t + string(src[i]) //从当前的src[i]开始，一直往后连续判断是否都是数字，将这串连续数字以string形式压入栈
			}
			dst[j] = t
			j++
		}

		//左括号 压入栈中
		if i < len(src) && string(src[i]) == "(" {
			op.push(string(src[i]))
			i++
		} else if i < len(src) && string(src[i]) == ")" {
			//遇到右括号 一直弹出栈中的符号到dst直到见到左括号
			for op.top != -1 && op.data[op.top] != "(" {
				dst[j] = op.pop()
				j++
			}
			op.pop() //pop"("
			i++
		} else if i < len(src) && op.top == -1 {
			op.push(string(src[i]))
			i++
			//栈中有(时，只有再遇)才会弹出，在此期间其他运算符直接入栈
		} else if i < len(src) && op.data[op.top] == "(" {
			op.push(string(src[i]))
			i++
			//原表达式中待进入栈的优先级更大时，直接入栈
		} else if i < len(src) && (string(src[i]) == "*" || string(src[i]) == "/") && (op.data[op.top] == "+" || op.data[op.top] == "-") {
			op.push(string(src[i]))
			i++
			//优先级更低，弹弹弹
		} else if i < len(src) {
			dst[j] = op.pop()
			j++
		}
	}
	//表达式读完后，栈区剩下的全部弹出来
	for op.top != -1 {
		dst[j] = op.pop()
		j++
	}
	return dst
}

func Calculate(RPN [Max_Size]string) {
	var num Stack
	num.init()
	for i := 0; i < len(RPN); i++ {
		if RPN[i] != "" {
			//因为dst是一串一串的
			if string(RPN[i][0]) >= "0" && string(RPN[i][0]) <= "9" {
				num.push(RPN[i]) //数字串直接入栈
			} else if RPN[i] == "+" {
				x, _ := strconv.Atoi(num.pop())
				y, _ := strconv.Atoi(num.pop())
				n := x + y
				num.push(strconv.Itoa(n))
			} else if RPN[i] == "-" {
				x, _ := strconv.Atoi(num.pop())
				y, _ := strconv.Atoi(num.pop())
				n := y - x
				num.push(strconv.Itoa(n))
			} else if RPN[i] == "*" {
				x, _ := strconv.Atoi(num.pop())
				y, _ := strconv.Atoi(num.pop())
				n := y * x
				num.push(strconv.Itoa(n))
			} else if RPN[i] == "/" {
				x, _ := strconv.Atoi(num.pop())
				y, _ := strconv.Atoi(num.pop())
				n := y / x
				num.push(strconv.Itoa(n))
			}
		}

	}
	fmt.Println(num.data[0])
}
