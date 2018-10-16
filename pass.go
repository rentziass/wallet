package wallet

type Pass struct {
	FormatVersion       int                 `json:"formatVersion,omitempty"`
	PassTypeIdentifier  string              `json:"passTypeIdentifier,omitempty"`
	SerialNumber        string              `json:"serialNumber,omitempty"`
	TeamIdentifier      string              `json:"teamIdentifier,omitempty"`
	WebServiceURL       string              `json:"webServiceURL,omitempty"`
	AuthenticationToken string              `json:"authenticationToken,omitempty"`
	RelevantDate        string              `json:"relevantDate,omitempty"`
	Barcode             *PassBarcode        `json:"barcode,omitempty"`
	OrganizationName    string              `json:"organizationName,omitempty"`
	Description         string              `json:"description,omitempty"`
	ForegroundColor     string              `json:"foregroundColor,omitempty"`
	BackgroundColor     string              `json:"backgroundColor,omitempty"`
	LogoText            string              `json:"logoText,omitempty"`
	LabelColor          string              `json:"labelColor,omitempty"`
	EventTicketDetails  *EventTicketDetails `json:"eventTicket,omitempty"`
}

type PassBarcode struct {
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

type EventTicketDetails struct {
	PrimaryFields   []*Field `json:"primaryFields,omitempty"`
	SecondaryFields []*Field `json:"secondaryFields,omitempty"`
	AuxiliaryFields []*Field `json:"auxiliaryFields,omitempty"`
}

type Field struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Value string `json:"value"`
}
