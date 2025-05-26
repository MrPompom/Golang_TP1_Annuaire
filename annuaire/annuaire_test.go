package annuaire

import "testing"

func TestAjouterContact(t *testing.T) {
	ann := NewAnnuaire()

	err := ann.AjouterContact("Alice", "0123456789")
	if err != nil {
		t.Errorf("Erreur lors de l'ajout: %v", err)
	}

	err = ann.AjouterContact("Alice", "0987654321")
	if err == nil {
		t.Error("Devrait retourner une erreur pour un contact existant")
	}
}

func TestRechercherContact(t *testing.T) {
	ann := NewAnnuaire()
	ann.AjouterContact("Bob", "0555666777")

	contact, existe := ann.RechercherContact("Bob")
	if !existe {
		t.Error("Contact devrait exister")
	}
	if contact.Telephone != "0555666777" {
		t.Errorf("Mauvais numéro: %s", contact.Telephone)
	}

	_, existe = ann.RechercherContact("Charlie")
	if existe {
		t.Error("Contact ne devrait pas exister")
	}
}
