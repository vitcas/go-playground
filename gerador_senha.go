package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/atotto/clipboard"
)

const (
	passwordLength = 16
	charset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!?@#$%&*-_"
)

func main() {
	for {
		password := generatePassword(passwordLength)
		fmt.Printf("Senha gerada: %s\n", password)

		var confirm string
		fmt.Print("Você quer usar essa senha? (s/n): ")
		fmt.Scanln(&confirm)

		if confirm == "s" {
			err := clipboard.WriteAll(password)
			if err != nil {
				log.Fatalf("Falha ao copiar para o clipboard: %v", err)
			}
			fmt.Println("Senha confirmada e copiada para o clipboard!")
			break
		} else if confirm == "n" {
			fmt.Println("Gerando nova senha...")
		} else {
			fmt.Println("Opção inválida. Por favor, insira 's' para sim ou 'n' para não.")
		}
	}
}

func generatePassword(length int) string {
	password := make([]byte, length)
	for i := range password {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[randomIndex.Int64()]
	}
	return string(password)
}
