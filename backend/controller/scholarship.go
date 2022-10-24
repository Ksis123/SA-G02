package controller

import (
	"net/http"

	"github.com/Ksis123/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /
func CreateScholarship(c *gin.Context) {

	var scholarship entity.Scholarship
	var scholarshipsstatus entity.Scholarships_status
	var scholarshipstype entity.Scholarships_type


	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Scholarship
	if err := c.ShouldBindJSON(&scholarship); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", scholarship.StatusID).First(&scholarshipsstatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	// 10: ค้นหา type ด้วย id
	if tx := entity.DB().Where("id = ?", scholarship.TypeID).First(&scholarshipstype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type not found"})
		return
	}

	// 12: สร้าง Scholarship
	wv := entity.Scholarship{

		Scholarship_Name:   scholarship.Scholarship_Name,
		StatusID:           scholarship.StatusID,
		TypeID:             scholarship.TypeID,
		Scholarship_detail: scholarship.Scholarship_detail,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": wv})
}

// GET /Scholarship/:id
func GetScholarship(c *gin.Context) {
	var scholarship entity.Scholarship
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&scholarship); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Scholarship not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": scholarship})
}

// GET /Scholarship
func ListScholarships(c *gin.Context) {
	var scholarships []entity.Scholarship
	if err := entity.DB().Preload("thetype").Preload("thestatus").Preload("theAdmin").Raw("SELECT * FROM scholarships").Find(&scholarships).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholarships})
}

// DELETE /Scholarship/:id
func DeleteScholarship(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM scholarships WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Scholarship not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Scholarship
func UpdateScholarship(c *gin.Context) {
	var scholarship entity.Scholarship
	if err := c.ShouldBindJSON(&scholarship); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", scholarship.ID).First(&scholarship); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Scholarship not found"})
		return
	}

	if err := entity.DB().Save(&scholarship).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": scholarship})
}