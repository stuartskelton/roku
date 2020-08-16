package roku

import "encoding/xml"

// App has all the information of an application.
// https://github.com/kinghrothgar/roku/blob/master/roku/roku.go#L63
type App struct {
	Name    string `xml:",chardata" json:"name,omitempty"`
	ID      string `xml:"id,attr" json:"id,omitempty"`
	Type    string `xml:"type,attr" json:"type,omitempty"`
	SubType string `xml:"subtype,attr" json:"sub_type,omitempty"`
	Version string `xml:"version,attr" json:"version,omitempty"`
}

// App has all the information of an application.
// https://github.com/kinghrothgar/roku/blob/master/roku/roku.go#L63
type Player struct {
	XMLName xml.Name `xml:"player"`
	Text    string   `xml:",chardata"`
	Error   string   `xml:"error,attr"`
	State   string   `xml:"state,attr"`
	Plugin  struct {
		Text      string `xml:",chardata"`
		Bandwidth string `xml:"bandwidth,attr"`
		ID        string `xml:"id,attr"`
		Name      string `xml:"name,attr"`
	} `xml:"plugin"`
	Format struct {
		Text      string `xml:",chardata"`
		Audio     string `xml:"audio,attr"`
		Captions  string `xml:"captions,attr"`
		Container string `xml:"container,attr"`
		Drm       string `xml:"drm,attr"`
		Video     string `xml:"video,attr"`
	} `xml:"format"`
	Buffering struct {
		Text    string `xml:",chardata"`
		Current string `xml:"current,attr"`
		Max     string `xml:"max,attr"`
		Target  string `xml:"target,attr"`
	} `xml:"buffering"`
	NewStream struct {
		Text  string `xml:",chardata"`
		Speed string `xml:"speed,attr"`
	} `xml:"new_stream"`
	Position      string `xml:"position"`
	Duration      string `xml:"duration"`
	IsLive        string `xml:"is_live"`
	StreamSegment struct {
		Text          string `xml:",chardata"`
		Bitrate       string `xml:"bitrate,attr"`
		MediaSequence string `xml:"media_sequence,attr"`
		SegmentType   string `xml:"segment_type,attr"`
		Time          string `xml:"time,attr"`
	} `xml:"stream_segment"`
}
