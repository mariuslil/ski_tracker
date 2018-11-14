package snowtrack


type SnowTracksMongoDB struct {
	DatabaseURL string
	DatabaseName string
	//TracksCollectionName string
}

type SnowTrackStorage interface {
	Init()
}

type SnowTrack struct {
	Mountain string `json:"Mountain"`

}