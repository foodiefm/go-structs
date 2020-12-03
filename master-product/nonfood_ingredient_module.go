package masterproduct

// NonfoodIngredientModule is a module providing Information on ingredients for
// items that are not food for example detergents, medicines.
type NonfoodIngredientModule struct {
	// Ingredient statement for non-food items.
	NonfoodIngredientStatements *[]MLString `json:"nonfoodIngredientStatement,omitempty"`
	// Specifies a non-food ingredient of concern for a trade item as a
	// code. Uses code list nonfoodIngredientOfConcernCode.
	NonfoodIngredientOfConcernCode *[]string `json:"nonfoodIngredientOfConcernCode,omitempty"`
	// Information on presence or absence of additives or genetic
	// modifications contained in the trade item.
	AdditiveInformations *[]NonFoodAdditiveInformation `json:"additiveInformation.omitempty"`
	// Information on ingredients for items that are not food for example
	// detergents, medicines.
	NonfoodIngredients *[]NonfoodIngredient `json:"nonfoodIngredient,omitempty"`
}

// NonFoodAdditiveInformation contains information on presence or absence of
// additives.
type NonFoodAdditiveInformation struct {
	// Name of additive ingredient.
	AdditiveName string `json:"additiveName"`
	// Code indicating the level of presence of the additive. Uses code
	// list levelOfContainmentCode.
	LevelOfContainmentCode string `json:"levelOfContainmentCode"`
}

// NonfoodIngredient contains information on ingredients for items that are not
// food for example detergents, medicines.
type NonfoodIngredient struct {
	// The name of the non-food ingredient.
	IngredientName *string `json:"ingredientName,omitempty"`
	// Denotes the nonfood ingredient that should have it's text emphasised
	// in some fashion on the item's packaging.
	IsNonfoodIngredientEmphasised *bool `json:"isNonfoodIngredientEmphasized,omitempty"`
	// Substring emphasis for ingredientName.
	XEmphases *[]XEmphasis `json:"x_emphasis"`
}
