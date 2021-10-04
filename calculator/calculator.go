package calculator

import (
	"calc-func/structures"
	"calc-func/utils"
	"fmt"
	"strconv"
)

// Calc -- Метод вычисления обратной польской записи
func Calc(items []string) (float64, error) {
	var stack []string
	/**
	В функциональной парадигме принято отдавать предпочтение рекурсии — для чистоты и прозрачности,
	вместо использовании простого перебора через for.
	*/
	res, err := processRpn(items, stack)

	if err != nil {
		return 0, err
	}

	return res, nil
}

func processRpn(items []string, stack []string) (float64, error) {
	stackOp := stack
	itemsOp := items
	res := 0.0
	var err error

	fmt.Println("-------------------")
	fmt.Println(items)
	fmt.Println(stack)

	if utils.IsFloat(itemsOp[0]) {
		stackOp = structures.Push(itemsOp[0], stackOp)
		itemsOp = itemsOp[1:len(itemsOp)]
	} else {
		//обрабатываем оператор
		var a, b string
		a, b, stackOp = structures.PopTwo(stackOp)
		fmt.Println("PopTwo")
		fmt.Println(a)
		fmt.Println(b)
		res, err = evalOperator(itemsOp[0], a, b)
		stackOp = structures.Push(fmt.Sprintf("%f", res), stackOp)
		itemsOp = itemsOp[1:len(itemsOp)]
	}
	if len(itemsOp) > 0 {
		res, err = processRpn(itemsOp, stackOp)
	}

	fmt.Println("res")
	fmt.Println(res)

	return res, err
}

func evalOperator(operator, left, right string) (float64, error) {

	bf, err := strconv.ParseFloat(right, 64)
	af, err := strconv.ParseFloat(left, 64)

	if err != nil {
		return 0, err
	}

	var operate func(x, y float64) float64
	operate = structures.GetOperation(operator)

	if operate != nil {
		return operate(af, bf), nil
	} else {
		return 0, fmt.Errorf("unexpected operator")
	}
}


