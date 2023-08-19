package repository

import (
	"fmt"

	"github.com/Vasiliy82/PoolingRouterEmul/internal/model"
)

type CommonStorage struct {
	requests  map[string]*model.Request
	responses map[string]*model.Results
}

func NewPoolingStorage() *CommonStorage {
	return &CommonStorage{
		requests:  make(map[string]*model.Request),
		responses: make(map[string]*model.Results),
	}
}

func (p *CommonStorage) PutRequest(request *model.Request) (*model.Response, error) {
	response := model.NewResponse()
	if _, found := p.requests[response.IDRequest]; found {
		return nil, fmt.Errorf("идентификатор запроса IDRequest=%s уже существует", response.IDRequest)
	}
	p.requests[response.IDRequest] = request
	return response, nil

}
func (p *CommonStorage) GetRequest(response *model.Response) (*model.Request, error) {
	var result *model.Request
	var found bool

	if result, found = p.requests[response.IDRequest]; !found {
		return nil, fmt.Errorf("идентификатор запроса IDRequest=%s не найден", response.IDRequest)
	}
	return result, nil

}

func (p *CommonStorage) PutResponse(response *model.Response, results *model.Results) error {
	if _, found := p.requests[response.IDRequest]; !found {
		return fmt.Errorf("идентификатор запроса IDRequest=%s не найден", response.IDRequest)
	}
	if _, found := p.responses[response.IDRequest]; found {
		return fmt.Errorf("ответ на запрос IDRequest=%s уже существует", response.IDRequest)
	}
	p.responses[response.IDRequest] = results

	return nil
}

func (p *CommonStorage) GetResponse(response *model.Response) (*model.Results, error) {
	if _, found := p.requests[response.IDRequest]; !found {
		return nil, fmt.Errorf("идентификатор запроса IDRequest=%s не найден", response.IDRequest)
	}
	if _, found := p.responses[response.IDRequest]; !found {
		return nil, fmt.Errorf("ответ на запрос IDRequest=%s не найден", response.IDRequest)
	}
	return p.responses[response.IDRequest], nil
}
