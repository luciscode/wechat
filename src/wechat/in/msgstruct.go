package in

type Usualmessage struct {
	MsgId        string `xml:"MsgId,omitempty"`
	ToUserName   string `xml:"ToUserName,omitempty"`
	FromUserName string `xml:"FromUserName,omitempty"`
	CreateTime   string `xml:"CreateTime,omitempty"`
	MsgType      string `xml:"MsgType,omitempty"`
	Content      string `xml:"Content,omitempty"`
	PicUrl       string `xml:"PicUrl,omitempty"`
	MediaId      string `xml:"MediaId,omitempty"`
	ThumbMediaId string `xml:"ThumbMediaId,omitempty"`
	Location_X   string `xml:"Location_X,omitempty"`
	Location_Y   string `xml:"Location_Y,omitempty"`
	Scale        string `xml:"Scale,omitempty"`
	Label        string `xml:"Label,omitempty"`
	Title        string `xml:"Title,omitempty"`
	Description  string `xml:"Description,omitempty"`
	Url          string `xml:"Url,omitempty"`
	Recognition  string `xml:"Recognition,omitempty"`
	Format       string `xml:"Format,omitempty"`
}
type Pushmessage struct {
	ToUserName   string `xml:"ToUserName,omitempty"`
	FromUserName string `xml:"FromUserName,omitempty"`
	CreateTime   string `xml:"CreateTime,omitempty"`
	MsgType      string `xml:"MsgType,omitempty"`
	Event        string `xml:"Event,omitempty"`
	EventKey     string `xml:"EventKey,omitempty"`
	Ticket       string `xml:"Ticket,omitempty"`
	Latitude     string `xml:"Latitude,omitempty"`
	Longitude    string `xml:"Longitude,omitempty"`
	WxPrecision  string `xml:"Precision,omitempty"`
}
