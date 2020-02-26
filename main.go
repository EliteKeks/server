package main

import (
	"fmt" // пакет для форматированного ввода вывода
	//"encoding/json"
	// пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола
	//"text/template"
	//"strings"  // пакет для работы с  UTF-8 строками
)

// type User struct {
// 	FName string
// 	LName string
// 	Age string
// }

func main() {
	http.HandleFunc("/", HomeRouterHandler) // установим роутер
	http.ListenAndServe(":9000", nil)       // задаем слушать порт
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Вы успешно зарегестрированы!")
}
