package main

import (
	"time"
)

type Genre struct {
	Name string
	Rows []Genre
}

type Setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgDB       string
	PgPass     string
	Image      string
	Data       string
	Assets     string
	HTML       string
}

type User struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
	Admin    bool   `json:"Admin"`
	Rows     []User
}

type News struct {
	Title       string
	TextContent string
	PostDate    time.Time
	Image       string
	Rows        []News
}
