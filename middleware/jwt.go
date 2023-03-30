package middleware

import (
	"fmt"
	"game-api/utils/structs"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(Users structs.Users) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(200 * time.Minute)
	claims["authorized"] = true
	claims["user"] = Users.Email
	claims["role"] = Users.Roles
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}

func VerifyJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("jalan")
		if ctx.Request.Header["Token"] != nil {
			token, err := jwt.Parse(ctx.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					ctx.Writer.WriteHeader(http.StatusUnauthorized)
					_, err := ctx.Writer.Write([]byte("You're Unauthorized"))
					if err != nil {
						return nil, err
					}
				}
				return "", nil
			})
			if err != nil {
				ctx.Writer.WriteHeader(http.StatusUnauthorized)
				_, err2 := ctx.Writer.Write([]byte("You're Unauthorized due to error parsing the JWT"))
				ctx.Abort()
				if err2 != nil {
					return
				}

			}
			if token.Valid {
				_, err2 := ctx.Writer.Write([]byte("token Valid"))
				if err2 != nil {
					return
				}
			} else {
				ctx.Writer.WriteHeader(http.StatusUnauthorized)
				_, err := ctx.Writer.Write([]byte("You're Unauthorized due to invalid token"))
				if err != nil {
					return
				}
			}
		} else {
			ctx.Writer.WriteHeader(http.StatusUnauthorized)
			_, err := ctx.Writer.Write([]byte("You're Unauthorized due to No token in the header"))
			if err != nil {
				return
			}
		}
	}
}

func ExtractClaims(ctx *gin.Context) (string, error) {
	if ctx.Request.Header["Token"] != nil {
		tokenString := ctx.Request.Header["Token"][0]
		fmt.Println(tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there's an error with the signing method")
			}
			return SecretKey, nil
		})
		if err != nil {
			return "Error Parsing Token: ", err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			user := claims["user"].(string)
			return user, nil
		}
	}

	return "unable to extract claims", nil
}

// func authPage() {
// 	token, _ := GenerateJWT()
// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
// 	req.Header.Set("Token", token)
// 	_, _ = client.Do(req)
// }

// func main() {
// 	http.HandleFunc("/home", verifyJWT(handlePage))
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		log.Println("There was an error listening on port :8080", err)
// 	}

// }
