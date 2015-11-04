package fileServers

import (
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