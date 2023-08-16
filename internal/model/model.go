package model

type Request struct {
	TrafficJams string `json:"traffic_jams"`
	Couriers    []struct {
		ID               string   `json:"id"`
		TravelModes      string   `json:"travel_modes"`
		Speed            string   `json:"speed"`
		MaxOrders        string   `json:"max_orders"`
		MaxDistance      string   `json:"max_distance"`
		MaxWeight        string   `json:"max_weight"`
		QueueNumber      string   `json:"queue_number"`
		HasTerminal      string   `json:"has_terminal"`
		PickupTime       string   `json:"pickup_time"`
		HandOverTime     string   `json:"hand_over_time"`
		DeliveryServices []string `json:"delivery_services"`
	} `json:"couriers"`
	Parcels []struct {
		ID     string `json:"id"`
		Orders []struct {
			ID          string `json:"id"`
			ShopNo      string `json:"shop_no"`
			PickupPoint struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
				Main      string `json:"main"`
			} `json:"pickup_point"`
		} `json:"orders"`
		DeliveryPoint struct {
			Location struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"location"`
		} `json:"delivery_point"`
		Weight              string   `json:"weight"`
		DeliveryServices    []string `json:"delivery_services"`
		NeedsTerminal       string   `json:"needs_terminal"`
		DateSupply          string   `json:"date_supply"`
		DateSupplyUntilSoft string   `json:"date_supply_until_soft"`
		DateSupplyUntilHard string   `json:"date_supply_until_hard"`
		ServiceID           string   `json:"service_id"`
	} `json:"parcels"`
}

type Response struct {
	IDRequest string `json:"id_request"`
}

type Results struct {
	IDRequest string `json:"id_request"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	Data      struct {
		RouterName       string `json:"router_name"`
		DroppedLocations []struct {
			ParcelID string `json:"parcel_id"`
		} `json:"dropped_locations"`
		Routes []struct {
			CourierID             string `json:"courier_id"`
			DateShipTarget        string `json:"date_ship_target"`
			DateShipTargetExtreme string `json:"date_ship_target_extreme"`
			CalculationTime       string `json:"calculation_time"`
			TotalDistance         string `json:"total_distance"`
			TotalDuration         string `json:"total_duration"`
			EstimatedTime         string `json:"estimated_time"`
			Nodes                 []struct {
				NodeID    string `json:"node_id"`
				Action    string `json:"action"`
				ParcelID  string `json:"parcel_id"`
				OrderID   string `json:"order_id"`
				PointType string `json:"point_type"`
				Location  struct {
					Latitude  string `json:"latitude"`
					Longitude string `json:"longitude"`
				} `json:"location"`
				DistanceFromPrevious string `json:"distance_from_previous"`
				DurationFromPrevious string `json:"duration_from_previous"`
			} `json:"nodes"`
		} `json:"routes"`
	} `json:"data"`
}
