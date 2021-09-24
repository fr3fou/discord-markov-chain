package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/fr3fou/polo/polo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type DM struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Content string `json:"content"`
}

func main() {
	f, err := os.Open("./dump.json")
	if err != nil {
		panic(err)
	}

	log.Println("parsing dump.json")
	var d DM
	if err := json.NewDecoder(f).Decode(&d); err != nil {
		panic(err)
	}
	f.Close()
	messages := []string{}
	for _, message := range d.Messages {
		messages = append(messages, message.Content)
	}

	order := 1
	chain := polo.NewFromText(order, messages)
	log.Println("finished building chain")

	web(chain)
}

func web(c polo.Chain) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		in := r.URL.Query().Get("in")
		if in != "" {
			w.Write([]byte(c.NextUntilEnd(in)))
			w.WriteHeader(200)
			return
		} else {
			w.Write([]byte(c.NextUntilEnd(c.RandomState())))
			w.WriteHeader(200)
			return
		}
	})
	log.Println("listening on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
