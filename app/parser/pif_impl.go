package parser

type pifImpl struct {
	elems []PIFElem
}

func NewPIF() ProgramInternalForm {
	return &pifImpl{elems: make([]PIFElem, 0)}
}

func (pif *pifImpl) Add(token string, index STIndex) {
	pif.elems = append(pif.elems, PIFElem{token: token, index: index})
}