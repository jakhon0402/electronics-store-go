package payload

type ProductDto struct {
	Name          string  `json:"name" validate:"required,min=5"`
	Title         string  `json:"title" validate:"required,min=10"`
	Price         float32 `json:"price" validate:"required"`
	Specification string  `json:"specification" validate:"required"`
}
