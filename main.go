package main

import (
	"api/API/src/config"
	"api/API/src/router"
	"fmt"
	"log"
	"net/http"
)

//cria uma chave de 64 bytes para ser usada como secret key
//func init() {
//	chave := make([]byte, 64)
//	if _, erro := rand.Read(chave); erro != nil {
//		log.Fatal(erro)
//
//	}
//
//	stringBase64 := base64.StdEncoding.EncodeToString(chave)
//	fmt.Println(stringBase64)
//}

func main() {
	config.Carregar()
	r := router.Gerar()

	//	fmt.Println(config.SecretKey)

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
