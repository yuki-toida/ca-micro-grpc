package interfaces

import (
	"server-client/application/usecase/usecase_email"
	"server-client/application/usecase/usecase_user"
)

type UseCases interface {
	NewUserUseCase() usecase_user.UseCase
	NewEmailUseCase() usecase_email.UseCase
}
