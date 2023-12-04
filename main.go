package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Gère les requêtes avec la fonction `handler`
	http.HandleFunc("/", handler)

	// Démarre le serveur sur le port 8080
	port := 8080
	fmt.Printf("Serveur en cours d'exécution sur le port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}

// Fonction handler qui gère toutes les requêtes HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	// Récupère le chemin demandé dans la requête
	path := r.URL.Path[1:]

	// Si le chemin est vide, renvoie index.html par défaut
	if path == "" {
		path = "index.html"
	}

	// Lit le fichier correspondant au chemin
	content, err := readFile(path)
	if err != nil {
		// Si le fichier n'est pas trouvé, renvoie une erreur 404
		http.NotFound(w, r)
		return
	}

	// Détermine le type de contenu en fonction de l'extension du fichier
	contentType := getContentType(path)

	// Définit le type de contenu dans l'en-tête de la réponse
	w.Header().Set("Content-Type", contentType)

	// Écrit le contenu du fichier dans la réponse
	w.Write(content)
}

// Fonction pour lire le contenu d'un fichier
func readFile(filename string) ([]byte, error) {
	// Remplacez cette logique par la lecture réelle de vos fichiers depuis le système de fichiers
	// Actuellement, elle renvoie simplement un message indiquant le nom du fichier
	return []byte(fmt.Sprintf("Contenu du fichier : %s", filename)), nil
}

// Fonction pour obtenir le type de contenu en fonction de l'extension du fichier
func getContentType(filename string) string {
	// Vous pouvez étendre cette fonction pour gérer d'autres types de fichiers
	if filename[len(filename)-5:] == ".html" {
		return "text/html"
	} else if filename[len(filename)-4:] == ".css" {
		return "text/css"
	} else {
		return "text/plain"
	}
}
