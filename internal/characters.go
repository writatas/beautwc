package internal

type DisplayCharacters struct {
	matches map[string]Match
}

type Match struct {
	filename  string
	match     string
	instances int
}
