package responses

type Address struct {
	AddressLine1 string `param:"address_line_1" query:"address_line_1" form:"address_line_1" json:"address_line_1" xml:"address_line_1" validate:"required"`
	AddressLine2 string `param:"address_line_2" query:"address_line_2" form:"address_line_2" json:"address_line_2" xml:"address_line_2" validate:""`
	City         string `param:"city" query:"city" form:"city" json:"city" xml:"city" validate:"required"`
	PostCode     string `param:"post_code" query:"post_code" form:"post_code" json:"post_code" xml:"post_code" validate:"required"`
}
