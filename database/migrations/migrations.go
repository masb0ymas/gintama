package migrations

import (
	"fmt"
	"strings"

	"gintama/database/entity"

	"gorm.io/gorm"
)

func Initial(db *gorm.DB) {
	// collect all migration
	collectMigrations := []string{
		BaseMigrations(),
	}

	schema := strings.Join(collectMigrations, ` `)
	result := db.Exec(schema)

	// auto migrate from entity
	db.AutoMigrate(&entity.Product{}, &entity.Order{}, &entity.OrderItem{})

	fmt.Println("GORM", "Migration Successfully", result)
}
