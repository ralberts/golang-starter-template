package database

import (
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// DB connection with gorm
	DB *gorm.DB
)

// Connect to the database
// Should look like this from Heroku: postgres://<username>:<password>@<host>/<dbname>
func Connect(dbURL string, sslmode string) *gorm.DB {
	log.SetOutput(os.Stdout)
	log.Print("Attempting to connect to db: ", dbURL)
	parsedURL, err := url.Parse(dbURL)
	hostAndPort := strings.Split(parsedURL.Host, ":")
	pgPassword := ""
	if password, ok := parsedURL.User.Password(); ok {
		pgPassword = password
	}

	if err != nil {
		panic(err)
	}

	connectionStr := "host=" + hostAndPort[0] + " port=" + hostAndPort[1] + " user=" + parsedURL.User.Username() + " dbname=" + parsedURL.Path[1:] + " password=" + pgPassword + " sslmode=" + sslmode
	log.Print("Postgres connection string: ", connectionStr)
	DB, err := gorm.Open("postgres", connectionStr)

	if err != nil {
		panic(err)
	}

	// defer DB.Close()
	CreateSchema(DB)
	Migrate(DB)
	return DB
}
