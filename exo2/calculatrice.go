package main
import (
	"fmt";
	"errors"
)

func main() {
	fmt.Printf("Effectuez une opération (exemple: 5 6 * => 30):")
	for {
		var a, b float64
		var op string
		fmt.Printf("\n>")
		fmt.Scan(&a, &b, &op)
		var result, err = operer(a,b,op)
		if (op == "quit") {
			break
		}
		if (err != nil) {
			fmt.Printf("Erreur : %s", err)
		} else {
			fmt.Printf("%.2f", result)
		}
	}
}

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Division par 0 impossible")
		}
		return a / b, nil
	default:
		return 0, errors.New("Opération inconnue")
	}
}