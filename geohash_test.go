package gogeohash

import "testing"

const (
	lat  = -33.863574
	lng  = 150.915070
	hash = "r3gr65m42h22"
)

func Test_Encode_1(t *testing.T) {
	result := Encode(lat, lng)
	if result != hash {
		t.Error(result, hash)
	}
}
func Test_Decode_1(t *testing.T) {
	lats, lngs := Decode("r3gr65m42h22")
	if lat < lats[0] || lat > lats[1] {
		t.Error("lat out of range", lat, lats)
	}
	if lng < lngs[0] || lng > lngs[1] {
		t.Error("lng out of range", lng, lngs)
	}
}
