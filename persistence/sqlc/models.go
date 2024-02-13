// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"database/sql"
)

type Product struct {
	ID            int32
	SubcategoryID int32
	Name          string
	Description   sql.NullString
	Price         string
	Cost          string
	Quantity      int32
	SizeCode      string
	ColorID       int32
	Brand         string
	Sku           string
	Upc           string
	ImageUrl      string
	SourceID      int32
	SourceUrl     sql.NullString
	IsOffered     bool
	OfferPercent  sql.NullInt32
	IsActive      bool
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
}

type ProductCategory struct {
	ID          int32
	Name        string
	Description string
}

type ProductColor struct {
	ID   int32
	Name string
	Hex  string
}

type ProductInOrder struct {
	OrderID   int32
	ProductID int32
}

type ProductSize struct {
	Code        string
	Description string
}

type ProductSource struct {
	ID   int32
	Name string
}

type ProductSubcategory struct {
	ID               int32
	ParentCategoryID int32
	Name             string
	Description      string
}

type ProductView struct {
	ProductID                   int32
	CategoryID                  int32
	CategoryName                string
	CategoryDescription         string
	SubcategoryID               int32
	SubcategoryParentCategoryID int32
	SubcategoryName             string
	SubcategoryDescription      string
	ProductName                 string
	ProductDescription          sql.NullString
	ProductPrice                string
	ProductCost                 string
	ProductQuantity             int32
	SizeCode                    string
	SizeDescription             string
	ColorID                     int32
	ColorName                   string
	ColorHex                    string
	ProductBrand                string
	ProductSku                  string
	ProductUpc                  string
	ProductImageUrl             string
	SourceID                    int32
	SourceName                  string
	ProductSourceUrl            sql.NullString
	ProductIsOffered            bool
	ProductOfferPercent         sql.NullInt32
	ProductIsActive             bool
	ProductCreatedAt            sql.NullTime
	ProductUpdatedAt            sql.NullTime
}

type ProductsOrder struct {
	ID            int32
	WeightPayment string
	Taxes         string
	TotalAmount   string
	ArrivedAt     sql.NullTime
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
}