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

type ConfigConnProvider struct {
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}

func NewPSQL(ctx context.Context, connString string, cfg *ConfigConnProvider, ormConfig *gorm.Config) (*Provider, error) {
	orm, err := gorm.Open(postgres.Open(connString), ormConfig)
	if err != nil {
		return nil, err
	}

	db, err := orm.WithContext(ctx).DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return &Provider{orm}, nil
}
