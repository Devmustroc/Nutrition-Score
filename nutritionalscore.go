package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	scoreTYpe ScoreType
}

type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcids float64
type SodiumMilligram float64
type FruitsPercent float64
type FibreGram float64
type ProteinGram float64

type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedfattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}

func GetnutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoint()
		fibrePoints := n.Fibre.GetPoint()

		negative = n.Energy.GetPoint() + n.Sugars.GetPoint() + n.SaturatedFattyAcids.GetPoint() + n.Sodium.GetPoint()
		positive = n.FibreGetPoint() + n.Protein.GetPoint() + fruitPoints + fibrePoints
	}
	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		Scoretype: st,
	}
}
