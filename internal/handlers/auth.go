package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type User struct {
    ID       uint   `gorm:"primarykey"`
    Email    string `gorm:"uniqueIndex"`
    Password string
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB, jwtSecret string) {
    if err := db.AutoMigrate(&User{}); err != nil {
        panic(err)
    }

    r.GET("/health", func(c *gin.Context) {
        sqlDB, err := db.DB()
        if err != nil || sqlDB.Ping() != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "db error"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    r.POST("/auth/register", func(c *gin.Context) {
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        user := User{Email: req.Email, Password: string(hashed)}
        if err := db.Create(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "registered"})
    })

    r.POST("/auth/login", func(c *gin.Context) {
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        var user User
        if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
            return
        }
        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
            return
        }
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "sub": user.ID,
            "exp": time.Now().Add(time.Hour).Unix(),
        })
        signed, err := token.SignedString([]byte(jwtSecret))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"access_token": signed})
    })
}

