package handlers

import (
	"database/sql"
	_ "github.com/Xoro-1337/AuthMe/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Failed to bind JSON data",
			})
			return
		}

		var existingUser User
		stmt := "SELECT * FROM users WHERE uid = ?;"
		err := db.QueryRow(stmt, user.UID).Scan(&existingUser.UID)
		if err == sql.ErrNoRows {
			stmt = "INSERT INTO users (uid) VALUES (?);"
			_, err := db.Exec(stmt, user.UID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "error",
					"error":  "Failed to create new user",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"user": gin.H{
					"uid": user.UID,
				},
			})
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "Failed to query database",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "error",
				"error":  "The UID already exists",
			})
		}
	}
}
