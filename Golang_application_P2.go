package main
/
import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rahul10-pu/CIGO0122/database"
	"github.com/rahul10-pu/CIGO0122/handler"
	"github.com/rahul10-pu/CIGO0122/routers"
)

func init() { 
	database.Setup() 
	database.Example()
	database.exam()
}
func init() { 
	fmt.Print(1)
}
func init() { 
	fmt.Print(2)
}
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func AuthMiddleware() gin.HandlerFunc {
	
	Foo()
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL)
		token := c.GetHeader("token")
		fmt.Println("got token:	" + token)
		isValid, err := handler.ValidateToken(token)
		if err != nil && !isValid {
			respondWithError(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}
func main() {
	engine := gin.Default() 
	api := handler.Handler{
		DB: database.GetDB(), 
	} 
	engine.Use(AuthMiddleware())
	routers.BookRouter(engine, api)
	routers.AuthRouter(engine, api)

	err := engine.Run("127.0.0.1:8080") 
	if err != nil {
		log.Fatal(err)
	}
}
*/