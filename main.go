package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var words = []string{"golang", "javascript", "python", "java", "ruby", "html", "css"}
var currentWord string
var currentWordState []string
var guesses []string
var maxGuesses = 6

func main() {
	// Générer un mot aléatoire
	rand.Seed(time.Now().Unix())
	currentWord = words[rand.Intn(len(words))]
	currentWordState = make([]string, len(currentWord))
		for i := range currentWordState {
		currentWordState[i] = "_"
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/guess", guessHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Vérifier si le joueur a gagné ou perdu
	if !strings.Contains(strings.Join(currentWordState, ""), "_") {
		fmt.Fprint(w, "<h1>Vous avez gagné!</h1>")
		return
	} else if len(guesses) >= maxGuesses {
		fmt.Fprint(w, "<h1>Vous avez perdu!</h1>")
		return
	}

	// Afficher l'état actuel du mot
	fmt.Fprint(w, "<h1>Mot: ")
		for _, letter := range currentWordState {
		fmt.Fprint(w, letter+" ")
		}
	fmt.Fprintln(w, "</h1>")

	// Afficher les lettres déjà devinées
	fmt.Fprintln(w, "<h2>Lettres devinées: "+strings.Join(guesses, ", ")+"</h2>")

	// Afficher le formulaire de devinette
	fmt.Fprint(w, `
	<form action="/guess" method="POST">
	<label for="guess">Devinez une lettre:</label>
	<input type="text" name="guess" maxlength="1">
	<input type="submit" value="Devinez">
	</form>
	`)
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer la lettre devinée
	guess := strings.ToLower(r.FormValue("guess"))

	// Vérifier si la lettre est déjà devinée
	for _, letter := range guesses {
		if letter == guess {
		fmt.Fprint(w, "<h1>Cette lettre a déjà été devinée. Essayez encore.</h1>")
		return
		}
	}

	// Vérifier si la lettre est présente dans le mot
	letterFound := false
	for i, letter := range currentWord {
		if string(letter) == guess {
			currentWordState[i] = guess
			letterFound = true
		}
	}

	if !letterFound && guess !=""{
		guesses = append(guesses,guess)
	}

	// Ajouter la lettre devinée à la liste des devinettes
	guesses = append(guesses, guess)

	// Rediriger vers la page d'accueil
	
	http.Redirect(w, r, "/", http.StatusFound)
}

//pour run faut aller sur http://localhost:8080