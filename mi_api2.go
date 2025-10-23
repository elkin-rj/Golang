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

// DELETE /users/{id} - Eliminar usuario
func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"message":"Usuario eliminado"}`)
			return
		}
	}
	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
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

	// Ruta para eliminar usuario específico
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			// Extraer ID de la URL
			path := r.URL.Path[len("/users/"):]
			id := 0
			fmt.Sscanf(path, "%d", &id)
			deleteUser(w, r, id)
		}
	})

	// Iniciar servidor
	fmt.Println("🚀 Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
