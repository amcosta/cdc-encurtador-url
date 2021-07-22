package url

import (
	"crypto/rand"
	"net/http"
	"time"
)

const (
	tamanho = 5
	simbolos = "abcdefghijklmnopqr...STUVWXYZ1234567890_+-"
)

type Url struct {
	Id      string
	Criacao time.Time
	Destino string
}

type Headers map[string]string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Encurtador(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responderCom(w, http.StatusMethodNotAllowed, Headers{
			"Allow": "POST"
		})

		return
	}

	url := extrairUrl(r)
}

func responderCom(w http.ResponseWriter, status int, headers Headers) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)
}

func extrairUrl(r *http.Request) string {
	url := make([]byte r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}