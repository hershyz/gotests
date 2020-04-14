	currentfile, err := os.Open(filepath)

	if err != nil {
		log.Fatal("failed to open file: " + filepath)
	}

	fileScanner := bufio.NewScanner(currentfile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	currentfile.Close()
