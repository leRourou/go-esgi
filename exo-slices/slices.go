package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	flotte := []string{
		"LAPTOP-01,laptop,2019",
		"TABLET-05,tablet,2021",
		"SRV-PROD,server,2020",
	}

	var recents []string
	var obsoletes []string

	for _, appareil := range flotte {
		parts := strings.Split(appareil, ",")
		annee, _ := strconv.Atoi(parts[2])
		if annee > 2019 {
			recents = append(recents, appareil)
		} else {
			obsoletes = append(obsoletes, appareil)
		}
	}

	fmt.Printf("Récents   : %d\n", len(recents))
	fmt.Printf("Obsolètes : %d\n\n", len(obsoletes))

	fmt.Println("=== Rapport récents ===")
	for _, appareil := range recents {
		parts := strings.Split(appareil, ",")
		nom := parts[0]
		typeApp := parts[1]

		fmt.Printf("%-10s → ", nom)
		switch typeApp {
		case "server":
			fmt.Println("Critique — SLA 24h")
			fmt.Printf("%-10s   ", "")
			fallthrough
		case "laptop":
			fmt.Println("Standard — ticket J+1")
		case "tablet", "phone":
			fmt.Println("Faible priorité")
		}
	}
}