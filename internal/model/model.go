package model

type Request struct {
	TrafficJams bool `json:"traffic_jams"`
	Couriers    []struct {
		ID               string   `json:"id"`
		TravelModes      string   `json:"travel_modes"`
		Speed            float64  `json:"speed"`
		MaxOrders        int      `json:"max_orders"`
		MaxDistance      float64  `json:"max_distance"`
		MaxWeight        float64  `json:"max_weight"`
		QueueNumber      int      `json:"queue_number"`
		HasTerminal      bool     `json:"has_terminal"`
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
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Main      bool    `json:"main"`
			} `json:"pickup_point"`
		} `json:"orders"`
		DeliveryPoint struct {
			Location struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"location"`
		} `json:"delivery_point"`
		Weight              float64  `json:"weight"`
		DeliveryServices    []string `json:"delivery_services"`
		NeedsTerminal       bool     `json:"needs_terminal"`
		DateSupply          string   `json:"date_supply"`
		DateSupplyUntilSoft string   `json:"date_supply_until_soft"`
		DateSupplyUntilHard string   `json:"date_supply_until_hard"`
		ServiceID           string   `json:"service_id"`
	} `json:"parcels"`
}

type Response struct {
	IDRequest string `json:"id_request"`
}

type ResultRouteNode struct {
	NodeID    int    `json:"node_id"`
	Action    string `json:"action"`
	ParcelID  string `json:"parcel_id"`
	OrderID   string `json:"order_id"`
	PointType string `json:"point_type"`
	Location  struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	DistanceFromPrevious float64 `json:"distance_from_previous"`
	DurationFromPrevious string  `json:"duration_from_previous"`
}

type ResultRoute struct {
	CourierID             string            `json:"courier_id"`
	DateShipTarget        string            `json:"date_ship_target"`
	DateShipTargetExtreme string            `json:"date_ship_target_extreme"`
	CalculationTime       string            `json:"calculation_time"`
	TotalDistance         float64           `json:"total_distance"`
	TotalDuration         float64           `json:"total_duration"`
	EstimatedTime         string            `json:"estimated_time"`
	Nodes                 []ResultRouteNode `json:"nodes"`
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
		Routes []ResultRoute `json:"routes"`
	} `json:"data"`
}
