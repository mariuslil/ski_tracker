package main


type SnowTrack struct {
	TrackID string `json:"trackid"`
	Name string `json:"name"`
	Mountain string `json:"mountain"`
	Length float32 `json: "length"`
}

type TrackService interface {
	AddTrack(t *SnowTrack) error
	GetTrack(track string) (*SnowTrack, error)
}

type SnowTracksMongoDB struct {
	DatabaseURL string
	DatabaseName string
	SnowTracksCollectionName string
}

type SnowTrackStorage interface {
	Init()
	Get(id string) (SnowTrack, bool)
	Add(t SnowTrack) int
	Count() int
	getField(field string, id int) (string, bool)
	GetAll() []SnowTrack
}
