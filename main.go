package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)


func main() {
	e := connect()
	if e != nil{
		panic(e.Error())
	}

	//addContent("Genre", "Name", "Триллер")
	prepareQueries()
	router := gin.Default()
	
	router.Use(static.Serve("/", static.LocalFile(cfg.Assets, false)))

	// создали куки с очень сложным приватным ключом
	store := sessions.NewCookieStore([]byte("TheMostSecretWord"))
	// Эту куку записываем в сам рутер с ключом сессии
	router.Use(sessions.Sessions("session", store))

	router.LoadHTMLGlob(cfg.HTML + "*.html")

	router.Static("assets", cfg.Assets)

	router.GET("/", index)
	router.GET("/news", news)
	router.POST("/login", login)
	router.POST("/logout", logout)
	router.POST("/news", addNews)
	router.DELETE("/news", deleteNews)
	router.Run(cfg.ServerHost + ":" + cfg.ServerPort)

}

func index(c *gin.Context) {
	s := sessions.Default(c)

	admin := false 
	login := false

	role := s.Get("MySecretKey")

	if role == true {
		admin = true
		login = true
	}	else if role == false {
		login = true
	}
	
	

	c.HTML(200, "index.html", gin.H{
			"Admin": admin,
			"isLogin": login,
	})
}

func news(c * gin.Context) {
	
	s := sessions.Default(c)
	admin := false 
	login := false

	role := s.Get("MySecretKey")

	if role == true {
		admin = true
		login = true
	}	else if role == false {
		login = true
	}

	var cat Genre
	e := cat.Select()
	if e != nil{
		fmt.Println(e.Error())
	}
	var m News
	e = m.Select()
	if e != nil {
		fmt.Println(e.Error())
	}
	c.HTML(200, "news.html", gin.H {
		"Admin": admin,
		"isLogin": login,
		"Genre": cat.Rows,
		"News": m.Rows,
	})
}

func addNews(c *gin.Context){
	var m News
	e := c.BindJSON(&m)
	//changing date to actual state since it was easier to make it on server side rather then on client side
	date := time.Now()
	YYYYMMDD := "2006-01-02"
	date.Format(YYYYMMDD)
	m.PostDate = date
	
	if e != nil {	
		fmt.Println(e.Error())
		c.JSON(400, gin.H{"Error": e.Error()})
	}
	// Inserting data to postgres
	e = m.Add() 
	if e != nil {
		fmt.Println(e.Error())
		c.JSON(400, gin.H{"Error": e.Error()})
	}
}

func deleteNews(c *gin.Context){
	var m News
	e := c.BindJSON(&m)
	fmt.Println("ToDeleteData=", m.Title)
	if e != nil {	
		fmt.Println(e.Error())
		c.JSON(400, gin.H{"Error": e.Error()})
	}

	e = m.Delete() 
	if e != nil {
		fmt.Println(e.Error())
		c.JSON(400, gin.H{"Error": e.Error()})
	}
}

func login(c *gin.Context) {

	var m User
	e := c.BindJSON(&m)
	if e != nil {	
		fmt.Println(e.Error())
		c.JSON(200, gin.H{"Error":e.Error()})
	}

	e = m.Select()
	if e != nil {
		fmt.Println(e.Error())

		c.JSON(400, gin.H {
			"Error": e.Error(),
		})
		return
	}

	// Инициализируем сессию
	s := sessions.Default(c)
	// Устанавливаем супер секретный ключ
	s.Set("MySecretKey", m.Admin)
	// Сохраняем сессию в куках
	e = s.Save()
	fmt.Println("session=", s)
	
	if e != nil{
		fmt.Println("SOMETHING WRONG WITH THE SESSION: ", e.Error())
	}

	c.JSON(200, gin.H{
		"Error": nil,
		"Login": m.Login,
	})
}

func logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
}