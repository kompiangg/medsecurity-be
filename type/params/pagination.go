package params

type RepositoryPaginationParam struct {
	Limit  uint64
	Offset uint64
}

type ServicePaginationParam struct {
	Limit  uint64 `query:"limit"`
	Offset uint64 `query:"offset"`
}

func (p ServicePaginationParam) ToRepositoryPaginationParam() RepositoryPaginationParam {
	if p.Limit == 0 {
		p.Limit = 10
	}

	if p.Offset == 0 {
		p.Offset = 0
	}

	return RepositoryPaginationParam(p)
}
