package pkg

type RobotUpdate struct {
	ReservationID     int32 `json:"reservation_id"`
	RobotID           int32 `json:"robot_id"`
	ReservationStatus int32 `json:"reservation_status"`
}
