package main

import (
	"flag"
	"fmt"

	"github.com/MrPompom/Golang_TP1_Annuaire/annuaire"
)

func main() {
	var action string
	var nom string
	var tel string

	flag.StringVar(&action, "action", "", "Action: ajouter, rechercher, lister, supprimer, modifier")
	flag.StringVar(&nom, "nom", "", "Nom et prénom")
	flag.StringVar(&tel, "tel", "", "Numéro de téléphone")
	flag.Parse()

	ann := annuaire.NewAnnuaire()

	switch action {
	case "ajouter":
		if nom == "" || tel == "" {
			fmt.Println("Nom et téléphone requis")
			return
		}
		err := ann.AjouterContact(nom, tel)
		if err != nil {
			fmt.Printf("Erreur: %v\n", err)
		} else {
			fmt.Printf("Contact '%s' ajouté\n", nom)
		}

	case "rechercher":
		if nom == "" {
			fmt.Println("Nom requis")
			return
		}
		contact, existe := ann.RechercherContact(nom)
		if existe {
			fmt.Printf("%s - %s\n", contact.Nom, contact.Telephone)
		} else {
			fmt.Printf("Contact '%s' non trouvé\n", nom)
		}

	case "lister":
		contacts := ann.ListerContacts()
		for _, contact := range contacts {
			fmt.Printf("%s: %s\n", contact.Nom, contact.Telephone)
		}

	case "supprimer":
		if nom == "" {
			fmt.Println("Nom requis")
			return
		}
		err := ann.SupprimerContact(nom)
		if err != nil {
			fmt.Printf("Erreur: %v\n", err)
		} else {
			fmt.Printf("Contact '%s' supprimé\n", nom)
		}

	case "modifier":
		if nom == "" || tel == "" {
			fmt.Println("Nom et nouveau téléphone requis")
			return
		}
		err := ann.ModifierContact(nom, tel)
		if err != nil {
			fmt.Printf("Erreur: %v\n", err)
		} else {
			fmt.Printf("Contact '%s' modifié\n", nom)
		}

	default:
		fmt.Println("Action non reconnue")
		flag.Usage()
	}
}
