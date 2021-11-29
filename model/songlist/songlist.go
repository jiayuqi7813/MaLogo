package songlist

type Songlist struct{
    Sid int `json:"sid"`
    Cover string `json:"cover"`
    Length int `json:"length"`
    Bpm float64 `json:"bpm"`
    Title string `json:"title"`
    Artist string `json:"artist"`
    Mode int `json:"mode"`
    Time int64 `json:"time"`
}
