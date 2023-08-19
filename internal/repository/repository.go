package repository

import "github.com/Vasiliy82/PoolingRouterEmul/internal/model"

type Storage interface {
	PutRequest(*model.Request) (*model.Response, error)
	GetRequest(*model.Response) (*model.Request, error)

	PutResponse(*model.Response, *model.Results) error
	GetResponse(*model.Response) (*model.Results, error)
}
