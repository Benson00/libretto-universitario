package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"text/template"
	"time"
)

type Esame struct {
	nome string
	data time.Time
	voto int
	cfu  int
}

type Libretto struct {
	esami []Esame
}

func (l *Libretto) AggiungiEsame(e Esame) {
	l.esami = append(l.esami, e)
	sort.Slice(l.esami, func(i, j int) bool {
		return l.esami[i].data.Before(l.esami[j].data)
	})
}

func (l *Libretto) MediaPonderata() float64 {
	var num float64
	var den int
	for _, esame := range l.esami {
		num += float64(esame.voto) * float64(esame.cfu)
		den += esame.cfu
	}
	return num / float64(den)
}

func (l *Libretto) VotoLaurea() float64 {
	return l.MediaPonderata() * 110 / 30
}

func main() {
	libretto := Libretto{}

	libretto.AggiungiEsame(Esame{"Matematica", time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC), 28, 6})
	libretto.AggiungiEsame(Esame{"Fisica", time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC), 30, 6})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libretto)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("template.html"))
		tmpl.Execute(w, struct {
			MediaPonderata float64
			VotoLaurea     float64
		}{
			MediaPonderata: libretto.MediaPonderata(),
			VotoLaurea:     libretto.VotoLaurea(),
		})
	})

	http.ListenAndServe(":8080", nil)
}
