package main

import (
	"database/sql"
	"net/http"
	"path"
	"time"

	"./controller"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("pq", "doadmin:flgkgp2q2mmxuumj@tcp(db-sg1-do-user-6494085-0.db.ondigitalocean.com:25060)/compose")
	if err != nil {
		return nil, err
	}

	return db, nil

}

type kategori []kategoriElement

type kategoriElement struct {
	MainKategori string        `json:"main_kategori"`
	SubKategori  []subKategori `json:"sub_kategori"`
}

type subKategori struct {
	Nama string `json:"nama"`
	ID   int64  `json:"id"`
}

type student struct {
	ID    string
	Name  string
	Grade int
}

var kat = []kategoriElement{
	kategoriElement{"Telekomunikasi", []subKategori{
		subKategori{"pulsa", 11},
		subKategori{"Pascabayar", 12},
	}},
	kategoriElement{"Tagihan", []subKategori{
		subKategori{"Listrik", 21},
		subKategori{"PDAM", 22},
		subKategori{"TV Kabel", 23},
		subKategori{"Internet", 24},
		subKategori{"Telepon", 25},
		subKategori{"Gas", 26},
		subKategori{"Rekening Virtual", 27},
		subKategori{"E-Commerce", 28},
	}},
	kategoriElement{"Transportasi", []subKategori{
		subKategori{"Kereta", 31},
		subKategori{"taksi", 32},
		subKategori{"Bis", 33},
	}},
}

var data = []student{
	student{"E001", "ethan", 21},
	student{"W001", "wick", 22},
	student{"B001", "bourne", 23},
	student{"B002", "bond", 23},
}

func getKategori(c echo.Context) error {
	return c.JSON(http.StatusOK, kat)
}

func users(c echo.Context) error {
	return c.JSON(http.StatusOK, data)
}
func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// func accessible(c echo.Context) error {
// 	var message = [...]string{
// 		"Silent but deadly",
// 		"Dummy sample REST API",
// 		"Built with Golang"}
// 	//	for _, mes := range message {
// 	//		return fmt.Printf(mes)
// 	//	}
// 	return c.(http.StatusOK, message)
// }

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func home(c echo.Context) error {
	return c.File("public/index.html")
}

func main() {
	fp := path.Join("/home/burhan/Documents/go/src/go_rest_api", "index.html")
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	// index
	e.GET("/", func(c echo.Context) error {
		return c.File(fp)
	})

	e.GET("/users", users)

	e.GET("/kategori", getKategori)
	e.GET("/nana", controller.GetKategori)
	e.GET("/nanana", controller.GetAllUser)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":8080"))
}
