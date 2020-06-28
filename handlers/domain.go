package handlers

import (
	"context"
	"fmt"
	"goglobalrank/helper"
	"goglobalrank/models"
	"math"
	"net/http"
	"strconv"
	"time"
)
func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GetDomainrank(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	if r.Method == "POST" || r.Method == "OPTIONS" {

		_ = r.ParseForm()

		ctx, cancel := context.WithCancel(context.Background())
		start := time.Now()
		defer cancel()

		limit, err := strconv.Atoi(r.FormValue("limit"))
		offset, err := strconv.Atoi(r.FormValue("offset"))
		if err != nil {
			limit = 100
			offset = 0
		}

		domains, err := models.GetAll(ctx, limit, offset)

		helper.ErrorCheck(err)

		helper.ResponseJSON(w, domains, http.StatusOK)
		duration := time.Since(start)
		fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
		return

	} else {
		helper.ResponseJSON(w, "Wrong Method", http.StatusMethodNotAllowed)
	}

}
