package client

import (
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

type Client struct {
	ContactEmail    string
	ContactPassword string
	ContactToken    string
	Server          string
	conn            *grpc.ClientConn
	Context         context.Context
}

func Initialize() *Client {
	return &Client{
		ContactEmail:    viper.GetString("Email"),
		ContactPassword: viper.GetString("Password"),
		Server:          viper.GetString("API.base"),
		Context:         context.Background(),
	}
}

func (c *Client) Connection() *grpc.ClientConn {
	if c.conn == nil {
		conn, err := grpc.Dial(c.Server, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		c.conn = conn
	}
	return c.conn
}

func (c *Client) Contact() *ContactClient {
	return &ContactClient{
		Client: c,
	}
}

func (c *Client) Mirror() *MirrorClient {
	return &MirrorClient{
		Client: c,
	}
}

func (c *Client) PrepareHeader() {
	if c.Context == nil {
		c.Context = context.Background()
	}
	md := make(metadata.MD)
	if len(c.ContactToken) > 0 {
		md = metadata.New(map[string]string{
			"ContactToken": c.ContactToken,
		})
	}
	grpc.SetHeader(c.Context, md)
}