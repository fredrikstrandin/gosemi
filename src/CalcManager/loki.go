package main

import (
	"log"
	"os"
	"time"

	"github.com/afiskon/promtail-client/promtail"
)

var (
	Loki promtail.Client
	err  error
)

func init() {
	lokiUrl := os.Getenv("LOKI_URL")
	format := os.Getenv("LOKI_FORMAT") // "proto" or "json"
	source_name := os.Getenv("LOKI_SORCE")
	job_name := os.Getenv("LOKI_JOB")

	
	labels := "{source=\"" + source_name + "\",job=\"" + job_name + "\"}"
	conf := promtail.ClientConfig{
		PushURL:            lokiUrl,
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}

	if format == "proto" {
		Loki, err = promtail.NewClientProto(conf)
	} else {
		Loki, err = promtail.NewClientJson(conf)
	}

	if err != nil {
		log.Printf("promtail.NewClient: %s\n", err)
		os.Exit(1)
	}
}
