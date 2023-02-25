package dao

type OrderStatus int32

const (
	UNDEFINED OrderStatus = 0
	INITATED  OrderStatus = 1
	CREATED   OrderStatus = 2
	FAILED    OrderStatus = 3
)

type Order struct {
	ID        string      `json:"id"`
	TotalCost uint32      `json:"totalCost"`
	VehicleNo string      `json:"vehicleNo"`
	Status    OrderStatus `json:"status"`
}
