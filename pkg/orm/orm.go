package orm

import (
	"context"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Provider struct {
	*gorm.DB
}

func NewPSQL(ctx context.Context, connString string) (*Provider, error) {
	orm, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := orm.WithContext(ctx).DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return &Provider{orm}, nil
}
