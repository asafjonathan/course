package main

import (
	"net/http"
	"os"
	"tailor/cmd"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	port := os.Getenv("PORT")
	rotes := cmd.Route()

	http.ListenAndServe(port, rotes)
}
