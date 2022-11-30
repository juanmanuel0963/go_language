package maps

func SearchMap(search_map map[int]string, search_key int) string {

	if val, isExists := search_map[search_key]; isExists {

		//do steps needed here
		return val
	}

	return ""
}
