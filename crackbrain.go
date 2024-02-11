package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	lp := filepath.Join(ExecutablePath+"/templates", "layout.html")
	table := NewCrackbrainTables(seed)
	scenario := CreateScenario(table)

	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		panic("Error parsing templates" + err.Error())
	}
	tmpl.ExecuteTemplate(w, "layout", scenario)
}

func perform_main() {

	n := 1
	if len(os.Args) > 1 {
		z, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("Usage: crackbrain <num_of_scenarios")
		} else {
			if n < 1 {
				panic("Please specify a positive number of scenarios")
			}
			n = z
		}
	}

	//rand.NewSource(time.Now().UnixNano())
	table := NewCrackbrainTables(1)
	builder := strings.Builder{}
	for h := 1; h <= n; h++ {
		scenario := CreateScenario(table)
		/*
			b, err := json.MarshalIndent(scenario, "", "  ")
			if err != nil {
				fmt.Printf("Error json" + err.Error())
			} else {
				res := string(b[:])
				fmt.Println(res)
			}
		*/
		res, err := template.ParseFiles("template.txt")
		if err != nil {
			fmt.Printf("Error template" + err.Error())

		}
		var bufString bytes.Buffer
		res.Execute(&bufString, scenario)
		builder.WriteString(bufString.String())
		if h < n {
			builder.WriteString("\n\n#pagebreak()\n\n")
		}

	}
	fmt.Println(builder.String())
}
