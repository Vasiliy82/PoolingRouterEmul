package main

import (
	"encoding/json"
	"fmt"

	"github.com/PoolingRouterEmul/internal/model"
)

func main() {

	request_arr := []byte(`{
		"traffic_jams":"true",
		"couriers":[
		   {
			  "id":"101",
			  "travel_modes":"driving",
			  "speed":"10",
			  "max_orders":"10",
			  "max_distance":"1000",
			  "max_weight":"10",
			  "queue_number":"1",
			  "has_terminal":"false",
			  "pickup_time":"10",
			  "hand_over_time":"10",
			  "delivery_services":["1", "2"]
		   },
		   {
			"id":"102",
			"travel_modes":"walking",
			"speed":"5",
			"max_orders":"10",
			"max_distance":"1000",
			"max_weight":"10",
			"queue_number":"1",
			"has_terminal":"true",
			"pickup_time":"10",
			"hand_over_time":"10",
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
					   "latitude":"37.5",
					   "longitude":"55.6",
					   "main":"true"
					}
				 },
				 {
					"id":"10002",
					"shop_no":"201",
					"pickup_point":{
					   "latitude":"37.5",
					   "longitude":"55.6",
					   "main":"false"
					}
				 },
				 {
					"id":"10003",
					"shop_no":"201",
					"pickup_point":{
					   "latitude":"37.5",
					   "longitude":"55.6",
					   "main":"false"
					}
				 }

			  ],
			  "delivery_point":{
				 "location":{
					"latitude":"37.6",
					"longitude":"55.5"
				 }
			  },
			  "weight":"3",
			  "delivery_services":["1"],
			  "needs_terminal":"false",
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

	fmt.Printf("%#v", request)

}
