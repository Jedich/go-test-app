package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	portMap := map[string]int64{
		"DEV":  8080,
		"PROD": 80,
	}

	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		panic("ENVIRONMENT env variable not set!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("I am a %s environment", env))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", portMap[env]), nil)
	if err != nil {
		panic(err)
	}
}
