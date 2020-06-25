package handlers

import (
	"context"
	"fmt"
	"goglobalrank/helper"
	"goglobalrank/models"
	"math"
	"net/http"
	"time"
)

func GetDomainrank(w http.ResponseWriter, r *http.Request)  {
	r.Method = "GET"
	ctx, cancel := context.WithCancel(context.Background())
	start := time.Now()
	defer cancel()

	domains, err := models.GetAll(ctx)
	helper.ErrorCheck(err)

	helper.ResponseJSON(w, domains, http.StatusOK)
	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
	return
}


