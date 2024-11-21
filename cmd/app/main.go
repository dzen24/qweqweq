package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Структура для обработки данных в POST-запросе
type Message struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Маршрут для обработки GET-запросов
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		// Ответ на GET-запрос
		response := map[string]string{"message": "Hello, GET request!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Маршрут для обработки POST-запросов
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		// Чтение данных из тела запроса
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
			return
		}

		// Разбор JSON в структуру
		var msg Message
		err = json.Unmarshal(body, &msg)
		if err != nil {
			http.Error(w, "Ошибка парсинга JSON", http.StatusBadRequest)
			return
		}

		// Формирование ответа
		response := map[string]string{
			"status": "success",
			"name":   msg.Name,
			"email":  msg.Email,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Запуск сервера на порту 8080
	port := ":8080"
	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
