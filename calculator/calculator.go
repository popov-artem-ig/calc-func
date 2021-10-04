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
	// создаем новые массивы для "неизменяемости" наших данных
	stackOp := stack
	itemsOp := items
	res := 0.0
	var err error

	if utils.IsFloat(itemsOp[0]) {
		//число
		stackOp = structures.Push(itemsOp[0], stackOp)
		itemsOp = itemsOp[1:len(itemsOp)]
	} else if utils.IsContainsFunc(itemsOp[0]){
		//обрабатываем оператор
		var a, b string
		a, b, stackOp = structures.PopTwo(stackOp)
		res, err = evalOperator(itemsOp[0], a, b)
		stackOp = structures.Push(fmt.Sprintf("%f", res), stackOp)
		itemsOp = itemsOp[1:len(itemsOp)]
	}
	if len(itemsOp) > 0 {
		res, err = processRpn(itemsOp, stackOp)
	}

	return res, err
}

/**
Функция вычисления
 */
func evalOperator(operator, a, b string) (float64, error) {

	bf, err := strconv.ParseFloat(b, 64)
	af, err := strconv.ParseFloat(a, 64)

	if err != nil {
		return 0, err
	}

	//получение функции вычисления по оператору
	var funcOperate func(x, y float64) float64
	funcOperate = structures.GetOperation(operator)

	if funcOperate != nil {
		return funcOperate(af, bf), nil
	} else {
		return 0, fmt.Errorf("Unexpected operator")
	}
}


