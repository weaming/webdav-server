package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

func getEnv(name, def string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return def
}

func main() {
	dir := getEnv("ROOT", ".") // root of webdav
	httpPort := getEnv("HTTP_ADDR", ":80")
	httpsPort := getEnv("HTTPS_ADDR", ":443")
	serveSecure := getEnv("HTTPS_ENABLE", "") != "" // enable https
	certKey := getEnv("CERT_KEY", "cert.pem")
	pubKey := getEnv("PUB_KEY", "key.pem")
	authUsers := getEnv("AUTH_PATH", "auth.json") // HTTP basic auth table in name-password format

	srv := &webdav.Handler{
		FileSystem: webdav.Dir(dir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Printf("WEBDAV [%s]: %s, ERROR: %s\n", r.Method, r.URL, err)
			} else {
				log.Printf("WEBDAV [%s]: %s \n", r.Method, r.URL)
			}
		},
	}

	http.Handle("/", basicAuth(srv.ServeHTTP, authUsers))

	if serveSecure {
		if _, err := os.Stat(certKey); err != nil {
			log.Fatalf("[x] cert key for https not found: %s", certKey)
		}
		if _, er := os.Stat(pubKey); er != nil {
			log.Fatalf("[x] public key for https not found: %s", pubKey)
		}

		log.Printf(fmt.Sprintf("serve https on port %s", httpsPort))
		go http.ListenAndServeTLS(httpsPort, certKey, pubKey, nil)
	}

	log.Printf(fmt.Sprintf("serve http on port %s", httpPort))
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatalf("Error with WebDAV server: %v", err)
	}
}
