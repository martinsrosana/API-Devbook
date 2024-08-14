package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e retorna a hash dessa string
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara uma senha hash com uma senha em texto plano e verifica se são iguais
func VerificarSenha(senhaHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
