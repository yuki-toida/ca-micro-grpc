package registry

import (
	"github.com/jinzhu/gorm"
	"server-client/adapter/repositories/repository_email"
	"server-client/adapter/repositories/repository_profile"
	"server-client/adapter/repositories/repository_user"
	"server-client/domain/entities/entity_email"
	"server-client/domain/entities/entity_profile"
	"server-client/domain/entities/entity_user"
	"server-client/registry/interfaces"
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
