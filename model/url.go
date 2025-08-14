package model

import (
	"errors"
	"fmt"
	"go-url-shrtr/db"
	"strings"
	"time"
)

type Url struct {
	Id         int64     `json:"id"`
	Code       string    `json:"code" binding:"required"`
	Url        string    `json:"url" binding:"required"`
	CreatedAt  time.Time `json:"createdAt" binding:"required" time_format:"2006-01-02T15:04:05Z"`
	ClickCount int64     `json:"clickCount"`
}

func (u *Url) Save() error {
	fmt.Println("url", u.Url)
	if !strings.HasPrefix(u.Url, "https://") && !strings.HasPrefix(u.Url, "http://") {
		return errors.New("invalid  url")
	}

	existingCodeQuery := `SELECT * FROM urls WHERE code = ?`

	row := db.DB.QueryRow(existingCodeQuery, u.Code)
	if row == nil {
		return errors.New("Your preferred short code already exists.")
	}

	saveQuery := "INSERT INTO urls(code , url,created_at , click_count) VALUES(?,?,?,?)"

	stmt, err := db.DB.Prepare(saveQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Code, u.Url, u.CreatedAt, u.ClickCount)
	if err != nil {
		return err
	}

	urlId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.Id = urlId

	return nil
}
