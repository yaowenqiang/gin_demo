package main
import (
	"time"
	"log"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

var loginCookies = map[string]*loginCookie{}
var identities =  []identity{
	{
		employeeNumber: "1234",
		password: "password",
	},
}

const loginCookieName = "Identity"


type identity  struct {
	employeeNumber string
	password string

}

type loginCookie struct {
	value string
	expiration time.Time
	origin  string
}


func loginMiddleware(c *gin.Context) {
	if strings.HasPrefix(c.Request.URL.Path, "/login") ||
	strings.HasPrefix(c.Request.URL.Path, "/public") {
		log.Print("url matched")
		return
	}

	cookieValue, err := c.Cookie(loginCookieName)
	log.Printf("cookieValue: %v", cookieValue)

	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	cookie, ok := loginCookies[cookieValue]
	log.Printf("cookie: %v", cookie)

	if !ok  || 
		cookie.expiration.Unix() < time.Now().Unix() ||
		cookie.origin != c.Request.RemoteAddr { 
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.Next()
}
