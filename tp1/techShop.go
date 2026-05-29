package main
import "fmt"

type Produit struct {
	ID int
	Nom string
	Marque string
	Prix float64
	Stock int
	Categorie string
}

type Catalogue struct {
	Produits []Produit
}

func (p Produit) Rapport() string {
	return fmt.Sprintf("ID: %d\nNom: %s\nMarque: %s\nPrix: %.2f€\nStock: %d\nCatégorie: %s\n", p.ID, p.Nom, p.Marque, p.Prix, p.Stock, p.Categorie)
}

func (e Catalogue) TrouverParID(id int) (Produit, error) {
	for _, p := range e.Produits {
		if p.ID == id {
			return p, nil
		}
	}
	return Produit{}, fmt.Errorf("Produit avec ID %d non trouvé", id)
}

func (e *Catalogue) AjouterProduit(p Produit) (error) {
	if _, err := e.TrouverParID(p.ID); err == nil {
		return fmt.Errorf("Produit avec ID %d existe déjà", p.ID)
	}
	e.Produits = append(e.Produits, p)
	return nil
}

func (e Catalogue) TrouverParCategorie(cat string) []Produit {
	var result []Produit
	for _, p := range e.Produits {
		if p.Categorie == cat {
			result = append(result, p)
		}
	}
	return result
}

func (e *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	count := 0
	for i, p := range e.Produits {
		if p.Categorie == categorie {
			e.Produits[i].Prix = p.Prix * (1 - pct/100)
			count++
		}
	}
	return count
}

func (e *Catalogue) Vendre(id int, quantite int) error {
	for i, p := range e.Produits {
		if p.ID == id {
			if p.Stock < quantite {
				return fmt.Errorf("Stock insuffisant pour le produit %d", id)
			}
			e.Produits[i].Stock -= quantite
		}
	}
	return nil
}

func (e *Catalogue) Rapport() {
	categories := make(map[string]int)
	for _, p := range e.Produits {
		categories[p.Categorie]++
	}

	fmt.Println("--- Rapport de stock ---")
	for cat, count := range categories {
		fmt.Printf("%s : %d produits\n", cat, count)
	}
}

func main() {

	var catalogue Catalogue

	var iPhone = Produit{ID: 1, Nom: "iPhone 13", Marque: "Apple", Prix: 999.99, Stock: 50, Categorie: "smartphone"}
	var macbook = Produit{ID: 2, Nom: "MacBook Pro", Marque: "Apple", Prix: 1999.99, Stock: 20, Categorie: "ordinateur"}
	var galaxy = Produit{ID: 3, Nom: "Galaxy S21", Marque: "Samsung", Prix: 799.99, Stock: 30, Categorie: "smartphone"}
	var surface = Produit{ID: 4, Nom: "Surface Laptop", Marque: "Microsoft", Prix: 1499.99, Stock: 15, Categorie: "ordinateur"}
	var iPad = Produit{ID: 5, Nom: "iPad Pro", Marque: "Apple", Prix: 1099.99, Stock: 25, Categorie: "tablette"}

	catalogue.AjouterProduit(iPhone)
	catalogue.AjouterProduit(macbook)
	catalogue.AjouterProduit(galaxy)
	catalogue.AjouterProduit(surface)
	catalogue.AjouterProduit(iPad)

	fmt.Println("Menu : [1] Ajouter [2] Chercher [3] Soldes [4] Vendre [5] Rapport [0] Quitter")

	for {
		var choix int
		fmt.Printf("\nEntrez votre choix: ")
		fmt.Scan(&choix)
		switch choix {
			case 1:
				var p Produit
				fmt.Println("Entrez les détails du produit (ID, Nom, Marque, Prix, Stock, Categorie):")
				fmt.Scan(&p.ID, &p.Nom, &p.Marque, &p.Prix, &p.Stock, &p.Categorie)
				catalogue.AjouterProduit(p)
			case 2:
				var id int
				fmt.Println("ID du produit à chercher:")
				fmt.Scan(&id)
				prod, err := catalogue.TrouverParID(id)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Produit trouvé: %+v\n", prod.Rapport())
				}
			case 3:
				var cat string
				var pct float64
				fmt.Println("Catégorie à mettre en réduction et pourcentage de réduction:")
				fmt.Scan(&cat, &pct)
				count := catalogue.AppliquerReduction(cat, pct)
				fmt.Printf("%d produits soldés dans la catégorie %s\n", count, cat)
			case 4:
				var id, quantite int
				fmt.Println("ID du produit à vendre et quantité:")
				fmt.Scan(&id, &quantite)
				err := catalogue.Vendre(id, quantite)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Vente réussie pour le produit %d\n", id)
				}
			case 5:
				catalogue.Rapport()
			case 0:
				return
			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}