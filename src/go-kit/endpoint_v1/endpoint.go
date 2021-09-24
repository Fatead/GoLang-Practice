package endpoint_v1

import (
	"awesomeProject/src/go-kit/service_v1"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type EndPointServer struct {
	AddEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc service_v1.Service) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
	}
	return EndPointServer{AddEndPoint: addEndPoint}
}

func (s EndPointServer) Add(ctx context.Context, in service_v1.Add) service_v1.AddAck {
	res, _ := s.AddEndPoint(ctx, in)
	return res.(service_v1.AddAck)
}

func MakeAddEndPoint(s service_v1.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service_v1.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}
