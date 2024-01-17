package metric

type UseCase interface {
}

type Metric struct {
}

func New() *Metric {
	return &Metric{}
}

func (u *Metric) PromHttp() error {
	return nil
}
