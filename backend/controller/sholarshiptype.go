package controller

import (
	"net/http"

	"github.com/Ksis123/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /types
func CreateScholarshipsType(c *gin.Context) {
	var scholarshipstype entity.Scholarships_type
	if err := c.ShouldBindJSON(&scholarshipstype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&scholarshipstype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": scholarshipstype})
}

// GET /types
func ListMyScholarshipsType(c *gin.Context) {
	var scholarshipstype []entity.Scholarships_type
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM scholarships_types").Find(&scholarshipstype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholarshipstype})
}

// GET /type/:id
func GetScholarshipsType(c *gin.Context) {
	var scholarshipstype entity.Scholarships_type
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM scholarships_types WHERE id = ?", id).Find(&scholarshipstype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if scholarshipstype.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholarshipstype})
}

// DELETE /types/:id
func DeleteScholarshipsType(c *gin.Context) {
	TypeID := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM scholarships_types WHERE id = ?", TypeID); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": TypeID})
}

// PATCH /types
func UpdateScholarshipsType(c *gin.Context) {
	var scholarshipstype entity.Scholarships_type
	if err := c.ShouldBindJSON(&scholarshipstype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", scholarshipstype.ID).First(&scholarshipstype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "scholarships_types not found"})
		return
	}

	if err := entity.DB().Save(&scholarshipstype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholarshipstype})
}