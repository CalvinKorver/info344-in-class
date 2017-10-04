package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/calvinkorver/info344-in-class/zipsvr/handlers"
	"github.com/calvinkorver/info344-in-class/zipsvr/models"
)

const zipsPath = "/zips/"

func helloHandler(w http.ResponseWriter, r *http.Request) { // Remember, its variable name , variable type
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello %s!", name)
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC()
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(stats)

}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatal("")
	}
	log.Printf("loaded %d zips", len(zips))

	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	// fmt.Println("Hellow World!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: zipsPath,
	}
	mux.Handle(zipsPath, cityHandler)

	fmt.Printf("server is listening at http://%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux)) // "Hostname : Port" passed in as a string
}
