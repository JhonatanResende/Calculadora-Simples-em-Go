package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Criamos um leitor para pegar a linha inteira (com espaços)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🧮 Simple Calculator")
	fmt.Println("Type an operation (ex: 12 + 3) or 'exit' to quit")

	for {
		fmt.Print("> ")

		// Lê a linha inteira digitada
		input, _ := reader.ReadString('\n')

		// Remove espaços extras e quebra de linha
		input = strings.TrimSpace(input)

		// Sair do programa
		if input == "exit" {
			fmt.Println("Bye 👋")
			break
		}

		// Remove todos os espaços: "12 + 3" -> "12+3"
		input = strings.ReplaceAll(input, " ", "")

		result, err := calculate(input)

		// Se deu erro, mostramos bonito
		if err != nil {
			fmt.Println("❌ Error:", err)
			continue
		}

		fmt.Println("✅ Result:", result)
	}
}

func calculate(input string) (float64, error) {

	operators := []string{"+", "-", "*", "/"}

	var op string
	var parts []string

	// Procurar operador
	for _, operator := range operators {
		if strings.Contains(input, operator) {
			op = operator
			parts = strings.Split(input, operator)
			break
		}
	}

	// Validação
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid format. Use something like 2+2")
	}

	// Converter números
	num1, err1 := strconv.ParseFloat(parts[0], 64)
	num2, err2 := strconv.ParseFloat(parts[1], 64)

	if err1 != nil || err2 != nil {
		return 0, fmt.Errorf("invalid numbers")
	}

	// Evitar divisão por zero
	if op == "/" && num2 == 0 {
		return 0, fmt.Errorf("division by zero")
	}

	// Calcular
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		return num1 / num2, nil
	}

	return 0, fmt.Errorf("unknown operator")
}
