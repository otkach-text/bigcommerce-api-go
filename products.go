package bigcommerce

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// include_fields
var productFields = []string{"name", "sku", "custom_url", "is_visible", "price"}

// include (subresources, like variants images custom_fields bulk_pricing_rules primary_image modifiers options videos)
var productInclude []string

// extra arguments for product interface
var productArgs map[string]string

// Product is a BigCommerce product object
type Product struct {
	ID                      int64         `json:"id,omitempty"`
	Name                    string        `json:"name,omitempty"`
	Type                    string        `json:"type,omitempty"`
	Sku                     string        `json:"sku,omitempty"`
	Description             string        `json:"description,omitempty"`
	Weight                  float64       `json:"weight,omitempty"`
	Width                   float64       `json:"width,omitempty"`
	Depth                   float64       `json:"depth,omitempty"`
	Height                  float64       `json:"height,omitempty"`
	Price                   float64       `json:"price,omitempty"`
	CostPrice               float64       `json:"cost_price,omitempty"`
	RetailPrice             float64       `json:"retail_price,omitempty"`
	SalePrice               float64       `json:"sale_price,omitempty"`
	MapPrice                float64       `json:"map_price,omitempty"`
	TaxClassID              int64         `json:"tax_class_id,omitempty"`
	ProductTaxCode          string        `json:"product_tax_code,omitempty"`
	CalculatedPrice         float64       `json:"calculated_price,omitempty"`
	Categories              []interface{} `json:"categories,omitempty"`
	BrandID                 int64         `json:"brand_id,omitempty"`
	OptionSetID             interface{}   `json:"option_set_id,omitempty"`
	OptionSetDisplay        string        `json:"option_set_display,omitempty"`
	InventoryLevel          int           `json:"inventory_level,omitempty"`
	InventoryWarningLevel   int           `json:"inventory_warning_level,omitempty"`
	InventoryTracking       string        `json:"inventory_tracking,omitempty"`
	ReviewsRatingSum        int           `json:"reviews_rating_sum,omitempty"`
	ReviewsCount            int           `json:"reviews_count,omitempty"`
	TotalSold               int           `json:"total_sold,omitempty"`
	FixedCostShippingPrice  float64       `json:"fixed_cost_shipping_price,omitempty"`
	IsFreeShipping          bool          `json:"is_free_shipping,omitempty"`
	IsVisible               bool          `json:"is_visible,omitempty"`
	IsFeatured              bool          `json:"is_featured,omitempty"`
	RelatedProducts         []int         `json:"related_products,omitempty"`
	Warranty                string        `json:"warranty,omitempty"`
	BinPickingNumber        string        `json:"bin_picking_number,omitempty"`
	LayoutFile              string        `json:"layout_file,omitempty"`
	Upc                     string        `json:"upc,omitempty"`
	Mpn                     string        `json:"mpn,omitempty"`
	Gtin                    string        `json:"gtin,omitempty"`
	SearchKeywords          string        `json:"search_keywords,omitempty"`
	Availability            string        `json:"availability,omitempty"`
	AvailabilityDescription string        `json:"availability_description,omitempty"`
	GiftWrappingOptionsType string        `json:"gift_wrapping_options_type,omitempty"`
	GiftWrappingOptionsList []interface{} `json:"gift_wrapping_options_list,omitempty"`
	SortOrder               int           `json:"sort_order,omitempty"`
	Condition               string        `json:"condition,omitempty"`
	IsConditionShown        bool          `json:"is_condition_shown,omitempty"`
	OrderQuantityMinimum    int           `json:"order_quantity_minimum,omitempty"`
	OrderQuantityMaximum    int           `json:"order_quantity_maximum,omitempty"`
	PageTitle               string        `json:"page_title,omitempty"`
	MetaKeywords            []interface{} `json:"meta_keywords,omitempty"`
	MetaDescription         string        `json:"meta_description,omitempty"`
	DateCreated             time.Time     `json:"date_created,omitempty"`
	DateModified            time.Time     `json:"date_modified,omitempty"`
	ViewCount               int           `json:"view_count,omitempty"`
	PreorderReleaseDate     interface{}   `json:"preorder_release_date,omitempty"`
	PreorderMessage         string        `json:"preorder_message,omitempty"`
	IsPreorderOnly          bool          `json:"is_preorder_only,omitempty"`
	IsPriceHidden           bool          `json:"is_price_hidden,omitempty"`
	PriceHiddenLabel        string        `json:"price_hidden_label,omitempty"`
	CustomURL               struct {
		URL          string `json:"url,omitempty"`
		IsCustomized bool   `json:"is_customized,omitempty"`
	} `json:"custom_url,omitempty"`
	BaseVariantID               int64  `json:"base_variant_id,omitempty"`
	OpenGraphType               string `json:"open_graph_type,omitempty"`
	OpenGraphTitle              string `json:"open_graph_title,omitempty"`
	OpenGraphDescription        string `json:"open_graph_description,omitempty"`
	OpenGraphUseMetaDescription bool   `json:"open_graph_use_meta_description,omitempty"`
	OpenGraphUseProductName     bool   `json:"open_graph_use_product_name,omitempty"`
	OpenGraphUseImage           bool   `json:"open_graph_use_image,omitempty"`
	Variants                    []struct {
		ID                        int64         `json:"id,omitempty"`
		ProductID                 int64         `json:"product_id,omitempty"`
		Sku                       string        `json:"sku,omitempty"`
		SkuID                     interface{}   `json:"sku_id,omitempty"`
		Price                     float64       `json:"price,omitempty"`
		CalculatedPrice           float64       `json:"calculated_price,omitempty"`
		SalePrice                 float64       `json:"sale_price,omitempty"`
		RetailPrice               float64       `json:"retail_price,omitempty"`
		MapPrice                  float64       `json:"map_price,omitempty"`
		Weight                    float64       `json:"weight,omitempty"`
		Width                     int           `json:"width,omitempty"`
		Height                    int           `json:"height,omitempty"`
		Depth                     int           `json:"depth,omitempty"`
		IsFreeShipping            bool          `json:"is_free_shipping,omitempty"`
		FixedCostShippingPrice    float64       `json:"fixed_cost_shipping_price,omitempty"`
		CalculatedWeight          float64       `json:"calculated_weight,omitempty"`
		PurchasingDisabled        bool          `json:"purchasing_disabled,omitempty"`
		PurchasingDisabledMessage string        `json:"purchasing_disabled_message,omitempty"`
		ImageURL                  string        `json:"image_url,omitempty"`
		CostPrice                 float64       `json:"cost_price,omitempty"`
		Upc                       string        `json:"upc,omitempty"`
		Mpn                       string        `json:"mpn,omitempty"`
		Gtin                      string        `json:"gtin,omitempty"`
		InventoryLevel            int           `json:"inventory_level,omitempty"`
		InventoryWarningLevel     int           `json:"inventory_warning_level,omitempty"`
		BinPickingNumber          string        `json:"bin_picking_number,omitempty"`
		OptionValues              []interface{} `json:"option_values,omitempty"`
	} `json:"variants,omitempty"`
	Images           []Image           `json:"images,omitempty"`
	PrimaryImage     interface{}       `json:"primary_image,omitempty"`
	Videos           []Video           `json:"videos,omitempty"`
	CustomFields     []CustomField     `json:"custom_fields,omitempty"`
	BulkPricingRules []BulkPricingRule `json:"bulk_pricing_rules,omitempty"`
	Options          []interface{}     `json:"options,omitempty"`
	Modifiers        []interface{}     `json:"modifiers,omitempty"`
}

type CustomField struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty" validate:"required"`
	Value string `json:"value,omitempty" validate:"required"`
}

type BulkPricingRule struct {
	ID          int64  `json:"id" validate:"required"`
	QuantityMin int    `json:"quantity_min" validate:"required"`
	QuantityMax int    `json:"quantity_max" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Amount      string `json:"amount" validate:"required"`
}

// Metafield is a struct representing a BigCommerce product metafield
type Metafield struct {
	ID            int64     `json:"id,omitempty"`
	Key           string    `json:"key,omitempty"`
	Value         string    `json:"value,omitempty"`
	ResourceID    int64     `json:"resource_id,omitempty"`
	ResourceType  string    `json:"resource_type,omitempty"`
	Description   string    `json:"description,omitempty"`
	DateCreated   time.Time `json:"date_created,omitempty"`
	DateModified  time.Time `json:"date_modified,omitempty"`
	Namespace     string    `json:"namespace,omitempty"`
	PermissionSet string    `json:"permission_set,omitempty"`
}

type CreateProductPayload struct {
	Name                     string            `json:"name" validate:"required"`
	Type                     string            `json:"type" validate:"required"`
	SKU                      *string           `json:"sku,omitempty"`
	Description              *string           `json:"description,omitempty"`
	Weight                   *float64          `json:"weight" validate:"required"`
	Width                    *float64          `json:"width,omitempty"`
	Depth                    *float64          `json:"depth,omitempty"`
	Height                   *float64          `json:"height,omitempty"`
	Price                    *float64          `json:"price" validate:"required"`
	CostPrice                *float64          `json:"cost_price,omitempty"`
	RetailPrice              *float64          `json:"retail_price,omitempty"`
	SalePrice                *float64          `json:"sale_price,omitempty"`
	InventoryLevel           *int              `json:"inventory_level,omitempty"`
	InventoryWarning         *int              `json:"inventory_warning_level,omitempty"`
	InventoryTracking        string            `json:"inventory_tracking,omitempty"`
	Availability             string            `json:"availability,omitempty"`
	AvailabilityDescription  *string           `json:"availability_description,omitempty"`
	GiftWrappingOptionsType  *string           `json:"gift_wrapping_options_type,omitempty"`
	GiftWrappingOptionsList  []int             `json:"gift_wrapping_options_list,omitempty"`
	SortOrder                *int              `json:"sort_order,omitempty"`
	Condition                *string           `json:"condition,omitempty"`
	IsConditionShown         *bool             `json:"is_condition_shown,omitempty"`
	Categories               []int             `json:"categories,omitempty"`
	BrandID                  *int              `json:"brand_id,omitempty"`
	MetaKeywords             *[]string         `json:"meta_keywords,omitempty"`
	MetaDescription          *string           `json:"meta_description,omitempty"`
	Images                   []Image           `json:"images,omitempty"`
	Videos                   []Video           `json:"videos,omitempty"`
	CustomFields             []CustomField     `json:"custom_fields,omitempty"`
	BulkPricingRules         []BulkPricingRule `json:"bulk_pricing_rules,omitempty"`
	OptionSetID              *int              `json:"option_set_id,omitempty"`
	OptionSetDisplay         *string           `json:"option_set_display,omitempty"`
	UPC                      *string           `json:"upc,omitempty"`
	SearchKeywords           *string           `json:"search_keywords,omitempty"`
	TaxClassID               *int              `json:"tax_class_id,omitempty"`
	ViewCount                *int              `json:"view_count,omitempty"`
	PreorderReleaseDate      *string           `json:"preorder_release_date,omitempty"`
	PreorderMessage          *string           `json:"preorder_message,omitempty"`
	OrderQuantityMinimum     *int              `json:"order_quantity_minimum,omitempty"`
	OrderQuantityMaximum     *int              `json:"order_quantity_maximum,omitempty"`
	PageTitle                *string           `json:"page_title,omitempty"`
	IsVisible                *bool             `json:"is_visible,omitempty"`
	IsFeatured               *bool             `json:"is_featured,omitempty"`
	Warranty                 *string           `json:"warranty,omitempty"`
	BinPickingNumber         *string           `json:"bin_picking_number,omitempty"`
	LayoutFile               *string           `json:"layout_file,omitempty"`
	UpSellingRelatedProducts *[]int            `json:"up_selling_related_products,omitempty"`
	EventDateFieldName       *string           `json:"event_date_field_name,omitempty"`
	EventDateType            *string           `json:"event_date_type,omitempty"`
	EventDateStart           *string           `json:"event_date_start,omitempty"`
	EventDateEnd             *string           `json:"event_date_end,omitempty"`
	MyobAssetAccount         *string           `json:"myob_asset_account,omitempty"`
	MyobExpenseAccount       *string           `json:"myob_expense_account,omitempty"`
	MyobIncomeAccount        *string           `json:"myob_income_account,omitempty"`
	XeroSalesAccount         *string           `json:"xero_sales_account,omitempty"`
	XeroSalesTaxType         *string           `json:"xero_sales_tax_type,omitempty"`
	XeroPurchaseAccount      *string           `json:"xero_purchase_account,omitempty"`
	XeroPurchaseTaxType      *string           `json:"xero_purchase_tax_type,omitempty"`
}

// GetAllProducts gets all products from BigCommerce
// args is a key-value map of additional arguments to pass to the API
func (bc *Client) CreateProduct(payload *CreateProductPayload) (*Product, error) {
	var b []byte
	b, _ = json.Marshal([]CreateProductPayload{*payload})
	req := bc.getAPIRequest(http.MethodPost, "/v3/catalog/products", bytes.NewBuffer(b))
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		if res.StatusCode == http.StatusUnprocessableEntity {
			var errResp ErrorResult
			err = json.Unmarshal(body, &errResp)
			if err != nil {
				log.Printf("Error: %s\nResult: %s", err, string(body))
				return nil, err
			}
			if len(errResp.Errors) > 0 {
				errors := []string{}
				for _, e := range errResp.Errors {
					errors = append(errors, e)
				}
				return nil, fmt.Errorf("%s", strings.Join(errors, ", "))
			}
			return nil, errors.New("unknown error")
		}
		log.Printf("Error: %s\nResult: %s", err, string(body))
		return nil, err
	}
	var ret struct {
		Product Product `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	return &ret.Product, nil
}

// GetAllProducts gets all products from BigCommerce
// args is a key-value map of additional arguments to pass to the API
func (bc *Client) GetAllProducts(args map[string]string) ([]Product, error) {
	ps := []Product{}
	var psp []Product
	page := 1
	more := true
	var err error
	retries := 0
	for more {
		psp, more, err = bc.GetProducts(args, page)
		// log.Printf("page %d entries %d", page, len(psp))
		if err != nil {
			retries++
			if retries > bc.MaxRetries {
				log.Println("Max retries reached")
				return ps, err
			}
			break
		}
		ps = append(ps, psp...)
		page++
	}
	return ps, err
}

// GetProducts gets a page of products from BigCommerce
// args is a key-value map of additional arguments to pass to the API
// page: the page number to download
func (bc *Client) GetProducts(args map[string]string, page int) ([]Product, bool, error) {
	fpart := ""
	for k, v := range args {
		fpart += "&" + k + "=" + v
	}
	url := "/v3/catalog/products?page=" + strconv.Itoa(page) + fpart
	// log.Printf("GET %s", url)

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNoContent {
		return nil, false, ErrNoContent
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, false, err
	}
	var pp struct {
		Status int       `json:"status"`
		Title  string    `json:"title"`
		Data   []Product `json:"data"`
		Meta   struct {
			Pagination Pagination `json:"pagination"`
		} `json:"meta"`
	}
	err = json.Unmarshal(body, &pp)
	if err != nil {
		return nil, false, err
	}
	//	log.Printf("%d products (%+v)", len(pp.Data), pp.Meta.Pagination)

	if pp.Status != 0 {
		return nil, false, errors.New(pp.Title)
	}
	return pp.Data, pp.Meta.Pagination.CurrentPage < pp.Meta.Pagination.TotalPages, nil
}

// GetProductByID gets a product from BigCommerce by ID
// productID: BigCommerce product ID to get
func (bc *Client) GetProductByID(productID int64) (*Product, error) {
	url := "/v3/catalog/products/" + strconv.FormatInt(productID, 10) + "?include=variants,images,custom_fields,bulk_pricing_rules,primary_image,modifiers,options,videos"
	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var productResponse struct {
		Data Product `json:"data"`
	}
	err = json.Unmarshal(body, &productResponse)
	if err != nil {
		return nil, err
	}
	return &productResponse.Data, nil
}

// GetProductMetafields gets metafields values for a product
// productID: BigCommerce product ID to get metafields for
func (bc *Client) GetProductMetafields(productID int64) (map[string]Metafield, error) {
	url := "/v3/catalog/products/" + strconv.FormatInt(productID, 10) + "/metafields"
	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var metafieldsResponse struct {
		Metafields []Metafield `json:"data,omitempty"`
	}
	err = json.Unmarshal(body, &metafieldsResponse)
	if err != nil {
		return nil, err
	}
	ret := map[string]Metafield{}
	for _, mf := range metafieldsResponse.Metafields {
		ret[mf.Key] = mf
	}
	return ret, nil
}
