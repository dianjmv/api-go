package client

type Client struct {
	Address           string          `json:"address" bson:"address"`
	Blocked           string          `json:"blocked" bson:"blocked"`
	CellPhone         string          `json:"cellPhone" bson:"cellPhone"`
	Channel           string          `json:"channel" bson:"channel"`
	ClientID          string          `json:"clientId" bson:"clientId"`
	Country           string          `json:"country" bson:"country"`
	CustomerGroup     Group           `json:"customerGroup" bson:"customerGroup"`
	CustomerSchema    int             `json:"customerSchema" bson:"customerSchema"`
	DistrChan         string          `json:"distrChan" bson:"distrChan"`
	Division          string          `json:"division" bson:"division"`
	FiscalCode1       string          `json:"fiscalCode1" bson:"fiscalCode1"`
	FiscalCode2       string          `json:"fiscalCode2" bson:"fiscalCode2"`
	Frequency         bool            `json:"frequency" bson:"frequency"`
	FrequencyDays     string          `json:"frequencyDays" bson:"frequencyDays"`
	IDPortfolio       string          `json:"idPortfolio" bson:"idPortfolio"`
	ImmediateDelivery bool            `json:"immediateDelivery" bson:"immediateDelivery"`
	IndustryCode      string          `json:"industryCode" bson:"industryCode"`
	IndustryCode1     string          `json:"industryCode1" bson:"industryCode1"`
	IsEnrollment      bool            `json:"isEnrollment" bson:"isEnrollment"`
	Name              string          `json:"name" bson:"name"`
	Name2             string          `json:"name2" bson:"name2"`
	Office            string          `json:"office" bson:"office"`
	PaymentCondition  string          `json:"paymentCondition" bson:"paymentCondition"`
	PaymentMethods    []PaymentMethod `json:"paymentMethods" bson:"paymentMethods"`
	PriceGroup        string          `json:"priceGroup" bson:"priceGroup"`
	PriceList         string          `json:"priceList" bson:"priceList"`
	RouteID           string          `json:"routeId" bson:"routeId"`
	SalesOrg          string          `json:"salesOrg" bson:"salesOrg"`
	SupplyCenter      string          `json:"supplyCenter" bson:"supplyCenter"`
	UpdateDate        Date            `json:"updateDate" bson:"updateDate"`
	VendorGroup       string          `json:"vendorGroup" bson:"vendorGroup"`
}

type Group struct {
	Group  string `json:"group" bson:"group"`
	Group1 string `json:"group1" bson:"group1"`
	Group2 string `json:"group2" bson:"group2"`
	Group3 string `json:"group3" bson:"group3"`
	Group4 string `json:"group4" bson:"group4"`
	Group5 string `json:"group5" bson:"group5"`
}

type PaymentMethod struct {
	ClientID        string `json:"clientId" bson:"clientId"`
	TypeCredit      string `json:"typeCredit" bson:"typeCredit"`
	CreditLimit     int    `json:"creditLimit" bson:"creditLimit"`
	CreditUsed      int    `json:"creditUsed" bson:"creditUsed"`
	AmountAvailable int    `json:"amountAvailable" bson:"amountAvailable"`
}

type Date struct {
	Date string `json:"$date" bson:"$date"`
}
