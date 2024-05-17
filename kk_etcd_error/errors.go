package kk_etcd_error

import "errors"

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
