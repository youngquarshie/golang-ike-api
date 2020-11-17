package main


import (

	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement:true;not_null"`
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
        {Name: "Isaac", Occupation: "driver"},
        {Name: "Lucy Smith", Occupation: "teacher"},
        {Name: "David Brown", Occupation: "programmer"},
	}
	
	result := db.Create(users) // pass pointer of data to Create
	//return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, result.Error )

  }