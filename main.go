package main

import (
	"crypto/tls"
	//"encoding/json"
	"fmt"
	//"github/go-backend/model"
	"github/go-backend/router"
	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := router.NewRouter()
	puerto := os.Getenv("PUERTO")
	fmt.Println("Serve on port ", puerto)
	server := http.ListenAndServe(":"+puerto, router)
	log.Fatal(server)
}
