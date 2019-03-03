package transport_email

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"github.com/yuki-toida/ca-micro-grpc/server-client/application/usecase/usecase_email"
)

func MakeUpdateEndpoint(u usecase_email.UseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		uid, _ := strconv.Atoi(req.Uid)
		email, err := u.Update(uint64(uid), req.Email)
		if err != nil {
			return nil, err
		}
		return email, nil
	}
}

func DecodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request updateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type updateRequest struct {
	Uid   string
	Email string
}
