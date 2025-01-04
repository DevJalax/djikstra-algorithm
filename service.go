package service

import (
	"context"

	pb "karplexus.com/janitor2/api/janitor/v1"
	"karplexus.com/janitor2/internal/handler"
)

type RuntimeConfigService struct {
	pb.UnimplementedRuntimeConfigServer
	hRuntimeCfg *handler.RuntimeCfgHandler
}

func NewRuntimeConfigService(hRuntimeCfg *handler.RuntimeCfgHandler) *RuntimeConfigService {
	return &RuntimeConfigService{hRuntimeCfg: hRuntimeCfg}
}

func (s *RuntimeConfigService) CreateRunTimeCfg(ctx context.Context, req *pb.CreateRuntimeCfgRq) (*pb.CreateRuntimeCfgRs, error) {
	return s.hRuntimeCfg.Create(ctx, req)
}
func (s *RuntimeConfigService) UpdateRunTimeCfg(ctx context.Context, req *pb.UpdateRuntimeCfgRq) (*pb.UpdateRuntimeCfgRs, error) {
	return s.hRuntimeCfg.Update(ctx, req)
}
func (s *RuntimeConfigService) GetRunTimeCfg(ctx context.Context, req *pb.GetRuntimeCfgRq) (*pb.GetRuntimeCfgRs, error) {
	return s.hRuntimeCfg.Get(ctx, req)
}
func (s *RuntimeConfigService) BootstrapRuntimeCfg(ctx context.Context, req *pb.BootstrapRuntimeCfgRq) (*pb.BootstrapRuntimeCfgRq, error) {
	return &pb.BootstrapRuntimeCfgRq{}, nil
}
