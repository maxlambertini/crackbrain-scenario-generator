package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var ExecutablePath = ""

func main() {
	ExecutablePath = GetExecutablePath()
	fmt.Println("Crackbrain running in ", ExecutablePath)
	ct := mime.TypeByExtension(".css")
	fmt.Printf("ct: %s\n", ct)

	router := httprouter.New()
	mime.AddExtensionType(".css", "text/css; charset=utf-8")

	//fs := http.FileServer(http.Dir("./static"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	staticDir := http.Dir(ExecutablePath + "/static")
	router.GET("/", serveTemplate)
	router.GET("/api", serveJson)
	router.GET("/typst", serveTypst)
	router.ServeFiles("/static/*filepath", staticDir)

	//rand.NewSource(time.Now().UnixNano())

	log.Print("Listening on :3069...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3069", router))

}

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func serveTemplate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	seed := extractSeedFromQueryString(r)
	lp := filepath.Join(ExecutablePath+"/templates", "layout.html")
	table := NewCrackbrainTables(seed)
	scenario := CreateScenario(table)
	host := r.Host
	scheme := "http"
	if r.URL.Scheme != "" {
		scheme = r.URL.Scheme
	}
	fmt.Println(r.RequestURI)
	scenario.Url = fmt.Sprintf("%s://%s/?seed=%d", scheme, host, seed)

	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		panic("Error parsing templates" + err.Error())
	}
	tmpl.ExecuteTemplate(w, "layout", scenario)
}

func serveJson(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	seed := extractSeedFromQueryString(r)
	table := NewCrackbrainTables(seed)
	scenario := CreateScenario(table)
	host := r.Host
	scheme := "http"
	if r.URL.Scheme != "" {
		scheme = r.URL.Scheme
	}
	fmt.Println(r.RequestURI)
	scenario.Url = fmt.Sprintf("%s://%s/api/?seed=%d", scheme, host, seed)
	scenario.Seed = seed
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scenario)
}

func serveTypst(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	seed := extractSeedFromQueryString(r)
	lp := filepath.Join(ExecutablePath+"/templates/typst", "template.txt")

	table := NewCrackbrainTables(seed)
	scenario := CreateScenario(table)
	host := r.Host
	scheme := "http"
	if r.URL.Scheme != "" {
		scheme = r.URL.Scheme
	}
	fmt.Println(r.RequestURI)
	scenario.Url = fmt.Sprintf("%s://%s/api/?seed=%d", scheme, host, seed)
	scenario.Seed = seed
	w.Header().Set("Content-Type", "text/plain")
	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		panic("Error parsing templates" + err.Error())
	}
	var bufString bytes.Buffer
	tmpl.Execute(&bufString, scenario)
	w.Write(bufString.Bytes())
}

func extractSeedFromQueryString(r *http.Request) int64 {
	seed := time.Now().UnixNano()
	sSeed := r.URL.Query().Get("seed")
	if sSeed != "" {
		fmt.Println("Setting seed: ", sSeed)
		seed2, err := strconv.ParseInt(sSeed, 10, 64)
		if err != nil {
			fmt.Println("[WARN] Error converting to int64 parameter seed", r.Form.Get("seed"))
		} else {
			if seed2 < 0 {
				seed = -seed2
			} else {
				seed = seed2
			}
		}
	} else {
		fmt.Println("SEED:", r.Form.Get("seed"))
	}
	return seed
}
