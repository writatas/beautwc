package pkg

func PrintAll(files ...string) error {
	wordsCount := WordsCount{}
	linesCount := LinesCount{}
	charactersCount := CharactersCount{}
	memoryBytes, err := GetFilesSize(files...)
	if err != nil {
		return err
	}
	err = memoryBytes.TotalBytes.CalculateBytesAndPrefixes()
	if err != nil {
		return err
	}
	err = wordsCount.GetWordsCount(files...)
	if err != nil {
		return err
	}
	err = linesCount.GetLinesCount(files...)
	if err != nil {
		return err
	}
	err = charactersCount.GetCharactersCount(files...)
	if err != nil {
		return err
	}
	PrintDefaultValue(&memoryBytes.TotalBytes)
	PrintDefaultValue(&wordsCount)
	PrintDefaultValue(&linesCount)
	PrintDefaultValue(&charactersCount)
	return nil
}

// interface to always print default string values if no flags are provided
type Default interface {
	PrintDefault()
}

func PrintDefaultValue(t Default) {
	t.PrintDefault()
}
