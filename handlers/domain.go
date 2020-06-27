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

func GetDomainrank(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

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
