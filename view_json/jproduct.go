package viewjson

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type (
	//JProduct contains json implementation of product
	JProduct struct {
		ID          string `json:"id"`
		URL         string `json:"url"`
		Title       string `json:"title"`
		Price       string `json:"price"`
		Currency    string `json:"currency"`
		ImgURL      string `json:"img_url"`
		CreatedAt   string `json:"created_at"`
		LastTrackAt string `json:"last_track_at"`
	}
)

//JSON returns jProduct to JSON Marshalling
func (jp *JProduct) JSON() []byte {
	b, e := json.Marshal(jp)
	if e != nil {
		log.Fatal(e)
	}
	return b
}
