package models

import (
	"context"
	"fmt"
	"goglobalrank/config"
	"goglobalrank/helper"
)

type Domain struct {
	Domain     string `json:"Domain"`
	GlobalRank int64  `json:"GlobalRank"`
	TLD        string `json:"TLD"`
	IDNDomain  string `json:"IDN_Domain"`
}

const (
	table = "domain"
	//layoutDatetime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context, limit int, offset int) ([]Domain, error) {

	var (
		domains []Domain
		e       error
		domain  Domain
	)

	l := limit
	o := offset

	db, err := config.MYSQL()
	helper.ErrorCheck(err)

	defer db.Close()
	queryText := fmt.Sprintf("SELECT GlobalRank, Domain, TLD, IDN_Domain FROM %v ORDER BY GlobalRank LIMIT %d OFFSET %d", table, l, o)
	fmt.Println(queryText)
	rowQuery, err := db.QueryContext(ctx, queryText)

	helper.ErrorCheck(err)

	defer rowQuery.Close()

	for rowQuery.Next() {
		e = rowQuery.Scan(&domain.GlobalRank, &domain.Domain, &domain.TLD, &domain.IDNDomain)
		helper.ErrorCheck(e)
		domains = append(domains, domain)
	}

	return domains, nil

}
