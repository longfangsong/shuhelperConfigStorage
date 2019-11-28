package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"shuhelperConfigStorage/model"
	"shuhelperConfigStorage/service/token"
)

func getConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenInHeader := r.Header.Get("Authorization")
	if len(tokenInHeader) <= 7 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id := token.StudentIdForToken(tokenInHeader[7:])
	config, err := model.Get(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	data, err := json.Marshal(config)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func setConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenInHeader := r.Header.Get("Authorization")
	if len(tokenInHeader) <= 7 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id := token.StudentIdForToken(tokenInHeader[7:])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	var inputForm struct {
		Mode           string `json:"mode"`
		SaveSettingsIn string `json:"saveSettingsIn"`
		SaveTodoIn     string `json:"saveTodoIn"`
	}
	err = json.Unmarshal(body, &inputForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	if inputForm.SaveSettingsIn != "server" {
		http.Error(w, "Should not upload this config!", http.StatusNotAcceptable)
		return
	}
	err = model.Save(model.Config{
		StudentId:  id,
		Mode:       inputForm.Mode,
		SaveToDoIn: inputForm.SaveTodoIn,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getConfigHandler(w, r)
	case "POST":
		setConfigHandler(w, r)
	case "PUT":
		setConfigHandler(w, r)
	}
}

func PingPongHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}
