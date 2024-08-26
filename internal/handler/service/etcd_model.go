package service

type InitKKEtcdConfig struct {
	Endpoints    []string
	RootPassword string
	UserName     string
	Password     string
	DebugMode    bool
}

func (x *InitKKEtcdConfig) Check() {
	if len(x.Endpoints) == 0 {
		panic("endpoints is empty")
	}
	if x.RootPassword == "" {
		panic("rootPassword is empty")
	}
	if x.UserName == "" {
		panic("userName is empty")
	}
	if x.Password == "" {
		panic("password is empty")
	}
}
