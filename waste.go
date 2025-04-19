package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type WasteCalculation struct {
	Type int
	Mass float64
	Cost float64
}

var ErrInvalidWaste = errors.New("invalid waste calculation input")

func NewWasteCalculation(waste_type int, mass float64) (WasteCalculation, error) {
	if waste_type < 0 || mass < 0 || waste_type > 2 {
		return WasteCalculation{}, ErrInvalidWaste
	}
	rate_first := 222907.36
	rate_second := 62468.26
	cost := mass
	if waste_type == 1 {
		cost *= rate_first
	} else {
		cost *= rate_second
	}
	return WasteCalculation{waste_type, mass, cost}, nil
}

func getWasteCalc(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	tp := q.Get("type")
	mass := q.Get("mass")
	if tp == "" || mass == "" {
		http.Error(w, "unspecified mass or waste type", http.StatusBadRequest)
		return
	}
	tp_i, err := strconv.Atoi(tp)
	if err != nil {
		http.Error(w, "invalid waste type format", http.StatusBadRequest)
		return
	}
	mass_f, err := strconv.ParseFloat(mass, 64)
	if err != nil {
		http.Error(w, "invalid waste mass format", http.StatusBadRequest)
		return
	}
	c, err := NewWasteCalculation(tp_i, mass_f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("%f", c.Cost)))
}
