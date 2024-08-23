package service

import (
	"context"
	"errors"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
)

type ExtPBUser struct {
	User *kk_etcd_models.PBUser
}

func NewExtPBUser(pbUser *kk_etcd_models.PBUser) *ExtPBUser {
	return &ExtPBUser{
		User: pbUser,
	}
}

func (e *ExtPBUser) Store() error {
	if err := toolUser.checkFields(e.User); err != nil {
		return err
	}
	err := serKV.PutJson(kk_etcd_const.User+e.User.GetUserName(), e.User)
	if err != nil {
		return err
	}
	_, err = kk_etcd_client.EtcdClient.UserAdd(context.Background(), e.User.GetUserName(), e.User.GetPassword())
	if err != nil && !errors.Is(err, rpctypes.ErrGRPCUserAlreadyExist) {
		return err
	}
	return nil
}

func (e *ExtPBUser) Update() error {
	if err := toolUser.checkFields(e.User); err != nil {
		return err
	}
	err := serKV.UpdateJson(kk_etcd_const.User+e.User.GetUserName(), e.User)
	return err
}

func (e *ExtPBUser) Load() error {
	if err := toolUser.checkFields(e.User); err != nil {
		return err
	}
	err := serKV.GetJson(kk_etcd_const.User+e.User.GetUserName(), e.User)
	return err
}
func (e *ExtPBUser) Delete() error {
	if err := toolUser.checkFields(e.User); err != nil {
		return err
	}
	err := serKV.KVDel(kk_etcd_const.User + e.User.GetUserName())
	if err != nil {
		return err
	}
	_, err = kk_etcd_client.EtcdClient.UserDelete(context.Background(), kk_etcd_const.User+e.User.GetUserName())
	if err != nil {
		return err
	}
	return nil
}
func (e *ExtPBUser) JwtGet() (string, error) {
	if err := toolUser.checkFields(e.User); err != nil {
		return "", err
	}
	err, value := serKV.KVGet(kk_etcd_const.Jwt + e.User.GetUserName())
	return string(value), err
}
func (e *ExtPBUser) JwtSet(token string) error {
	if err := toolUser.checkFields(e.User); err != nil {
		return err
	}
	err := serKV.KVPut(kk_etcd_const.Jwt+e.User.GetUserName(), token)
	return err
}

func (e *ExtPBUser) JwtDel() error {
	if err := toolUser.checkFields(e.User); err != nil {
		return err
	}
	err := serKV.KVDel(kk_etcd_const.Jwt + e.User.GetUserName())
	return err
}
