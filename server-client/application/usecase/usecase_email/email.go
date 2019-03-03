package usecase_email

import (
	"github.com/yuki-toida/ca-micro-grpc/server-client/domain/entities/entity_email"
)

type UseCase interface {
	Update(emailID uint64, emailAddr string) (*entity_email.Email, error)
}
