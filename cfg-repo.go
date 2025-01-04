package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"karplexus.com/janitor2/internal/data/model"
)

type RuntimeCfgRepo struct {
	db     *gorm.DB
	logger *log.Helper
}

func NewRuntimeCfgRepo(db *gorm.DB, logger log.Logger) *RuntimeCfgRepo {
	return &RuntimeCfgRepo{
		db:     db,
		logger: log.NewHelper(logger),
	}
}

func (r *RuntimeCfgRepo) Create(ctx context.Context, req *model.RuntimeCfg) error {
	if dbc := r.db.Create(req); dbc.Error != nil {
		return dbc.Error
	} else {
		return nil
	}
}

func (r *RuntimeCfgRepo) Get(id string) ([]model.RuntimeCfg, error) {
	found := make([]model.RuntimeCfg, 0)
	res := r.db.Where("institution_id = ? ", id).Find(&found)
	if res.RowsAffected < 1 {
		return nil, res.Error
	} else {
		return found, nil
	}
}

func (r *RuntimeCfgRepo) Update(ctx context.Context, req *model.RuntimeCfg) error {

	return r.db.Save(req).Error
}

func (r *RuntimeCfgRepo) GetByValue(iid string, parentkey string, cfgkey string) (*model.RuntimeCfg, error) {

	var cfg model.RuntimeCfg
	if dbc := r.db.Where("institution_id = ? AND parent_key = ? AND cfg_key = ?", iid, parentkey, cfgkey).Find(&cfg); dbc.Error == nil {
		return &cfg, nil
	} else {
		return nil, dbc.Error
	}

}

// Heart Beat
func (r *RuntimeCfgRepo) GetDefaultOrgId() (string, error) {
	return r.getCfgValue("default", "org", "id")
}

func (r *RuntimeCfgRepo) GetUpiBankCode(iid string) (string, error) {
	return r.getCfgValue(iid, "upi", "bankCode")
}

func (r *RuntimeCfgRepo) getCfgValue(iid, parentCfgKey, cfgKey string) (string, error) {
	var cfg model.RuntimeCfg
	if dbc := r.db.Where("institution_id = ? AND parent_key = ? AND cfg_key = ?", iid, parentCfgKey, cfgKey).Find(&cfg); dbc.Error == nil {
		return cfg.CfgValue, nil
	} else {
		return "", dbc.Error
	}
}
