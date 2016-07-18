package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

type WandWWW interface {
	Host() bool
}

type WandScriptResult struct {
	Value string `json:"value"`
}

func WandFileHandler(w http.ResponseWriter, r *http.Request, config WandConfig) {
	w.Header().Set("Content-Type", config.ContentType)
	http.ServeFile(w, r, config.Filepath)
}

func WandScriptHandler(w http.ResponseWriter, r *http.Request, config WandConfig) {
	output, err := exec.Command(config.Filepath, "run").Output()

	if err != nil {
		log.Fatal(err)
	}

	result := WandScriptResult{
		Value: string(output),
	}
	json, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (config WandConfig) Host() bool {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if config.Script {
			WandScriptHandler(w, r, config)
		} else {
			WandFileHandler(w, r, config)
		}
	})
	port := fmt.Sprintf(":%s", config.Port)

	if config.Script {
		color.Blue("Wand: Hosting '%s' as a script on localhost:%s.", config.Filename, config.Port)
	} else {
		color.Green("Wand: Hosting '%s' on localhost:%s.", config.Filename, config.Port)
	}
	log.Fatal(http.ListenAndServe(port, r))

	return true
}
