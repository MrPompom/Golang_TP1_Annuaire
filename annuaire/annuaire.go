package annuaire

import "fmt"

type Contact struct {
	Nom       string
	Telephone string
}

type Annuaire struct {
	contacts map[string]Contact
}

func NewAnnuaire() *Annuaire {
	ann := &Annuaire{
		contacts: make(map[string]Contact),
	}
	ann.contacts["Alice Dupont"] = Contact{Nom: "Alice Dupont", Telephone: "0123456789"}
	ann.contacts["Bob Martin"] = Contact{Nom: "Bob Martin", Telephone: "0987654321"}
	ann.contacts["Charlie Brown"] = Contact{Nom: "Charlie Brown", Telephone: "0555666777"}

	return ann
}

func (a *Annuaire) AjouterContact(nom, telephone string) error {
	if _, existe := a.contacts[nom]; existe {
		return fmt.Errorf("contact '%s' existe déjà", nom)
	}

	a.contacts[nom] = Contact{
		Nom:       nom,
		Telephone: telephone,
	}
	return nil
}

func (a *Annuaire) RechercherContact(nom string) (Contact, bool) {
	contact, existe := a.contacts[nom]
	return contact, existe
}

func (a *Annuaire) ListerContacts() []Contact {
	contacts := make([]Contact, 0)
	for _, contact := range a.contacts {
		contacts = append(contacts, contact)
	}
	return contacts
}

func (a *Annuaire) SupprimerContact(nom string) error {
	if _, existe := a.contacts[nom]; !existe {
		return fmt.Errorf("contact '%s' non trouvé", nom)
	}
	delete(a.contacts, nom)
	return nil
}

func (a *Annuaire) ModifierContact(nom, nouveauTel string) error {
	contact, existe := a.contacts[nom]
	if !existe {
		return fmt.Errorf("contact '%s' non trouvé", nom)
	}
	contact.Telephone = nouveauTel
	a.contacts[nom] = contact
	return nil
}
