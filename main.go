package main

import (
	"bufio"
	"calc-func/calculator"
	"calc-func/utils"
	"fmt"
	"os"
	"strings"
)


func main() {

	for true {
		start()
		str, err := readInputStr()
		if err != nil {
			fmt.Println(err)
			continue
		} else if strings.Contains(str, "exit") {
			fmt.Println("exit...")
			os.Exit(1)
		}
		fmt.Println("'"+str+"'")

		rpnArr := utils.RpnStrToArr(str)
		fmt.Println(rpnArr)

		fmt.Println("calculate...")
		result, err := calculator.Calc(rpnArr)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}
}

func start() {
	fmt.Println("Evaluating an expression in reverse polish notation")
	fmt.Println("example:\n 10 2 5 + * 14 1 2 1 2 + * + / - \n 1 2 + 4 * 3 + \n 2.2 10 * 15 - \n 5 - 5 6 * +")
	fmt.Println("input expression and press Enter or input 'exit'")
}

/**
Read input string
*/
func readInputStr() (string, error) {
	in := bufio.NewReader(os.Stdin)
	scanStr, err := in.ReadString('\n')
	scanStr = strings.TrimSuffix(scanStr, "\n")
	fmt.Println("replace ',' to '.'...")
	scanStr = strings.Replace(scanStr, ",", ".", -1)
	if err != nil {
		return "", err
	}

	return scanStr, nil
}
