package main

import (
	"fmt"
	"math/rand"
)

func roulettGame(namber, stavka int) (int, bool) {
	var valueI int

	for i := 0; i < 38; i++ {

		if i == 37 {
			i = 0
		}

		//fmt.Println(i)

		ball := rand.Intn(37)
		if ball == 0 {
			valueI = i
			break
		}
	}

	if valueI == namber {
		return stavka*36 + stavka, true
	} else {
		//fmt.Println("valueI: ", valueI)
		return 0, false
	}
}

// i := 0
// valueI := 0
// for {
// 	if i == 37 {
// 		i = 0
// 	}
// 	ball := rand.Intn(37)
// 	if ball == 0 {
// 		valueI = i
// 		break
// 	}
// 	fmt.Println(i)
// 	i++
// }

// if valueI == namber {
// 	return stavka * 36 + stavka, true
// } else {
// 	fmt.Println("valueI: ", valueI)
// 	return 0, false
// }
// }

func main() {

	budjet := 500000
	games := 100
	stavka := 1000
	maxStavka := 20000

	for numberGame := 0; numberGame < games; numberGame++ {
		budjet = budjet - stavka
		number := rand.Intn(37)
		//number := 12
		fmt.Println(stavka)
		result, ok := roulettGame(number, stavka)

		switch {
		case ok:
			budjet = budjet + result
			stavka = 1000
		default:
			stavka = stavka * 2
			if stavka >= maxStavka {
				stavka = 20000
			}
		}

		// if budjet > 500000 {
		// 	fmt.Println(budjet, numberGame)
		// 	break
		// }

		fmt.Println(budjet)
		if budjet <= 0 {
			fmt.Println(numberGame)
			break
		}
	}

	fmt.Println(budjet)

}
