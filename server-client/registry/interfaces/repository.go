package interfaces

import (
	"github.com/yuki-toida/grpc-clean/server-client/domain/entities/entity_email"
	"github.com/yuki-toida/grpc-clean/server-client/domain/entities/entity_profile"
	"github.com/yuki-toida/grpc-clean/server-client/domain/entities/entity_user"
)

type Repositories interface {
	NewUserRepository() entity_user.Repository
	NewProfileRepository() entity_profile.Repository
	NewEmailRepository() entity_email.Repository
}
