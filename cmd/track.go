package main


type SnowTrack struct {
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
}
