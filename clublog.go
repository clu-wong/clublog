package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"clublog/lib/models"
	"io"
	"html/template"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
	db, err := gorm.Open("mysql", "root:12345678@/goblogdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		panic("db error: "+err.Error())
	}
	db.AutoMigrate(&models.User{})
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	// Echo instance
	e := echo.New()
	// render template
	e.Renderer = t

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/new_sessions", new_session)
	e.Static("/assets", "assets")


	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(200, "hello")
	//return c.Render(http.StatusOK, "hello.html", "World")
}
func new_session(c echo.Context) error{
	return c.Render(http.StatusOK,"login.html", "World")
}
