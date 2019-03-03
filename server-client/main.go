package main

import (
	"net/http"
	"os"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/go-kit/kit/log"
	"github.com/yuki-toida/ca-micro-grpc/server-client/adapter/repositories/repository_email"
	"github.com/yuki-toida/ca-micro-grpc/server-client/adapter/repositories/repository_profile"
	"github.com/yuki-toida/ca-micro-grpc/server-client/adapter/repositories/repository_user"
	"github.com/yuki-toida/ca-micro-grpc/server-client/external/config"
	"github.com/yuki-toida/ca-micro-grpc/server-client/external/mysql"
	"github.com/yuki-toida/ca-micro-grpc/server-client/external/web"
	"github.com/yuki-toida/ca-micro-grpc/server-client/registry"
)

func main() {
	c := config.Load()
	db := mysql.Connect(c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)
	defer mysql.Close()

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
