package parser

type PIFElem struct {
	token string
	index STIndex
}

type ProgramInternalForm interface {
	Add(t string, i STIndex)
}
