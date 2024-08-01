package pkg

const (
	// Advertisement statuses
	AdvertisementCreated             = 1  // On creation, probably unused
	AdvertisementPendingPayment      = 2  // After creation, waiting for user to pay
	AdvertisementPendingUpload       = 3  // After payment, waiting for media upload
	AdvertisementPendingApproval     = 4  // After upload, waiting for admin or whoever to approve it to SCAP guidelines
	AdvertisementPendingResubmission = 5  // On rejection, user is required to reupload the media file
	AdvertisementPendingDistribution = 10 // On approval, distribution to ETV machines is required
	AdvertisementDistributed         = 11 // On distribution, media *should* play
	AdvertisementFinished            = 12 // After the advertisement campaign has ended
	AdvertisementRefundRequested     = 89 // User requested refund
	AdvertisementRefunded            = 90 // Refund went through
	AdvertisementCancelled           = 99 // Advertisement cancelled due to various reasons

	// Service Pricing Model pricing types
	ServicePricingTypePerBooking = 1 // Price is static for the reservation
	ServicePricingTypePerMinute  = 2 // Price depends on how long the reservation is

	// Coupon discount types
	CouponDiscountFlat    = 1 // Flat fee discount
	CouponDiscountPercent = 2 // Percentage discount

	// Coupon statuses
	CouponDisabled = -1
	CouponEnabled  = 1

	// Payment statuses
	PaymentSucceeded = 1
	PaymentFailed    = -1

	// Possible payment types
	PaymentTypePayment       = 1
	PaymentTypePartialRefund = 2
	PaymentTypeFullRefund    = 3

	// Possible statuses for the `order` rows
	//
	OrderCreated           = 1  // Order just created. This status probably won't be used
	OrderPendingPayment    = 2  // Order created, awaiting payment
	OrderFailedPayment     = 3  // Payment attempted but failed
	OrderPaid              = 4  // Order has been paid for
	OrderPendingRefund     = 90 // User requested refund for order
	OrderRefunded          = 91 // Order has been refunded
	OrderPartiallyRefunded = 92 // Order has been partially refunded
	OrderRefundFailed      = 93 // Order refund request failed
	OrderCancelled         = 99 // Order cancelled

	// Possible statuses for the `reservation` rows
	//
	ReservationCreated                = 1  // Just created reservation row
	ReservationOrderFailed            = 2  // Failed to create order row
	ReservationPendingOrderSuccess    = 3  // Awaiting payment
	ReservationPendingRobotAssignment = 4  // Order paid, awaiting to assign job to robot
	ReservationPendingFindingLot      = 5  // Robot assigned, waiting for it to start finding a free parking space
	ReservationFindingLot             = 6  // Robot assigned and is finding free parking space
	ReservationAwaitingUser           = 7  // Robot found parking space, awaiting user to arrive
	ReservationAwaitingUserPark       = 8  // Robot moved aside for user to park, waiting for parking confirmation
	ReservationPendingUserCharge      = 9  // If user arrived and has charging booked, wait for QR code scan
	ReservationCharging               = 10 // QR code scanned, charging active for duration booked
	ReservationReturnGun              = 11 // Charging completed, awaiting user to return charging gun
	ReservationRobotOnTheWay          = 20 // If booked charging only, robot is now on the way to the station lot
	ReservationCompleted              = 50 // Robot has finished its duties
	ReservationReviewed               = 51 // User has made a review
	ReservationFailed                 = 90 // Something unexpected happened and caused the reservation to fail
	ReservationCancelledByUser        = 98 // User cancelled reservation
	ReservationCancelledBySystem      = 99 // Fatal error caught that caused reservation to be cancelled

	// Task IDs to be sent to the robutt from the microservice
	//
	RobotTaskFindParking   = 1  // Instruct robot to find a parking space
	RobotTaskUserArrived   = 2  // Instruct robot that user has arrived at parking space (Move out of the way)
	RobotTaskUserParked    = 3  // Instruct robot that user has parked in the parking space (Move close to the car)
	RobotTaskUnlockGun     = 4  // Instruct robot to unlock charging gun
	RobotTaskBeginCharging = 5  // Instruct robot to allow charging
	RobotTaskStopCharging  = 6  // Instruct robot to stop charging
	RobotTaskReturn        = 7  // Instruct robot to return to holding area
	RobotTaskGenQRCode     = 10 // Instruct robot to generate QR code to begin charging

	// Statuses expected to receive from the robutt
	//
	RobotDisabled         = -1 // Robot indicates it is disabled (Probably unused)
	RobotAvailable        = 1  // Robot indicates it is available for jobs
	RobotAssigned         = 2  // Robot indicates it has been assigned a reservation and is pending action
	RobotFindingParking   = 3  // Robot indicates it is currently finding a parking space
	RobotParked           = 4  // Robot indicates it has found a parking space and is parked
	RobotMakingWayParking = 5  // Robot indicates it is making way for the user to park
	RobotMadeWay          = 6  // Robot indicates it has made way for user to park
	RobotMovingClose      = 7  // Robot indicates that it is moving close to the user's vehicle for charging
	RobotReadyForCharging = 8  // Robot indicates that it is near the user's vehicle and is awaiting QR code scan
	RobotCharging         = 9  // Robot indicates it is currently charging the vehicle
	RobotFinishedCharging = 10 // Robot indicates it has finished charging
	RobotReturning        = 11 // Robot indicates it is currently returning to the holding area
	RobotMaintenance      = 90 // Robot indicates it is under maintenance (Probably unused)
	RobotDisconnected     = 98 // Robot unexpectedly disconnected from MQTT. Used in the last will payload

	// MQTT Topic types
	MQTTTopicTypeIdle         = 0 // Empty status
	MQTTTopicTypeRobotStatus  = 1 // Topic is used for robot status updates
	MQTTTopicTypeRobotMetrics = 2 // Topic is used for robot metrics

	// Push notification types
	PushNotificationToken = 1 // Send notification to a FCM token/user
	PushNotificationTopic = 2 // Send notification to a user subscribed topic

	OrderTypeBooking       = "booking"
	OrderTypeAdvertisement = "advertisement"

	ServiceTypeIDParking         = 1
	ServiceTypeIDCharging        = 2
	ServiceTypeIDParkingCharging = 3
	ServiceTypeIDAdvertisement   = 4
)
