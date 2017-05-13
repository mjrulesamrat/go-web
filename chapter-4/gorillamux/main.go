package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Creation time.Time `json:"creation"`
}

//Store for the Notes collection
var noteStore = make(map[string]Note)

// variable to generate key for the collection
var id int = 0

func GetNoteHandler(w http.ResponseWriter, r *http.Request){
	var notes []Note
	for _, v := range noteStore{
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Post a new note source
func PostNoteHandler(w http.ResponseWriter, r *http.Request){
	var note Note
	// decode incoming note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}

	note.Creation = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// Update saved one using PUT method
func PutNoteHandler(w http.ResponseWriter, r *http.Request){
	var err error
	vars := mux.Vars(r)
	k := vars["id"]

	var noteToUpdate Note
	err = json.NewDecoder(r.Body).Decode(&noteToUpdate)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok {
		noteToUpdate.Creation = note.Creation
		delete(noteStore, k)
		noteStore[k] = noteToUpdate
	} else {
		log.Printf("Could not fnd key for Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}


func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}