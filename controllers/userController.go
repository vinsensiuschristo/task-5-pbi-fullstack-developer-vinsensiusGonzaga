package controllers

import (
	"net/http"
	"os"
	"time"

	"example.com/practice/initializers"
	"example.com/practice/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})

		return
	}

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
	})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to create token",
		})

		return

	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func PutUser(c *gin.Context) {
	idParam := c.Param("id")

	var updatePhoto struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&updatePhoto) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read input",
		})

		return
	}

	// tambahin if else dulu

	result := initializers.DB.Model(&models.User{}).Where("id = ?", idParam).Updates(&models.User{
		Username: updatePhoto.Username,
		Email:    updatePhoto.Email,
		Password: updatePhoto.Password,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
	})
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	// cek di table photo, photoId nya udah ada softDelete blom
	test := initializers.DB.Unscoped().Where("id = " + idParam).Find(&models.User{})

	if test != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak bisa dihapus karena data tidak ada",
		})

		return
	}

	println(test)

	initializers.DB.Delete(&models.User{}, idParam)

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}
