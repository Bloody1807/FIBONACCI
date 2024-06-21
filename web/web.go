package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"main.go/logic"
)

func render(w http.ResponseWriter, r *http.Request, filePath string) {
	file, err := os.ReadFile(filePath)

	if err != nil {
		return
	}

	fmt.Fprintf(w, "%s", file)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "./resources/index.html")
}

func resultHandler(w http.ResponseWriter, r *http.Request) {

	render(w, r, "./resources/result.html")

	fs := logic.FibonacciService{}

	number, err := strconv.Atoi(r.FormValue("numberValue"))

	if err != nil || number < 0 {

		fmt.Fprintf(w, "<h2>Некорректное число!</h2>")

	} else {

		fmt.Fprintf(w, "<p><a>Введённое число: %d\n</a></p>", number)

		if fs.IsFibonacci(number) {
			prev, next := fs.GetAdjacentFibonacci(number)
			fmt.Fprintf(w, "<p><a>Предыдущее число Фибоначчи: %d</a></p>", prev)
			fmt.Fprintf(w, "<p><a>Следующее число Фибоначчи: %d</a></p>", next)
		} else {
			closest := fs.GetNearestFibonacci(number)
			fmt.Fprintf(w, "<p><a>Ближайшее число Фибоначчи: %d</a></p>", closest)
		}

	}

	fmt.Fprintf(w, "<p><a href=\"/\">Ввести новое число</a></p>")
}

func Start() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/result", resultHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
