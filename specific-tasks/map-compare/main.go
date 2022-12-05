package main

import (
	"time"
)

type Object struct {
	ID         string     `json:"id" db:"id"`
	Name       *string    `json:"name" db:"name"`
	Address    *string    `json:"address" db:"address"`
	Categories Categories `json:"categories" db:"categories"`
	Point      *Point     `json:"point" db:"point"`
	UpdatedAt  time.Time  `db:"updated_at"`
}
type Categories []string

type Point struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func RemoveDuplWithOk() {
	objectDupl := append(getObject(), getObject()[len(getObject())/3:]...)

	name := "sdflhjn"
	address := "askdjhasdhgiustyd736w478134ajksdh"

	noDuplicates := make(map[string]Object)

	for _, v := range objectDupl {
		_, ok := noDuplicates[v.ID]
		if !ok {
			noDuplicates[v.ID] = Object{
				ID:      "asd",
				Name:    &name,
				Address: &address,
				Point: &Point{
					Latitude:  12.12398712,
					Longitude: 129.128973612435,
				},
				UpdatedAt:  time.Now(),
				Categories: []string{"asd", "kqywged"},
			}
		}
	}
}

func RemoveDupl() {
	objectDupl := append(getObject(), getObject()[len(getObject())/3:]...)

	noDuplicates := make(map[string]Object)

	for i, v := range objectDupl {
		noDuplicates[v.ID] = objectDupl[i]
	}
}

func main() {
	RemoveDuplWithOk()
	RemoveDupl()
}

func getObject() []Object {
	name := "qasasdqsadad"
	address := "alishdasduohasohdoahsdhouqye823y479yedhd2389"

	o := []Object{{ID: "79317063646", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "209832055133", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "38744539350", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "192602870517", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "129228125236", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "120582801305", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "53790808879", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "220598002304", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "12490342623", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "196743108935", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "24802214082", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "65700038405", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "243083506667", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "28235256535", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "178547430873", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "17347017202", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "2471145385", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "8843420952", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "200574131361", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "3166424715", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "197376313761", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "154377591314", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "178675338665", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "37459801772", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "15169243769", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "178454900411", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "37524005631", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "29079738959", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "193384751330", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "14393125382", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "133106190348", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "89036851723", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "26642563545", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "20829697766", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "22457491930", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "168572455359", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "132758475692", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "168720940116", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "174288089199", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "4907577245", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "147265108967", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "121474046309", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "231636881367", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "1354961225", Name: &name, Address: &address, Categories: Categories{"Управление городским транспортом и его обслуживание "}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "136383412989", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "208428374627", Name: &name, Address: &address, Categories: Categories{"Железнодорожная станция"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "99235743880", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "1130032794", Name: &name, Address: &address, Categories: Categories{"Управление городским транспортом и его обслуживание "}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "76279155500", Name: &name, Address: &address, Categories: Categories{"Станция метро"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "1345235194", Name: &name, Address: &address, Categories: Categories{"Управление городским транспортом и его обслуживание ", "Информационная служба"}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "1072240200", Name: &name, Address: &address, Categories: Categories{"Управление городским транспортом и его обслуживание ", "Управление железными дорогами и их обслуживание "}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "145209195268", Name: &name, Address: &address, Categories: Categories{"Управление городским транспортом и его обслуживание "}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}, {ID: "1191185025", Name: &name, Address: &address, Categories: Categories{"Управление городским транспортом и его обслуживание "}, Point: &Point{Latitude: 123.1298736213, Longitude: 28.1928637}, UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}}
	return o
}
