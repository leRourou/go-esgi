package main
import "fmt"

type Personne struct {
	Prenom string
	Nom string
	Age int
	Email string
}

type Adresse struct {
	Rue string
	Ville string
	CodePostal string
}

type Employe struct {
	Personne
	Adresse
	Poste string
	Salaire float64
}

type Etudiant struct {
	Personne
	Promo string
	Moyenne float64
}

func (p Personne) NomComplet() string {
	return fmt.Sprintf("%s %s", p.Prenom, p.Nom)
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("Nom : %s\nAge : %d\nEmail : %s", p.NomComplet(), p.Age, p.Email)
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf("%s\nPoste : %s\nSalaire : %.2f€\nAdresse : %s, %s %s\n", e.Presentation(), e.Poste, e.Salaire, e.Rue, e.CodePostal, e.Ville)
}

func (e Etudiant) MentionObtenue() string {
	switch {
		case e.Moyenne >= 16: 
			return "TB"
		case e.Moyenne >= 14:
			return "B"
		case e.Moyenne >= 12:
			return "AB"
	}
	return "P"
}

func (e Etudiant) FicheEtudiant() string {
	return fmt.Sprintf("%s\nPromo : %s\nMoyenne : %.2f\nMention : %s\n", e.Presentation(), e.Promo, e.Moyenne, e.MentionObtenue())
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire + pct
}

func main() {
	employe1 := Employe {
		Personne: Personne {
			Prenom:	"Axel",
			Nom:	"ROUQUETTE",
			Age:	23,
			Email:	"axel.rouquette@outlook.fr",
		},
		Adresse: Adresse {
			Rue:		"2 Avenue de la Libération",
			Ville:		"Meyzieu",
			CodePostal:	"69330",
		},
		Poste:		"Développeur",
		Salaire:	1800.00,
	}

	employe2 := Employe {
		Personne: Personne {
			Prenom:	"Maxence",
			Nom:	"DUBOIS",
			Age:	23,
			Email:	"maxdub@gmail.com",
		},
		Adresse: Adresse {
			Rue:		"45 rue des Lilas",
			Ville:		"Montagny-Les-Lanches",
			CodePostal:	"74350",
		},
		Poste:		"Développeur",
		Salaire:	2300.00,
	}
	
	fmt.Println(employe1.FicheEmploye())
	employe1.AugmenterSalaire(500.00)
	fmt.Println(employe1.FicheEmploye())
	fmt.Println(employe2.FicheEmploye())

	etudiant1 := Etudiant {
		Personne: Personne {
			Prenom:	"Jean",
			Nom:	"DUJARDIN",
			Age:	22,
			Email:	"jean.dujardin@gmail.com",
		},
		Promo:		"2022",
		Moyenne:	19.5,
	}

	etudiant2 := Etudiant {
		Personne: Personne {
			Prenom:	"Marie",
			Nom:	"CURIE",
			Age:	21,
			Email:	"marie.curie@gmail.com",
		},
		Promo:		"2022",
		Moyenne:	14.0,
	}

	fmt.Println(etudiant1.FicheEtudiant())
	fmt.Println(etudiant2.FicheEtudiant())
}