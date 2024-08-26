package kk_etcd_error

import (
	"errors"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
)

// key
var (
	KeyNotFound      = errors.New("key not found")
	KeyAlreadyExists = errors.New("key already exists")
)

// role
var (
	NoRootRole = errors.New("no root role")
)

// Permission
var (
	PermissionDenied = errors.New("permission denied")
)

// ErrorIs temporary solution for https://github.com/etcd-io/etcd/issues/18493
func ErrorIs(err error, target error) bool {
	if err == nil || target == nil {
		return errors.Is(err, target)
	}
	return err.Error() == rpctypes.ErrorDesc(target)
}
