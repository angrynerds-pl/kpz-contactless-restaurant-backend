package requests

type AddFoodToMenuRequest struct {
	Food struct {
		Name        string `param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required"`
		Description string `param:"description" query:"description" form:"description" json:"description" xml:"description" validate:""`

		Address Address `param:"address" query:"address" form:"address" json:"address" xml:"address" validate:""`
	} `json:"restaurant"`
}

func (f AddFoodToMenuRequest) bind() {

}
