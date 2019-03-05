package interfaces

import (
	"server-client/domain/entities/entity_email"
	"server-client/domain/entities/entity_profile"
	"server-client/domain/entities/entity_user"
)

type Repositories interface {
	NewUserRepository() entity_user.Repository
	NewProfileRepository() entity_profile.Repository
	NewEmailRepository() entity_email.Repository
}
