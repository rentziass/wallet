package wallet

const (
	PKDataDetectorTypePhoneNumber   = "PKDataDetectorTypePhoneNumber"
	PKDataDetectorTypeLink          = "PKDataDetectorTypeLink"
	PKDataDetectorTypeAddress       = "PKDataDetectorTypeAddress"
	PKDataDetectorTypeCalendarEvent = "PKDataDetectorTypeCalendarEvent"

	PKDateStyleNone   = "PKDateStyleNone"
	PKDateStyleShort  = "PKDateStyleShort"
	PKDateStyleMedium = "PKDateStyleMedium"
	PKDateStyleLong   = "PKDateStyleLong"
	PKDateStyleFull   = "PKDateStyleFull"

	PKNumberStyleDecimal    = "PKNumberStyleDecimal"
	PKNumberStylePercent    = "PKNumberStylePercent"
	PKNumberStyleScientific = "PKNumberStyleScientific"
	PKNumberStyleSpellOut   = "PKNumberStyleSpellOut"

	PKTextAlignmentLeft    = "PKTextAlignmentLeft"
	PKTextAlignmentCenter  = "PKTextAlignmentCenter"
	PKTextAlignmentRight   = "PKTextAlignmentRight"
	PKTextAlignmentNatural = "PKTextAlignmentNatural"

	PKTransitTypeAir     = "PKTransitTypeAir"
	PKTransitTypeBoat    = "PKTransitTypeBoat"
	PKTransitTypeBus     = "PKTransitTypeBus"
	PKTransitTypeGeneric = "PKTransitTypeGeneric"
	PKTransitTypeTrain   = "PKTransitTypeTrain"
)

type Pass struct {
	// Refer to https://developer.apple.com/library/archive/documentation/UserExperience/Reference/PassKit_Bundle/Chapters/TopLevel.html#//apple_ref/doc/uid/TP40012026-CH2-SW1
	// for more information.

	// Standard keys
	// Information that is required for all passes.
	Description        string `json:"description"`
	FormatVersion      int    `json:"formatVersion"`
	OrganizationName   string `json:"organizationName"`
	PassTypeIdentifier string `json:"passTypeIdentifier"`
	SerialNumber       string `json:"serialNumber"`
	TeamIdentifier     string `json:"teamIdentifier"`

	// Associated App Keys
	// Information about an app that is associated with a pass.
	AppLaunchURL               string `json:"appLaunchURL,omitempty"`
	AssociatedStoreIdentifiers []int  `json:"associatedStoreIdentifiers,omitempty"`

	// Companion App Keys
	// Custom information about a pass provided for a companion app to use.
	UserInfo map[string]interface{} `json:"userInfo,omitempty"`

	// Expiration Keys
	ExpirationDate string `json:"expirationDate,omitempty"`
	Voided         *bool  `json:"voided,omitempty"`

	// Relevance Keys
	// Information about where and when a pass is relevant.
	Beacons      []*Beacon   `json:"beacons,omitempty"`
	Locations    []*Location `json:"locations,omitempty"`
	MaxDistance  int         `json:"maxDistance,omitempty"`
	RelevantDate string      `json:"relevantDate,omitempty"`

	// Style Keys
	// Keys that specify the pass style
	CouponDetails       *CouponDetails       `json:"coupon,omitempty"`
	EventTicketDetails  *EventTicketDetails  `json:"eventTicket,omitempty"`
	GenericDetails      *GenericDetails      `json:"generic,omitempty"`
	StoreCardDetails    *StoreCardDetails    `json:"storeCard,omitempty"`
	BoardingPassDetails *BoardingPassDetails `json:"boardingPass,omitempty"`

	// Visual Appearance Keys
	// Keys that define the visual style and appearance of the pass.
	Barcode            *PassBarcode   `json:"barcode,omitempty"`
	Barcodes           []*PassBarcode `json:"barcodes,omitempty"`
	BackgroundColor    string         `json:"backgroundColor,omitempty"`
	ForegroundColor    string         `json:"foregroundColor,omitempty"`
	GroupingIdentifier string         `json:"groupingIdentifier,omitempty"`
	LabelColor         string         `json:"labelColor,omitempty"`
	LogoText           string         `json:"logoText,omitempty"`
	SuppressStripShine *bool          `json:"suppressStripShine,omitempty"`

	// Web Service Keys
	// Information used to update passes using the web service.
	WebServiceURL       string `json:"webServiceURL,omitempty"`
	AuthenticationToken string `json:"authenticationToken,omitempty"`

	// NFC-Enabled Pass Keys
	// NFC-enabled pass keys support sending reward
	// card information as part of an Apple Pay transaction.
	NFC *NFC `json:"nfc,omitempty"`
}

type PassBarcode struct {
	AltText         string `json:"altText,omitempty"`
	Message         string `json:"message"`
	Format          string `json:"format"`
	MessageEncoding string `json:"messageEncoding"`
}

func NewPassBarcode(message string, format string) *PassBarcode {
	return &PassBarcode{
		Message:         message,
		Format:          format,
		MessageEncoding: "iso-8859-1",
	}
}

type Beacon struct {
	Major         int16  `json:"major,omitempty"`
	Minor         int16  `json:"minor,omitempty"`
	ProximityUUID string `json:"proximityUUID"`
	RelevantText  string `json:"relevantText,omitempty"`
}

type Location struct {
	Altitude     float32 `json:"altitude,omitempty"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
	RelevantText string  `json:"relevantText,omitempty"`
}

type NFC struct {
	Message             string `json:"message"`
	EncryptionPublicKey string `json:"encryptionPublicKey,empty"`
}

type EventTicketDetails struct {
	HeaderFields    []*Field `json:"headerFields,omitempty"`
	PrimaryFields   []*Field `json:"primaryFields,omitempty"`
	SecondaryFields []*Field `json:"secondaryFields,omitempty"`
	AuxiliaryFields []*Field `json:"auxiliaryFields,omitempty"`
	BackFields      []*Field `json:"backFields,omitempty"`
}

type CouponDetails struct {
	HeaderFields    []*Field `json:"headerFields,omitempty"`
	PrimaryFields   []*Field `json:"primaryFields,omitempty"`
	SecondaryFields []*Field `json:"secondaryFields,omitempty"`
	AuxiliaryFields []*Field `json:"auxiliaryFields,omitempty"`
	BackFields      []*Field `json:"backFields,omitempty"`
}

type GenericDetails struct {
	HeaderFields    []*Field `json:"headerFields,omitempty"`
	PrimaryFields   []*Field `json:"primaryFields,omitempty"`
	SecondaryFields []*Field `json:"secondaryFields,omitempty"`
	AuxiliaryFields []*Field `json:"auxiliaryFields,omitempty"`
	BackFields      []*Field `json:"backFields,omitempty"`
}

type StoreCardDetails struct {
	HeaderFields    []*Field `json:"headerFields,omitempty"`
	PrimaryFields   []*Field `json:"primaryFields,omitempty"`
	SecondaryFields []*Field `json:"secondaryFields,omitempty"`
	AuxiliaryFields []*Field `json:"auxiliaryFields,omitempty"`
	BackFields      []*Field `json:"backFields,omitempty"`
}

type BoardingPassDetails struct {
	HeaderFields    []*Field `json:"headerFields,omitempty"`
	PrimaryFields   []*Field `json:"primaryFields,omitempty"`
	SecondaryFields []*Field `json:"secondaryFields,omitempty"`
	AuxiliaryFields []*Field `json:"auxiliaryFields,omitempty"`
	BackFields      []*Field `json:"backFields,omitempty"`
	TransitType     string   `json:"transitType,omitempty"`
}

type Field struct {
	// Standard Field Dictionary Keys
	// Information about a field.
	// These keys are used for all dictionaries that define a field.
	AttributedValue   string   `json:"attributedValue,omitempty"`
	ChangeMessage     string   `json:"changeMessage,omitempty"`
	DataDetectorTypes []string `json:"dataDetectorTypes,omitempty"`
	Key               string   `json:"key"`
	Label             string   `json:"label,omitempty"`
	TextAlignment     string   `json:"textAlignment,optional"`
	Value             string   `json:"value"`

	// Date Style Keys
	// Information about how a date should be displayed in a field.
	DateStyle       string `json:"dateStyle,omitempty"`
	IgnoresTimeZone *bool  `json:"ignoresTimeZone,omitempty"`
	IsRelative      *bool  `json:"isRelative,omitempty"`
	TimeStyle       string `json:"timeStyle,omitempty"`

	// Number Style Keys
	// Information about how a number should be displayed in a field.
	CurrencyCode string `json:"currencyCode,omitempty"`
	NumberStyle  string `json:"numberStyle,omitempty"`
}
