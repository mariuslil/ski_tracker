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

type TrackDB struct {
	tracks map[string]SnowTrack
}

func (db *TrackDB) Init() {
	db.tracks = make(map[string]SnowTrack)
}

func (db *TrackDB) Add(t SnowTrack) error {
	db.tracks[t.TrackID] = t
	return nil
}

func (db *TrackDB) Count() int {
	return len(db.tracks)
}

func (db *TrackDB) Get(ID string) (SnowTrack, bool) {
	t, ok := db.tracks[ID]
	return t, ok
}

func (db *TrackDB) GetAll() []SnowTrack {
	all := make([]SnowTrack, 0, db.Count())
	for _, s := range db.tracks {
		all = append(all, s)
	}

}