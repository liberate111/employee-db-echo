package main

import (
	"emp/config"
	"emp/employees"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Template can use any template engine
type Template struct {
	templates *template.Template
}

// Render implement echo.Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo instance
	e := echo.New()

	// Register templates
	t := &Template{
		templates: config.TPL,
	}
	e.Renderer = t

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1323"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Setup route
	setupRoute(e)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func setupRoute(e *echo.Echo) {
	e.GET("/", index)

	emp := e.Group("/emps")
	emp.GET("", employees.Index)
	emp.GET("/show", employees.Show)
	emp.GET("/create", employees.Create)
	emp.POST("/create/process", employees.CreateProcess)
	emp.GET("/update", employees.Update)
	emp.POST("/update/process", employees.UpdateProcess)
	emp.GET("/delete/process", employees.DeleteProcess)
}

func index(c echo.Context) error {
	return c.Redirect(303, "/emps")
}
