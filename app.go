// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fernandojr999/go-learning/delivery"
	"github.com/fernandojr999/go-learning/repository"
	"github.com/fernandojr999/go-learning/usecase"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "my_bi"
)

var db *sql.DB

func main() {
	// Configurar a conexão com o banco de dados PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Criar tabela de usuários se não existir
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
						id SERIAL PRIMARY KEY,
						username VARCHAR(50) UNIQUE NOT NULL,
						password VARCHAR(100) NOT NULL
					)`)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	r := mux.NewRouter()

	// Handlers
	userHandler := delivery.NewUserHandler(userUsecase)
	r.HandleFunc("/api/user/register", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/user/authenticate", userHandler.AuthenticateUser).Methods("POST")
	r.HandleFunc("/api/user/all", userHandler.GetAllUsers).Methods("GET")

	// Iniciar o servidor
	serverAddr := "localhost:8081"
	srv := &http.Server{
		Handler:      r,
		Addr:         serverAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Servidor iniciado em http://%s\n", serverAddr)
	log.Fatal(srv.ListenAndServe())
}
