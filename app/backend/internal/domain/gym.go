package domain

type Gym struct {
	ID        string  `bson:"_id,omitempty" json:"id"`
	Name      string  `bson:"name" json:"name"`
	Lat       float64 `bson:"lat" json:"lat"`
	Lng       float64 `bson:"lng" json:"lng"`
	Province  string  `bson:"province" json:"province"`
	CreatedAt int64   `bson:"created_at" json:"createdAt"`
}
