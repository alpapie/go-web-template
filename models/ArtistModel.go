package models


type Artist struct{
	Id              int         `json:"id"`
    Image			string      `json:"image"`
    Name			string      `json:"name"`
    Members			[]string    `json:"members"`
    CreationDate	int         `json:"creationDate"`
    FirstAlbum		string      `json:"firstAlbum"`
    Locations		string      `json:"locations"`
    ConcertDates	string      `json:"concertDates"`
    Relations		string      `json:"relations"`
}
type ArtistOne struct{
    ArtistOne       Artist
    Locations		Location
    ConcertDates	Date
    Relations		Relation
}
// "id": 1,
// "image": "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
// "name": "Queen",
// "members": [
//   "Freddie Mercury",
//   "Brian May",
//   "John Daecon",
//   "Roger Meddows-Taylor",
//   "Mike Grose",
//   "Barry Mitchell",
//   "Doug Fogie"
// ],
// "creationDate": 1970,
// "firstAlbum": "14-12-1973",
// "locations": "https://groupietrackers.herokuapp.com/api/locations/1",
// "concertDates": "https://groupietrackers.herokuapp.com/api/dates/1",
// "relations": "https://groupietrackers.herokuapp.com/api/relation/1"