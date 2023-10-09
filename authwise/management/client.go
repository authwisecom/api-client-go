package management

import (
	"github.com/authwisecom/api-client-go/authwise"
	v1 "github.com/authwisecom/api-client-go/authwise/management/v1alpha1"
	"google.golang.org/grpc"
)

// NewDefaultClient returns default client with authentication support
func NewDefaultClient(address string, options ...grpc.DialOption) (v1.AuthwiseManagementServiceClient, error) {
	opts := append(options, grpc.WithUnaryInterceptor(authwise.NewDefaultBearerCredentialsInterceptor()))
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}
	return v1.NewAuthwiseManagementServiceClient(conn), nil
}
