package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var e_Reg = regexp.MustCompile(`[0-9]+[\+\-\*\:][0-9]+[=][?]`)
var op_Reg = regexp.MustCompile(`[0-9]+[\+\-\*\:][0-9]+`)
var operator = []string{"+", "-", "*", ":"}

func main() {

	filenameout := "./out.txt"

	_ = os.Remove(filenameout)
	file, err := os.OpenFile(filenameout, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	filename_ent := "./ent.txt"

	file1, err := os.OpenFile(filename_ent, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file1.Close()

	outputFile, err := os.OpenFile(string(filenameout), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer outputFile.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file1)
	writer := bufio.NewWriter(outputFile)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		equasion := e_Reg.FindString(string(line))
		if equasion == "" {
			continue
		}

		operation := op_Reg.FindString(equasion)
		if operation == "" {
			continue
		}

		numbers := []string{}
		var currentOperator string

		for _, op := range operator {
			numbers = strings.Split(operation, op)
			if len(numbers) > 1 {
				currentOperator = op
				break
			}
		}

		if currentOperator == "" || len(numbers) < 2 {
			continue
		}

		var answer float32

		first, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		second, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		switch currentOperator {
		case "+":
			answer = float32(first + second)
		case "-":
			answer = float32(first - second)
		case "*":
			answer = float32(first * second)
		case ":":
			answer = float32(first / second)
		default:
			fmt.Println("Данный оператор не поддерживается:\n", equasion)
			return
		}

		outputString := fmt.Sprintf("%d%v%d=%0.f", first, currentOperator, second, answer)

		_, err = writer.WriteString(outputString + "\n")

		if err != nil {
			fmt.Println(err)
			break
		}
	}

	writer.Flush()
	fmt.Println("Результат выведен в файл:", string(filenameout))
}
