package ping

type ServiceItf interface {
	Ping() string
}

type Ping struct {
}

func New() Ping {
	return Ping{}
}
