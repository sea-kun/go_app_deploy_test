package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()
	e.File("/", "template/index.html")
	e.Logger.Fatal(e.Start(":8000"))
}
