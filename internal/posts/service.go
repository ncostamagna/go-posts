package posts

import "log/slog"

type (
	Filters struct {
		Name string
	}

	Service interface {
	}

	service struct {
		log  *slog.Logger
		repo Repository
	}
)
