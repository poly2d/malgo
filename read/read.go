package read

import (
	"strconv"

	"github.com/poly2d/malgo/model"
)

func readAtom(r reader) model.MalForm {
	token := r.next()

	if num, err := strconv.Atoi(token); err == nil {
		return model.MalForm{
			Type:  model.MalTypeNumber,
			Value: num,
		}
	}

	switch token {
	case "nil":
		return model.MalForm{
			Type: model.MalTypeNil,
		}
	case "true", "false":
		if b, err := strconv.ParseBool(token); err == nil {
			return model.MalForm{
				Type:  model.MalTypeBool,
				Value: b,
			}
		}
	}

	return model.MalForm{
		Type:  model.MalTypeSymbol,
		Value: token,
	}
}

func readList(r reader) model.MalForm {
	list := []model.MalForm{}
	for {
		mf := readForm(r)

		switch mf.Type {
		case model.MalTypeSymbol:
			str := mf.ValString()
			if str == ")" {
				return model.MalForm{
					Type:  model.MalTypeList,
					Value: list,
				}
			}
		}
		list = append(list, mf)
	}
	return model.MalForm{}
}

func readForm(r reader) model.MalForm {
	top := r.peek()
	switch top {
	case "(":
		r.next() // Discard open paren.
		return readList(r)
	}
	return readAtom(r)
}

func ReadStr(in string) model.MalForm {
	reader := newReader(in)
	return readForm(reader)
}
