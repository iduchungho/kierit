package middleware

import (
	"context"
	"example.com/m/database"
	"example.com/m/modules/entities/user"
	"example.com/m/modules/session"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
)

func RequiredUser(c *fiber.Ctx) error {
	// Get the cookie off req
	sess, err := session.GetStore().Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	id := c.Query("id", "none")
	if id == "none" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?id = ...",
			"success": false,
		})
	}
	tkStr := sess.Get("Authorization")
	if tkStr == nil || tkStr == -1 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"data":    "no token in session",
			"success": false,
		})
	}

	tokenString := tkStr.(string)

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		// find the user with token sub
		filter := bson.D{{"id", claims["sub"]}}
		var db = new(database.Database)
		collection := db.GetConn().Database("MessengerDB").Collection("users")

		var users = new(user.User)
		errFind := collection.FindOne(context.TODO(), filter).Decode(users)
		if errFind != nil {
			// c.AbortWithStatus(http.StatusUnauthorized)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "not authorized",
				"success": false,
			})
		}

		if id != claims["sub"] {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "not authorized",
				"success": false,
				"id":      id,
			})
		}
		// attach to req
		// cookie := new(fiber.Cookie)
		// cookie.Name = "john"
		// cookie.Value =
		// cookie.Expires = time.Now().Add(24 * time.Hour)
		// c.Cookie("user", user)

		// continue
		errNext := c.Next()
		if errNext != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"data":    "no token in session",
			"success": false,
		})
	}
	return nil
}
