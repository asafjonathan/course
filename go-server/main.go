package main

import (
	"fmt"
	"net/http"
	"os"
	"tailor/cmd"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	port := os.Getenv("PORT")
	rotes := cmd.Route()
	fmt.Println("ssss")
	http.ListenAndServe(port, rotes)
}
