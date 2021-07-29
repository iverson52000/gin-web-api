package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  int
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Restaurant struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Type      string
	Rating    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("sqliteDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Restaurant{})

	// Create
	// db.Create(&Restaurant{Name: "Chiptole", Type: "Maxican", Rating: 5})

	// // Read
	// var restaurant Restaurant
	// db.First(&restaurant, 1)                 // find restaurant with integer primary key
	// db.First(&restaurant, "code = ?", "D42") // find restaurant with code D42

	// // Update - update restaurant's price to 200
	// db.Model(&restaurant).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&restaurant).Updates(Restaurant{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&restaurant).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete restaurant
	// db.Delete(&restaurant, 1)
}
