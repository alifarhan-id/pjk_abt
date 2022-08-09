package module

import (
	"net/http"

	models "github.com/alifarhan1230/pjk_abt/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindOwner(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var owner []models.Owner
	db.Raw("SELECT * FROM owner").Scan(&owner)
	c.JSON(http.StatusOK, gin.H{"data": owner})
}
