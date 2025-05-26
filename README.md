# Golang_TP1_Annuaire


# Annuaire Go

## Description
Programme d'annuaire en Go pour gérer des contacts.

## Utilisation

### Ajouter un contact
```bash
go run main.go --action ajouter --nom "Charlie Puth" --tel "0811223344"
```

### Rechercher un contact
```bash
go run main.go --action rechercher --nom "Alice"
```

### Lister tous les contacts
```bash
go run main.go --action lister
```

### Supprimer un contact
```bash
go run main.go --action supprimer --nom "Charlie Puth"
```

### Modifier un contact
```bash
go run main.go --action modifier --nom "Alice" --tel "0999888777"
```

## Tests
```bash
go test ./annuaire
```

## Structure
- `main.go` : Point d'entrée
- `annuaire/annuaire.go` : Logique métier
- `annuaire/annuaire_test.go` : Tests unitaires