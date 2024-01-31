package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	jwtSecret := os.Getenv("JWT_SECRET")

	fmt.Println(jwtSecret)
}
