package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(rdr io.Reader) (league League, err error) {
	err = json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("解析联盟json出错:%v", err)
	}
	return
}
