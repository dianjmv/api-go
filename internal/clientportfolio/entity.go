package clientportfolio

type ClientPortfolio struct {
	ID           string `json:"_id" bson:"_id"`
	Channel      string `json:"channel" bson:"channel"`
	Country      string `json:"country" bson:"country"`
	CreatedDate  Date   `json:"createdDate" bson:"createdDate"`
	CustomerCode string `json:"customerCode" bson:"customerCode"`
	Items        []Item `json:"items" bson:"items"`
	Route        string `json:"route" bson:"route"`
}

type Item struct {
	SKU                    string `json:"sku" bson:"sku"`
	Title                  string `json:"title" bson:"title"`
	CategoryID             string `json:"categoryId" bson:"categoryId"`
	Category               string `json:"category" bson:"category"`
	Brand                  string `json:"brand" bson:"brand"`
	Classification         string `json:"classification" bson:"classification"`
	UnitsPerBox            string `json:"unitsPerBox" bson:"unitsPerBox"`
	MinOrderUnits          string `json:"minOrderUnits" bson:"minOrderUnits"`
	PackageDescription     string `json:"packageDescription" bson:"packageDescription"`
	PackageUnitDescription string `json:"packageUnitDescription" bson:"packageUnitDescription"`
	QuantityMaxRedeem      int    `json:"quantity_max_redeem" bson:"quantity_max_redeem"`
	RedeemUnit             string `json:"redeem_unit" bson:"redeem_unit"`
	OrderReasonRedeem      string `json:"order_reason_redeem" bson:"order_reason_redeem"`
	SKURedeem              bool   `json:"sku_redeem" bson:"sku_redeem"`
	Price                  Price  `json:"price" bson:"price"`
	Points                 int    `json:"points" bson:"points"`
}

type Price struct {
	FullPrice int   `json:"full_price" bson:"full_price"`
	Taxes     []Tax `json:"taxes" bson:"taxes"`
}

type Tax struct {
	TaxType string `json:"taxType" bson:"taxType"`
	TaxID   string `json:"taxId" bson:"taxId"`
	Rate    int    `json:"rate" bson:"rate"`
}

type Date struct {
	Date string `json:"$date" bson:"$date"`
}
