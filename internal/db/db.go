package db

import (
	"context"
	"database/sql"
	"time"
)

func New (addr string, maxOpenConns,maxIdleConns int, maxIdleTime string) (*sql.DB,error) {

	db, err := sql.Open("postgres",addr)
	if err!=nil {
		return nil,err
	}
	db.SetMaxOpenConns(maxOpenConns)
	duration,err := time.ParseDuration(maxIdleTime)
	db.SetConnMaxIdleTime(duration)
	db.SetMaxIdleConns(maxIdleConns)
	

	if err !=nil {
		return nil,err
	}
		ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err !=nil {
		return nil,err
	}
	return db,nil

}