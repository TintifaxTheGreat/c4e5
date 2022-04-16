package cmd

func hasNext(a []string) bool {
	return len(a) > 0
}

func getNext(a []string) (string, []string) {
	n, a := a[0], a[1:]
	return n, a
}
