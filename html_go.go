package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8") // ✅ importante
		html := `
		<html>
			<head>
				<meta charset="UTF-8">
				<title>Generador de número</title>
			</head>
			<body>
				<h3>Hola Elkin</h3>
				<form action="/numero" method="post">
					<button type="submit">Generar número</button>
				</form>
			</body>
		</html>`
		fmt.Fprint(w, html)
	})

	http.HandleFunc("/numero", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    r := rand.New(rand.NewSource(time.Now().UnixNano())) // generador local
    num := r.Intn(100)
    fmt.Fprintf(w, "<h2>Número aleatorio: %d</h2>", num)
    fmt.Fprint(w, `<a href="/">Volver</a>`)
})

	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
