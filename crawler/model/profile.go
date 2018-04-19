package model

import "encoding/json"

// Profile define a user profile
type Profile struct {
	Name          string
	Gender        string
	Age           int
	Height        int
	Weight        int
	Income        string
	Marriage      string
	Education     string
	Occupation    string
	Hukou         string // 户口
	Constellation string // 星座
	House         string
	Car           string
}

// FromJSONObj set profile from json object
func FromJSONObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
