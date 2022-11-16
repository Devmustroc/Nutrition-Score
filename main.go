package main

import "fmt"

func main() {
	ns = GetnutritionalScore(nutritionalData{
		Energy:              EnergyKJ(),
		Sugars:              SugarGram(),
		SaturatedfattyAcids: SaturatedFattyAcids(),
		Sodium:              SodiumMilligram(),
		Fruits:              FruitsPercent(),
		Fibre:               FibreGram(),
		Protein:             ProteinGram(),
	}, food)
	fmt.Printf("Nutritional Score:%d", ns.Value)
}
