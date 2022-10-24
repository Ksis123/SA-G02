package controller

import (
	"net/http"

	"github.com/Ksis123/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

func CreateScholarshipsStatus(c *gin.Context) {
	var scholarshipsStatus entity.Scholarships_status
	if err := c.ShouldBindJSON(&scholarshipsStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&scholarshipsStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": scholarshipsStatus})
}

// GET /stat
func ListMyscholarshipsStatus(c *gin.Context) {
	var scholarships_statuses []entity.Scholarships_status
	if err := entity.DB().Raw("SELECT * FROM scholarships_statuses").Scan(&scholarships_statuses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholarships_statuses})
}

// GET /paymentstatus/:id
func GetScholarshipsStatus(c *gin.Context) {
	var scholarshipsStatus entity.Scholarships_status
	StatusID := c.Param("id")
	if tx := entity.DB().Where("id = ?", StatusID).First(&scholarshipsStatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "scholarships_statuses not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholarshipsStatus})
}

// DELETE /stat/:id
func DeleteScholarshipsStatus(c *gin.Context) {
	StatusID := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM scholarships_statuses WHERE id = ?", StatusID); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": StatusID})
}

// PATCH /stat
func UpdateScholarshipsStatus(c *gin.Context) {
	var scholarshipsStatus entity.Scholarships_status
	if err := c.ShouldBindJSON(&scholarshipsStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", scholarshipsStatus.ID).First(&scholarshipsStatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stat not found"})
		return
	}

	if err := entity.DB().Save(&scholarshipsStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholarshipsStatus})
}