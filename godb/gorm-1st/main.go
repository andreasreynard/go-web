package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Title  string
	Artist string
	Price  float32
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:250204@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Album{})

	// Create
	result := db.Create(&Album{Title: "Semester 5 Bikin Mati Rasa", Artist: "Andrey Sam", Price: 99.99})
	if result.Error != nil {
		panic(result.Error) // Handle error
	}

	// Check inserted data
	rowsAffected := result.RowsAffected // returns inserted records count
	println("Rows affected:", rowsAffected)

	// Read
	var album Album
	// db.First(&album, 1)                                         // find album with integer primary key
	db.First(&album, "title = ?", "Semester 5 Bikin Mati Rasa") // find album with title Semester 5 Bikin Mati Rasa

	// Update - update product's price to 199.99
	db.Model(&album).Update("Price", 199.99)
	// Update - update multiple fields
	db.Model(&album).Updates(Album{Price: 199.99, Title: "Semester 5 ITS Bikin Mati Rasa"}) // non-zero fields
	db.Model(&album).Updates(map[string]interface{}{"Price": 199.99, "Title": "Semester 5 ITS Bikin Mati Rasa"})

	// Delete - delete product
	db.Delete(&album, 1)
}
