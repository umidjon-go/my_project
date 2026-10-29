package models

import "time"

type Branch struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	Location  string     `json:"location"`
	CreatedAt *time.Time `json:"created_at"`
}
