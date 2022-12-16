package abixjwt

import (
	"abix360/shared"
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v2"
)

type configApp struct {
	Jwt struct {
		Secret string `yaml:"secret"`
	}
}

type ResponseLogin struct {
	Id    int64
	Email string
	Token string
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println("shared / HashAndSalt / bcrypt.GenerateFromPassword: " + err.Error())
	}
	return string(hash)
}

func ComparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println(err)
		log.Println("shared / ComparePassword / bcrypt.CompareHashAndPassword: " + err.Error())
		return false
	}
	return true
}

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(getKeySecret())

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Println("shared / GenerateJWT / token.SignedString: ", err.Error())
		return "", err
	}
	return tokenString, nil
}

func CheckPasswordHash(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}

func isValidToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])
		}
		return []byte(getKeySecret()), nil
	})
}

func getKeySecret() string {
	content, err := os.ReadFile(shared.GetRootPath() + "/app/cfg-app.yml")
	if err != nil {
		log.Fatal("abix-jwt / getKeySecret() / os.ReadFile: ", err)
	}

	var config configApp
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("abix-jwt / getKeySecret() / yaml.Unmarshal: ", err)
	}

	return config.Jwt.Secret
}

func VerifyToken(tokenString string) bool {
	token, err := isValidToken(tokenString)
	if err != nil {
		log.Println("abix-jwt / VerifyToken() / isValidToken: ", err)
		return false
	}
	return token.Valid
}

func GetTokenRequest(c *gin.Context) string {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		return ""
	}
	tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])
	return tokenString
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetTokenRequest(c)
		if !VerifyToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token no válido"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var repository domain.UserRepository = mysql.NewUserDao()
		user := domain.FindUserByToken(token, repository)
		if !user.Exists() {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token no válido"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}
