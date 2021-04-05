package mail

import (
	"os"
	grpc "google.golang.org/grpc"
)

type ConnectionOption struct {
	Host *string
}

func Connect(connectionOption ConnectionOption) (MailClient, *grpc.ClientConn) {
	defaultGRPCAddress := os.Getenv("MAIL_GRPC")

	grpcAddress := connectionOption.Host
	if grpcAddress == nil {
		grpcAddress = &defaultGRPCAddress
	}

	conn, err := grpc.Dial(*grpcAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return NewMailClient(conn), conn
}
