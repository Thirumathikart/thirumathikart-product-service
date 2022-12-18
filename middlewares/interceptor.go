package middlewares

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func WithClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	var err error
	ctx = metadata.AppendToOutgoingContext(ctx, "secret", "xxxx")
	var header metadata.MD
	opts = append(opts, grpc.Header(&header))
	start := time.Now()
	// Calls the invoker to execute RPC
	err = invoker(ctx, method, req, reply, cc, opts...)
	// Logic after invoking the invoker
	GrpcLog.Infof("Invoked RPC method=%s; Duration=%s; Error=%v", method,
		time.Since(start), err)
	return err
}
