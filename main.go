package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		panic("ENVIRONMENT env variable not set!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("I am a %s environment", env))
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
