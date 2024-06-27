package main

import (
    "encoding/json"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

// Estrutura para representar a vovó
type Vovo struct {
    Nome   string `json:"nome"`
    Idade  int    `json:"idade"`
    Cidade string `json:"cidade"`
}

// Função handler para o verbo GET
func getVovo(w http.ResponseWriter, r *http.Request) {
    vovo := Vovo{"Tia Claudia", 91, "Araguari"}
    jsonResponse, err := json.Marshal(vovo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

// Função handler para o verbo POST
func postVovo(w http.ResponseWriter, r *http.Request) {
    var vovo Vovo
    err := json.NewDecoder(r.Body).Decode(&vovo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    jsonResponse, err := json.Marshal(vovo)
	w.Write([]byte(`{"message":"}`))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

// Função handler para o verbo DELETE
func deleteVovo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message":"Vovo deleted successfully"}`))
}

func main() {
    r := mux.NewRouter()
    
    r.HandleFunc("/vovo", getVovo).Methods("GET")
    r.HandleFunc("/vovo", postVovo).Methods("POST")
    r.HandleFunc("/vovo", deleteVovo).Methods("DELETE")

    log.Print("Servidor rodando na porta 8080")
    err := http.ListenAndServe(":8080", r)
    if err != nil {
        log.Fatal(err)
    }
}
