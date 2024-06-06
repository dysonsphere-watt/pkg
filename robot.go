package pkg

type RobotUpdate struct {
	ReservationID     int32 `json:"reservation_id"`
	RobotID           int32 `json:"robot_id"`
	ReservationStatus int32 `json:"reservation_status"`
}

type ToRobotMQTTMessage struct {
	TaskID  int32  `json:"task_id"`
	Payload []byte `json:"payload"`
}

type FromRobotMQTTMessage struct {
	Status  int32  `json:"status"`
	Payload []byte `json:"payload"`
}
