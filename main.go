package main

import "time"

type Attachment struct {
	Name     string    `json:"Name"`
	Date     time.Time `json:"Date"`
	Contents string    `json:"Contents"`
}

type Task struct {
	ID          int           `json:"Id"`
	Text        string        `json:"Text"`
	Tags        []string      `json:"Tags"`
	Due         time.Time     `json:"Due"`
	Attachments []*Attachment `json:"Attachments"`
}
