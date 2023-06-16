package models


type Relation struct{
	Id				int					 `json:"id"`
	DatesLocations	map[string][]string	 `json:"datesLocations"`
}

type Index struct{
	Indexes	[]Relation	`json:"index"`
}