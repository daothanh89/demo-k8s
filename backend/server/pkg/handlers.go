package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type handlers struct {
	users map[string]string
}

func (h *handlers) login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	loginBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.Unmarshal(loginBody, &body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if password, ok := h.users[body.Username]; ok {
		if body.Password == password {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(generateToken(body.Username)))
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect password"))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("User not found!"))
}

func (h *handlers) register(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	loginBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.Unmarshal(loginBody, &body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := h.users[body.Username]; ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User already exsit!"))
		return
	}

	h.users[body.Username] = body.Password
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Resigter OK!"))
}

func (h *handlers) quoteGet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
