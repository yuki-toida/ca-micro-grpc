package interfaces

import (
	"github.com/yuki-toida/grpc-clean/server-client/application/usecase/usecase_email"
	"github.com/yuki-toida/grpc-clean/server-client/application/usecase/usecase_user"
)

type UseCase interface {
	NewUserUseCase() usecase_user.UseCase
	NewEmailUseCase() usecase_email.UseCase
}
