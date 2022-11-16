package main

import "fmt"

func main() {
	ns := GetnutritionalScore(nutritionalData{
		Energy:              EnergyFromkcal(),
		Sugars:              SugarGram(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium:              SodiumMIlligram(),
		Fruits:              FruitsPercent(),
		Fibre:               FibreGram(),
		Protein:             Protein(),
	}, food)

	fmt.Printf("Nutritional Score:%d", ns.Value)
}
