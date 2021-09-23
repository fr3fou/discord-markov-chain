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

	var d DM
	if err := json.NewDecoder(f).Decode(&d); err != nil {
		panic(err)
	}
	messages := []string{}
	for _, message := range d.Messages {
		messages = append(messages, message.Content)
	}

	order := 1
	chain := polo.NewFromText(order, messages)

	/*
		fmt.Println("Press enter for the next generated message")
		fmt.Println("	You can also enter a starting word")
		fmt.Println("	Type 'quit' to quit")
		rl, err := readline.New("> ")
		if err != nil {
			panic(err)
		}
		defer rl.Close()
		in := ""
		for {
			in, err = rl.ReadlineWithDefault(in)
			if err != nil {
				fmt.Println(err)
				return
			}
			if in == "quit" {
				return
			}
			fmt.Print("< ")
			if in == "" {
				fmt.Println(chain.NextUntilEnd(chain.RandomState()))
			} else {
				fmt.Println(chain.NextUntilEnd(in))
			}
		}
	*/
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
