package pkg

const (
	OrderCreated           = 1  // Order just created. This status probably won't be used
	OrderPendingPayment    = 2  // Order created, awaiting payment
	OrderFailedPayment     = 3  // Payment attempted but failed
	OrderPaid              = 4  // Order has been paid for
	OrderPendingRefund     = 90 // User requested refund for order
	OrderRefunded          = 91 // Order has been refunded
	OrderPartiallyRefunded = 92 // Order has been partially refunded
	OrderRefundFailed      = 93 // Order refund request failed
	OrderCancalled         = 99 // Order cancelled

	ReservationCreated                     = 1  // Just created reservation row
	ReservationOrderFailed                 = 2  // Failed to create order row
	ReservationPendingOrderSuccess         = 3  // Awaiting payment
	ReservationPendingRobotAssignment      = 4  // Order paid, awaiting to assign job to robot
	ReservationFindingLot                  = 5  // Robot assigned and is finding free parking space
	ReservationAwaitingUser                = 6  // Robot found parking space, awaiting user to arrive
	ReservationPendingChargingConfirmation = 7  // If user arrived and has charging booked, wait for QR code scan
	ReservationCharging                    = 8  // QR code scanned, charging active for duration booked
	ReservationCompleted                   = 9  // Robot has finished its duties
	ReservationReviewed                    = 10 // User has made a review
	ReservationFailed                      = 90 // Something unexpected happened and caused the reservation to fail
	ReservationCancelledByUser             = 98 // User cancelled reservation
	ReservationCancelledBySystem           = 99 // Fatal error caught that caused reservation to be cancelled

	RobotTaskFindParking   = 1  // Instruct robot to find a parking space
	RobotTaskUserArrived   = 2  // Instruct robot that user has arrived at parking space
	RobotTaskBeginCharging = 3  // Instruct robot to allow charging
	RobotTaskStopCharging  = 4  // Instruct robot to stop charging
	RobotTaskReturn        = 5  // Instruct robot to return to holding area
	RobotTaskGenQRCode     = 10 // Instruct robot to generate QR code to begin charging

	RobotDisabled         = 0 // Robot indicates it is disabled (Probably unused)
	RobotAvailable        = 1 // Robot indicates it is available for jobs
	RobotMaintenance      = 2 // Robot indicates it is under maintenance (Probably unused)
	RobotFindingParking   = 3 // Robot indicates it is currently finding a parking space
	RobotParked           = 4 // Robot indicates it has found a parking space and is parked
	RobotCharging         = 5 // Robot indicates it is currently charging the vehicle
	RobotFinishedCharging = 6 // Robot indicates it has finished charging
	RobotReturning        = 7 // Robot indicates it is currently returning to the holding area
)
