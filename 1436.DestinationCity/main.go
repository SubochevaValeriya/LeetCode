package main

func main() {

}

func destCity(paths [][]string) string {
	possibleDestinations := map[string]struct{}{}
	notDestination := map[string]struct{}{}

	for _, path := range paths {
		possibleDestinations[path[1]] = struct{}{}
		notDestination[path[0]] = struct{}{}
	}

	for city := range possibleDestinations {
		if _, ok := notDestination[city]; !ok {
			return city
		}
	}

	return ""
}
