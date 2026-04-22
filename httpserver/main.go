package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "admin" {
		// A senha é "admin123"
		return "$1$phxcRdHK$Hs5LE0DvxYlm1dlKHM09k0" // hash da senha "admin123"
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}

	httpDir := os.Args[1]
	porta := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	http.Handle("/", authenticator.Wrap(func(w http.ResponseWriter, ar *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &ar.Request)
	}))

	log.Fatal(http.ListenAndServe(":"+porta, nil))
}