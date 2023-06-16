package models

type Location struct{
	Id			int
  	Locations	[]string
  	Dates		string
}
type LocationOne struct{
	LocationOne	Location
  	Date		Date		
}
