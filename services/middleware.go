package services

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	g "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CheckPasswordHashing(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(user models.User) models.AuthSession {
	db := configs.Store
	var obj models.AuthSession
	var profile models.Profile
	db.
		Preload("User").
		Preload("Area").
		Preload("Whs").
		Preload("Factory").
		Preload("Position").
		Preload("Department").
		Preload("PrefixName").
		Where("user_id=?", &user.ID).
		First(&profile)
		// As variable
	profile.AvatarURL = "emp/user01.png"
	obj.Profile = &profile
	obj.User = &user
	obj.IsAdmin = false
	var admin models.Administrator
	db.First(&admin, "user_id=?", &user.ID)
	if admin.ID != "" {
		obj.IsAdmin = true
	}
	obj.Header = "Authorization"
	obj.JwtType = "Bearer"
	obj.JwtToken, _ = g.New(60)
	secret_key := os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = obj.JwtToken
	claims["name"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenKey, err := token.SignedString([]byte(secret_key))
	if err != nil {
		panic(err)
	}

	/// Insert Token Key to DB
	t := new(models.JwtToken)
	t.ID = obj.JwtToken
	t.UserID = &user.ID
	t.Token = tokenKey
	// Delete UserID before creating TokenID
	err = db.Where("user_id=?", t.UserID).Delete(&models.JwtToken{}).Error
	if err != nil {
		panic(err)
	}

	err = db.Create(&t).Error
	if err != nil {
		panic(err)
	}
	return obj
}

func ValidateToken(tokenKey string) (interface{}, error) {
	token, err := jwt.Parse(tokenKey, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}
	return claims["name"], nil
}

func AuthorizationRequired(c *fiber.Ctx) error {
	var r models.Response
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	if token == "" {
		r.Message = MessageNotFoundTokenKey
		return c.Status(fiber.StatusUnauthorized).JSON(&r)
	}

	// Check Token On DB
	db := configs.Store
	var jwtToken models.JwtToken
	err := db.Where("id=?", token).Find(&jwtToken).Error
	if err != nil {
		r.Message = MessageSystemError
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if jwtToken.ID == "" {
		r.Message = MessageUserNotAuthenticated
		return c.Status(fiber.StatusUnauthorized).JSON(&r)
	}

	_, er := ValidateToken(jwtToken.Token)
	if er != nil {
		r.Message = MessageTokenIsExpired
		db.Delete(&jwtToken)
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	return c.Next()
}
