package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var chaveSecreta = []byte("chave-secreta") // Substitua por uma chave secreta forte em um ambiente de produção.

type Mensagem struct {
	Texto string `json:"texto"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Olá, este é o meu webservice em Go!")
}

func mensagemHandler(w http.ResponseWriter, r *http.Request) {
	mensagem := Mensagem{Texto: "Esta é uma mensagem JSON."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mensagem)
}

func autenticarHandler(w http.ResponseWriter, r *http.Request) {
	// Aqui você autenticaria o usuário e geraria um token JWT válido

	

	usuario := "usuario"
	token := criarToken(usuario)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func criarToken(usuario string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	reivindicacoes := token.Claims.(jwt.MapClaims)
	reivindicacoes["usuario"] = usuario
	reivindicacoes["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expira em 1 hora

	tokenString, _ := token.SignedString(chaveSecreta)
	return tokenString
}

func middlewareAutenticacao(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extrairToken(r)
		if tokenString == "" {
			http.Error(w, "Token de autenticação ausente", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return chaveSecreta, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token de autenticação inválido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func extrairToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(bearerToken) > 7 && bearerToken[:6] == "Bearer" {
		return bearerToken[7:]
	}
	return ""
}

func main() {
	// Configurar handlers para diferentes rotas
	http.HandleFunc("/", handler)
	http.HandleFunc("/mensagem", mensagemHandler)
	http.HandleFunc("/autenticar", autenticarHandler)

	// Configurar middleware de autenticação para proteger rotas
	http.Handle("/restrito", middlewareAutenticacao(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Esta é uma área restrita.")
	})))

	// Iniciar o servidor na porta 8080
	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
