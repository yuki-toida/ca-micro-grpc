package registry

import (
	"server-client/application/interactors/interactor_email"
	"server-client/application/interactors/interactor_user"
	"server-client/application/usecase/usecase_email"
	"server-client/application/usecase/usecase_user"
	"server-client/registry/interfaces"
)

type useCases struct {
	repositories interfaces.Repositories
}

func NewUseCases(r interfaces.Repositories) interfaces.UseCases {
	return &useCases{repositories: r}
}

func (u *useCases) NewUserUseCase() usecase_user.UseCase {
	ur := u.repositories.NewUserRepository()
	pr := u.repositories.NewProfileRepository()
	er := u.repositories.NewEmailRepository()
	return interactor_user.New(ur, pr, er)
}

func (u *useCases) NewEmailUseCase() usecase_email.UseCase {
	er := u.repositories.NewEmailRepository()
	return interactor_email.New(er)
}
