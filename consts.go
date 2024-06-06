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

	RobotDisabled    = 0 // Robot is unusable
	RobotEnabled     = 1 // Robot is ready for jobs
	RobotBusy        = 2 // Robot is assigned a job
	RobotMaintenance = 3 // Robot is down for maintenance
)
