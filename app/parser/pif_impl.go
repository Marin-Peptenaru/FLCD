package parser

import (
	"fmt"
	"strings"
)

type pifImpl struct {
	elems []PIFElem
}

func NewPIF() ProgramInternalForm {
	return &pifImpl{elems: make([]PIFElem, 0)}
}

func (pif *pifImpl) Add(token string, index STIndex) {
	pif.elems = append(pif.elems, PIFElem{token: token, index: index})
}

func (pif *pifImpl) String() string {
	builder := strings.Builder{}

	for _, elem := range pif.elems {
		builder.WriteString(fmt.Sprintln(elem))
	}

	return builder.String()
}