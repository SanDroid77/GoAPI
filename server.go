package main

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/SanDroid77/REST/connect"
  "github.com/SanDroid77/REST/structs"
)

func main() {
  connect.InitializaDatabase()
  defer connect.CloseConnection()
  r := mux.NewRouter()
  r.HandleFunc("/user/{id}", GetUser).Methods("GET")
  r.HandleFunc("/user/new", NewUser).Methods("POST")
  r.HandleFunc("/user/update/{id}", UpdateUser).Methods("PATCH")
  r.HandleFunc("/user/delete/{id}", DeleteUser).Methods("DELETE")

  log.Println("El servidor se encuentra en el puerto 8000")
  log.Fatal(http.ListenAndServe(":8000", r))
}

func GetUser(w http.ResponseWriter, r* http.Request) {
  vars := mux.Vars(r)
  user_id := vars["id"]

  status := "success"
  var message string

  user := connect.GetUser(user_id)

  if user.Id <= 0 {
    status = "error"
    message = "User not found."
  }

  response := structs.Response{status, user, message}
  json.NewEncoder(w).Encode(response)
}

func NewUser(w http.ResponseWriter, r* http.Request) {
  user := GetUserRequest(r)
  connect.CreateUser(user)
  response := structs.Response{"success", connect.CreateUser(user), ""}
  json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r* http.Request) {
  vars := mux.Vars(r)
  user_id := vars["id"]

  user := GetUserRequest(r)
  
  response := structs.Response{"success", connect.UpdateUser(user_id, user), ""}
  json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r* http.Request) {
  vars := mux.Vars(r)
  user_id := vars["id"]
  
  var user structs.User

  connect.DeleteUser(user_id)

  response := structs.Response{"success", user, ""}
  json.NewEncoder(w).Encode(response)
}

func GetUserRequest(r* http.Request) structs.User{
  var user structs.User

  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&user)
  if err != nil {
    log.Fatal(err)
  }
  return user
}