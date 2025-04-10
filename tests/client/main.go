package main

import (
	"context"
	"time"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type Client struct {
	cli    pb.LibrarianSephirahServiceClient
	userID *librarian.InternalID
}

func main() {
	ctx := context.Background()

	c := Client{
		cli:    NewSephirahClient(),
		userID: nil,
	}
	c.WaitServerOnline(ctx)
	log.Info("[Client] Server online, login in via default admin")
	ctx = c.LoginViaDefaultAdmin(ctx)
	log.Info("[Client] Login successful")
	log.Info("[Client] Test tiphereth begin")
	c.TestTiphereth(ctx)
	log.Info("[Client] Test tiphereth finish")
	log.Info("[Client] Test gebura begin")
	c.TestGebura(ctx)
	log.Info("[Client] Test gebura finish")
	log.Info("[Client] Test yesod begin")
	c.TestYesod(ctx)
	log.Info("[Client] Test yesod finish")

	log.Info("[Client] All test finished")
}

func NewSephirahClient() pb.LibrarianSephirahServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:10000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	cli := pb.NewLibrarianSephirahServiceClient(conn)
	return cli
}

func (c *Client) WaitServerOnline(ctx context.Context) {
	const maxRetry = 30
	_, err := c.cli.GetServerInformation(ctx, new(pb.GetServerInformationRequest))
	i := 1
	for errors.IsServiceUnavailable(err) && i < maxRetry {
		time.Sleep(time.Second)
		i += 1
		log.Infof("[Client] Waiting server online %s", err.Error())
		_, err = c.cli.GetServerInformation(ctx, new(pb.GetServerInformationRequest))
	}
	if i == maxRetry {
		log.Fatal("Server unavailable")
	}
}

var defaultPaging = &librarian.PagingRequest{ //nolint:gochecknoglobals // no need
	PageNum:  1,
	PageSize: 10, //nolint:mnd // no need
}
