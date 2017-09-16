package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"fmt"
	"os"
	"strconv"
	"time"
)

const defaultMaxConns = 20

var (
	db *gorm.DB
)

type User struct {
	gorm.Model

	Name string
	Mail string
}

type UserPost struct {
	Name string `form:"name" json:"name" binding:"required"`
	Mail string `form:"email" json:"email" binding:"required"`
}

func main() {
	r := gin.Default()

	init_db()

	r.GET("/", hello_world)
	r.GET("/users", list_users)
	r.GET("/users/:id", show_user)
	r.POST("/users", add_user)
	r.PUT("/users/:id", mod_user)
	r.DELETE("/users/:id", del_user)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func init_db() {
	// get database url from env
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		panic("Environment variable 'DATABASE_URL' not defined.")
	}

	// connect to postgres
	var err error
	for trial := 0; ; trial++ {
		db, err = gorm.Open("postgres", url)
		if err == nil {
			break
		}
		fmt.Fprintf(os.Stderr, "Failed to connect database: %s.\n", err)
		if trial >= 5 {
			panic("")
		}
		time.Sleep(3 * time.Second)
	}

	// get max connections from env
	max := defaultMaxConns
	env := os.Getenv("DATABASE_MAX_CONNS")
	if env != "" {
		tmp, err := strconv.Atoi(env)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid 'DATABASE_MAX_CONNS'. Ignored.")
		} else {
			max = tmp
		}
	}

	// set max connection
	db.DB().SetMaxIdleConns(max / 5)
	db.DB().SetMaxOpenConns(max)

	// migrate
	if err := db.AutoMigrate(&User{}).Error; err != nil {
		panic(err)
	}
}

func hello_world(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!",
	})
}

func on_error(c *gin.Context, err error, code int) {
	c.String(code, http.StatusText(code))
	c.AbortWithError(code, err)
}

func list_users(c *gin.Context) {
	var user []User
	if err := db.Find(&user).Error; err != nil {
		on_error(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func show_user(c *gin.Context) {
	id := c.Param("id")

	var user User
	if err := db.First(&user, id).Error; err != nil {
		on_error(c, err, http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)
}

func add_user(c *gin.Context) {
	var json UserPost
	if err := c.BindJSON(&json); err != nil {
		on_error(c, err, http.StatusBadRequest)
		return
	}

	user := User {
		Name: json.Name,
		Mail: json.Mail,
	}
	if err := db.Create(&user).Error; err != nil {
		on_error(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func mod_user(c *gin.Context) {
	id := c.Param("id")

	// get current entry
	var user User
	if err := db.First(&user, id).Error; err != nil {
		on_error(c, err, http.StatusNotFound)
		return
	}

	// get new info
	var json UserPost
	if err := c.BindJSON(&json); err != nil {
		on_error(c, err, http.StatusBadRequest)
		return
	}

	// save new info
	user.Name = json.Name
	user.Mail = json.Mail
	if err := db.Save(&user).Error; err != nil {
		on_error(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func del_user(c *gin.Context) {
	id := c.Param("id")

	var user User
	if err := db.First(&user, id).Error; err != nil {
		on_error(c, err, http.StatusNotFound)
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		on_error(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusNoContent, user)
}
