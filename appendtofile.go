	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if _, err := file.WriteString(str + "\n"); err != nil {
		log.Fatal(err)
	}
	if err != nil {

	}

	file.Close()
