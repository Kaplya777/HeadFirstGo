// guess игра в которой нужно угадать число
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // Получаем текущую дату и время в формате целого числа
	rand.Seed(seconds)           //Функция генератора случайных счисел
	target := rand.Intn(100) + 1
	fmt.Println("I've choose a random number beetwen 1 and 100")
	fmt.Println("Can you guess it?")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin) //для чтения ввода с клавиатуры

	success := false
	for guesses := 0; guesses < 10; guesses++ { //в переменной guesses хранится кол-во попыток
		fmt.Println("You have", 10-guesses, "guesses left")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n') //прочитать данные введенные пользователем
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  //удаление символа новой строки
		guess, err := strconv.Atoi(input) //входная строка преобразовывается в числыо
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("Oops! You guess was high")
		} else if target > guess {
			fmt.Println("Oops! You guess was low")
		} else {
			success = true
			fmt.Println("WOW! You guessed it!")
			break //выход из цикла если пользователь угадал число
		}
	}
	if !success { // если игрок не угадал (если переменная succsess содержит false) то вывести сообщение
		fmt.Println("Sorry, you didnt guess number( It was:", target)
	}
}
