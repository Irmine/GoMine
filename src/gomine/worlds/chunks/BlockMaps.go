package chunks

var LightFilterMap = map[int]byte{
	0: 0,
	1: 15,
	2: 15,
	3: 15,
	7: 15,
}

var LightDiffusionMap = map[int]bool{
	0: false,
	1: true,
	2: true,
	3: true,
	7: true,
}

func GetLightFilter(id int) byte {
	var filter, ok = LightFilterMap[id]
	if !ok {
		return 0
	}
	return filter
}

func DiffusesLight(id int) bool {
	var filter, ok = LightDiffusionMap[id]
	if !ok {
		return true
	}
	return filter
}