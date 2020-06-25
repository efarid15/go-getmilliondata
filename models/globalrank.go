package models

import (
	"context"
	"fmt"
	"goglobalrank/config"
	"log"
)

type Domain struct {
	Domain	string `json:"Domain"`
	GlobalRank int64 `json:"GlobalRank"`
	TLD	string `json:"TLD"`
	IDNDomain string `json:"IDN_Domain"`
}

const (
	table          = "domain"
	//layoutDatetime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]Domain, error) {
	var domains []Domain
	db, err := config.MYSQL()
	if err != nil {
		log.Fatal("Error Database Connection", err)
	}
	defer db.Close()
	queryText := fmt.Sprintf("SELECT GlobalRank, Domain, TLD, IDN_Domain FROM %v ORDER BY GlobalRank", table)

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	defer rowQuery.Close()

	for rowQuery.Next() {
		var domain Domain
		if err = rowQuery.Scan(&domain.GlobalRank, &domain.Domain, &domain.TLD, &domain.IDNDomain); err != nil {
			fmt.Printf("%s \n", err)
			return nil, err
		}
		domains = append(domains, domain)
	}

	return domains, nil
}
