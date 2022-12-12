package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// the server will retrieve the user from the body, and randomly generate a status to return
func server(w http.ResponseWriter, r *http.Request) {
	var status string
	var user string

	userStatus.WithLabelValues(user, status).Inc()

	var mr MyRequest
	json.NewDecoder(r.Body).Decode(&mr)

	if rand.Float32() > 0.8 {
		status = "4xx"
	} else {
		status = "2xx"
	}

	user = mr.User
	Loki.Infof(user, status)
	w.Write([]byte(status))
}

// the producer will randomly select a user from a pool of users and send it to the server
func producer() {
	userPool := []string{"bob", "alice", "jack"}
	for {
		postBody, _ := json.Marshal(MyRequest{
			User: userPool[rand.Intn(len(userPool))],
		})
		requestBody := bytes.NewBuffer(postBody)
		http.Post("http://calcmanager:8080", "application/json", requestBody)
		time.Sleep(time.Second * 2)
	}
}

func main() {

	// the producer runs on its own goroutine
	go producer()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", server)

	http.ListenAndServe(":5000", nil)

	Loki.Infof("starting web server")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		Loki.Errorf(err.Error())
	}

	Loki.Shutdown()
}
