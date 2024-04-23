package lasagna

func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime == 0 {
		preparationTime = 2
	}
	return len(layers) * preparationTime
}

func Quantities(layers []string) (int, float64) {
	sauceLayer := int(0)
	noodlesLayer := float64(0)

	for _, layer := range layers {
		if layer == "sauce" {
			sauceLayer += 50
		}
		if layer == "noodles" {
			noodlesLayer += 0.2
		}
	}
	return sauceLayer, noodlesLayer
}

func AddSecretIngredient(friendIngredients []string, myIngredients []string) {
	myIngredients = append(myIngredients, friendIngredients...)
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	for i, amount := range amounts {
		amounts[i] = amount * float64(portions)
	}
	return amounts
}
