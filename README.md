# Bigcommerce API Client for Go

## Usage

```go
package main

import (
    "github.com/gpmd/bigcommerce-api-go"
    "log"
)

func main() {
    // for app development, you need to provide arguments
    // for CLI and other web apps you can use empty strings
    client := bigcommerce.NewClient("** my store's hash like '123abcdefg' **", "**my X-Auth-Token generated in BigCommerce admin**")
    products, err := client.GetAllProducts()
    if err != nil {
        log.Fatalf("Error while getting products: %v", err)
    }
    for _, product := range products {
        log.Println(product.Name)
    }
}
```

## Errors

```go
var ErrNoContent = errors.New("no content 204 from BigCommerce API")
var ErrNoMainThumbnail = errors.New("no main thumbnail")
var ErrNotFound = errors.New("404 not found")
```

## Types

#### type Address

```go
type Address struct {
	ID              int64  `json:"id"`
	CustomerID      int64  `json:"customer_id"`
	Address1        string `json:"address1"`
	Address2        string `json:"address2"`
	AddressType     string `json:"address_type"`
	City            string `json:"city"`
	Company         string `json:"company"`
	Country         string `json:"country"`
	CountryCode     string `json:"country_code"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Phone           string `json:"phone"`
	PostalCode      string `json:"postal_code"`
	StateOrProvince string `json:"state_or_province"`
}
```

Address is for Customer Address endpoint

#### type AddressClient

```go
type AddressClient interface {
	CreateAddress(customerID int64, address *Address) (*Address, error)
	UpdateAddress(customerID int64, address *Address) (*Address, error)
	DeleteAddress(customerID int64, addressID int64) error
	GetAddresses(customerID int64) ([]Address, error)
}
```


#### type App

```go
type App struct {
	Hostname        string
	AppClientID     string
	AppClientSecret string
	HTTPClient      HTTPClient
	MaxRetries      int
	ChannelID       int
}
```

BigCommerce is the BigCommerce API client object for BigCommerce Apps holds no
client specific information

#### func  NewApp

```go
func NewApp(hostname, appClientID, appClientSecret string) *App
```
New returns a new BigCommerce API object with the given hostname, client ID, and
client secret The client ID and secret are the App's client ID and secret from
the BigCommerce My Apps dashboard The hostname is the domain name of the app
from the same page (e.g. app.exampledomain.com)

#### func (*App) CheckSignature

```go
func (bc *App) CheckSignature(signedPayload string) ([]byte, error)
```
CheckSignature checks the signature of the request whith SHA256 HMAC

#### func (*App) GetAuthContext

```go
func (bc *App) GetAuthContext(requestURLQuery url.Values) (*AuthContext, error)
```
GetAuthContext returns an AuthContext object from the BigCommerce API Call it
with r.URL.Query() - will return BigCommerce Auth Context or error

#### func (*App) GetClientRequest

```go
func (bc *App) GetClientRequest(requestURLQuery url.Values) (*ClientRequest, error)
```
GetClientRequest returns a ClientRequest object from the BigCommerce API Call it
with r.URL.Query() - will return BigCommerce Client Request or error

#### type AuthContext

```go
type AuthContext struct {
	AccessToken string `json:"access_token"` // used later as X-Auth-Token header
	Scope       string `json:"scope"`
	User        BCUser `json:"user"`
	Context     string `json:"context"`
	URL         string `json:"url"`
	Error       string `json:"error"`
}
```

AuthContext is a BigCommerce auth context object

#### type AuthContexter

```go
type AuthContexter interface {
	GetAuthContext(clientID, clientSecret string, q url.Values) (*AuthContext, error)
}
```

AuthContexter interface for GetAuthContext

#### type AuthTokenRequest

```go
type AuthTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	Scope        string `json:"scope"`
	GrantType    string `json:"grant_type"`
	RedirectURI  string `json:"redirect_uri"`
	Context      string `json:"context"`
}
```

AuthTokenRequest is sent to BigCommerce to get AuthContext

#### type Authentication

```go
type Authentication struct {
	ForcePasswordReset bool   `json:"force_password_reset"`
	Password           string `json:"new_password"`
}
```

AccountAuthentication is for CreateAccountPayload's authentication field

#### type BCUser

```go
type BCUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
```

BCUser is a BigCommerce shorthand object type that's in many other responses

#### type BlogClient

```go
type BlogClient interface {
	GetAllPosts(context, xAuthToken string) ([]Post, error)
	GetPosts(page int) ([]Post, bool, error)
}
```

BlogClient interface handles blog-related requests

#### type Brand

```go
type Brand struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	PageTitle       string   `json:"page_title"`
	MetaKeywords    []string `json:"meta_keywords"`
	MetaDescription string   `json:"meta_description"`
	ImageURL        string   `json:"image_url"`
	SearchKeywords  string   `json:"search_keywords"`
	CustomURL       struct {
		URL          string `json:"url"`
		IsCustomized bool   `json:"is_customized"`
	} `json:"custom_url"`
	URL string `json:"-"`
}
```

Brand is BigCommerce brand object

#### type Cart

```go
type Cart struct {
	ID          string `json:"id"`
	CheckoutURL string `json:"checkout_url"`
	CustomerID  int64  `json:"customer_id"`
	ChannelID   int64  `json:"channel_id"`
	Email       string `json:"email"`
	Currency    struct {
		Code string `json:"code"`
	} `json:"currency"`
	TaxIncluded    bool         `json:"tax_included"`
	BaseAmount     float64      `json:"base_amount"`
	DiscountAmount float64      `json:"discount_amount"`
	CartAmount     float64      `json:"cart_amount"`
	Discounts      []Discount   `json:"discounts"`
	Coupons        []CartCoupon `json:"coupons"`
	LineItems      struct {
		PhysicalItems    []LineItem `json:"physical_items"`
		DigitalItems     []LineItem `json:"digital_items"`
		GiftCertificates []LineItem `json:"gift_certificates"`
		CustomItems      []LineItem `json:"custom_items"`
	} `json:"line_items"`
	CreatedTime  time.Time `json:"created_time"`
	UpdatedTime  time.Time `json:"updated_time"`
	RedirectUrls struct {
		CartURL             string `json:"cart_url"`
		CheckoutURL         string `json:"checkout_url"`
		EmbeddedCheckoutURL string `json:"embedded_checkout_url"`
	} `json:"redirect_urls"`
	Locale string `json:"locale"`
}
```

Cart is a BigCommerce cart object

#### type CartClient

```go
type CartClient interface {
	CreateCart(items []LineItem) (*Cart, error)
	GetCart(cartID string) (*Cart, error)
	CartAddItems(cartID string, items []LineItem) (*Cart, error)
	CartEditItem(cartID string, item LineItem) (*Cart, error)
	CartDeleteItem(cartID string, item LineItem) (*Cart, error)
	CartUpdateCustomerID(cartID, customerID string) (*Cart, error)
	DeleteCart(cartID string) error
}
```

CartClient interface handles cart and login related requests

#### type CartCoupon

```go
type CartCoupon struct {
	Code             string  `json:"code"`
	ID               string  `json:"id"`
	CouponType       string  `json:"coupon_type"`
	DiscountedAmount float64 `json:"discounted_amount"`
}
```


#### type CartURLs

```go
type CartURLs struct {
	CartURL             string `json:"cart_url"`
	CheckoutURL         string `json:"checkout_url"`
	EmbeddedCheckoutURL string `json:"embedded_checkout_url"`
}
```


#### type CatalogClient

```go
type CatalogClient interface {
	GetAllBrands() ([]Brand, error)
	GetBrands(page int) ([]Brand, bool, error)
	GetAllCategories() ([]Category, error)
	GetCategories(page int) ([]Category, bool, error)
	GetClientRequest(requestURLQuery url.Values) (*ClientRequest, error)
	GetMainThumbnailURL(productID int64) (string, error)
	SetProductFields(fields []string)
	SetProductInclude(subresources []string)
	GetAllProducts(map[string]string) ([]Product, error)
	GetProducts(page int) ([]Product, bool, error)
	GetProductByID(productID int64) (*Product, error)
}
```

CatalogClient interface handles catalog-related requests

#### type Category

```go
type Category struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ParentID  int64  `json:"parent_id"`
	Visible   bool   `json:"is_visible"`
	FullName  string `json:"-"`
	CustomURL struct {
		URL        string `json:"url"`
		Customized bool   `json:"is_customized"`
	} `json:"custom_url"`
	URL string `json:"-"`
}
```

Category is a BC category object

#### type Channel

```go
type Channel struct {
	IconURL          string    `json:"icon_url"`
	IsListableFromUI bool      `json:"is_listable_from_ui"`
	IsVisible        bool      `json:"is_visible"`
	DateCreated      time.Time `json:"date_created"`
	ExternalID       string    `json:"external_id"`
	Type             string    `json:"type"`
	Platform         string    `json:"platform"`
	IsEnabled        bool      `json:"is_enabled"`
	DateModified     time.Time `json:"date_modified"`
	Name             string    `json:"name"`
	ID               int       `json:"id"`
	Status           string    `json:"status"`
}
```


#### type Client

```go
type Client struct {
	StoreHash  string `json:"store-hash"`
	XAuthToken string `json:"x-auth-token"`
	MaxRetries int
	HTTPClient *http.Client
	ChannelID  int
}
```


#### func  NewClient

```go
func NewClient(storeHash, xAuthToken string) *Client
```

#### func (*Client) CartAddItems

```go
func (bc *Client) CartAddItems(cartID string, items []LineItem) (*Cart, error)
```
CartAddItem adds line items to a cart

#### func (*Client) CartDeleteItem

```go
func (bc *Client) CartDeleteItem(cartID string, item LineItem) (*Cart, error)
```
DeleteItem deletes a line item from a cart, returns the updated cart Arguments:

    cartID: the cart ID
    item: the line item, must have an existing line item ID

returns nil for empty cart

#### func (*Client) CartEditItem

```go
func (bc *Client) CartEditItem(cartID string, item LineItem) (*Cart, error)
```
EditItem edits a line item in a cart, returns the updated cart Arguments:

    cartID: the cart ID
    item: the line item to edit. Must have an ID, quantity, and product ID

#### func (*Client) CartUpdateCustomerID

```go
func (bc *Client) CartUpdateCustomerID(cartID, customerID string) (*Cart, error)
```
CartUpdateCustomerID updates the customer ID for a cart Arguments: cartID: the
BigCommerce cart ID customerID: the new BigCommerce customer ID

#### func (*Client) CreateAccount

```go
func (bc *Client) CreateAccount(payload *CreateAccountPayload) (*Customer, error)
```
CreateAccount creates a new customer account in BigCommerce and returns the
customer or error

#### func (*Client) CreateAddress

```go
func (bc *Client) CreateAddress(customerID int64, address *Address) (*Address, error)
```
CreateAddress creates a new address for a customer from given data, ignoring ID
(duplicating address)

#### func (*Client) CreateCart

```go
func (bc *Client) CreateCart(items []LineItem) (*Cart, error)
```
CreateCart creates a new cart in BigCommerce and returns it

#### func (*Client) CreateCoupon

```go
func (bc *Client) CreateCoupon(coupon Coupon) (*Coupon, error)
```

#### func (*Client) CreateWebhook

```go
func (bc *Client) CreateWebhook(scope, destination string, headers map[string]string) (int64, error)
```
CreateWebhook creates a new webhook or activates it if it already exists but
inactive

#### func (*Client) CustomerGetFormFields

```go
func (bc *Client) CustomerGetFormFields(customerID int64) ([]FormField, error)
```

#### func (*Client) CustomerSetFormFields

```go
func (bc *Client) CustomerSetFormFields(customerID int64, formFields []FormField) error
```
CustomerSetFormFields sets the form fields for a customer

#### func (*Client) DeleteAddress

```go
func (bc *Client) DeleteAddress(customerID, addressID int64) error
```
DeleteAddress deletes an existing address, address ID is required

#### func (*Client) DeleteCart

```go
func (bc *Client) DeleteCart(cartID string) error
```
DeleteCart deletes a cart by ID from BigCommerce

#### func (*Client) DeleteCoupon

```go
func (bc *Client) DeleteCoupon(couponID int64) error
```

#### func (*Client) GetActiveThemeConfig

```go
func (bc *Client) GetActiveThemeConfig() (*ThemeConfig, error)
```
GetActiveThemeConfig returns the active theme config (not handling variations
yet)

#### func (*Client) GetAddressPage

```go
func (bc *Client) GetAddressPage(customerID int64, page int) ([]Address, bool, error)
```
GetAddressPage returns all addresses for a curstomer, handling pagination
customerID is bigcommerce customer id page: the page number to download

#### func (*Client) GetAddresses

```go
func (bc *Client) GetAddresses(customerID int64) ([]Address, error)
```
GetAddresses returns all addresses for a curstomer, handling pagination
customerID is bigcommerce customer id

#### func (*Client) GetAllBrands

```go
func (bc *Client) GetAllBrands(args map[string]string) ([]Brand, error)
```
GetAllBrands returns all brands, handling pagination args is a map of arguments
to pass to the API

#### func (*Client) GetAllCategories

```go
func (bc *Client) GetAllCategories(args map[string]string) ([]Category, error)
```
GetAllCategories returns a list of categories, handling pagination args is a map
of arguments to pass to the API

#### func (*Client) GetAllChannels

```go
func (bc *Client) GetAllChannels() ([]Channel, error)
```

#### func (*Client) GetAllCoupons

```go
func (bc *Client) GetAllCoupons(args map[string]string) ([]Coupon, error)
```

#### func (*Client) GetAllPosts

```go
func (bc *Client) GetAllPosts() ([]Post, error)
```
GetAllPosts downloads all posts from BigCommerce, handling pagination

#### func (*Client) GetAllProducts

```go
func (bc *Client) GetAllProducts(args map[string]string) ([]Product, error)
```
GetAllProducts gets all products from BigCommerce args is a key-value map of
additional arguments to pass to the API

#### func (*Client) GetBrands

```go
func (bc *Client) GetBrands(args map[string]string, page int) ([]Brand, bool, error)
```
GetBrands returns all brands, handling pagination args is a map of arguments to
pass to the API page: the page number to download

#### func (*Client) GetCart

```go
func (bc *Client) GetCart(cartID string) (*Cart, error)
```
GetCart gets a cart by ID from BigCommerce and returns it

#### func (*Client) GetCategories

```go
func (bc *Client) GetCategories(args map[string]string, page int) ([]Category, bool, error)
```
GetCategories returns a list of categories, handling pagination args is a map of
arguments to pass to the API page: the page number to download

#### func (*Client) GetChannels

```go
func (bc *Client) GetChannels(page int) ([]Channel, bool, error)
```

#### func (*Client) GetCoupon

```go
func (bc *Client) GetCoupon(couponID int64) (*Coupon, error)
```

#### func (*Client) GetCoupons

```go
func (bc *Client) GetCoupons(args map[string]string, page int) ([]Coupon, bool, error)
```

#### func (*Client) GetCurrencies

```go
func (bc *Client) GetCurrencies() ([]Currency, error)
```
GetCurrencies returns the store's defined currencies

#### func (*Client) GetCustomerByEmail

```go
func (bc *Client) GetCustomerByEmail(email string) (*Customer, error)
```

#### func (*Client) GetCustomerByID

```go
func (bc *Client) GetCustomerByID(customerID int64) (*Customer, error)
```

#### func (*Client) GetMainThumbnailURL

```go
func (bc *Client) GetMainThumbnailURL(productID int64) (string, error)
```
GetMainThumbnailURL returns the main thumbnail URL for a product this is due to
the fact that the Product API does not return the main thumbnail URL

#### func (*Client) GetOrder

```go
func (bc *Client) GetOrder(orderID int64) (*Order, error)
```
GetOrder returns a given order filters: request query parameters for BigCommerce
orders endpoint, for example {"customer_id": "41"}

#### func (*Client) GetOrderCoupons

```go
func (bc *Client) GetOrderCoupons(orderID int64) ([]OrderCoupon, error)
```
GetOrderCoupons returns all coupons for a given order

#### func (*Client) GetOrderProducts

```go
func (bc *Client) GetOrderProducts(orderID int64) ([]OrderProduct, error)
```
GetOrderProducts returns all products for a given order

#### func (*Client) GetOrderShippingAddresses

```go
func (bc *Client) GetOrderShippingAddresses(orderID int64) ([]OrderShippingAddress, error)
```
GetOrderShippingAddresses returns all shipping addresses for a given order

#### func (*Client) GetOrders

```go
func (bc *Client) GetOrders(filters map[string]string) ([]Order, error)
```
GetOrders returns all orders using filters filters: request query parameters for
BigCommerce orders endpoint, for example {"customer_id": "41"}

#### func (*Client) GetPosts

```go
func (bc *Client) GetPosts(page int) ([]Post, bool, error)
```
GetPosts downloads all posts from BigCommerce, handling pagination page: the
page number to download

#### func (*Client) GetProductByID

```go
func (bc *Client) GetProductByID(productID int64) (*Product, error)
```
GetProductByID gets a product from BigCommerce by ID productID: BigCommerce
product ID to get

#### func (*Client) GetProductMetafields

```go
func (bc *Client) GetProductMetafields(productID int64) (map[string]Metafield, error)
```
GetProductMetafields gets metafields values for a product productID: BigCommerce
product ID to get metafields for

#### func (*Client) GetProducts

```go
func (bc *Client) GetProducts(args map[string]string, page int) ([]Product, bool, error)
```
GetProducts gets a page of products from BigCommerce args is a key-value map of
additional arguments to pass to the API page: the page number to download

#### func (*Client) GetStoreInfo

```go
func (bc *Client) GetStoreInfo() (StoreInfo, error)
```
GetStoreInfo returns the store info for the current store page: the page number
to download

#### func (*Client) GetThemeConfig

```go
func (bc *Client) GetThemeConfig(uuid string) (*ThemeConfig, error)
```
GetThemeConfig returns the configuration for a specific theme by theme UUID

#### func (*Client) GetThemes

```go
func (bc *Client) GetThemes() ([]Theme, error)
```
GetThemes returns a list of all store themes

#### func (*Client) GetWebhooks

```go
func (bc *Client) GetWebhooks() ([]Webhook, error)
```

#### func (*Client) UpdateAddress

```go
func (bc *Client) UpdateAddress(customerID int64, address *Address) (*Address, error)
```
UpdateAddress updates an existing address, address ID is required

#### func (*Client) UpdateCoupon

```go
func (bc *Client) UpdateCoupon(couponID int64, coupon Coupon) (*Coupon, error)
```

#### func (*Client) ValidateCredentials

```go
func (bc *Client) ValidateCredentials(email, password string) (int64, error)
```
ValidateCredentials returns customer ID or error (i.e. ErrNotfound) if the
provided credentials are valid in BigCommerce

#### type ClientRequest

```go
type ClientRequest struct {
	User      UserPart `json:"user"`
	Owner     UserPart `json:"owner"`
	Context   string   `json:"context"`
	StoreHash string   `json:"store_hash"`
}
```

ClientRequest is a BigCommerce client request object that comes with most App
callbacks in the GET request signed_payload parameter

#### type Coupon

```go
type Coupon struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Amount      string `json:"amount"`
	MinPurchase string `json:"min_purchase"`
	Expires     string `json:"expires"`
	Enabled     bool   `json:"enabled"`
	Code        string `json:"code"`
	AppliesTo   struct {
		Entity string  `json:"entity"`
		Ids    []int64 `json:"ids"`
	} `json:"applies_to"`
	NumUses            int           `json:"num_uses"`
	MaxUses            int           `json:"max_uses"`
	MaxUsesPerCustomer int           `json:"max_uses_per_customer"`
	RestrictedTo       []interface{} `json:"restricted_to"`
	ShippingMethods    struct {
	} `json:"shipping_methods"`
	DateCreated string `json:"date_created"`
}
```


#### type CreateAccountPayload

```go
type CreateAccountPayload struct {
	Company                                 string         `json:"company"`
	FirstName                               string         `json:"first_name"`
	LastName                                string         `json:"last_name"`
	Email                                   string         `json:"email"`
	Phone                                   string         `json:"phone"`
	Notes                                   string         `json:"notes"`
	TaxExemptCategory                       string         `json:"tax_exempt_category"`
	CustomerGroupID                         int64          `json:"customer_group_id"`
	Addresses                               []Address      `json:"addresses"`
	Authentication                          Authentication `json:"authentication"`
	AcceptsProductReviewAbandonedCartEmails bool           `json:"accepts_product_review_abandoned_cart_emails"`
	StoreCreditAmounts                      []StoreCredit  `json:"store_credit_amounts"`
	OriginChannelID                         int            `json:"origin_channel_id"`
	ChannelIds                              []int          `json:"channel_ids"`
}
```


#### type Currency

```go
type Currency struct {
	ID                     int      `json:"id"`
	IsDefault              bool     `json:"is_default"`
	LastUpdated            string   `json:"last_updated"`
	CountryIso2            string   `json:"country_iso2"`
	DefaultForCountryCodes []string `json:"default_for_country_codes"`
	CurrencyCode           string   `json:"currency_code"`
	CurrencyExchangeRate   string   `json:"currency_exchange_rate"`
	Name                   string   `json:"name"`
	Token                  string   `json:"token"`
	AutoUpdate             bool     `json:"auto_update"`
	TokenLocation          string   `json:"token_location"`
	DecimalToken           string   `json:"decimal_token"`
	ThousandsToken         string   `json:"thousands_token"`
	DecimalPlaces          int      `json:"decimal_places"`
	Enabled                bool     `json:"enabled"`
	IsTransactional        bool     `json:"is_transactional"`
	UseDefaultName         bool     `json:"use_default_name"`
}
```

Currency is entry for BC currency API

#### type Customer

```go
type Customer struct {
	ID               int64       `json:"id"`
	Company          string      `json:"company"`
	Firstname        string      `json:"first_name"`
	Lastname         string      `json:"last_name"`
	Email            string      `json:"email"`
	Phone            string      `json:"phone"`
	FormFields       interface{} `json:"form_fields"`
	DateCreated      string      `json:"date_created"`
	DateModified     string      `json:"date_modified"`
	StoreCredit      string      `json:"store_credit"`
	RegistrationIP   string      `json:"registration_ip_address"`
	CustomerGroup    int64       `json:"customer_group_id"`
	Notes            string      `json:"notes"`
	TaxExempt        string      `json:"tax_exempt_category"`
	ResetPassword    bool        `json:"reset_pass_on_login"`
	AcceptsMarketing bool        `json:"accepts_marketing"`
	Addresses        []Address   `json:"addresses"`
}
```

Customer is a struct for the BigCommerce Customer API

#### type CustomerClient

```go
type CustomerClient interface {
	ValidateCredentials(email, password string) (int64, error)
	CreateAccount(customer *CreateAccountPayload) (*Customer, error)
	CustomerSetFormFields(customerID int64, formFields []FormField) error
	CustomerGetFormFields(customerID int64) ([]FormField, error)
	GetCustomerByID(customerID int64) (*Customer, error)
	GetCustomerByEmail(email string) (*Customer, error)
}
```


#### type Discount

```go
type Discount struct {
	ID               string  `json:"id"`
	DiscountedAmount float64 `json:"discounted_amount"`
}
```


#### type ErrorResult

```go
type ErrorResult struct {
	Status int               `json:"status"`
	Title  string            `json:"title"`
	Type   string            `json:"type"`
	Errors map[string]string `json:"errors"`
}
```


#### type FormField

```go
type FormField struct {
	CustomerID int64  `json:"customer_id"`
	Name       string `json:"name"`
	Value      string `json:"value"`
}
```

FormField is a struct for the BigCommerce Customer API Form Fiel values

#### type HTTPClient

```go
type HTTPClient interface {
	Do(req *http.Request) (res *http.Response, err error)
	Get(url string) (res *http.Response, err error)
	Post(urstring, bodyType string, body io.Reader) (res *http.Response, err error)
}
```


#### type Image

```go
type Image struct {
	ID           int64  `json:"id"`
	ProductID    int64  `json:"product_id"`
	IsThumbnail  bool   `json:"is_thumbnail"`
	SortOrder    int64  `json:"sort_order"`
	Description  string `json:"description"`
	ImageFile    string `json:"image_file"`
	URLZoom      string `json:"url_zoom"`
	URLStandard  string `json:"url_standard"`
	URLThumbnail string `json:"url_thumbnail"`
	URLTiny      string `json:"url_tiny"`
	DateModified string `json:"date_modified"`
}
```

Image is entry for BC product images

#### type InventoryEntry

```go
type InventoryEntry struct {
	ProductID int64   `json:"product_id"`
	Method    string  `json:"method"`
	Value     float64 `json:"value"`
	VariantID int64   `json:"variant_id"`
}
```


#### type LineItem

```go
type LineItem struct {
	ID                string     `json:"id"`
	ParentID          int64      `json:"parent_id"`
	VariantID         int64      `json:"variant_id"`
	ProductID         int64      `json:"product_id"`
	Sku               string     `json:"sku"`
	Name              string     `json:"name"`
	URL               string     `json:"url"`
	Quantity          float64    `json:"quantity"`
	Taxable           bool       `json:"taxable"`
	ImageURL          string     `json:"image_url"`
	Discounts         []Discount `json:"discounts"`
	Coupons           []Coupon   `json:"coupons"`
	DiscountAmount    float64    `json:"discount_amount"`
	CouponAmount      float64    `json:"coupon_amount"`
	ListPrice         float64    `json:"list_price"`
	SalePrice         float64    `json:"sale_price"`
	ExtendedListPrice float64    `json:"extended_list_price"`
	ExtendedSalePrice float64    `json:"extended_sale_price"`
	IsRequireShipping bool       `json:"is_require_shipping"`
	IsMutable         bool       `json:"is_mutable"`
}
```

LineItem is a BigCommerce line item object for cart

#### type LoadContext

```go
type LoadContext struct {
	User      BCUser  `json:"user"`
	Owner     BCUser  `json:"owner"`
	Context   string  `json:"context"`
	StoreHash string  `json:"store_hash"`
	Timestamp float64 `json:"timestamp"`
	URL       string  `json:"url"`
}
```

LoadContext is a BigCommerce load context object

#### type Metafield

```go
type Metafield struct {
	ID            int64     `json:"id"`
	Key           string    `json:"key"`
	Value         string    `json:"value"`
	ResourceID    int64     `json:"resource_id"`
	ResourceType  string    `json:"resource_type"`
	Description   string    `json:"description"`
	DateCreated   time.Time `json:"date_created"`
	DateModified  time.Time `json:"date_modified"`
	Namespace     string    `json:"namespace"`
	PermissionSet string    `json:"permission_set"`
}
```

Metafield is a struct representing a BigCommerce product metafield

#### type Order

```go
type Order struct {
	ID                                      int64        `json:"id"`
	CustomerID                              int64        `json:"customer_id"`
	DateCreated                             string       `json:"date_created"`
	DateModified                            string       `json:"date_modified"`
	DateShipped                             string       `json:"date_shipped"`
	StatusID                                int64        `json:"status_id"`
	Status                                  string       `json:"status"`
	SubtotalExTax                           string       `json:"subtotal_ex_tax"`
	SubtotalIncTax                          string       `json:"subtotal_inc_tax"`
	SubtotalTax                             string       `json:"subtotal_tax"`
	BaseShippingCost                        string       `json:"base_shipping_cost"`
	ShippingCostExTax                       string       `json:"shipping_cost_ex_tax"`
	ShippingCostIncTax                      string       `json:"shipping_cost_inc_tax"`
	ShippingCostTax                         string       `json:"shipping_cost_tax"`
	ShippingCostTaxClassID                  int64        `json:"shipping_cost_tax_class_id"`
	BaseHandlingCost                        string       `json:"base_handling_cost"`
	HandlingCostExTax                       string       `json:"handling_cost_ex_tax"`
	HandlingCostIncTax                      string       `json:"handling_cost_inc_tax"`
	HandlingCostTax                         string       `json:"handling_cost_tax"`
	HandlingCostTaxClassID                  int64        `json:"handling_cost_tax_class_id"`
	BaseWrappingCost                        string       `json:"base_wrapping_cost"`
	WrappingCostExTax                       string       `json:"wrapping_cost_ex_tax"`
	WrappingCostIncTax                      string       `json:"wrapping_cost_inc_tax"`
	WrappingCostTax                         string       `json:"wrapping_cost_tax"`
	WrappingCostTaxClassID                  int64        `json:"wrapping_cost_tax_class_id"`
	TotalExTax                              string       `json:"total_ex_tax"`
	TotalIncTax                             string       `json:"total_inc_tax"`
	TotalTax                                string       `json:"total_tax"`
	ItemsTotal                              int          `json:"items_total"`
	ItemsShipped                            int          `json:"items_shipped"`
	PaymentMethod                           string       `json:"payment_method"`
	PaymentProviderID                       string       `json:"payment_provider_id"`
	PaymentStatus                           string       `json:"payment_status"`
	RefundedAmount                          string       `json:"refunded_amount"`
	OrderIsDigital                          bool         `json:"order_is_digital"`
	StoreCreditAmount                       string       `json:"store_credit_amount"`
	GiftCertificateAmount                   string       `json:"gift_certificate_amount"`
	IPAddress                               string       `json:"ip_address"`
	IPAddressV6                             string       `json:"ip_address_v6"`
	GeoipCountry                            string       `json:"geoip_country"`
	GeoipCountryIso2                        string       `json:"geoip_country_iso2"`
	CurrencyID                              int64        `json:"currency_id"`
	CurrencyCode                            string       `json:"currency_code"`
	CurrencyExchangeRate                    string       `json:"currency_exchange_rate"`
	DefaultCurrencyID                       int64        `json:"default_currency_id"`
	DefaultCurrencyCode                     string       `json:"default_currency_code"`
	StaffNotes                              string       `json:"staff_notes"`
	CustomerMessage                         string       `json:"customer_message"`
	DiscountAmount                          string       `json:"discount_amount"`
	CouponDiscount                          string       `json:"coupon_discount"`
	ShippingAddressCount                    int          `json:"shipping_address_count"`
	IsDeleted                               bool         `json:"is_deleted"`
	EbayOrderID                             string       `json:"ebay_order_id"`
	CartID                                  string       `json:"cart_id"`
	BillingAddress                          OrderAddress `json:"billing_address"`
	IsEmailOptIn                            bool         `json:"is_email_opt_in"`
	CreditCardType                          interface{}  `json:"credit_card_type"`
	OrderSource                             string       `json:"order_source"`
	ChannelID                               int64        `json:"channel_id"`
	ExternalSource                          string       `json:"external_source"`
	Products                                interface{}  `json:"products"`
	ShippingAddresses                       interface{}  `json:"shipping_addresses"`
	Coupons                                 interface{}  `json:"coupons"`
	ExternalID                              interface{}  `json:"external_id"`
	ExternalMerchantID                      interface{}  `json:"external_merchant_id"`
	TaxProviderID                           string       `json:"tax_provider_id"`
	StoreDefaultCurrencyCode                string       `json:"store_default_currency_code"`
	StoreDefaultToTransactionalExchangeRate string       `json:"store_default_to_transactional_exchange_rate"`
	CustomStatus                            string       `json:"custom_status"`
	CustomerLocale                          string       `json:"customer_locale"`
}
```


#### type OrderAddress

```go
type OrderAddress struct {
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	Company     string        `json:"company"`
	Street1     string        `json:"street_1"`
	Street2     string        `json:"street_2"`
	City        string        `json:"city"`
	State       string        `json:"state"`
	Zip         string        `json:"zip"`
	Country     string        `json:"country"`
	CountryIso2 string        `json:"country_iso2"`
	Phone       string        `json:"phone"`
	Email       string        `json:"email"`
	FormFields  []interface{} `json:"form_fields"`
}
```


#### type OrderCoupon

```go
type OrderCoupon struct {
	ID       int64  `json:"id"`
	CouponID int64  `json:"coupon_id"`
	OrderID  int64  `json:"order_id"`
	Code     string `json:"code"`
	Amount   int    `json:"amount"`
	Type     int    `json:"type"`
	Discount int    `json:"discount"`
}
```


#### type OrderProduct

```go
type OrderProduct struct {
	ID                   int64             `json:"id"`
	OrderID              int64             `json:"order_id"`
	ProductID            int64             `json:"product_id"`
	OrderAddressID       int64             `json:"order_address_id"`
	Name                 string            `json:"name"`
	NameCustomer         string            `json:"name_customer"`
	NameMerchant         string            `json:"name_merchant"`
	Sku                  string            `json:"sku"`
	Upc                  string            `json:"upc"`
	Type                 string            `json:"type"`
	BasePrice            string            `json:"base_price"`
	PriceExTax           string            `json:"price_ex_tax"`
	PriceIncTax          string            `json:"price_inc_tax"`
	PriceTax             string            `json:"price_tax"`
	BaseTotal            string            `json:"base_total"`
	TotalExTax           string            `json:"total_ex_tax"`
	TotalIncTax          string            `json:"total_inc_tax"`
	TotalTax             string            `json:"total_tax"`
	Weight               string            `json:"weight"`
	Quantity             int               `json:"quantity"`
	BaseCostPrice        string            `json:"base_cost_price"`
	CostPriceIncTax      string            `json:"cost_price_inc_tax"`
	CostPriceExTax       string            `json:"cost_price_ex_tax"`
	CostPriceTax         string            `json:"cost_price_tax"`
	IsRefunded           bool              `json:"is_refunded"`
	QuantityRefunded     int               `json:"quantity_refunded"`
	RefundAmount         string            `json:"refund_amount"`
	ReturnID             int64             `json:"return_id"`
	WrappingName         string            `json:"wrapping_name"`
	BaseWrappingCost     string            `json:"base_wrapping_cost"`
	WrappingCostExTax    string            `json:"wrapping_cost_ex_tax"`
	WrappingCostIncTax   string            `json:"wrapping_cost_inc_tax"`
	WrappingCostTax      string            `json:"wrapping_cost_tax"`
	WrappingMessage      string            `json:"wrapping_message"`
	QuantityShipped      int               `json:"quantity_shipped"`
	FixedShippingCost    string            `json:"fixed_shipping_cost"`
	EbayItemID           string            `json:"ebay_item_id"`
	EbayTransactionID    string            `json:"ebay_transaction_id"`
	OptionSetID          int64             `json:"option_set_id"`
	ParentOrderProductID interface{}       `json:"parent_order_product_id"`
	IsBundledProduct     bool              `json:"is_bundled_product"`
	BinPickingNumber     string            `json:"bin_picking_number"`
	ExternalID           interface{}       `json:"external_id"`
	FulfillmentSource    string            `json:"fulfillment_source"`
	AppliedDiscounts     []ProductDiscount `json:"applied_discounts"`
	ProductOptions       []ProductOption   `json:"product_options"`
	ConfigurableFields   []interface{}     `json:"configurable_fields"`
	EventName            interface{}       `json:"event_name"`
	EventDate            interface{}       `json:"event_date"`
}
```


#### type OrderShippingAddress

```go
type OrderShippingAddress struct {
	ID                     int64         `json:"id"`
	OrderID                int64         `json:"order_id"`
	FirstName              string        `json:"first_name"`
	LastName               string        `json:"last_name"`
	Company                string        `json:"company"`
	Street1                string        `json:"street_1"`
	Street2                string        `json:"street_2"`
	City                   string        `json:"city"`
	Zip                    string        `json:"zip"`
	Country                string        `json:"country"`
	CountryIso2            string        `json:"country_iso2"`
	State                  string        `json:"state"`
	Email                  string        `json:"email"`
	Phone                  string        `json:"phone"`
	ItemsTotal             int           `json:"items_total"`
	ItemsShipped           int           `json:"items_shipped"`
	ShippingMethod         string        `json:"shipping_method"`
	BaseCost               string        `json:"base_cost"`
	CostExTax              string        `json:"cost_ex_tax"`
	CostIncTax             string        `json:"cost_inc_tax"`
	CostTax                string        `json:"cost_tax"`
	CostTaxClassID         int64         `json:"cost_tax_class_id"`
	BaseHandlingCost       string        `json:"base_handling_cost"`
	HandlingCostExTax      string        `json:"handling_cost_ex_tax"`
	HandlingCostIncTax     string        `json:"handling_cost_inc_tax"`
	HandlingCostTax        string        `json:"handling_cost_tax"`
	HandlingCostTaxClassID int64         `json:"handling_cost_tax_class_id"`
	ShippingZoneID         int64         `json:"shipping_zone_id"`
	ShippingZoneName       string        `json:"shipping_zone_name"`
	ShippingQuotes         interface{}   `json:"shipping_quotes"`
	FormFields             []interface{} `json:"form_fields"`
}
```


#### type Pagination

```go
type Pagination struct {
	Count       int `json:"count"`
	CurrentPage int `json:"current_page"`
	Links       struct {
		Current  string `json:"current"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
	} `json:"links"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}
```


#### type Post

```go
type Post struct {
	ID                   int64       `json:"id"`
	Title                string      `json:"title"`
	URL                  string      `json:"url"`
	PreviewURL           string      `json:"preview_url"`
	Body                 string      `json:"body"`
	Tags                 []string    `json:"tags"`
	Summary              string      `json:"summary"`
	IsPublished          bool        `json:"is_published"`
	PublishedDate        interface{} `json:"publisheddate"`
	PublishedDateISO8601 string      `json:"publisheddate_iso8601"`
	MetaDescription      string      `json:"meta_description"`
	MetaKeywords         string      `json:"meta_keywords"`
	Author               string      `json:"author"`
	ThumbnailPath        string      `json:"thumbnail_path"`
}
```

Post is a BC blog post

#### type Product

```go
type Product struct {
	ID                      int64         `json:"id"`
	Name                    string        `json:"name"`
	Type                    string        `json:"type"`
	Sku                     string        `json:"sku"`
	Description             string        `json:"description"`
	Weight                  float64       `json:"weight"`
	Width                   int           `json:"width"`
	Depth                   int           `json:"depth"`
	Height                  int           `json:"height"`
	Price                   float64       `json:"price"`
	CostPrice               float64       `json:"cost_price"`
	RetailPrice             float64       `json:"retail_price"`
	SalePrice               float64       `json:"sale_price"`
	MapPrice                float64       `json:"map_price"`
	TaxClassID              int64         `json:"tax_class_id"`
	ProductTaxCode          string        `json:"product_tax_code"`
	CalculatedPrice         float64       `json:"calculated_price"`
	Categories              []interface{} `json:"categories"`
	BrandID                 int64         `json:"brand_id"`
	OptionSetID             interface{}   `json:"option_set_id"`
	OptionSetDisplay        string        `json:"option_set_display"`
	InventoryLevel          int           `json:"inventory_level"`
	InventoryWarningLevel   int           `json:"inventory_warning_level"`
	InventoryTracking       string        `json:"inventory_tracking"`
	ReviewsRatingSum        int           `json:"reviews_rating_sum"`
	ReviewsCount            int           `json:"reviews_count"`
	TotalSold               int           `json:"total_sold"`
	FixedCostShippingPrice  float64       `json:"fixed_cost_shipping_price"`
	IsFreeShipping          bool          `json:"is_free_shipping"`
	IsVisible               bool          `json:"is_visible"`
	IsFeatured              bool          `json:"is_featured"`
	RelatedProducts         []int         `json:"related_products"`
	Warranty                string        `json:"warranty"`
	BinPickingNumber        string        `json:"bin_picking_number"`
	LayoutFile              string        `json:"layout_file"`
	Upc                     string        `json:"upc"`
	Mpn                     string        `json:"mpn"`
	Gtin                    string        `json:"gtin"`
	SearchKeywords          string        `json:"search_keywords"`
	Availability            string        `json:"availability"`
	AvailabilityDescription string        `json:"availability_description"`
	GiftWrappingOptionsType string        `json:"gift_wrapping_options_type"`
	GiftWrappingOptionsList []interface{} `json:"gift_wrapping_options_list"`
	SortOrder               int           `json:"sort_order"`
	Condition               string        `json:"condition"`
	IsConditionShown        bool          `json:"is_condition_shown"`
	OrderQuantityMinimum    int           `json:"order_quantity_minimum"`
	OrderQuantityMaximum    int           `json:"order_quantity_maximum"`
	PageTitle               string        `json:"page_title"`
	MetaKeywords            []interface{} `json:"meta_keywords"`
	MetaDescription         string        `json:"meta_description"`
	DateCreated             time.Time     `json:"date_created"`
	DateModified            time.Time     `json:"date_modified"`
	ViewCount               int           `json:"view_count"`
	PreorderReleaseDate     interface{}   `json:"preorder_release_date"`
	PreorderMessage         string        `json:"preorder_message"`
	IsPreorderOnly          bool          `json:"is_preorder_only"`
	IsPriceHidden           bool          `json:"is_price_hidden"`
	PriceHiddenLabel        string        `json:"price_hidden_label"`
	CustomURL               struct {
		URL          string `json:"url"`
		IsCustomized bool   `json:"is_customized"`
	} `json:"custom_url"`
	BaseVariantID               int64  `json:"base_variant_id"`
	OpenGraphType               string `json:"open_graph_type"`
	OpenGraphTitle              string `json:"open_graph_title"`
	OpenGraphDescription        string `json:"open_graph_description"`
	OpenGraphUseMetaDescription bool   `json:"open_graph_use_meta_description"`
	OpenGraphUseProductName     bool   `json:"open_graph_use_product_name"`
	OpenGraphUseImage           bool   `json:"open_graph_use_image"`
	Variants                    []struct {
		ID                        int64         `json:"id"`
		ProductID                 int64         `json:"product_id"`
		Sku                       string        `json:"sku"`
		SkuID                     interface{}   `json:"sku_id"`
		Price                     float64       `json:"price"`
		CalculatedPrice           float64       `json:"calculated_price"`
		SalePrice                 float64       `json:"sale_price"`
		RetailPrice               float64       `json:"retail_price"`
		MapPrice                  float64       `json:"map_price"`
		Weight                    float64       `json:"weight"`
		Width                     int           `json:"width"`
		Height                    int           `json:"height"`
		Depth                     int           `json:"depth"`
		IsFreeShipping            bool          `json:"is_free_shipping"`
		FixedCostShippingPrice    float64       `json:"fixed_cost_shipping_price"`
		CalculatedWeight          float64       `json:"calculated_weight"`
		PurchasingDisabled        bool          `json:"purchasing_disabled"`
		PurchasingDisabledMessage string        `json:"purchasing_disabled_message"`
		ImageURL                  string        `json:"image_url"`
		CostPrice                 float64       `json:"cost_price"`
		Upc                       string        `json:"upc"`
		Mpn                       string        `json:"mpn"`
		Gtin                      string        `json:"gtin"`
		InventoryLevel            int           `json:"inventory_level"`
		InventoryWarningLevel     int           `json:"inventory_warning_level"`
		BinPickingNumber          string        `json:"bin_picking_number"`
		OptionValues              []interface{} `json:"option_values"`
	} `json:"variants"`
	Images       []interface{} `json:"images"`
	PrimaryImage interface{}   `json:"primary_image"`
	Videos       []interface{} `json:"videos"`
	CustomFields []struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"custom_fields"`
	BulkPricingRules []interface{} `json:"bulk_pricing_rules"`
	Options          []interface{} `json:"options"`
	Modifiers        []interface{} `json:"modifiers"`
}
```

Product is a BigCommerce product object

#### type ProductDiscount

```go
type ProductDiscount struct {
	ID     string      `json:"id"`
	Amount string      `json:"amount"`
	Name   string      `json:"name"`
	Code   interface{} `json:"code"`
	Target string      `json:"target"`
}
```


#### type ProductOption

```go
type ProductOption struct {
	ID                   int64  `json:"id"`
	OptionID             int64  `json:"option_id"`
	OrderProductID       int64  `json:"order_product_id"`
	ProductOptionID      int64  `json:"product_option_id"`
	DisplayName          string `json:"display_name"`
	DisplayNameCustomer  string `json:"display_name_customer"`
	DisplayNameMerchant  string `json:"display_name_merchant"`
	DisplayValue         string `json:"display_value"`
	DisplayValueCustomer string `json:"display_value_customer"`
	DisplayValueMerchant string `json:"display_value_merchant"`
	Value                string `json:"value"`
	Type                 string `json:"type"`
	Name                 string `json:"name"`
	DisplayStyle         string `json:"display_style"`
}
```


#### type StoreClient

```go
type StoreClient interface {
	GetAllChannels() ([]Channel, error)
	GetChannels(page int) ([]Channel, bool, error)
	GetClientRequest(requestURLQuery url.Values) (*ClientRequest, error)
	GetStoreInfo() (StoreInfo, error)
}
```

StoreClient interface handles generic store requests

#### type StoreCredit

```go
type StoreCredit struct {
	Amount float64 `json:"amount"`
}
```

StoreCredit is for CreateAccountPayload's store_credit_ammounts field

#### type StoreInfo

```go
type StoreInfo struct {
	ID          string `json:"id"`
	Domain      string `json:"domain"`
	SecureURL   string `json:"secure_url"`
	Status      string `json:"status"`
	Name        string `json:"name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	AdminEmail  string `json:"admin_email"`
	OrderEmail  string `json:"order_email"`
	FaviconURL  string `json:"favicon_url"`
	Timezone    struct {
		Name          string `json:"name"`
		RawOffset     int    `json:"raw_offset"`
		DstOffset     int    `json:"dst_offset"`
		DstCorrection bool   `json:"dst_correction"`
		DateFormat    struct {
			Display         string `json:"display"`
			Export          string `json:"export"`
			ExtendedDisplay string `json:"extended_display"`
		} `json:"date_format"`
	} `json:"timezone"`
	Language                string        `json:"language"`
	Currency                string        `json:"currency"`
	CurrencySymbol          string        `json:"currency_symbol"`
	DecimalSeparator        string        `json:"decimal_separator"`
	ThousandsSeparator      string        `json:"thousands_separator"`
	DecimalPlaces           int           `json:"decimal_places"`
	CurrencySymbolLocation  string        `json:"currency_symbol_location"`
	WeightUnits             string        `json:"weight_units"`
	DimensionUnits          string        `json:"dimension_units"`
	DimensionDecimalPlaces  int           `json:"dimension_decimal_places"`
	DimensionDecimalToken   string        `json:"dimension_decimal_token"`
	DimensionThousandsToken string        `json:"dimension_thousands_token"`
	PlanName                string        `json:"plan_name"`
	PlanLevel               string        `json:"plan_level"`
	PlanIsTrial             bool          `json:"plan_is_trial"`
	Industry                string        `json:"industry"`
	Logo                    interface{}   `json:"logo"`
	IsPriceEnteredWithTax   bool          `json:"is_price_entered_with_tax"`
	ActiveComparisonModules []interface{} `json:"active_comparison_modules"`
	Features                struct {
		StencilEnabled       bool   `json:"stencil_enabled"`
		SitewidehttpsEnabled bool   `json:"sitewidehttps_enabled"`
		FacebookCatalogID    string `json:"facebook_catalog_id"`
		CheckoutType         string `json:"checkout_type"`
		WishlistsEnabled     bool   `json:"wishlists_enabled"`
	} `json:"features"`
}
```

StoreInfo is a BigCommerce store info object

#### type StoreSettings

```go
type StoreSettings struct {
	HideBreadcrumbs                                       bool     `json:"hide_breadcrumbs"`
	HidePageHeading                                       bool     `json:"hide_page_heading"`
	HideCategoryPageHeading                               bool     `json:"hide_category_page_heading"`
	HideBlogPageHeading                                   bool     `json:"hide_blog_page_heading"`
	HideContactUsPageHeading                              bool     `json:"hide_contact_us_page_heading"`
	HomepageNewProductsCount                              int      `json:"homepage_new_products_count"`
	HomepageFeaturedProductsCount                         int      `json:"homepage_featured_products_count"`
	HomepageTopProductsCount                              int      `json:"homepage_top_products_count"`
	HomepageShowCarousel                                  bool     `json:"homepage_show_carousel"`
	HomepageShowCarouselArrows                            bool     `json:"homepage_show_carousel_arrows"`
	HomepageShowCarouselPlayPauseButton                   bool     `json:"homepage_show_carousel_play_pause_button"`
	HomepageStretchCarouselImages                         bool     `json:"homepage_stretch_carousel_images"`
	HomepageNewProductsColumnCount                        int      `json:"homepage_new_products_column_count"`
	HomepageFeaturedProductsColumnCount                   int      `json:"homepage_featured_products_column_count"`
	HomepageTopProductsColumnCount                        int      `json:"homepage_top_products_column_count"`
	HomepageBlogPostsCount                                int      `json:"homepage_blog_posts_count"`
	ProductpageVideosCount                                int      `json:"productpage_videos_count"`
	ProductpageReviewsCount                               int      `json:"productpage_reviews_count"`
	ProductpageRelatedProductsCount                       int      `json:"productpage_related_products_count"`
	ProductpageSimilarByViewsCount                        int      `json:"productpage_similar_by_views_count"`
	CategorypageProductsPerPage                           int      `json:"categorypage_products_per_page"`
	ShopByPriceVisibility                                 bool     `json:"shop_by_price_visibility"`
	BrandpageProductsPerPage                              int      `json:"brandpage_products_per_page"`
	SearchpageProductsPerPage                             int      `json:"searchpage_products_per_page"`
	ShowProductQuickView                                  bool     `json:"show_product_quick_view"`
	ShowProductQuantityBox                                bool     `json:"show_product_quantity_box"`
	ShowPoweredBy                                         bool     `json:"show_powered_by"`
	ShopByBrandShowFooter                                 bool     `json:"shop_by_brand_show_footer"`
	ShowCopyrightFooter                                   bool     `json:"show_copyright_footer"`
	ShowAcceptAmex                                        bool     `json:"show_accept_amex"`
	ShowAcceptDiscover                                    bool     `json:"show_accept_discover"`
	ShowAcceptMastercard                                  bool     `json:"show_accept_mastercard"`
	ShowAcceptPaypal                                      bool     `json:"show_accept_paypal"`
	ShowAcceptVisa                                        bool     `json:"show_accept_visa"`
	ShowAcceptAmazonpay                                   bool     `json:"show_accept_amazonpay"`
	ShowAcceptGooglepay                                   bool     `json:"show_accept_googlepay"`
	ShowAcceptKlarna                                      bool     `json:"show_accept_klarna"`
	ShowProductDetailsTabs                                bool     `json:"show_product_details_tabs"`
	ShowProductReviews                                    bool     `json:"show_product_reviews"`
	ShowCustomFieldsTabs                                  bool     `json:"show_custom_fields_tabs"`
	ShowProductWeight                                     bool     `json:"show_product_weight"`
	ShowProductDimensions                                 bool     `json:"show_product_dimensions"`
	ShowProductSwatchNames                                bool     `json:"show_product_swatch_names"`
	ProductListDisplayMode                                string   `json:"product_list_display_mode"`
	LogoPosition                                          string   `json:"logo-position"`
	LogoSize                                              string   `json:"logo_size"`
	LogoFontSize                                          int      `json:"logo_fontSize"`
	BrandSize                                             string   `json:"brand_size"`
	GallerySize                                           string   `json:"gallery_size"`
	ProductgallerySize                                    string   `json:"productgallery_size"`
	ProductSize                                           string   `json:"product_size"`
	ProductviewThumbSize                                  string   `json:"productview_thumb_size"`
	ProductthumbSize                                      string   `json:"productthumb_size"`
	ThumbSize                                             string   `json:"thumb_size"`
	ZoomSize                                              string   `json:"zoom_size"`
	BlogSize                                              string   `json:"blog_size"`
	DefaultImageBrand                                     string   `json:"default_image_brand"`
	DefaultImageProduct                                   string   `json:"default_image_product"`
	DefaultImageGiftCertificate                           string   `json:"default_image_gift_certificate"`
	BodyFont                                              string   `json:"body-font"`
	HeadingsFont                                          string   `json:"headings-font"`
	FontSizeRoot                                          int      `json:"fontSize-root"`
	FontSizeH1                                            int      `json:"fontSize-h1"`
	FontSizeH2                                            int      `json:"fontSize-h2"`
	FontSizeH3                                            int      `json:"fontSize-h3"`
	FontSizeH4                                            int      `json:"fontSize-h4"`
	FontSizeH5                                            int      `json:"fontSize-h5"`
	FontSizeH6                                            int      `json:"fontSize-h6"`
	ApplePayButton                                        string   `json:"applePay-button"`
	ColorTextBase                                         string   `json:"color-textBase"`
	ColorTextBaseHover                                    string   `json:"color-textBase--hover"`
	ColorTextBaseActive                                   string   `json:"color-textBase--active"`
	ColorTextSecondary                                    string   `json:"color-textSecondary"`
	ColorTextSecondaryHover                               string   `json:"color-textSecondary--hover"`
	ColorTextSecondaryActive                              string   `json:"color-textSecondary--active"`
	ColorTextLink                                         string   `json:"color-textLink"`
	ColorTextLinkHover                                    string   `json:"color-textLink--hover"`
	ColorTextLinkActive                                   string   `json:"color-textLink--active"`
	ColorTextHeading                                      string   `json:"color-textHeading"`
	ColorPrimary                                          string   `json:"color-primary"`
	ColorPrimaryDark                                      string   `json:"color-primaryDark"`
	ColorPrimaryDarker                                    string   `json:"color-primaryDarker"`
	ColorPrimaryLight                                     string   `json:"color-primaryLight"`
	ColorSecondary                                        string   `json:"color-secondary"`
	ColorSecondaryDark                                    string   `json:"color-secondaryDark"`
	ColorSecondaryDarker                                  string   `json:"color-secondaryDarker"`
	ColorError                                            string   `json:"color-error"`
	ColorErrorLight                                       string   `json:"color-errorLight"`
	ColorInfo                                             string   `json:"color-info"`
	ColorInfoLight                                        string   `json:"color-infoLight"`
	ColorSuccess                                          string   `json:"color-success"`
	ColorSuccessLight                                     string   `json:"color-successLight"`
	ColorWarning                                          string   `json:"color-warning"`
	ColorWarningLight                                     string   `json:"color-warningLight"`
	ColorBlack                                            string   `json:"color-black"`
	ColorWhite                                            string   `json:"color-white"`
	ColorWhitesBase                                       string   `json:"color-whitesBase"`
	ColorGrey                                             string   `json:"color-grey"`
	ColorGreyDarkest                                      string   `json:"color-greyDarkest"`
	ColorGreyDarker                                       string   `json:"color-greyDarker"`
	ColorGreyDark                                         string   `json:"color-greyDark"`
	ColorGreyMedium                                       string   `json:"color-greyMedium"`
	ColorGreyLight                                        string   `json:"color-greyLight"`
	ColorGreyLighter                                      string   `json:"color-greyLighter"`
	ColorGreyLightest                                     string   `json:"color-greyLightest"`
	BannerDeafaultBackgroundColor                         string   `json:"banner--deafault-backgroundColor"`
	ButtonDefaultColor                                    string   `json:"button--default-color"`
	ButtonDefaultColorHover                               string   `json:"button--default-colorHover"`
	ButtonDefaultColorActive                              string   `json:"button--default-colorActive"`
	ButtonDefaultBorderColor                              string   `json:"button--default-borderColor"`
	ButtonDefaultBorderColorHover                         string   `json:"button--default-borderColorHover"`
	ButtonDefaultBorderColorActive                        string   `json:"button--default-borderColorActive"`
	ButtonPrimaryColor                                    string   `json:"button--primary-color"`
	ButtonPrimaryColorHover                               string   `json:"button--primary-colorHover"`
	ButtonPrimaryColorActive                              string   `json:"button--primary-colorActive"`
	ButtonPrimaryBackgroundColor                          string   `json:"button--primary-backgroundColor"`
	ButtonPrimaryBackgroundColorHover                     string   `json:"button--primary-backgroundColorHover"`
	ButtonPrimaryBackgroundColorActive                    string   `json:"button--primary-backgroundColorActive"`
	ButtonDisabledColor                                   string   `json:"button--disabled-color"`
	ButtonDisabledBackgroundColor                         string   `json:"button--disabled-backgroundColor"`
	ButtonDisabledBorderColor                             string   `json:"button--disabled-borderColor"`
	IconColor                                             string   `json:"icon-color"`
	IconColorHover                                        string   `json:"icon-color-hover"`
	ButtonIconSvgColor                                    string   `json:"button--icon-svg-color"`
	IconRatingEmpty                                       string   `json:"icon-ratingEmpty"`
	IconRatingFull                                        string   `json:"icon-ratingFull"`
	CarouselBgColor                                       string   `json:"carousel-bgColor"`
	CarouselTitleColor                                    string   `json:"carousel-title-color"`
	CarouselDescriptionColor                              string   `json:"carousel-description-color"`
	CarouselDotColor                                      string   `json:"carousel-dot-color"`
	CarouselDotColorActive                                string   `json:"carousel-dot-color-active"`
	CarouselDotBgColor                                    string   `json:"carousel-dot-bgColor"`
	CarouselArrowColor                                    string   `json:"carousel-arrow-color"`
	CarouselArrowColorHover                               string   `json:"carousel-arrow-color--hover"`
	CarouselArrowBgColor                                  string   `json:"carousel-arrow-bgColor"`
	CarouselArrowBorderColor                              string   `json:"carousel-arrow-borderColor"`
	CarouselPlayPauseButtonTextColor                      string   `json:"carousel-play-pause-button-textColor"`
	CarouselPlayPauseButtonTextColorHover                 string   `json:"carousel-play-pause-button-textColor--hover"`
	CarouselPlayPauseButtonBgColor                        string   `json:"carousel-play-pause-button-bgColor"`
	CarouselPlayPauseButtonBorderColor                    string   `json:"carousel-play-pause-button-borderColor"`
	CardTitleColor                                        string   `json:"card-title-color"`
	CardTitleColorHover                                   string   `json:"card-title-color-hover"`
	CardFigcaptionButtonBackground                        string   `json:"card-figcaption-button-background"`
	CardFigcaptionButtonColor                             string   `json:"card-figcaption-button-color"`
	CardAlternateBackgroundColor                          string   `json:"card--alternate-backgroundColor"`
	CardAlternateBorderColor                              string   `json:"card--alternate-borderColor"`
	CardAlternateColorHover                               string   `json:"card--alternate-color--hover"`
	FormLabelFontColor                                    string   `json:"form-label-font-color"`
	InputFontColor                                        string   `json:"input-font-color"`
	InputBorderColor                                      string   `json:"input-border-color"`
	InputBorderColorActive                                string   `json:"input-border-color-active"`
	InputBgColor                                          string   `json:"input-bg-color"`
	InputDisabledBg                                       string   `json:"input-disabled-bg"`
	SelectBgColor                                         string   `json:"select-bg-color"`
	SelectArrowColor                                      string   `json:"select-arrow-color"`
	CheckRadioColor                                       string   `json:"checkRadio-color"`
	CheckRadioBackgroundColor                             string   `json:"checkRadio-backgroundColor"`
	CheckRadioBorderColor                                 string   `json:"checkRadio-borderColor"`
	AlertBackgroundColor                                  string   `json:"alert-backgroundColor"`
	AlertColor                                            string   `json:"alert-color"`
	AlertColorAlt                                         string   `json:"alert-color-alt"`
	StoreNameColor                                        string   `json:"storeName-color"`
	BodyBg                                                string   `json:"body-bg"`
	HeaderBackgroundColor                                 string   `json:"header-backgroundColor"`
	FooterBackgroundColor                                 string   `json:"footer-backgroundColor"`
	NavUserColor                                          string   `json:"navUser-color"`
	NavUserColorHover                                     string   `json:"navUser-color-hover"`
	NavUserDropdownBackgroundColor                        string   `json:"navUser-dropdown-backgroundColor"`
	NavUserDropdownBorderColor                            string   `json:"navUser-dropdown-borderColor"`
	NavUserIndicatorBackgroundColor                       string   `json:"navUser-indicator-backgroundColor"`
	NavPagesColor                                         string   `json:"navPages-color"`
	NavPagesColorHover                                    string   `json:"navPages-color-hover"`
	NavPagesSubMenuBackgroundColor                        string   `json:"navPages-subMenu-backgroundColor"`
	NavPagesSubMenuSeparatorColor                         string   `json:"navPages-subMenu-separatorColor"`
	DropdownQuickSearchBackgroundColor                    string   `json:"dropdown--quickSearch-backgroundColor"`
	DropdownWishListBackgroundColor                       string   `json:"dropdown--wishList-backgroundColor"`
	BlockquoteCiteFontColor                               string   `json:"blockquote-cite-font-color"`
	ContainerBorderGlobalColorBase                        string   `json:"container-border-global-color-base"`
	ContainerFillBase                                     string   `json:"container-fill-base"`
	ContainerFillDark                                     string   `json:"container-fill-dark"`
	LabelBackgroundColor                                  string   `json:"label-backgroundColor"`
	LabelColor                                            string   `json:"label-color"`
	OverlayBackgroundColor                                string   `json:"overlay-backgroundColor"`
	LoadingOverlayBackgroundColor                         string   `json:"loadingOverlay-backgroundColor"`
	PaceProgressBackgroundColor                           string   `json:"pace-progress-backgroundColor"`
	SpinnerBorderColorDark                                string   `json:"spinner-borderColor-dark"`
	SpinnerBorderColorLight                               string   `json:"spinner-borderColor-light"`
	HideContentNavigation                                 bool     `json:"hide_content_navigation"`
	OptimizedCheckoutHeaderBackgroundColor                string   `json:"optimizedCheckout-header-backgroundColor"`
	OptimizedCheckoutShowBackgroundImage                  bool     `json:"optimizedCheckout-show-backgroundImage"`
	OptimizedCheckoutBackgroundImage                      string   `json:"optimizedCheckout-backgroundImage"`
	OptimizedCheckoutBackgroundImageSize                  string   `json:"optimizedCheckout-backgroundImage-size"`
	OptimizedCheckoutShowLogo                             string   `json:"optimizedCheckout-show-logo"`
	OptimizedCheckoutLogo                                 string   `json:"optimizedCheckout-logo"`
	OptimizedCheckoutLogoSize                             string   `json:"optimizedCheckout-logo-size"`
	OptimizedCheckoutLogoPosition                         string   `json:"optimizedCheckout-logo-position"`
	OptimizedCheckoutHeadingPrimaryColor                  string   `json:"optimizedCheckout-headingPrimary-color"`
	OptimizedCheckoutHeadingPrimaryFont                   string   `json:"optimizedCheckout-headingPrimary-font"`
	OptimizedCheckoutHeadingSecondaryColor                string   `json:"optimizedCheckout-headingSecondary-color"`
	OptimizedCheckoutHeadingSecondaryFont                 string   `json:"optimizedCheckout-headingSecondary-font"`
	OptimizedCheckoutBodyBackgroundColor                  string   `json:"optimizedCheckout-body-backgroundColor"`
	OptimizedCheckoutColorFocus                           string   `json:"optimizedCheckout-colorFocus"`
	OptimizedCheckoutContentPrimaryColor                  string   `json:"optimizedCheckout-contentPrimary-color"`
	OptimizedCheckoutContentPrimaryFont                   string   `json:"optimizedCheckout-contentPrimary-font"`
	OptimizedCheckoutContentSecondaryColor                string   `json:"optimizedCheckout-contentSecondary-color"`
	OptimizedCheckoutContentSecondaryFont                 string   `json:"optimizedCheckout-contentSecondary-font"`
	OptimizedCheckoutButtonPrimaryFont                    string   `json:"optimizedCheckout-buttonPrimary-font"`
	OptimizedCheckoutButtonPrimaryColor                   string   `json:"optimizedCheckout-buttonPrimary-color"`
	OptimizedCheckoutButtonPrimaryColorHover              string   `json:"optimizedCheckout-buttonPrimary-colorHover"`
	OptimizedCheckoutButtonPrimaryColorActive             string   `json:"optimizedCheckout-buttonPrimary-colorActive"`
	OptimizedCheckoutButtonPrimaryBackgroundColor         string   `json:"optimizedCheckout-buttonPrimary-backgroundColor"`
	OptimizedCheckoutButtonPrimaryBackgroundColorHover    string   `json:"optimizedCheckout-buttonPrimary-backgroundColorHover"`
	OptimizedCheckoutButtonPrimaryBackgroundColorActive   string   `json:"optimizedCheckout-buttonPrimary-backgroundColorActive"`
	OptimizedCheckoutButtonPrimaryBorderColor             string   `json:"optimizedCheckout-buttonPrimary-borderColor"`
	OptimizedCheckoutButtonPrimaryBorderColorHover        string   `json:"optimizedCheckout-buttonPrimary-borderColorHover"`
	OptimizedCheckoutButtonPrimaryBorderColorActive       string   `json:"optimizedCheckout-buttonPrimary-borderColorActive"`
	OptimizedCheckoutButtonPrimaryBorderColorDisabled     string   `json:"optimizedCheckout-buttonPrimary-borderColorDisabled"`
	OptimizedCheckoutButtonPrimaryBackgroundColorDisabled string   `json:"optimizedCheckout-buttonPrimary-backgroundColorDisabled"`
	OptimizedCheckoutButtonPrimaryColorDisabled           string   `json:"optimizedCheckout-buttonPrimary-colorDisabled"`
	OptimizedCheckoutFormChecklistBackgroundColor         string   `json:"optimizedCheckout-formChecklist-backgroundColor"`
	OptimizedCheckoutFormChecklistColor                   string   `json:"optimizedCheckout-formChecklist-color"`
	OptimizedCheckoutFormChecklistBorderColor             string   `json:"optimizedCheckout-formChecklist-borderColor"`
	OptimizedCheckoutFormChecklistBackgroundColorSelected string   `json:"optimizedCheckout-formChecklist-backgroundColorSelected"`
	OptimizedCheckoutButtonSecondaryFont                  string   `json:"optimizedCheckout-buttonSecondary-font"`
	OptimizedCheckoutButtonSecondaryColor                 string   `json:"optimizedCheckout-buttonSecondary-color"`
	OptimizedCheckoutButtonSecondaryColorHover            string   `json:"optimizedCheckout-buttonSecondary-colorHover"`
	OptimizedCheckoutButtonSecondaryColorActive           string   `json:"optimizedCheckout-buttonSecondary-colorActive"`
	OptimizedCheckoutButtonSecondaryBackgroundColor       string   `json:"optimizedCheckout-buttonSecondary-backgroundColor"`
	OptimizedCheckoutButtonSecondaryBorderColor           string   `json:"optimizedCheckout-buttonSecondary-borderColor"`
	OptimizedCheckoutButtonSecondaryBackgroundColorHover  string   `json:"optimizedCheckout-buttonSecondary-backgroundColorHover"`
	OptimizedCheckoutButtonSecondaryBackgroundColorActive string   `json:"optimizedCheckout-buttonSecondary-backgroundColorActive"`
	OptimizedCheckoutButtonSecondaryBorderColorHover      string   `json:"optimizedCheckout-buttonSecondary-borderColorHover"`
	OptimizedCheckoutButtonSecondaryBorderColorActive     string   `json:"optimizedCheckout-buttonSecondary-borderColorActive"`
	OptimizedCheckoutLinkColor                            string   `json:"optimizedCheckout-link-color"`
	OptimizedCheckoutLinkFont                             string   `json:"optimizedCheckout-link-font"`
	OptimizedCheckoutDiscountBannerBackgroundColor        string   `json:"optimizedCheckout-discountBanner-backgroundColor"`
	OptimizedCheckoutDiscountBannerTextColor              string   `json:"optimizedCheckout-discountBanner-textColor"`
	OptimizedCheckoutDiscountBannerIconColor              string   `json:"optimizedCheckout-discountBanner-iconColor"`
	OptimizedCheckoutOrderSummaryBackgroundColor          string   `json:"optimizedCheckout-orderSummary-backgroundColor"`
	OptimizedCheckoutOrderSummaryBorderColor              string   `json:"optimizedCheckout-orderSummary-borderColor"`
	OptimizedCheckoutStepBackgroundColor                  string   `json:"optimizedCheckout-step-backgroundColor"`
	OptimizedCheckoutStepTextColor                        string   `json:"optimizedCheckout-step-textColor"`
	OptimizedCheckoutFormTextColor                        string   `json:"optimizedCheckout-form-textColor"`
	OptimizedCheckoutFormFieldBorderColor                 string   `json:"optimizedCheckout-formField-borderColor"`
	OptimizedCheckoutFormFieldTextColor                   string   `json:"optimizedCheckout-formField-textColor"`
	OptimizedCheckoutFormFieldShadowColor                 string   `json:"optimizedCheckout-formField-shadowColor"`
	OptimizedCheckoutFormFieldPlaceholderColor            string   `json:"optimizedCheckout-formField-placeholderColor"`
	OptimizedCheckoutFormFieldBackgroundColor             string   `json:"optimizedCheckout-formField-backgroundColor"`
	OptimizedCheckoutFormFieldErrorColor                  string   `json:"optimizedCheckout-formField-errorColor"`
	OptimizedCheckoutFormFieldInputControlColor           string   `json:"optimizedCheckout-formField-inputControlColor"`
	OptimizedCheckoutStepBorderColor                      string   `json:"optimizedCheckout-step-borderColor"`
	OptimizedCheckoutHeaderBorderColor                    string   `json:"optimizedCheckout-header-borderColor"`
	OptimizedCheckoutHeaderTextColor                      string   `json:"optimizedCheckout-header-textColor"`
	OptimizedCheckoutLoadingToasterBackgroundColor        string   `json:"optimizedCheckout-loadingToaster-backgroundColor"`
	OptimizedCheckoutLoadingToasterTextColor              string   `json:"optimizedCheckout-loadingToaster-textColor"`
	OptimizedCheckoutLinkHoverColor                       string   `json:"optimizedCheckout-link-hoverColor"`
	ProductSaleBadges                                     string   `json:"product_sale_badges"`
	ColorBadgeProductSaleBadges                           string   `json:"color_badge_product_sale_badges"`
	ColorTextProductSaleBadges                            string   `json:"color_text_product_sale_badges"`
	ColorHoverProductSaleBadges                           string   `json:"color_hover_product_sale_badges"`
	ProductSoldOutBadges                                  string   `json:"product_sold_out_badges"`
	ColorBadgeProductSoldOutBadges                        string   `json:"color_badge_product_sold_out_badges"`
	ColorTextProductSoldOutBadges                         string   `json:"color_text_product_sold_out_badges"`
	ColorHoverProductSoldOutBadges                        string   `json:"color_hover_product_sold_out_badges"`
	FocusTooltipTextColor                                 string   `json:"focusTooltip-textColor"`
	FocusTooltipBackgroundColor                           string   `json:"focusTooltip-backgroundColor"`
	RestrictToLogin                                       bool     `json:"restrict_to_login"`
	SwatchOptionSize                                      string   `json:"swatch_option_size"`
	SocialIconPlacementTop                                bool     `json:"social_icon_placement_top"`
	SocialIconPlacementBottom                             string   `json:"social_icon_placement_bottom"`
	NavigationDesign                                      string   `json:"navigation_design"`
	PriceRanges                                           bool     `json:"price_ranges"`
	PdpPriceLabel                                         string   `json:"pdp-price-label"`
	PdpSaleBadgeLabel                                     string   `json:"pdp_sale_badge_label"`
	PdpSoldOutLabel                                       string   `json:"pdp_sold_out_label"`
	PdpSalePriceLabel                                     string   `json:"pdp-sale-price-label"`
	PdpNonSalePriceLabel                                  string   `json:"pdp-non-sale-price-label"`
	PdpRetailPriceLabel                                   string   `json:"pdp-retail-price-label"`
	PdpCustomFieldsTabLabel                               string   `json:"pdp-custom-fields-tab-label"`
	PaymentbuttonsPaypalLayout                            string   `json:"paymentbuttons-paypal-layout"`
	PaymentbuttonsPaypalColor                             string   `json:"paymentbuttons-paypal-color"`
	PaymentbuttonsPaypalShape                             string   `json:"paymentbuttons-paypal-shape"`
	PaymentbuttonsPaypalLabel                             string   `json:"paymentbuttons-paypal-label"`
	PaymentbannersHomepageColor                           string   `json:"paymentbanners-homepage-color"`
	PaymentbannersHomepageRatio                           string   `json:"paymentbanners-homepage-ratio"`
	PaymentbannersCartpageTextColor                       string   `json:"paymentbanners-cartpage-text-color"`
	PaymentbannersCartpageLogoPosition                    string   `json:"paymentbanners-cartpage-logo-position"`
	PaymentbannersCartpageLogoType                        string   `json:"paymentbanners-cartpage-logo-type"`
	PaymentbannersProddetailspageColor                    string   `json:"paymentbanners-proddetailspage-color"`
	PaymentbannersProddetailspageRatio                    string   `json:"paymentbanners-proddetailspage-ratio"`
	PaymentbuttonsContainer                               string   `json:"paymentbuttons-container"`
	SupportedCardTypeIcons                                []string `json:"supported_card_type_icons"`
	SupportedPaymentMethods                               []string `json:"supported_payment_methods"`
	LazyloadMode                                          string   `json:"lazyload_mode"`
	CheckoutPaymentbuttonsPaypalColor                     string   `json:"checkout-paymentbuttons-paypal-color"`
	CheckoutPaymentbuttonsPaypalShape                     string   `json:"checkout-paymentbuttons-paypal-shape"`
	CheckoutPaymentbuttonsPaypalSize                      string   `json:"checkout-paymentbuttons-paypal-size"`
	CheckoutPaymentbuttonsPaypalLabel                     string   `json:"checkout-paymentbuttons-paypal-label"`
}
```

StoreSettings are the settings for a store look and feel

#### type Theme

```go
type Theme struct {
	UUID       string `json:"uuid"`
	Variations []struct {
		UUID        string `json:"uuid"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ExternalID  string `json:"external_id"`
	} `json:"variations"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
	IsActive  bool   `json:"is_active"`
}
```

Theme is the theme object for BigCommerce stores

#### type ThemeConfig

```go
type ThemeConfig struct {
	UUID          string        `json:"uuid"`
	StoreHash     string        `json:"store_hash"`
	ChannelID     int64         `json:"channel_id"`
	Settings      StoreSettings `json:"settings"`
	ThemeUUID     string        `json:"theme_uuid"`
	VersionUUID   string        `json:"version_uuid"`
	VariationUUID string        `json:"variation_uuid"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
```

ThemeConfig represents the configuration for a BigCommerce theme

#### type UserPart

```go
type UserPart struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
```

UserPart is a BigCommerce user shorthand object type that's in many other
responses

#### type Webhook

```go
type Webhook struct {
	ID          int64             `json:"id"`
	ClientID    string            `json:"client_id"`
	StoreHash   string            `json:"store_hash"`
	CreatedAt   int64             `json:"created_at"`
	UpdatedAt   int64             `json:"updated_at"`
	Scope       string            `json:"scope"`
	Destination string            `json:"destination"`
	IsActive    bool              `json:"is_active"`
	Headers     map[string]string `json:"headers"`
}
```


#### type WebhookPayload

```go
type WebhookPayload struct {
	Scope   string `json:"scope"`
	StoreID string `json:"store_id"`
	Data    struct {
		Type     string `json:"type"`
		ID       int64  `json:"id"`
		CouponID string `json:"couponId"`
		CartID   string `json:"cartId"`
		OrderID  int64  `json:"orderId"`
		Address  struct {
			CustomerID int64 `json:"customer_id"`
		} `json:"address"`
		Inventory InventoryEntry `json:"inventory"`
		Message   struct {
			OrderMessageID int64 `json:"order_message_id"`
		} `json:"message"`
		Sku struct {
			ProductID int64 `json:"product_id"`
			VariantID int64 `json:"variant_id"`
		} `json:"sku"`
		Status struct {
			PreviousStatusID int64 `json:"previous_status_id"`
			NewStatusID      int64 `json:"new_status_id"`
		} `json:"status"`
	} `json:"data"`
	Hash      string `json:"hash"`
	CreatedAt int64  `json:"created_at"`
	Producer  string `json:"producer"`
}
```


#### func  GetWebhookPayload

```go
func GetWebhookPayload(r *http.Request) (*WebhookPayload, []byte, error)
```
GetWebhookPayload returns a WebhookPayload object and the raw payload from the
BigCommerce API Arguments: r - the http.Request object Returns: *WebhookPayload
- the WebhookPayload object []byte - the raw payload from the BigCommerce API
error - the error, if any
