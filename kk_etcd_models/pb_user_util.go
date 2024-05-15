package kk_etcd_models

func NewPBUser(userName string) *PBUser {
	return &PBUser{
		UserName: userName,
	}
}
