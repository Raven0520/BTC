package app

import (
	"fmt"
	"math/rand"
	"time"

	consul "github.com/hashicorp/consul/api"
	"github.com/raven0520/btc/util"
	"github.com/raven0520/proto/btc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Consul *ConsulServer

// ConsulServer Consul Configture
type ConsulServer struct {
	Env    string
	Host   string
	Port   int
	Addr   string
	Client *consul.Client
}

// InitConsulServer Initialization Consul
func InitConsulServer() (err error) {
	env := BaseConf.Consul.Env
	host := BaseConf.Consul.Host
	port := BaseConf.Consul.Port
	addr := fmt.Sprintf("%s:%d", host, port)
	Consul = &ConsulServer{
		Env:  env,
		Host: host,
		Port: port,
		Addr: addr,
	}
	config := consul.DefaultConfig()
	config.Address = addr
	Consul.Client, err = consul.NewClient(config)
	return
}

// Connect Connect Service
func Connect(tag string) (*grpc.ClientConn, error) {
	services, _, err := Consul.Client.Health().Service(fmt.Sprintf("%s-%s", tag, Consul.Env), Consul.Env, true, nil)
	if err != nil {
		return nil, err
	}
	if len(services) == 0 {
		return nil, util.ErrServiceNotFound
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(services))
	service := services[index]
	return grpc.Dial(service.Service.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// GetBTC Return BTC Service
func GetBTC() (btc.BTCServiceClient, *grpc.ClientConn, error) {
	Connect, err := Connect("BTC")
	if err != nil {
		return nil, nil, err
	}
	return btc.NewBTCServiceClient(Connect), Connect, nil
}
