package recipies

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MrKren/quantchef/src/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setup(t *testing.T) (*gin.Engine, *gorm.DB) {
	test_DB_url := "root:root@tcp(127.0.0.1:3307)/quantchef?parseTime=true"
	db, err := gorm.Open(mysql.Open(test_DB_url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("An error '%s' occured when opening test DB", err)
	}

	err = db.AutoMigrate(
		&models.Recipie{},
		&models.Ingredient{},
		&models.Step{},
		&models.Dish{},
	)
	if err != nil {
		t.Fatalf("An error while auto-migrating: '%s'", err)
	}

	router := gin.Default()
	RegisterRoutes(router, db)

	return router, db
}

func teardown(db *gorm.DB) {
	var tables []string
	table_query := `
	SELECT Concat('DROP TABLE ', TABLE_NAME) 
	FROM INFORMATION_SCHEMA.TABLES 
	WHERE table_schema = 'quantchef' order by Create_time desc;
	`
	db.Raw(table_query).Scan(&tables)
	for i := 0; i < len(tables); i++ {
		db.Exec(tables[i])
	}
}

func TestGetRecipies(t *testing.T) {
	r, db := setup(t)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/recipies", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	teardown(db)
}
