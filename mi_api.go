package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estructura simple
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Lista de usuarios en memoria
var users = []User{
	{ID: 1, Name: "Juan"},
	{ID: 2, Name: "Ana"},
}

// GET /users - Obtener todos los usuarios
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// POST /users - Crear usuario
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	user.ID = len(users) + 1
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Configurar rutas
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getUsers(w, r)
		} else if r.Method == "POST" {
			createUser(w, r)
		}
	})

	// Iniciar servidor
	fmt.Println("🚀 Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
