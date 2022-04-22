package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	javaHome := os.Getenv("JAVA_HOME")
	ip := ReadUserIP(r)
	w.Header().Set("RequestClient", userAgent)
	w.Header().Set("javaHome", javaHome)
	w.WriteHeader(200)
	fmt.Fprintf(w, "URL.Path = %q\n userAgent = %q\n javaHome = %q\n ip = %q\n",
	r.URL.Path, userAgent, javaHome, ip)
}

func ReadUserIP(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-Ip")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarded-For")
    }
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
    }
    return IPAddress
}
