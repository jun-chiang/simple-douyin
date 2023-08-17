package models

import (
	"fmt"
	"time"
)

type DateTime struct {
	time.Time
}

func (t *DateTime) MarshaJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2023-08-25 15:15:15"))
	return []byte(output), nil
}

// TODO: json yyyy-MM-dd HH:MM::ss   gorm time.Time
