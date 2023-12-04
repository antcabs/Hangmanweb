package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

func hangman() {
    rand.Seed(time.Now().UnixNano())

    dico := []string{"AVION", "ARBRE", "CAMIONS", "CHEVEUX", "CAMPING", "CHAUSSURE", "FOUET", "CHAMPS", "HORLOGE", "CHIPS", "ORDINATEUR", "TABLE", "APERO", "TABLEAU", "WAGON", "CAILLOUX"}
    wordToGuess := dico[rand.Intn(len(dico))]
    guessedWord := make([]string, len(wordToGuess))

    for i := range guessedWord {
        guessedWord[i] = "_"
    }

    attemptsLeft := 8
    guessedLetters := make([]string, 0)

    for {
        if attemptsLeft == 0 || strings.Join(guessedWord, "") == wordToGuess {
            break
        }

        fmt.Printf("Vies restantes : %d\n", attemptsLeft)
        fmt.Printf("Lettres utilisées : %s\n", strings.Join(guessedLetters, " "))
        fmt.Printf("Mot caché : %s\n", strings.Join(guessedWord, " "))

        var guess string
        fmt.Print("Devine une lettre : ")
        fmt.Scanln(&guess)

        if len(guess) != 1 {
            fmt.Println("Une seul lettre est demandé.")
            continue
        }
        guess = strings.ToUpper(guess)

        if strings.Contains(wordToGuess, guess) {
            for i, letter := range wordToGuess {
                if string(letter) == guess {
                    guessedWord[i] = guess
                }
            }
        } else {
            attemptsLeft--
        }
        guessedLetters = append(guessedLetters, guess)
    }

    if strings.Join(guessedWord, "") == wordToGuess {
        fmt.Printf("gg ! T'as trouvé")
    }
}
