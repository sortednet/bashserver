package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"os/exec"
)

var cmd string
var args []string
var version = "latest"

func handler(w http.ResponseWriter, r *http.Request) {
	if strings.Compare(r.URL.Path, "/version") == 0 {
		fmt.Fprintf(w, version)
	} else {
		// log.Printf("Call %s with args", cmd)	
		// fmt.Printf("Call %s with args", cmd)	
		// fmt.Fprintf(w, "Call %s with args", cmd)	
		cmdOut, err := exec.Command(cmd, args...).Output();
		output := string(cmdOut)
		if err == nil {
			log.Printf("%s returned %s", cmd, output)
			fmt.Fprintf(w, output)
		} else {
			fmt.Fprintf(w, "There was an error running %s : %s", cmd, err)
		}
	}
	
}

func main() {

	cmd = os.Args[1]
	args = os.Args[2:]

	envVersion, versionSet := os.LookupEnv("VERSION")
	if versionSet {
		version = envVersion
	}	

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
