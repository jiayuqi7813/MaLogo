
package charts

type Charts struct{
    Sid int     `json:"sid"`
    Cid int     `json:"cid"`
    Uid int     `json:"uid"`
    Creator string `json:"creator"`
    Version string  `json:"version"`
    Level string    `json:"level"`
    Type int        `json:"type"`
    Size int    `json:"size"`
    Mode int    `json:"mode"`
}
