package main

import (
	"github.com/gin-gonic/gin"
)

type Trivia struct {
	Nombre string
	Titulo string
	Preguntas []Pregunta
}

type Pregunta struct {
	Pregunta string
	Respuestas map[string]string
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./client/templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", home)
	r.POST("/trivia", triviaClient)
	r.Run(":24361")
}

func home(c *gin.Context){
	c.HTML(200, "home.html", nil)
}

func triviaClient(c *gin.Context){
	c.Request.ParseForm()
	name := c.PostForm("name")
	preg1 := Pregunta{
		Pregunta: "¿Cuando se creó la UMA?",
		Respuestas: map[string]string{
			"1": "1990",
			"2": "1980",
			"3": "1970",
			"4": "1960",
		},
	}
	pregs := []Pregunta{preg1}
	obj := Trivia{
		Nombre: name,
		Titulo: "Trivia UMA",
		Preguntas: pregs,
	}
	c.HTML(200, "trivia.html", obj)
}