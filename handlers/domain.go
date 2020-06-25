package handlers

import (
	"context"
	"goglobalrank/helper"
	"goglobalrank/models"
	"net/http"
)

func GetDomainrank(w http.ResponseWriter, r *http.Request)  {
	r.Method = "GET"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	domains, err := models.GetAll(ctx)

	if err != nil {
		println(err)
	}

	helper.ResponseJSON(w, domains, http.StatusOK)
	return
}


