package kk_etcd

import (
	"errors"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
)

// key
var (
	ErrKeyNotFound      = errors.New("key not found")
	ErrKeyAlreadyExists = errors.New("key already exists")
)

// role
var (
	ErrNoRootRole = errors.New("no root role")
)

var (
	ErrNoAvailableConn = errors.New("no available conn")
)

// Permission
var (
	ErrPermissionDenied = errors.New("permission denied")
)

// ErrorIs temporary solution for https://github.com/etcd-io/etcd/issues/18493
func ErrorIs(err error, target error) bool {
	if err == nil || target == nil {
		return errors.Is(err, target)
	}
	return err.Error() == rpctypes.ErrorDesc(target)
}
