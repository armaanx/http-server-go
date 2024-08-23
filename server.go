package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	fmt.Fprint(w, GetPlayerScore(player))
// }

func GetPlayerScore(name string) string {
	if name == "Armaan" {
		return "20"
	}
	if name == "Blake" {
		return "10"
	}
	return ""
}
