package main

func join(separator string, strList ...string) string {
	var r string = ""
	for _, str := range strList {
		if r == "" {
			r = str
		} else {
			r = r + separator + str
		}
	}
	return r
}
