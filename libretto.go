package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type Esame struct {
	Nome    string
	Data    time.Time
	Voto    int
	Crediti int
}

type Libretto struct {
	Esami []Esame
}

func (l *Libretto) AggiungiEsame(e Esame) {
	l.Esami = append(l.Esami, e)
	sort.Slice(l.Esami, func(i, j int) bool {
		return l.Esami[i].Data.Before(l.Esami[j].Data)
	})
}

func (l *Libretto) MediaPonderata() float64 {
	var sommaPesata float64
	var creditiTotali int
	for _, esame := range l.Esami {
		sommaPesata += float64(esame.Voto) * float64(esame.Crediti)
		creditiTotali += esame.Crediti
	}
	return sommaPesata / float64(creditiTotali)
}

func (l *Libretto) VotoLaurea() float64 {
	return l.MediaPonderata() * 11 / 3
}

func formatNumber(n float64) string {
	return fmt.Sprintf("%.2f", n)
}

func main() {
	libretto := Libretto{}
	libretto.AggiungiEsame(Esame{"Matematica", time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC), 28, 6})
	libretto.AggiungiEsame(Esame{"Fisica", time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC), 30, 6})
	libretto.AggiungiEsame(Esame{"Fisica", time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC), 20, 6})
	libretto.AggiungiEsame(Esame{"Fisica", time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC), 26, 6})
	libretto.AggiungiEsame(Esame{"Fisica", time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC), 18, 6})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(libretto); err != nil {
			log.Println("Error encoding JSON:", err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("template.html"))
		if err := tmpl.Execute(w, struct {
			MediaPonderata string
			VotoLaurea     string
			Esami          []Esame
		}{
			MediaPonderata: formatNumber(libretto.MediaPonderata()),
			VotoLaurea:     formatNumber(libretto.VotoLaurea()),
			Esami:          libretto.Esami,
		}); err != nil {
			log.Println("Error executing template:", err)
		}
	})

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
