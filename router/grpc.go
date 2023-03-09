package router

import (
	"context"
	"fmt"
	"log"
	"strconv"

	consul "github.com/hashicorp/consul/api"
	"github.com/raven0520/btc/app"
	"github.com/raven0520/btc/service"
	"github.com/raven0520/proto/btc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// var grpcServerList = []*GrpcServer{}

// HealthImpl Check Grpc Health
type HealthImpl struct {
	Status grpc_health_v1.HealthCheckResponse_ServingStatus
	Reason string
}

// ConsulServer Consul Server
type ConsulServer struct {
	Host string
	Port int
	Addr string
}

// GrpcServer GRPC Server
type GrpcServer struct {
	Env    string
	Name   string
	Host   string
	Port   int
	Addr   string
	Consul ConsulServer
	*grpc.Server
}

// NewGrpcServer Generate Grpc Server
func NewGrpcServer() *GrpcServer {
	server := &GrpcServer{
		Env:  app.BaseConf.Consul.Env,
		Name: app.BaseConf.Grpc.Name,
		Host: app.BaseConf.Grpc.Host,
		Port: app.BaseConf.Grpc.Port,
		Consul: ConsulServer{
			Host: app.BaseConf.Consul.Host,
			Port: app.BaseConf.Consul.Port,
		},
		Server: &grpc.Server{},
	}
	server.Addr = fmt.Sprintf("%s:%d", server.Host, server.Port)
	server.Consul.Addr = fmt.Sprintf("%s:%d", server.Consul.Host, server.Consul.Port)
	return server
}

// where the health status is returned directly, and there can also be more complex health check strategies, such as returning according to the server load.
// Check Implement the health check interface
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch Implement RegisterHealthServer
func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

// OnLine Be Online
func (h *HealthImpl) OnLine(reason string) {
	h.Status = grpc_health_v1.HealthCheckResponse_SERVING
	h.Reason = reason
	fmt.Println(reason)
}

// OffLine Off Line
func (h *HealthImpl) OffLine(reason string) {
	h.Status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	h.Reason = reason
	fmt.Println(reason)
}

// grpcRegister grpc register to consul
func (g *GrpcServer) GrpcRegister() {
	config := consul.DefaultConfig()
	config.Address = g.Consul.Addr
	client, err := consul.NewClient(config)
	if err != nil {
		panic(err)
	}
	agent := client.Agent()
	reg := &consul.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", g.Name, g.Env, g.Port), // Server Node ID
		Name:    fmt.Sprintf("%s-%s", g.Name, g.Env),            // Server Name
		Tags:    []string{g.Name, g.Env},                        // Server Tag
		Port:    g.Port,                                         // Server Port
		Address: g.Addr,                                         // Server Addr
		Check: &consul.AgentServiceCheck{
			Interval: "30s",
			GRPC:     fmt.Sprintf("%v/%s", g.Addr, ""),

			DeregisterCriticalServiceAfter: "5s",
		},
	}
	if err := agent.ServiceRegister(reg); err != nil {
		panic(err)
	}
}

// GrpcUnregister Stop rpc service
func (g *GrpcServer) GrpcUnregister() {
	config := consul.DefaultConfig()
	config.Address = g.Consul.Addr
	client, err := consul.NewClient(config)
	if err != nil {
		panic(err)
	}
	err = client.Agent().ServiceDeregister(fmt.Sprintf("%v-%v-%v", g.Name, g.Env, g.Port))
	if err != nil {
		panic(err)
	}
}

// GrpcServerStart Start Grpc service
func (g *GrpcServer) GrpcServerStart() {
	var server = new(service.BtcService)
	var options []grpc.ServerOption
	s := grpc.NewServer(options...)
	btc.RegisterBTCServiceServer(s, server)               // register service
	reflection.Register(s)                                // register
	grpc_health_v1.RegisterHealthServer(s, &HealthImpl{}) // register health server
	grace, err := app.NewGraceGrpc(s, "tcp", ":"+strconv.Itoa(g.Port), app.BaseConf.Path.Pid)
	if err != nil {
		log.Printf("Failed Server: %s ", err.Error())
	}
	if err := grace.Serve(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
