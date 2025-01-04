package handler

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	pb "karplexus.com/janitor2/api/janitor/v1"
	"karplexus.com/janitor2/api/payhub/upi"
	"karplexus.com/janitor2/internal/conf"
	"karplexus.com/janitor2/internal/data"
	"karplexus.com/janitor2/internal/data/model"
	"karplexus.com/janitor2/internal/util"
)

type RuntimeCfgHandler struct {
	hRuntimeCfgRepo *data.RuntimeCfgRepo
	logger          *log.Helper
	cfg             *conf.MessageBroker
}

func NewRuntimeCfgHandler(hRuntimeCfgRepo *data.RuntimeCfgRepo, cfg *conf.MessageBroker, logger log.Logger) *RuntimeCfgHandler {
	return &RuntimeCfgHandler{hRuntimeCfgRepo: hRuntimeCfgRepo, cfg: cfg, logger: log.NewHelper(logger)}
}

func (h *RuntimeCfgHandler) Create(ctx context.Context, in *pb.CreateRuntimeCfgRq) (*pb.CreateRuntimeCfgRs, error) {

	if err := in.Validate(); err != nil {
		return &pb.CreateRuntimeCfgRs{Status: util.BadRequest(err.Error())}, nil
	}

	runTimeCfg := model.RuntimeCfg{
		InstitutionId: in.InstitutionId,
		ParentKey:     in.ParentKey,
		CfgKey:        in.CfgKey,
		CfgValue:      in.CfgValue,
	}

	if err := h.hRuntimeCfgRepo.Create(ctx, &runTimeCfg); err != nil {

		return &pb.CreateRuntimeCfgRs{Status: util.Failed(upi.ProcessingErrorCode_INTERNAL_ERROR, "RUNTIME CFG CREATION FAILED", err.Error())}, nil
	}

	return &pb.CreateRuntimeCfgRs{Status: util.Success(), Id: runTimeCfg.Id}, nil
}

func (h *RuntimeCfgHandler) Get(ctx context.Context, in *pb.GetRuntimeCfgRq) (*pb.GetRuntimeCfgRs, error) {

	if err := in.Validate(); err != nil {
		return &pb.GetRuntimeCfgRs{Status: util.BadRequest(err.Error())}, nil
	}

	if runtimeCfg, err := h.hRuntimeCfgRepo.Get(in.GetInstitutionId()); err != nil {
		return &pb.GetRuntimeCfgRs{Status: util.Failed(upi.ProcessingErrorCode_FAILURE, "Record not found", err.Error())}, nil
	} else {
		var rtcfg = make([]*pb.RuntimeCfg, 0)
		for _, cfg := range runtimeCfg {
			pbRunTimeCfg := pb.RuntimeCfg{
				Id:               cfg.Id,
				InstitutionId:    in.InstitutionId,
				CfgKey:           cfg.CfgKey,
				ParentKey:        cfg.ParentKey,
				CfgValue:         cfg.CfgValue,
				Version:          uint32(cfg.Version),
				CreatedTime:      cfg.CreatedTime.String(),
				LastModifiedTime: cfg.LastModifiedTime.String(),
			}
			rtcfg = append(rtcfg, &pbRunTimeCfg)
		}
		return &pb.GetRuntimeCfgRs{Status: util.Success(), RuntimeCfg: rtcfg}, nil
	}
}

func (h *RuntimeCfgHandler) Update(ctx context.Context, in *pb.UpdateRuntimeCfgRq) (*pb.UpdateRuntimeCfgRs, error) {

	if err := in.Validate(); err != nil {
		return &pb.UpdateRuntimeCfgRs{Status: util.BadRequest(err.Error())}, nil
	}

	runTimeCfg, err := h.hRuntimeCfgRepo.GetByValue(in.GetInstitutionId(), in.GetParentKey(), in.GetCfgKey())
	if err != nil {
		return &pb.UpdateRuntimeCfgRs{Status: util.BadRequest(err.Error())}, nil
	}

	if len(runTimeCfg.InstitutionId) == 0 || len(runTimeCfg.ParentKey) == 0 || len(runTimeCfg.CfgKey) == 0 {
		return &pb.UpdateRuntimeCfgRs{Status: util.NotFound(upi.ProcessingErrorCode_FAILURE, "Record not Found")}, nil
	} else {
		if len(in.CfgValue) > 0 {
			runTimeCfg.CfgValue = in.CfgValue
			runTimeCfg.Version = runTimeCfg.Version + 1
			runTimeCfg.LastModifiedTime = time.Now()
		}

		if err := h.hRuntimeCfgRepo.Update(ctx, runTimeCfg); err != nil {
			return &pb.UpdateRuntimeCfgRs{Status: util.Failed(upi.ProcessingErrorCode_INTERNAL_ERROR, "failed while updating runtime details ", err.Error())}, nil
		}
		return &pb.UpdateRuntimeCfgRs{Status: util.Success(), Id: runTimeCfg.Id}, nil
	}
}
