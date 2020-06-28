package models

import (
	"context"
	"fmt"
	"goglobalrank/config"
	"goglobalrank/helper"
)

type Domain struct {
	ID 			int64 `json:"id"`
	Domain     string `json:"domain"`
	GlobalRank int64  `json:"globalrank"`
	TLD        string `json:"tld"`
	IDNDomain  string `json:"idndomain"`
}

const (
	table = "domain"
)

func GetAll(ctx context.Context, limit int, offset int) ([]Domain, error) {
	var (
		domains []Domain
		e       error
		domain  Domain
		l  = limit
		o  = offset
	)
	db, err := config.MYSQL()
	helper.ErrorCheck(err)
	defer db.Close()

	queryText := fmt.Sprintf("SELECT id, globalrank, domain, tld, idndomain FROM %v ORDER BY globalrank LIMIT %d OFFSET %d", table, l, o)
	fmt.Println(queryText)
	rowQuery, err := db.QueryContext(ctx, queryText)
	helper.ErrorCheck(err)

	defer rowQuery.Close()
	for rowQuery.Next() {
		e = rowQuery.Scan(&domain.ID, &domain.GlobalRank, &domain.Domain, &domain.TLD, &domain.IDNDomain)
		helper.ErrorCheck(e)
		domains = append(domains, domain)
	}

	return domains, nil
}
