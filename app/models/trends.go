package models

type ResponseTrends struct {
    Labels  []int
    Goals   map[string]int64
    Weights map[string]float64
    History map[string]int64
}
