package pkg

type RobotUpdate struct {
	ReservationID     int32 `json:"reservation_id"`
	RobotID           int32 `json:"robot_id"`
	ReservationStatus int32 `json:"reservation_status"`
}

type QRCodeMessage struct {
	TaskID        int32  `json:"task_id"`
	QRCodeContent string `json:"qr_code_content"`
}

type ToRobotCommandMessage struct {
	Command string `json:"command"`
	Data    []byte `json:"data_b64"` // Has b64 suffix because it's converted to base64 in json.Marshal()
}

type ToRobotMQTTMessage struct {
	TaskID  int32  `json:"task_id"`
	Payload []byte `json:"payload_b64"`
}

type FromRobotMQTTMessage struct {
	Status  int32  `json:"status"`
	Payload []byte `json:"payload"`
}
