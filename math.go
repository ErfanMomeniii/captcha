package captcha

import (
	"fmt"
	"math/rand"
	"strconv"
)

// evalExpr computes a op b for op in {+,-,*}.
func evalExpr(a, b int, op byte) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	default:
		return 0
	}
}

// Math generates an arithmetic captcha. The image shows an expression such as
// "3 + 5 =" and Result.Text holds the answer ("8").
func (c *Captcha) Math() (*Result, error) {
	ops := []byte{'+', '-', '*'}
	op := ops[rand.Intn(len(ops))]

	var a, b int
	switch op {
	case '*':
		a, b = rand.Intn(9)+1, rand.Intn(9)+1
	default:
		a, b = rand.Intn(20)+1, rand.Intn(20)+1
		if op == '-' && b > a {
			a, b = b, a // keep the answer non-negative
		}
	}

	answer := strconv.Itoa(evalExpr(a, b, op))
	display := fmt.Sprintf("%d %c %d =", a, op, b)

	im, err := c.draw(display)
	if err != nil {
		return nil, err
	}
	return &Result{Text: answer, Image: im}, nil
}
