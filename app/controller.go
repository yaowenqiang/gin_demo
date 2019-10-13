package main

import (
	"net/http"
	"time"
	"log"
	"github.com/gin-gonic/gin"
)


func RegisterRoutes() *gin.Engine {

	r := gin.Default()
	r.Use(loginMiddleware)
	//gin.New
	r.LoadHTMLGlob("templates/**/*.html")

	r.Any("/", func (c *gin.Context) {
		//c.String(http.StatusOK, "Hello from %v", "Gin")
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/employees/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, id, nil)
	})


	r.GET("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {
			c.String(http.StatusOK, "add", nil)
			return
		}
		employee, ok := employees[id]
		if !ok {
			c.String(http.StatusNotFound, "404 - Not Found")
		}
		//c.String(http.StatusOK, id, nil)
		//c.HTML(http.StatusOK, "admin-employee-edit.html", map[string]interface{}{
		c.HTML(http.StatusOK, "admin-employee-edit.html", gin.H{
			"Employee": employee,
		})
	})

	/*
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	*/

	r.Any("/login", func(c *gin.Context) {
		employeeNumber := c.PostForm("employeeNumber")
		password  := c.PostForm("password")
		log.Printf("identities: %v", identities)

		for _, identity := range identities {
			if identity.employeeNumber == employeeNumber  && 
			identity.password == password {
				log.Print("employ matched")
				lc := loginCookie{
					value: employeeNumber,
					expiration: time.Now().Add(24 * time.Hour),
					origin: c.Request.RemoteAddr,
				}

				loginCookies[lc.value] = &lc
				maxAge := lc.expiration.Unix() - time.Now().Unix()
				c.SetCookie(loginCookieName,  lc.value, int(maxAge), "", "", false, true, )
				log.Print("redirect to /")
				c.Redirect(http.StatusTemporaryRedirect, "/")
				return
			} else {
				log.Printf("identity = %v:", identity)
				log.Printf("employeeNumber = %v:", employeeNumber)
				log.Printf("password = %v:", password)
				c.String(http.StatusBadRequest, "login failed", nil)
				return
			}
		}
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/employees/:id/vacation/add", func(c *gin.Context) {
		var timeOff TimeOff
		err := c.BindJSON(&timeOff)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		id := c.Param("id")
		timesOff, ok  := TimesOff[id]
		if !ok {
			TimesOff[id] = []TimeOff{}
		}

		TimesOff[id] = append(timesOff, timeOff)

		c.JSON(http.StatusCreated, &timeOff)
	})

	r.POST("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {

			/*
			startDate, err := time.Parse("2006-01-02", c.PostForm("startDate"))
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			log.Print(startDate)
			*/

			var emp Employee
			err := c.Bind(&emp)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			emp.ID = 42
			emp.Status = "Active"
			//emp.StartDate = startDate
			employees["42"] = emp
			
			c.Redirect(http.StatusPermanentRedirect, "/employees/42")
			return



		}
		c.String(http.StatusOK, id, nil)
	})

	/*
	this route won't work, conflic with /employees/:id
	r.GET("/employees/add", func(c *gin.Context) {
		c.String(http.StatusOK, "add", nil)
	})
	*/

	//http --auth admin:admin --auth-type basic "localhost:3001/admin/"
	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})


	//r.Static("/static", "./static")
	r.StaticFS("/static", http.Dir("static"))

	//r.StaticFS("/publc",http.Dir("./public"))
	r.Run(":3001")
	return r
}

