package main

import (
        "encoding/json"
        "io/ioutil"
        "net/http"
        "fmt"
        "github.com/gorilla/mux"
)
 
type NameStr struct {
         Name string
 }
 

 func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to httprouter server!\n")
}

 func getMembersHandler(res http.ResponseWriter, req *http.Request) {
 		vars := mux.Vars(req)
 		inputName := vars["name"]
        fmt.Fprintf(res, "Hello, %s!\n", inputName)
 }
 
func postMembersHandler(w http.ResponseWriter, r *http.Request) {
         w.Header().Set("Content-Type", "application/json")
        var m NameStr
         b, _ := ioutil.ReadAll(r.Body)
         json.Unmarshal(b, &m)
 
 		resString:="Hello, "
 		resString+=string(m.Name)
 		resString+="!"

 		mapD := map[string]string{"greetings":resString}
		mapB, _ := json.Marshal(mapD)

        w.Write(mapB)
 }
 
func main() {
         r := mux.NewRouter()
         r.HandleFunc("/", Index).Methods("GET")
         r.HandleFunc("/hello/{name}", getMembersHandler).Methods("GET")
         r.HandleFunc("/hello", postMembersHandler).Methods("POST")

         http.Handle("/", r)
         http.ListenAndServe(":8080", nil)
 }