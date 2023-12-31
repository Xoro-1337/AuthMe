package main

import (
	"database/sql"
	"github.com/Xoro-1337/AuthMe/server/api/handlers"
	"github.com/Xoro-1337/AuthMe/server/api/middleware"
	"github.com/Xoro-1337/AuthMe/server/api/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConfig := utils.GetDBConfigFromEnv()
	db, err := utils.ConnectToDB(dbConfig)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	router := gin.Default()

	router.Use(middleware.CORS())
	// gin.SetMode(gin.ReleaseMode)

	// API routes
	router.POST("/login", handlers.Authenticate(db))
	router.GET("/download", handlers.DownloadHandler)
	router.POST("/register", handlers.RegisterHandler(db))

	err = router.Run(":8080")
	if err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
