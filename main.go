package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type commandStruct struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var cmd commandStruct

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("This is bad input"))
		return
	}

	fmt.Println("Command : ", cmd.Command)
	fmt.Println("Args : ", cmd.Args)

	out, err1 := exec.Command(cmd.Command, cmd.Args...).CombinedOutput()
	if err1 != nil {
		w.Write([]byte(err1.Error()))
		return
	}
	st := fmt.Sprintf("Executed command.\nOutput is :=\n%s\n", string(out))
	w.Write([]byte(st))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /run", handleRequest)

	log.Print("Starting server at port 4000")
	http.ListenAndServe(":4000", mux)
}
