package models

// ──────────────────────────────────────────────
// USER STATUS & ROLE
// ──────────────────────────────────────────────

// UserStatus replaces the separate is_active / is_blocked flags.
type UserStatus string

const (
	UserStatusPending UserStatus = "PENDING" // registered, not yet verified/approved
	UserStatusActive  UserStatus = "ACTIVE"  // fully active
	UserStatusBlocked UserStatus = "BLOCKED" // admin-blocked
	UserStatusDeleted UserStatus = "DELETED" // soft-deleted
)

type Role string

const (
	RoleSuperAdmin Role = "SUPER_ADMIN"
	RoleAdmin      Role = "ADMIN"
	RoleRider      Role = "RIDER"
	RoleMerchant   Role = "MERCHANT"
)

// ──────────────────────────────────────────────
// DIVISION  (8 divisions of Bangladesh)
// ──────────────────────────────────────────────

type Division string

const (
	DivisionDhaka      Division = "Dhaka"
	DivisionChattogram Division = "Chattogram"
	DivisionRajshahi   Division = "Rajshahi"
	DivisionKhulna     Division = "Khulna"
	DivisionBarishal   Division = "Barishal"
	DivisionSylhet     Division = "Sylhet"
	DivisionRangpur    Division = "Rangpur"
	DivisionMymensingh Division = "Mymensingh"
)

// ──────────────────────────────────────────────
// DISTRICT  (64 districts of Bangladesh)
// ──────────────────────────────────────────────

type District string

const (
	// Dhaka Division
	DistrictDhaka       District = "Dhaka"
	DistrictGazipur     District = "Gazipur"
	DistrictNarayanganj District = "Narayanganj"
	DistrictNarsingdi   District = "Narsingdi"
	DistrictManikganj   District = "Manikganj"
	DistrictMunshiganj  District = "Munshiganj"
	DistrictRajbari     District = "Rajbari"
	DistrictFaridpur    District = "Faridpur"
	DistrictShariatpur  District = "Shariatpur"
	DistrictMadaripur   District = "Madaripur"
	DistrictGopalganj   District = "Gopalganj"
	DistrictTangail     District = "Tangail"
	DistrictKishoreganj District = "Kishoreganj"

	// Chattogram Division
	DistrictChattogram   District = "Chattogram"
	DistrictCoxsBazar    District = "Cox's Bazar"
	DistrictBandarban    District = "Bandarban"
	DistrictRangamati    District = "Rangamati"
	DistrictKhagrachhari District = "Khagrachhari"
	DistrictFeni         District = "Feni"
	DistrictNoakhali     District = "Noakhali"
	DistrictLakshmipur   District = "Lakshmipur"
	DistrictComilla      District = "Comilla"
	DistrictChandpur     District = "Chandpur"
	DistrictBrahmanbaria District = "Brahmanbaria"

	// Rajshahi Division
	DistrictRajshahi        District = "Rajshahi"
	DistrictNatore          District = "Natore"
	DistrictNaogaon         District = "Naogaon"
	DistrictChapaiNawabganj District = "Chapai Nawabganj"
	DistrictPabna           District = "Pabna"
	DistrictSirajganj       District = "Sirajganj"
	DistrictBogura          District = "Bogura"
	DistrictJoypurhat       District = "Joypurhat"

	// Khulna Division
	DistrictKhulna    District = "Khulna"
	DistrictBagerhat  District = "Bagerhat"
	DistrictSatkhira  District = "Satkhira"
	DistrictJessore   District = "Jessore"
	DistrictNarail    District = "Narail"
	DistrictMagura    District = "Magura"
	DistrictJhenaidah District = "Jhenaidah"
	DistrictKushtia   District = "Kushtia"
	DistrictChuadanga District = "Chuadanga"
	DistrictMeherpur  District = "Meherpur"

	// Barishal Division
	DistrictBarishal   District = "Barishal"
	DistrictPatuakhali District = "Patuakhali"
	DistrictBhola      District = "Bhola"
	DistrictPirojpur   District = "Pirojpur"
	DistrictJhalokati  District = "Jhalokati"
	DistrictBarguna    District = "Barguna"

	// Sylhet Division
	DistrictSylhet      District = "Sylhet"
	DistrictMoulvibazar District = "Moulvibazar"
	DistrictHabiganj    District = "Habiganj"
	DistrictSunamganj   District = "Sunamganj"

	// Rangpur Division
	DistrictRangpur     District = "Rangpur"
	DistrictDinajpur    District = "Dinajpur"
	DistrictThakurgaon  District = "Thakurgaon"
	DistrictPanchagarh  District = "Panchagarh"
	DistrictNilphamari  District = "Nilphamari"
	DistrictLalmonirhat District = "Lalmonirhat"
	DistrictKurigram    District = "Kurigram"
	DistrictGaibandha   District = "Gaibandha"

	// Mymensingh Division
	DistrictMymensingh District = "Mymensingh"
	DistrictJamalpur   District = "Jamalpur"
	DistrictSherpur    District = "Sherpur"
	DistrictNetrokona  District = "Netrokona"
)

// DistrictDivisionMap maps every district to its division.
// Use DivisionOf(district) for safe lookups.
var DistrictDivisionMap = map[District]Division{
	// Dhaka Division (13)
	DistrictDhaka:       DivisionDhaka,
	DistrictGazipur:     DivisionDhaka,
	DistrictNarayanganj: DivisionDhaka,
	DistrictNarsingdi:   DivisionDhaka,
	DistrictManikganj:   DivisionDhaka,
	DistrictMunshiganj:  DivisionDhaka,
	DistrictRajbari:     DivisionDhaka,
	DistrictFaridpur:    DivisionDhaka,
	DistrictShariatpur:  DivisionDhaka,
	DistrictMadaripur:   DivisionDhaka,
	DistrictGopalganj:   DivisionDhaka,
	DistrictTangail:     DivisionDhaka,
	DistrictKishoreganj: DivisionDhaka,
	// Chattogram Division (11)
	DistrictChattogram:   DivisionChattogram,
	DistrictCoxsBazar:    DivisionChattogram,
	DistrictBandarban:    DivisionChattogram,
	DistrictRangamati:    DivisionChattogram,
	DistrictKhagrachhari: DivisionChattogram,
	DistrictFeni:         DivisionChattogram,
	DistrictNoakhali:     DivisionChattogram,
	DistrictLakshmipur:   DivisionChattogram,
	DistrictComilla:      DivisionChattogram,
	DistrictChandpur:     DivisionChattogram,
	DistrictBrahmanbaria: DivisionChattogram,
	// Rajshahi Division (8)
	DistrictRajshahi:        DivisionRajshahi,
	DistrictNatore:          DivisionRajshahi,
	DistrictNaogaon:         DivisionRajshahi,
	DistrictChapaiNawabganj: DivisionRajshahi,
	DistrictPabna:           DivisionRajshahi,
	DistrictSirajganj:       DivisionRajshahi,
	DistrictBogura:          DivisionRajshahi,
	DistrictJoypurhat:       DivisionRajshahi,
	// Khulna Division (10)
	DistrictKhulna:    DivisionKhulna,
	DistrictBagerhat:  DivisionKhulna,
	DistrictSatkhira:  DivisionKhulna,
	DistrictJessore:   DivisionKhulna,
	DistrictNarail:    DivisionKhulna,
	DistrictMagura:    DivisionKhulna,
	DistrictJhenaidah: DivisionKhulna,
	DistrictKushtia:   DivisionKhulna,
	DistrictChuadanga: DivisionKhulna,
	DistrictMeherpur:  DivisionKhulna,
	// Barishal Division (6)
	DistrictBarishal:   DivisionBarishal,
	DistrictPatuakhali: DivisionBarishal,
	DistrictBhola:      DivisionBarishal,
	DistrictPirojpur:   DivisionBarishal,
	DistrictJhalokati:  DivisionBarishal,
	DistrictBarguna:    DivisionBarishal,
	// Sylhet Division (4)
	DistrictSylhet:      DivisionSylhet,
	DistrictMoulvibazar: DivisionSylhet,
	DistrictHabiganj:    DivisionSylhet,
	DistrictSunamganj:   DivisionSylhet,
	// Rangpur Division (8)
	DistrictRangpur:     DivisionRangpur,
	DistrictDinajpur:    DivisionRangpur,
	DistrictThakurgaon:  DivisionRangpur,
	DistrictPanchagarh:  DivisionRangpur,
	DistrictNilphamari:  DivisionRangpur,
	DistrictLalmonirhat: DivisionRangpur,
	DistrictKurigram:    DivisionRangpur,
	DistrictGaibandha:   DivisionRangpur,
	// Mymensingh Division (4)
	DistrictMymensingh: DivisionMymensingh,
	DistrictJamalpur:   DivisionMymensingh,
	DistrictSherpur:    DivisionMymensingh,
	DistrictNetrokona:  DivisionMymensingh,
}

// DivisionOf returns the Division for a given District.
// The second return value is false if the district is unknown.
func DivisionOf(d District) (Division, bool) {
	div, ok := DistrictDivisionMap[d]
	return div, ok
}

// ──────────────────────────────────────────────
// SENDER TYPE / CONFIRMATION
// ──────────────────────────────────────────────

type SenderType string

const (
	SenderTypeGuest    SenderType = "GUEST"
	SenderTypeMerchant SenderType = "MERCHANT"
)

type ConfirmationMethod string

const (
	ConfirmationMethodOTP          ConfirmationMethod = "OTP"
	ConfirmationMethodCashHandover ConfirmationMethod = "CASH_HANDOVER"
)

// ──────────────────────────────────────────────
// PERSONAL INFO
// ──────────────────────────────────────────────

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
	GenderOther  Gender = "OTHER"
)

type BloodGroup string

const (
	BloodAPos  BloodGroup = "A_POS"
	BloodANeg  BloodGroup = "A_NEG"
	BloodBPos  BloodGroup = "B_POS"
	BloodBNeg  BloodGroup = "B_NEG"
	BloodABPos BloodGroup = "AB_POS"
	BloodABNeg BloodGroup = "AB_NEG"
	BloodOPos  BloodGroup = "O_POS"
	BloodONeg  BloodGroup = "O_NEG"
)

// ──────────────────────────────────────────────
// KYC & VEHICLE
// ──────────────────────────────────────────────

type RiderKycStatus string

const (
	RiderKycPendingReview      RiderKycStatus = "PENDING"
	RiderKycDocumentsRequested RiderKycStatus = "REQUESTED"
	RiderKycApproved           RiderKycStatus = "APPROVED"
	RiderKycRejected           RiderKycStatus = "REJECTED"
	RiderKycSuspended          RiderKycStatus = "SUSPENDED"
)

type VehicleCategory string

const (
	VehicleTwoWheeler   VehicleCategory = "TWO_WHEELER"
	VehicleThreeWheeler VehicleCategory = "THREE_WHEELER"
	VehicleFourWheeler  VehicleCategory = "FOUR_WHEELER"
)

type VehicleSubType string

const (
	VehicleBicycle      VehicleSubType = "BICYCLE"
	VehicleMotorcycle   VehicleSubType = "MOTORCYCLE"
	VehicleElectricBike VehicleSubType = "ELECTRIC_BIKE"
	VehicleAutoRickshaw VehicleSubType = "AUTO_RICKSHAW"
	VehicleCNGAuto      VehicleSubType = "CNG_AUTO"
	VehicleElectricAuto VehicleSubType = "ELECTRIC_AUTO"
	VehicleCar          VehicleSubType = "CAR"
	VehiclePickupVan    VehicleSubType = "PICKUP_VAN"
)

// ──────────────────────────────────────────────
// SHIPMENT
// ──────────────────────────────────────────────

type ShipmentStatus string

const (
	ShipmentPending                    ShipmentStatus = "PENDING"
	ShipmentPickupRequested            ShipmentStatus = "PICKUP_REQUESTED"
	ShipmentPickupRiderAssigned        ShipmentStatus = "PICKUP_RIDER_ASSIGNED"
	ShipmentPickedUp                   ShipmentStatus = "PICKED_UP"
	ShipmentAtMerchantHub              ShipmentStatus = "AT_MERCHANT_HUB"
	ShipmentReceivedAtHub              ShipmentStatus = "RECEIVED_AT_HUB"
	ShipmentInTransit                  ShipmentStatus = "IN_TRANSIT"
	ShipmentAtDeliveryHub              ShipmentStatus = "AT_DELIVERY_HUB"
	ShipmentOutForDelivery             ShipmentStatus = "OUT_FOR_DELIVERY"
	ShipmentDelivered                  ShipmentStatus = "DELIVERED"
	ShipmentFailedDelivery             ShipmentStatus = "FAILED_DELIVERY"
	ShipmentReturnInitiated            ShipmentStatus = "RETURN_INITIATED"
	ShipmentReturnInTransit            ShipmentStatus = "RETURN_IN_TRANSIT"
	ShipmentReturnedToHub              ShipmentStatus = "RETURNED_TO_HUB"
	ShipmentReturnToMerchantHub        ShipmentStatus = "RETURN_TO_MERCHANT_HUB"
	ShipmentAwaitingMerchantCollection ShipmentStatus = "AWAITING_MERCHANT_COLLECTION"
	ShipmentReturnedToMerchant         ShipmentStatus = "RETURNED_TO_MERCHANT"
	ShipmentCancelled                  ShipmentStatus = "CANCELLED"
)

type ParcelType string

const (
	ParcelDocument    ParcelType = "DOCUMENT"
	ParcelPackage     ParcelType = "PACKAGE"
	ParcelFragile     ParcelType = "FRAGILE"
	ParcelElectronics ParcelType = "ELECTRONICS"
	ParcelFood        ParcelType = "FOOD"
	ParcelHeavy       ParcelType = "HEAVY"
	ParcelOther       ParcelType = "OTHER"
)

// ──────────────────────────────────────────────
// PRICING / ZONE / SPLIT
// ──────────────────────────────────────────────

type SplitZone string

const (
	SplitSameCity        SplitZone = "SAME_CITY"
	SplitSameDistrict    SplitZone = "SAME_DISTRICT"
	SplitOutsideDistrict SplitZone = "OUTSIDE_DISTRICT"
)

// ZoneType is used by the pricing engine to determine delivery charge.
type ZoneType string

const (
	ZoneInsideDhaka            ZoneType = "INSIDE_DHAKA"
	ZoneDhakaSuburb            ZoneType = "DHAKA_SUBURB"
	ZoneOutsideDhakaDivisional ZoneType = "OUTSIDE_DHAKA_DIVISIONAL"
	ZoneOutsideDhakaDistrict   ZoneType = "OUTSIDE_DHAKA_DISTRICT"
	ZoneOutsideDhakaUpazila    ZoneType = "OUTSIDE_DHAKA_UPAZILA"
	ZoneSameCityNonDhaka       ZoneType = "SAME_CITY_NON_DHAKA"
)

// ──────────────────────────────────────────────
// PAYMENT
// ──────────────────────────────────────────────

type PaymentMethod string

const (
	PaymentCash         PaymentMethod = "CASH"
	PaymentBkash        PaymentMethod = "BKASH"
	PaymentNagad        PaymentMethod = "NAGAD"
	PaymentRocket       PaymentMethod = "ROCKET"
	PaymentBankTransfer PaymentMethod = "BANK_TRANSFER"
	PaymentHubPickup    PaymentMethod = "HUB_PICKUP"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "PENDING"
	PaymentCompleted PaymentStatus = "COMPLETED"
	PaymentFailed    PaymentStatus = "FAILED"
	PaymentRefunded  PaymentStatus = "REFUNDED"
)

type WalkInPaymentStatus string

const (
	WalkInCollected WalkInPaymentStatus = "COLLECTED"
	WalkInRefunded  WalkInPaymentStatus = "REFUNDED"
)

// ──────────────────────────────────────────────
// FRAUD / RISK
// ──────────────────────────────────────────────

// RiskLevel for receiver fraud score system.
type RiskLevel string

const (
	RiskLevelNewCustomer RiskLevel = "NEW_CUSTOMER"
	RiskLevelTrusted     RiskLevel = "TRUSTED"
	RiskLevelLow         RiskLevel = "LOW_RISK"
	RiskLevelMedium      RiskLevel = "MEDIUM_RISK"
	RiskLevelHigh        RiskLevel = "HIGH_RISK"
	RiskLevelBlacklisted RiskLevel = "BLACKLISTED"
)

// ──────────────────────────────────────────────
// WITHDRAWAL / CASHOUT
// ──────────────────────────────────────────────

// WithdrawalStatus for rider/merchant cashout requests.
type WithdrawalStatus string

const (
	WithdrawalRequested  WithdrawalStatus = "REQUESTED"
	WithdrawalApproved   WithdrawalStatus = "APPROVED"
	WithdrawalProcessing WithdrawalStatus = "PROCESSING"
	WithdrawalPaid       WithdrawalStatus = "PAID"
	WithdrawalRejected   WithdrawalStatus = "REJECTED"
)

// DepositStatus for COD cash deposit by rider at hub.
type DepositStatus string

const (
	DepositPending   DepositStatus = "PENDING"
	DepositSubmitted DepositStatus = "SUBMITTED"
	DepositConfirmed DepositStatus = "CONFIRMED"
)

// RevenueStatus for system revenue records.
type RevenueStatus string

const (
	RevenuePending   RevenueStatus = "PENDING"
	RevenueConfirmed RevenueStatus = "CONFIRMED"
)

// ──────────────────────────────────────────────
// PICKUP / RETURN  (Phase 2)
// ──────────────────────────────────────────────

// PickupType: merchant can drop at hub or request home pickup (Phase 2).
type PickupType string

const (
	PickupTypeHubDropOff PickupType = "HUB_DROP_OFF"
	PickupTypeHomePickup PickupType = "HOME_PICKUP"
)

// ReturnDeliveryType: how merchant wants their failed parcel returned (Phase 2).
type ReturnDeliveryType string

const (
	ReturnHubSelfCollect ReturnDeliveryType = "HUB_SELF_COLLECT"
	ReturnHomeRedelivery ReturnDeliveryType = "HOME_REDELIVERY"
)

// ReturnChargeStatus for return charge deduction tracking (Phase 2).
type ReturnChargeStatus string

const (
	ReturnChargePending  ReturnChargeStatus = "PENDING"
	ReturnChargeDeducted ReturnChargeStatus = "DEDUCTED"
	ReturnChargeWaived   ReturnChargeStatus = "WAIVED"
	ReturnChargeBlocked  ReturnChargeStatus = "BLOCKED"
)

// ReturnScenario classifies why a return happened.
type ReturnScenario string

const (
	ReturnScenarioCustomerPaid    ReturnScenario = "CUSTOMER_PAID"
	ReturnScenarioMerchantCharged ReturnScenario = "MERCHANT_CHARGED"
	ReturnScenarioWalkinReturn    ReturnScenario = "WALKIN_RETURN"
)

// ──────────────────────────────────────────────
// RATINGS
// ──────────────────────────────────────────────

// RiderRatingBadge reflects overall rider quality.
type RiderRatingBadge string

const (
	RiderBadgeNewRider    RiderRatingBadge = "NEW_RIDER"
	RiderBadgeTopRider    RiderRatingBadge = "TOP_RIDER"
	RiderBadgeGood        RiderRatingBadge = "GOOD"
	RiderBadgeAverage     RiderRatingBadge = "AVERAGE"
	RiderBadgeUnderReview RiderRatingBadge = "UNDER_REVIEW"
)

// MerchantRatingTag for Phase 2 merchant feedback.
type MerchantRatingTag string

const (
	MerchantTagLateDelivery       MerchantRatingTag = "LATE_DELIVERY"
	MerchantTagDamagedParcel      MerchantRatingTag = "DAMAGED_PARCEL"
	MerchantTagWrongArea          MerchantRatingTag = "WRONG_AREA"
	MerchantTagGoodService        MerchantRatingTag = "GOOD_SERVICE"
	MerchantTagCommunicationIssue MerchantRatingTag = "COMMUNICATION_ISSUE"
	MerchantTagOther              MerchantRatingTag = "OTHER"
)
