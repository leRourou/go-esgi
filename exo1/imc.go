package main
import (
  "fmt"
)

func main() {
	var poids, taille = 70.5, 1.75
	const IMCMaigreur, IMCNormal, IMCSurpoids = 18.5, 25.0, 30.0
	const nom = "Axel ROUQUETTE"

	var imc = poids / (taille * taille)

	fmt.Printf("Bonjour %s!\n", nom)
	fmt.Printf("Votre IMC (à 2 décimales près) est : %.2f\n", imc)

	var categorie = "surpoids"
	if (imc < IMCMaigreur) {
		categorie = "maigreur"
	} else if (imc < IMCNormal) {
		categorie = "normal"
	} else if (imc > 30) {
		categorie = "obésité"
	}

	fmt.Printf("Votre catégorie de poids est : %s\n", categorie)

}