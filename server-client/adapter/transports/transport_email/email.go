package transport_email

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"net/http"
	"server-client/adapter/transports/transport_email/pb"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"server-client/application/usecase/usecase_email"
)

func MakeUpdateEndpoint(u usecase_email.UseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)

		// gRPC client
		conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		defer conn.Close()
		endpoint := grpctransport.NewClient(
			conn,
			"proto.Address",
			"Get",
			func(_ context.Context, request interface{}) (interface{}, error) {
				return request.(*pb.Request), nil
			},
			func(_ context.Context, response interface{}) (interface{}, error) {
				return response.(*pb.Response), nil
			},
			pb.Response{},
		).Endpoint()
		res, err := endpoint(context.Background(), &pb.Request{Email: req.Email})
		if err != nil {
			return nil, err
		}
		emailAddr := res.(*pb.Response)

		emailID, _ := strconv.Atoi(req.ID)
		email, err := u.Update(uint64(emailID), emailAddr.EmailAddress)
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
	ID    string
	Email string
}
