package stack

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	return NewCalculator(s).Calculate()
}

// 操作符的优先级
var operatorPriority = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
	"(": 2,
	")": 2,
}

// Calculator 计算器
type Calculator struct {
	nums      []int
	operators []string
	exp       string
}

// NewCalculator NewCalculator
func NewCalculator(exp string) *Calculator {
	return &Calculator{
		nums:      []int{},
		operators: []string{},
		exp:       exp,
	}
}

func (c *Calculator) PushNums(v int) {
	c.nums = append(c.nums, v)
}

func (c *Calculator) PushOp(v string) {
	c.operators = append(c.operators, v)
}

func (c *Calculator) PopNums() int {
	if len(c.nums) == 0 {
		return 0
	}
	v := c.nums[len(c.nums)-1]
	c.nums = c.nums[:len(c.nums)-1]
	return v
}

func (c *Calculator) PopOp() string {
	if len(c.operators) == 0 {
		return ""
	}
	v := c.operators[len(c.operators)-1]
	c.operators = c.operators[:len(c.operators)-1]
	return v
}

// Calculate 获取计算结果
func (c *Calculator) Calculate() int {
	l := len(c.exp)
	for i := 0; i < l; i++ {
		switch e := c.exp[i]; e {
		case ' ':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			j := i
			for j < l && c.exp[j] <= '9' && c.exp[j] >= '0' {
				j++
			}
			n, _ := strconv.Atoi(c.exp[i:j])
			i = j - 1
			c.PushNums(n)
		case '+', '-', '*', '/':
			pre := c.PopOp()
			for pre != "" && pre != "(" && operatorPriority[string(e)] <= operatorPriority[pre] {
				c.PushNums(c.calc(pre))
				pre = c.PopOp()
			}
			if pre != "" {
				c.PushOp(pre)
			}
			c.PushOp(string(e))
		case '(':
			c.PushOp(string(e))
		case ')':
			for o := c.PopOp(); o != "(" && o != ""; o = c.PopOp() {
				c.PushNums(c.calc(o))
			}
		default:
			panic("invalid exp")
		}
	}
	o := c.PopOp()
	if o == "" {
		return c.PopNums()
	}
	return c.calc(o)
}

// calc 单次计算操作，o: 计算符
func (c *Calculator) calc(o string) int {
	b := c.PopNums()
	a := c.PopNums()

	fmt.Printf("%d %s %d\n", a, o, b)

	switch o {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}

	return 0
}
