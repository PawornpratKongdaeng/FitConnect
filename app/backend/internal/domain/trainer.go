package domain

type Trainer struct {
	ID          string   `bson:"_id,omitempty" json:"id"`
	UserID      string   `bson:"user_id" json:"userId"`
	Bio         string   `bson:"bio" json:"bio"`
	Specialties []string `bson:"specialties" json:"specialties"`
	Rate        int      `bson:"rate" json:"rate"`
	CertStatus  string   `bson:"cert_status" json:"certStatus"`
	RegionCode  string   `bson:"region_code" json:"regionCode"`
	Location    any      `bson:"location" json:"location"`
	CreatedAt   int64    `bson:"created_at" json:"createdAt"`
}
