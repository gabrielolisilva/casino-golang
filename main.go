package main

import (
	"fmt"
	"math/rand"
)

func definePlayerName() {
	name := ""

	fmt.Println("Welcome to the Casino's Game !")
	fmt.Printf("Please enter your name: ")
	_, err := fmt.Scan(&name)

	/* Null value in GO is nil */
	if err != nil {
		return
	}

	fmt.Printf("Welcome to the game %s\n", name)
}

func generateSymbolsArray() []string {
	symbols := map[string]uint {
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}

	var symbolsArray []string

	for symbol, count := range symbols{
		for i := uint(0); i < count; i++ {
			symbolsArray = append(symbolsArray, symbol)
		}
	}

	return symbolsArray
}

func definePlayerBet(balance uint) uint {
	typedBet := uint(0)

	/* Logic below same as while true */
	for true {
		fmt.Printf("Please enter your bet: ")
		_, err := fmt.Scan(&typedBet)
		if err != nil {
			continue
		}

		if typedBet == 0 {
			return 0
		} else if typedBet > balance {
			fmt.Println("You can't bet more than your balance")
			continue
		}

		break
	}

	return typedBet
}

func generateRandomValueByArray(array []string) uint {
	max := len(array) - 1
	min := 0

	randomValue := rand.Intn(max - min + 1) + min
	return uint(randomValue)
}

func containsValueInArray(array []uint, value uint) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

/** 
* Each array inside is row and each value inside is a column
* [ [A B C] [A B C] [A B C] ]
*/ 
func getSpinData(symbolsArray []string, rows uint, columns uint) [][]string {
	resultArray := [][]string{}

	for i := uint(0); i < rows; i++ {
		resultArray = append(resultArray, []string{})
	}

	for r := uint(0); r < rows; r++ {
		usedRowSymbolsIndex := []uint{}
		for c := uint(0); c < columns; c++ {
			var randomSymbolsIndexValue uint;
			for {
				randomSymbolsIndexValue = generateRandomValueByArray(symbolsArray);
				if !containsValueInArray(usedRowSymbolsIndex, randomSymbolsIndexValue) {
					break
				}
			}

			resultArray[r] = append(resultArray[r], symbolsArray[randomSymbolsIndexValue])
			usedRowSymbolsIndex = append(usedRowSymbolsIndex, randomSymbolsIndexValue)
		}
	}

	return resultArray
}

func checkWin(spinResult [][]string, balance uint) uint {
	multiplies := map[string]uint {
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	for _, row := range spinResult {
		if row[0] == row[1] && row[1] == row[2] {
			multiplierValue := multiplies[row[0]]
			balance *= multiplierValue
		}
	}

	return balance
}

func main() {
	definePlayerName()
	symboldArray := generateSymbolsArray()

	/* uint is a number greater than or equal 0 */
	starterBetValue := uint(200)
	for starterBetValue > 0 {
		playerBet := definePlayerBet(starterBetValue)
		if playerBet == 0 {
			fmt.Println("You left the game")
			break
		}

		starterBetValue -= playerBet
		fmt.Printf("Your balance new is %d, starting game\n", starterBetValue)

		spinResult := getSpinData(symboldArray, 3, 3)
		fmt.Printf("Spin result is %v\n", spinResult)

		starterBetValue = checkWin(spinResult, starterBetValue)
		fmt.Printf("Your balance final is %d\n\n", starterBetValue)
	}
}