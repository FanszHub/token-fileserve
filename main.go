package main

import (
	"os"
	"bufio"
	"log"
	"flag"
    "net/http"
)

var dir string
var tokensFile string
var listen string

func init() {
	flag.StringVar(&dir, "directory", ".", "Directory to serve files from")
	flag.StringVar(&tokensFile, "tokens", "", "File containing tokens")
	flag.StringVar(&listen, "listen", ":5760", "IP/Port to listen on")
	flag.Parse()
}

type TokenFileServer struct {
	fileServer http.Handler
	tokens []string
}

func (tokenFileServer *TokenFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpToken := r.Header.Get("Token");
	
	
	for _, token := range tokenFileServer.tokens {
		if(token == httpToken){
			log.Printf("File %s served", r.URL.String());
			tokenFileServer.fileServer.ServeHTTP(w,r);
			return;
		}
	}
	
	log.Printf("File %s rejected", r.URL.String());
	http.Error(w, "Invalid token", 403);
}

func NewTokenFileServer(tokens []string, dir string) http.Handler{
	return &TokenFileServer{
		fileServer: http.FileServer(http.Dir(dir)),
		tokens: tokens,
	}
}

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
	
	handler := NewTokenFileServer(tokens, dir);
    http.ListenAndServe(listen, handler)
}