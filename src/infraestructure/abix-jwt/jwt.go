package abixjwt

import (
	"abix360/shared"
	"fmt"
	"log"
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

func VerifyToken(c *gin.Context) bool {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])

	token, err := isValidToken(tokenString)
	if err != nil {
		log.Fatal("abix-jwt / VerifyToken() / isValidToken: ", err)
	}
	return token.Valid
}
