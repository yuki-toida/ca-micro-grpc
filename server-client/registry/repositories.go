package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/ca-micro-grpc/server-client/adapter/repositories/repository_email"
	"github.com/yuki-toida/ca-micro-grpc/server-client/adapter/repositories/repository_profile"
	"github.com/yuki-toida/ca-micro-grpc/server-client/adapter/repositories/repository_user"
	"github.com/yuki-toida/ca-micro-grpc/server-client/domain/entities/entity_email"
	"github.com/yuki-toida/ca-micro-grpc/server-client/domain/entities/entity_profile"
	"github.com/yuki-toida/ca-micro-grpc/server-client/domain/entities/entity_user"
	"github.com/yuki-toida/ca-micro-grpc/server-client/registry/interfaces"
)

type repositories struct {
	db *gorm.DB
}

func NewRepositories(db *gorm.DB) interfaces.Repositories {
	return &repositories{db: db}
}

func (r *repositories) NewUserRepository() entity_user.Repository {
	return repository_user.New(r.db)
}

func (r *repositories) NewProfileRepository() entity_profile.Repository {
	return repository_profile.New(r.db)
}

func (r *repositories) NewEmailRepository() entity_email.Repository {
	return repository_email.New(r.db)
}
