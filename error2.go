package main

import (
	"errors"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return a / b, nil
}

func main() {
	a := app.New()
	w := a.NewWindow("Manejo de errores gráfico")

	boton := widget.NewButton("Dividir 10 / 0", func() {
		_, err := dividir(10, 0)
		if err != nil {
			widget.ShowPopUp(widget.NewLabel("❌ Error: "+err.Error()), w.Canvas())
		}
	})

	w.SetContent(boton)
	w.ShowAndRun()
}
