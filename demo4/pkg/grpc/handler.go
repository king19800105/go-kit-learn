package grpc

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/king19800105/go-kit-learn/demo4/pkg/endpoint"
	"github.com/king19800105/go-kit-learn/demo4/pkg/grpc/pb"
	netcontext "golang.org/x/net/context"
)

func makeCreateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)
}

// 请求参数解析
func decodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateRequest)
	return endpoint.CreateRequest{OrderId: req.OrderId}, nil
}

func encodeCreateResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.CreateResponse)
	if nil != reply.Failed() {
		return nil, reply.Failed()
	}

	return &pb.CreateReply{Id: reply.Id, Source: reply.Source, IsPlay: reply.IsPay}, nil
}

func (g *grpcServer) Create(ctx netcontext.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	_, rep, err := g.create.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	return rep.(*pb.CreateReply), nil
}
