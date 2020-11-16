package main


import (

	"net/http"
	
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id         int
    Name       string
    Occupation string
}

func main() {
	e := echo.New()

	e.GET("/echo",getData)
		//return c.String(http.StatusOK, "Hello, World!)
	e.Logger.Fatal(e.Start(":8000"))
}



func getData (c echo.Context) (err error) {

	dsn := "user=postgres password=admin DB.name=postgres port=5432 host=localhost sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})
	

	users := []User{
        {Id: 2, Name: "Isaac", Occupation: "driver"},
        {Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
        {Id: 4, Name: "David Brown", Occupation: "programmer"},
	}
	

	result := db.Create(users) // pass pointer of data to Create
	//return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, result.RowsAffected)
  }