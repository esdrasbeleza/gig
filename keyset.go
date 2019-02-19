package main

type KeySet map[Key]interface{}

func NewKeySet() KeySet {
	return KeySet{}
}

func (ks KeySet) Add(keys ...Key) {
	for _, key := range keys {
		ks[key] = nil
	}
}

func (ks KeySet) Keys() []Key {
	keys := []Key{}

	for k := range ks {
		keys = append(keys, k)
	}

	return keys
}

func (ks KeySet) Strings() []string {
	strgs := []string{}

	for s := range ks {
		strgs = append(strgs, s.String())
	}

	return strgs
}