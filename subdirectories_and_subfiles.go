func test(path string) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	names, err := file.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
	}

	var i int = 0
	for i < len(names) {
		fmt.Println(names[i])
		i++
	}
}
