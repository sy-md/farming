package main

type RDY_plants struct {
	RdyPlt     string
	RdyPlt_amt int
}
type GRW_plants struct {
	GrwPlt     string
	GrwPlt_amt int
}
type Store struct {
	Iventory []RDY_plants `json:"Iventory"`
	SeedIvn  []GRW_plants `json:"seedIvn"`
}
