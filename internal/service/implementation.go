package service

import (
	"math/rand"
	"time"

	"github.com/Vasiliy82/PoolingRouterEmul/internal/logger"
	"github.com/Vasiliy82/PoolingRouterEmul/internal/model"
	"github.com/Vasiliy82/PoolingRouterEmul/internal/repository"
	"github.com/pkg/errors"
)

type CommonPoolingRouterService struct {
	storage repository.Storage
}

func NewPoolingRouterService(storage repository.Storage) *CommonPoolingRouterService {
	return &CommonPoolingRouterService{storage: storage}
}

func (p *CommonPoolingRouterService) Request(request *model.Request) (*model.Response, error) {
	response, err := p.storage.PutRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "PoolingRouterService.Request: ошибка обработки запроса")
	}

	p.asyncProcess(response)

	return response, nil

}
func (p *CommonPoolingRouterService) Results(response *model.Response) (*model.Results, error) {
	result, err := p.storage.GetResponse(response)
	if err != nil {
		return nil, errors.Wrap(err, "PoolingRouterService.Results: ошибка обработки запроса")
	}
	return result, nil
}

func (p *CommonPoolingRouterService) asyncProcess(response *model.Response) {

	go func() {
		request, err := p.storage.GetRequest(response)
		if err != nil {
			logger.Logger().Warnf("PoolingRouterService.asyncProcess: ошибка при обработке запроса IDRequest=%s: %v", response.IDRequest, err)
		}

		logger.Logger().Debugf("Начата обработка запроса IDRequest=%s", response.IDRequest)
		randomDelay := time.Duration(10+rand.Intn(5)) * time.Second

		time.Sleep(randomDelay)

		results := model.Results{
			Status:    "Ok",
			Message:   "Ok",
			IDRequest: response.IDRequest,
		}

		droppedLocations := map[string]any{}
		for _, p := range request.Parcels {
			droppedLocations[p.ID] = p
		}

		for _, c := range request.Couriers {
			orders := 0
			nodesPickup := []model.ResultRouteNode{}
			nodesDelivery := []model.ResultRouteNode{}

			for _, p := range request.Parcels {
				// проверяем, не отдали ли кому-то этот парцел
				_, present := droppedLocations[p.ID]
				if present {
					// если курьеру можно взять еще один парцел
					if orders < c.MaxOrders {
						// начинаем докидывать заказы в массивы "взять-отдать"
						for _, o := range p.Orders {
							nodePickup := model.ResultRouteNode{
								Action:    "pickup",
								PointType: "pickup",
								ParcelID:  p.ID,
								OrderID:   o.ID,
								Location: struct {
									Latitude  float64 "json:\"latitude\""
									Longitude float64 "json:\"longitude\""
								}{
									Latitude:  o.PickupPoint.Latitude,
									Longitude: o.PickupPoint.Longitude,
								},
								DistanceFromPrevious: 1,
								DurationFromPrevious: "00:01:00",
							}
							nodeDelivery := model.ResultRouteNode{
								Action:    "delivery",
								PointType: "delivery",
								ParcelID:  p.ID,
								OrderID:   o.ID,
								Location: struct {
									Latitude  float64 "json:\"latitude\""
									Longitude float64 "json:\"longitude\""
								}{
									Latitude:  p.DeliveryPoint.Location.Latitude,
									Longitude: p.DeliveryPoint.Location.Longitude,
								},
								DistanceFromPrevious: 1,
								DurationFromPrevious: "00:01:00",
							}
							nodesPickup = append(nodesPickup, nodePickup)
							nodesDelivery = append(nodesDelivery, nodeDelivery)

						}
						// считаем, что парцел взяли, т.е. убираем его из droppedLocations[p.ID]
						delete(droppedLocations, p.ID)
						orders++
					}
				}
			}
			// Если по итогу цикла по парцелам курьер схватил хоть что-то, надо добавить элемент в список Routes
			if orders > 0 {
				// соединяем массивы воедино и чиним нумерацию
				nodes := append(nodesPickup, nodesDelivery...)
				for i := range nodes {
					nodes[i].NodeID = 1 + i
				}

				elem := model.ResultRoute{
					CourierID:             c.ID,
					DateShipTarget:        time.Now().Add(10 * time.Minute).Format(time.RFC3339),
					DateShipTargetExtreme: time.Now().Add(15 * time.Minute).Format(time.RFC3339),
					EstimatedTime:         time.Now().Add(120 * time.Minute).Format(time.RFC3339),
					TotalDistance:         1,
					TotalDuration:         1,
					Nodes:                 nodes,
				}
				results.Data.Routes = append(results.Data.Routes, elem)
			}
		}

		if err := p.storage.PutResponse(response, &results); err != nil {
			logger.Logger().Warnf("PoolingRouterService.asyncProcess: ошибка при обработке запроса IDRequest=%s: %v", response.IDRequest, err)
		}

		logger.Logger().Debugf("Завершена обработка запроса IDRequest=%s", response.IDRequest)

		/*

			request_arr := []byte(`{
				"traffic_jams":true,
				"couriers":[
				   {
					  "id":"101",
					  "travel_modes":"driving",
					  "speed":10,
					  "max_orders":2,
					  "max_distance":1000,
					  "max_weight":10,
					  "queue_number":1,
					  "has_terminal":false,
					  "pickup_time":"00:10:00",
					  "hand_over_time":"00:05:00",
					  "delivery_services":["1", "2"]
				   },
				   {
					"id":"102",
					"travel_modes":"walking",
					"speed":5,
					"max_orders":2,
					"max_distance":1000,
					"max_weight":10,
					"queue_number":2,
					"has_terminal":true,
					"pickup_time":"00:10:00",
					"hand_over_time":"00:05:00",
					"delivery_services":["1", "3"]
				 }

				],
				"parcels":[
				   {
					  "id":"10001",
					  "orders":[
						 {
							"id":"10001",
							"shop_no":"201",
							"pickup_point":{
							   "latitude":37.5,
							   "longitude":55.6,
							   "main":true
							}
						 },
						 {
							"id":"10002",
							"shop_no":"201",
							"pickup_point":{
							   "latitude":37.5,
							   "longitude":55.6,
							   "main":false
							}
						 },
						 {
							"id":"10003",
							"shop_no":"201",
							"pickup_point":{
							   "latitude":37.5,
							   "longitude":55.6,
							   "main":false
							}
						 }

					  ],
					  "delivery_point":{
						 "location":{
							"latitude":37.6,
							"longitude":55.5
						 }
					  },
					  "weight":3,
					  "delivery_services":["1"],
					  "needs_terminal":false,
					  "date_supply":"2023-08-10 07:55:12",
					  "date_supply_until_soft":"2023-08-10 08:55:12",
					  "date_supply_until_hard":"2023-08-10 09:55:12",
					  "service_id":"1"
				   },
				   {
					"id":"10004",
					"orders":[
					   {
						  "id":"10004",
						  "shop_no":"201",
						  "pickup_point":{
							 "latitude":37.5,
							 "longitude":55.6,
							 "main":true
						  }
					   },
					   {
						  "id":"10005",
						  "shop_no":"201",
						  "pickup_point":{
							 "latitude":37.5,
							 "longitude":55.6,
							 "main":false
						  }
					   },
					   {
						  "id":"10006",
						  "shop_no":"201",
						  "pickup_point":{
							 "latitude":37.5,
							 "longitude":55.6,
							 "main":false
						  }
					   }

					],
					"delivery_point":{
					   "location":{
						  "latitude":37.6,
						  "longitude":55.5
					   }
					},
					"weight":3,
					"delivery_services":["1"],
					"needs_terminal":false,
					"date_supply":"2023-08-10 07:55:12",
					"date_supply_until_soft":"2023-08-10 08:55:12",
					"date_supply_until_hard":"2023-08-10 09:55:12",
					"service_id":"1"
				 },
				 {
					"id":"10007",
					"orders":[
					   {
						  "id":"10007",
						  "shop_no":"201",
						  "pickup_point":{
							 "latitude":37.5,
							 "longitude":55.6,
							 "main":true
						  }
					   },
					   {
						  "id":"10008",
						  "shop_no":"201",
						  "pickup_point":{
							 "latitude":37.5,
							 "longitude":55.6,
							 "main":false
						  }
					   },
					   {
						  "id":"10009",
						  "shop_no":"201",
						  "pickup_point":{
							 "latitude":37.5,
							 "longitude":55.6,
							 "main":false
						  }
					   }

					],
					"delivery_point":{
					   "location":{
						  "latitude":37.6,
						  "longitude":55.5
					   }
					},
					"weight":3,
					"delivery_services":["1"],
					"needs_terminal":false,
					"date_supply":"2023-08-10 07:55:12",
					"date_supply_until_soft":"2023-08-10 08:55:12",
					"date_supply_until_hard":"2023-08-10 09:55:12",
					"service_id":"1"
				 }
				]
			 }`)

			var request model.Request

			err = json.Unmarshal(request_arr, &request)

			if err != nil {
				fmt.Printf("error: %s", err)
				return
			}
		*/

	}()

}
