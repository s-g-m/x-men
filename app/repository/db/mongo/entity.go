package mongo

type Statistic struct {
	Id    bool `bson:"_id,omitempty"`
	Count int  `bson:"count,omitempty"`
}
