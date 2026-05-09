package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type commandStruct struct {
	Command string `json:"command"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var cmd commandStruct

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("This is a bad input"))
	}

	fmt.Println("Command", cmd.Command)

	out, err1 := exec.Command(cmd.Command).Output()
	if err1 != nil {
		w.Write([]byte("Invalid command"))
	}
	st := fmt.Sprintf("Executed command. Output is := %v", string(out))
	w.Write([]byte(st))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /run", handleRequest)

	log.Print("Starting server at port 4000")
	http.ListenAndServe(":4000", mux)
}
