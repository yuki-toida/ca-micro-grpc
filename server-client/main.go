package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/go-kit/kit/log"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"server-client/adapter/repositories/repository_email"
	"server-client/adapter/repositories/repository_profile"
	"server-client/adapter/repositories/repository_user"
	"server-client/external/config"
	"server-client/external/mysql"
	"server-client/external/web"
	"server-client/registry"
)

func main() {
	c := config.Load()
	db := mysql.Connect(c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)
	defer db.Close()

	db.LogMode(true)
	db.DropTableIfExists(&repository_user.User{}, &repository_profile.Profile{}, &repository_email.Email{})
	db.AutoMigrate(&repository_user.User{}, &repository_profile.Profile{}, &repository_email.Email{})

	for i := 1; i < 3; i++ {
		p := repository_profile.Profile{Name: strconv.Itoa(i)}
		db.Create(&p)
		u := repository_user.User{ProfileID: p.ID}
		db.Create(&u)
		for j := 1; j < 4; j++ {
			e := strconv.Itoa(j) + "@hacobu.jp"
			db.Create(&repository_email.Email{Email: e, UserID: u.ID})
		}
	}

	repositories := registry.NewRepositories(db)
	usecases := registry.NewUseCases(repositories)
	handler := web.Handle(usecases)

	port := ":8080"
	logger := log.NewLogfmtLogger(os.Stderr)
	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(port, handler))

}
