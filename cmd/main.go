package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/PoolingRouterEmul/internal/model"
)

func main() {

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

	err := json.Unmarshal(request_arr, &request)

	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	results := model.Results{
		Status:  "Ok",
		Message: "Ok",
	}

	droppedLocations := map[string]any{}
	for _, p := range request.Parcels {
		droppedLocations[p.ID] = p
	}

	fmt.Printf("%#v", request)

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
	fmt.Println("Ready")
	fmt.Printf("%#v\n", results)
}
