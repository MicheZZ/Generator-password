package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	// Наборы символов
	app := []rune("0123456789")
	bad := []rune("abcdefghijklmnopqrstuvwxyz")
	cap := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	den := []rune("!@#№$%^&?*_=+")

	guf := [][]rune{app, bad, cap, den}
	fuck := [][]rune{app, bad, cap}

	// Ввод длины пароля
	var length int
	for {
		fmt.Print("Введи длину пароля: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		n, err := strconv.Atoi(input)
		if err == nil && n > 0 {
			length = n
			break
		}
		fmt.Println("Ошибка! Введите целое число")
	}

	// Нужно ли использовать спец символы
	var useSpecial bool
	fmt.Println("Надо использовать спец символы? Такие как '!#$@%&*?^_+=-'")
	for {
		fmt.Print("> ")
		answer, _ := reader.ReadString('\n')
		answer = strings.ToLower(strings.TrimSpace(answer))
		if answer == "да" {
			useSpecial = true
			break
		} else if answer == "нет" {
			useSpecial = false
			break
		} else {
			fmt.Println("Напиши да или нет")
		}
	}

	// Генерация пароля
	password := make([]rune, 0, length)
	charSets := guf
	if !useSpecial {
		charSets = fuck
	}

	for i := 0; i < length; i++ {
		set := charSets[rand.Intn(len(charSets))]
		char := set[rand.Intn(len(set))]
		password = append(password, char)
	}

	// Вывод пароля
	fmt.Println("Сгенерированный пароль:")
	fmt.Println(string(password))

	// Пауза перед выходом
	fmt.Println("\nНажмите ENTER для выхода.")
	reader.ReadString('\n')
}
