package interfaces

import (
	"github.com/yuki-toida/ca-micro-grpc/server-client/application/usecase/usecase_email"
	"github.com/yuki-toida/ca-micro-grpc/server-client/application/usecase/usecase_user"
)

type UseCases interface {
	NewUserUseCase() usecase_user.UseCase
	NewEmailUseCase() usecase_email.UseCase
}
