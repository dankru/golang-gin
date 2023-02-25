package main

import (
	"database/sql"
	"errors"
	"fmt"
)

var db *sql.DB
// Словарь Queries. Ключ - строка, значение - sql statement, иначе говоря запрос
var Queries map[string]*sql.Stmt

func connect() error {
	var e error
	
	Queries = make(map[string]*sql.Stmt)

	// Генерируем строку подключения со значениями из кфг файла
	db, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgDB))
	if e != nil {
		return e
	}

	return nil
}

func addContent(table string, row string, value string) {
	_, e := db.Exec(fmt.Sprintf(`INSERT INTO "%s" ("%s") VALUES('%s')`, table, row, value))
	if e != nil{
		panic(e.Error())
	}
}

func prepareQueries() {
	var e error
	
	Queries["Select#Genre"], e = db.Prepare(`Select "Name" from "Genre" order by "Name"`)
	Queries["Select#User"], e = db.Prepare(`Select "Login", "Password", "Admin" from "User" where "Login"=$1 and "Password"=$2 order by "Admin"`)
	Queries["Select#News"], e = db.Prepare(`Select "Title", "TextContent", "PostDate", "Image" from "News"`)
	Queries["Insert#News"], e = db.Prepare(`Insert Into "News" ("Title", "TextContent", "PostDate", "Image") values($1, $2, $3, $4)`)
	Queries["Delete#News"], e = db.Prepare(`Delete from "News" where "Title"=$1`)
	
	if e != nil {
		panic(e.Error())
	}

}

func (m *Genre) Select() error {
	stmt, ok := Queries["Select#Genre"]
	if !ok {
		return errors.New("Chosen query doesn't exist")
	}
	
	rows, e := stmt.Query()
	if e != nil{
		return e
	}

	for rows.Next() {
		e = rows.Scan(&m.Name)

		if e != nil {	
			return e
		}

		m.Rows = append(m.Rows, Genre{Name: m.Name})
	}
	return nil
}

func (m *User) Select() error {
	stmt, ok := Queries["Select#User"]
	if !ok {
		return errors.New("Chosen query doesn't exist")
	}
 
	r := stmt.QueryRow(m.Login, m.Password)
	e := r.Scan(&m.Login, &m.Password, &m.Admin)
	if e != nil { 
		fmt.Println(e.Error())
		return errors.New("Invalid login or password")
	}
	return nil
}

func (m *News) Select() error {
	stmt, ok := Queries["Select#News"]
	if !ok {
		return errors.New("News query doesn't exist")
	}
	
	rows, e := stmt.Query()
	if e != nil{
		return e
	}

	for rows.Next() {
		e = rows.Scan(&m.Title, &m.TextContent, &m.PostDate, &m.Image)
		if e != nil {	
			return e
		}
		if m.Image == ""{
			m.Image = "replacement.png"
		}
		
		m.Rows = append(m.Rows, News{Title: m.Title, TextContent: m.TextContent, PostDate: m.PostDate, Image: m.Image})

	}
	return nil
}

func (m *News) Add() error {
	stmt, ok := Queries["Insert#News"]
	if !ok {
		return errors.New("News query doesn't exist")
	}
	r := stmt.QueryRow(m.Title, m.TextContent, m.PostDate, m.Image)
	e := r.Scan(&m.Title, &m.TextContent, &m.PostDate, &m.Image)
	fmt.Println("After querying object state is= ", m)
	if e != nil { 
		fmt.Println(e.Error())
		return errors.New("something went wrong while adding an Article")
	}

	return nil
}

func (m *News) Delete() error{
	stmt, ok := Queries["Delete#News"]
	if !ok {
		return errors.New("News query doesn't exist")
	}
	var r *sql.Row
	stmt.QueryRow(m.Title).Scan(r)
	fmt.Println(r)
	return nil
}