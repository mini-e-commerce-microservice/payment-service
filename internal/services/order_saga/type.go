package order_saga

type PayloadCreateOrderProductOrderItem struct {
	ID                int64   `db:"id" json:"id"`
	Name              string  `db:"name" json:"name"`
	Description       string  `db:"description" json:"description"`
	OrderID           int64   `db:"order_id" json:"order_id"`
	ProductItemID     int64   `db:"product_item_id" json:"product_item_id"`
	Qty               int32   `db:"qty" json:"qty"`
	UnitPrice         float64 `db:"unit_price" json:"unit_price"`
	TotalPrice        float64 `db:"total_price" json:"total_price"`
	Discount          float64 `db:"discount" json:"discount"`
	Weight            int32   `db:"weight" json:"weight"`
	PackageLength     int32   `db:"package_length" json:"package_length"`
	PackageWidth      int32   `db:"package_width" json:"package_width"`
	PackageHeight     int32   `db:"package_height" json:"package_height"`
	DimensionalWeight float64 `db:"dimensional_weight" json:"dimensional_weight"`
}

type PayloadCreateOrderProduct struct {
	OrderID           int64                                `json:"order_id"`
	UserID            int64                                `json:"user_id"`
	Courier           PayloadCreateOrderProductCourier     `json:"courier"`
	Origin            PayloadCreateOrderProductLocation    `json:"origin"`
	Destination       PayloadCreateOrderProductLocation    `json:"destination"`
	TotalAmount       float64                              `json:"total_amount"`
	PaymentMethodCode string                               `json:"payment_method_code"`
	Items             []PayloadCreateOrderProductOrderItem `json:"items"`
}

type PayloadCreateOrderProductLocation struct {
	LocationID string  `json:"location_id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	PostalCode int32   `json:"postal_code"`
	Address    string  `json:"address"`
}

type PayloadCreateOrderProductCourier struct {
	CourierCode        string `json:"courier_code"`
	CourierServiceCode string `json:"courier_service_code"`
	CourierCompany     string `json:"courier_company"`
	CourierType        string `json:"courier_type"`
	DeliveryType       string `json:"delivery_type"`
	DeliveryDate       string `json:"delivery_date"`
}
