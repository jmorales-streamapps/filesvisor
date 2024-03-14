package models

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Jon-MC-dev/files_copy/bd"
)

type Preferences struct {
	SizeChunk int32 `json:"sizeChunk,omitempty"`
	ModeLine  bool  `json:"modeLine,omitempty"`
}

func DefaultPreferences() *Preferences {
	return &Preferences{
		SizeChunk: 100000,
		ModeLine:  false,
	}
}

func LoadPreferences(bdPreferences *sql.DB) *Preferences {

	pref := DefaultPreferences()

	prefSizeChunk, _ := bd.GetPreference(bdPreferences, "PrefSizeChunk")
	if prefSizeChunk != "" {
		fmt.Println("prefSizeChunk :: ", prefSizeChunk)
		num, err := strconv.Atoi(prefSizeChunk)
		if err == nil {
			pref.SizeChunk = int32(num)
		}

	}

	PrefModeLine, _ := bd.GetPreference(bdPreferences, "PrefModeLine")
	if PrefModeLine != "" {
		fmt.Println("prefSizeChunk :: ", PrefModeLine)
		prefBool, err := strconv.ParseBool(PrefModeLine)
		if err == nil {
			pref.ModeLine = prefBool
		}

	}

	return pref
}
