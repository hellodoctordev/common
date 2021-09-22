package clients

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

func GetDB() *pg.DB {
	return pg.Connect(&pg.Options{
		Network:  "unix",
		Addr:     fmt.Sprintf("%s/%s/.s.PGSQL.5432", os.Getenv("CLOUD_SQL_SOCKET_PREFIX"), os.Getenv("CLOUD_SQL_INSTANCE_NAME")),
		User:     os.Getenv("CLOUD_SQL_USERNAME"),
		Password: os.Getenv("CLOUD_SQL_PASSWORD"),
		Database: os.Getenv("CLOUD_SQL_DATABASE"),
	})
}
