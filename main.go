package main

import (
	"os"
	"bufio"
	"log"
    "net/http"
	"github.com/mattdotmatt/token-fileserve/fileServers"
)

var dir string
var tokensFile string
var listen string

func getTokens(path string) []string {
	var tokens []string

	log.Println("Loading tokens");

	file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		text := scanner.Text();
		if text != "" {
			tokens = append(tokens, scanner.Text());
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return tokens
}

func main() {

	var tokens []string;
	if(tokensFile != ""){
		tokens = getTokens(tokensFile);
	}
	
	log.Println("Starting server");

	handler := token_fileserve.NewTokenFileServer(tokens, dir);

    http.ListenAndServe(listen, handler)
}