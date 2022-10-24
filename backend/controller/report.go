package controller

import (
	"net/http"

	"github.com/Ksis123/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /Report
func CreateReport(c *gin.Context) {
	var report entity.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&report).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": report})
}

// GET /report/:id
func GetReport(c *gin.Context) {
	var report entity.Report

	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&report); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "report not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": report})
}

// GET /reports
func ListReports(c *gin.Context) {
	var reports []entity.Report
	if err := entity.DB().Raw("SELECT * FROM reports").Scan(&reports).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reports})
}
// DELETE /reports/:id

func DeleteReport(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM reports WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "report not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /reports
func UpdateReport(c *gin.Context) {
	var report entity.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", report.ID).First(&report); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "report not found"})
		return
	}

	if err := entity.DB().Save(&report).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": report})
}		

