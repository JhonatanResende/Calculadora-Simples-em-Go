# 🧮 Calculadora Simples (Go)

Uma calculadora simples feita em **Golang** que lê operações matemáticas digitadas pelo usuário no terminal.

---

## 🚀 Funcionalidades

* ➕ Soma (`+`)
* ➖ Subtração (`-`)
* ✖️ Multiplicação (`*`)
* ➗ Divisão (`/`)
* ❌ Tratamento de erros
* 🔁 Loop contínuo até o usuário sair

---

## 💻 Como funciona

O usuário digita uma operação no terminal, por exemplo:

```
12 + 3
```

O programa:

1. Lê o texto digitado
2. Remove espaços
3. Identifica o operador matemático
4. Separa os números
5. Converte para número (float)
6. Realiza o cálculo
7. Mostra o resultado

---

## 🧪 Exemplo de uso

```
🧮 Simple Calculator
Type an operation (ex: 12 + 3) or 'exit' to quit

> 10 + 5
✅ Result: 15

> 7 * 3
✅ Result: 21

> 10 / 0
❌ Error: division by zero

> exit
Bye 👋
```

---

## ⚠️ Validações implementadas

* Formato inválido (ex: `2++2`)
* Valores não numéricos (ex: `a+b`)
* Divisão por zero
* Operadores desconhecidos

---

## 🧠 Estrutura do código

### `main()`

Responsável por:

* Interação com o usuário
* Leitura do input
* Loop contínuo

### `calculate()`

Responsável por:

* Identificar operador
* Separar números
* Converter valores
* Executar cálculo
* Retornar resultado ou erro

---

## 🛠️ Tecnologias utilizadas

* Go (Golang)
* Pacotes padrão:

  * `fmt`
  * `bufio`
  * `os`
  * `strconv`
  * `strings`

---

## 📌 Objetivo

Projeto criado para aprendizado, com foco em:

* Manipulação de strings
* Tratamento de erros
* Estrutura básica em Go
* Entrada e saída no terminal

---

## 👨‍💻 Autor

Jhonatan Resende 🚀

---

## 📄 Licença

Este projeto está sob a licença MIT.
