file, err := os.Open("names.txt")
	if err != nil {
		log.Panicln(err)

	}

	defer file.Close()

	readFileByte, err := io.ReadAll(file)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(readFileByte))