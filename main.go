package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Digite o tamanho da senha que deseja:")
	var tamanho int
	fmt.Scanln(&tamanho)

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*()_+[]{}"

	senha := gerarSenha(tamanho, charset)

	fmt.Printf("Sua senha gerada é: %s\n", senha)
}

// Função para gerar senha
func gerarSenha(tamanho int, charset string) string {
	rand.Seed(time.Now().UnixNano())

	senha := make([]byte, tamanho)

	for i := range senha {
		senha[i] = charset[rand.Intn(len(charset))]
	}

	return string(senha)
}
