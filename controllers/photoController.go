package controllers

import (
	"fmt"
	"net/http"

	"example.com/practice/initializers"
	"example.com/practice/models"
	"github.com/gin-gonic/gin"
)

/**

catatan untuk photos adalah, di beberapa validasi ketika melakukan aksi masih terdapat error
DELETE =  Belum bisa error jika menemukan / tidak menemukan data sebelum delete

ALL belum menambahkan validasi hanya user yang dapat menghapus photonya sendiri

*/

func PostPhoto(c *gin.Context) {
	var inputPhoto struct {
		Title    string
		Caption  string
		PhotoUrl string
	}

	if c.Bind(&inputPhoto) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read input",
		})

		return
	}

	user, _ := c.Get("user")

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User Not Identified",
		})

		return
	}

	userID := user.(models.User).ID

	userPost := models.Photo{
		Title:    inputPhoto.Title,
		Caption:  inputPhoto.Caption,
		PhotoUrl: inputPhoto.PhotoUrl,
		UserId:   userID,
	}

	result := initializers.DB.Create(&userPost)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "photo created successfully",
	})
}

func GetPhoto(c *gin.Context) {
	user, _ := c.Get("user")

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User Not Identified",
		})

		return
	}

	userID := user.(models.User).ID

	fmt.Println(userID)

	result := initializers.DB.Where("user_id = ?", userID).Find(&models.Photo{})

	fmt.Println(result)

	if result == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data Not Found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

func PutPhoto(c *gin.Context) {
	idParam := c.Param("id")

	var updatePhoto struct {
		Title    string
		Caption  string
		PhotoUrl string
	}

	if c.Bind(&updatePhoto) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read input",
		})

		return
	}

	// tambahin if else dulu

	result := initializers.DB.Model(&models.Photo{}).Where("id = ?", idParam).Updates(&models.Photo{
		Title:    updatePhoto.Title,
		Caption:  updatePhoto.Caption,
		PhotoUrl: updatePhoto.PhotoUrl,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to update photo",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "photo updated successfully",
	})
}

func DeletePhoto(c *gin.Context) {
	idParam := c.Param("id")

	// cek di table photo, photoId nya udah ada softDelete blom
	test := initializers.DB.Unscoped().Where("id = " + idParam).Find(&models.Photo{})

	if test != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak bisa dihapus karena data tidak ada",
		})

		return
	}

	println(test)

	initializers.DB.Delete(&models.Photo{}, idParam)

	c.JSON(http.StatusOK, gin.H{
		"message": "photo deleted successfully",
	})
}
