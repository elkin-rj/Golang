package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<html>
			<head><title>Botones en Go</title></head>
			<body style="font-family:sans-serif; text-align:center;">
				<h1>Panel de Control</h1>
				<form action="/accion" method="post">
					<button type="submit">Encender Motor</button>
				</form>
			</body>
		</html>`
		fmt.Fprint(w, html)
	})

	http.HandleFunc("/accion", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "✅ Motor encendido correctamente.")
	})

	fmt.Println("Servidor web en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
