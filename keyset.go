package main

import "sort"

type KeySet map[Key]interface{}

func NewKeySet() KeySet {
	return KeySet{}
}

func (ks KeySet) Add(keys ...Key) {
	for _, key := range keys {
		ks[key] = nil
	}
}

func (ks1 KeySet) Merge(ks2 KeySet) {
	for k, _ := range ks2 {
		ks1[k] = nil
	}
}

func (ks KeySet) Keys() []Key {
	var (
		input  = ks.Strings()
		output = make([]Key, len(input))
	)

	for pos, value := range input {
		output[pos] = Key(value)
	}

	return output
}

func (ks KeySet) Strings() []string {
	strgs := sort.StringSlice{}

	for s := range ks {
		strgs = append(strgs, s.String())
	}

	strgs.Sort()

	return strgs
}

func (ks KeySet) IsEmpty() bool {
	return len(ks) == 0
}
