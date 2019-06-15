package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var liveHero = 100
var liveDrag = 100
var nameHero string
var nameDrag string
var flag = true
var actionDrag string

func main() {
	fmt.Println("                         Попробуй победить Дракона!")
	menu()
}
func menu() {
	for {
		fmt.Println("Главное меню")
		fmt.Println("-----------------")
		fmt.Println("1.Начать игру")
		fmt.Println("2.Выход")
		var menu string
		Menu := bufio.NewScanner(os.Stdin)
		Menu.Scan()
		menu = Menu.Text()
		if menu == "1" {
			initialization()
			break
		} else if menu == "2" {
			fmt.Println("До скорой встречи")
			break
		} else {
			fmt.Println()
			fmt.Println()
			fmt.Println("Неправильно ввели значение, пожалйста будьте по внимательнее и повторите попытку!")
			fmt.Println()
		}

	}
}
func initialization() {
	fmt.Print("Введите имя героя:    ")
	fmt.Fscan(os.Stdin, &nameHero)
	for {
		fmt.Println("Вы будете задавать имя Дракону?")
		fmt.Println("1.Да")
		fmt.Println("2.Нет(Остается по умолчанию)")
		nameD := bufio.NewScanner(os.Stdin)
		nameD.Scan()
		nameQuest := nameD.Text()
		if nameQuest == "1" {
			fmt.Print("Введите имя дракона:   ")
			fmt.Fscan(os.Stdin, &nameDrag)
			fmt.Println("Игра началось!")
			start()
			break
		} else if nameQuest == "2" {
			nameDrag = "Дракон"
			fmt.Println("Игра началось!")
			start()
			break
		} else {
			fmt.Println()
			fmt.Println()
			fmt.Println("Неправильно ввели значение, пожалйста будьте по внимательнее и повторите попытку!")
			fmt.Println()
		}

	}
}
func start() {

	for {
		if flag {
			fmt.Println(nameHero, "   ", liveHero, "       Vs      ", liveDrag, "   ", nameDrag)
			fmt.Println()
			fmt.Println()
			fmt.Println("Ваш ход")
			fmt.Println("1.Нанести удар")
			fmt.Println("2.Подлечиться")
			choice := bufio.NewScanner(os.Stdin)
			choice.Scan()
			choiceOfAction := choice.Text()
			if choiceOfAction == "1" {
				liveDrag -= rand.Intn(10)
				if liveDrag <= 0 {
					fmt.Println("ПОЗДРАВЛЯЕМ ВЫ ВЫИГРАЛИ")
					break
				}
				flag = false
			} else if choiceOfAction == "2" {
				if liveHero < 100 {
					liveHero += rand.Intn(10)
					if liveHero > 100 {
						liveHero = 100
					}

				}
				flag = false

			} else {
				fmt.Println()
				fmt.Println()
				fmt.Println("Неправильно ввели значение, пожалйста будьте по внимательнее и повторите попытку!")
				fmt.Println()
			}
		} else {

			fmt.Println()
			fmt.Println()
			fmt.Println("ХОД ДРАКОНА")

			if rand.Intn(2) == 1 {
				liveHero -= rand.Intn(10)
				actionDrag = "Дракон нанес удар"
				fmt.Println(actionDrag)
				if liveHero <= 0 {
					fmt.Println("УВЫ ВЫ ПОТЕРПЕЛИ ПОРАЖЕНИЯ!")
					break
				}

			} else {
				if liveDrag < 100 {
					liveDrag += rand.Intn(10)
					if liveDrag > 100 {
						liveDrag = 100
					}

				}
				actionDrag = "Дракон подлечился"
				fmt.Println(actionDrag)
			}

			fmt.Println()
			fmt.Println()

			flag = true
		}

	}
}
