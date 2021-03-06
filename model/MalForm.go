package model

import (
	"fmt"
	"strconv"
	"strings"
)

type MalForm struct {
	Type  MalType
	Value interface{}
	Err   error
}

func (mf MalForm) Sprint() string {
	var sb strings.Builder

	switch mf.Type {
	case MalTypeNil:
		sb.WriteString("<nil>")
	case MalTypeFunc, MalTypeClosure:
		sb.WriteString("#<function>")
	case MalTypeBool:
		sb.WriteString(strconv.FormatBool(mf.ValBool()))
	case MalTypeSymbol:
		sb.WriteString(mf.ValString())
	case MalTypeNumber:
		sb.WriteString(strconv.Itoa(mf.ValInt()))
	case MalTypeList:
		sb.WriteString("(")
		vals := mf.ValList()
		if len(vals) == 0 {
			sb.WriteString(")")
			return sb.String()
		}
		for i, val := range vals {
			sb.WriteString(val.Sprint())

			if i == len(vals)-1 {
				sb.WriteString(")")
			} else {
				sb.WriteString(" ")
			}
		}
	default:
		panic(fmt.Sprintf("Invalid MalType, mf=%v\n", mf))
	}
	return sb.String()
}

func (mf MalForm) Print() {
	fmt.Print(mf.Sprint())
}

func (mf MalForm) Error() string {
	if mf.Err == nil {
		return ""
	}
	return mf.Err.Error()
}

func (mf MalForm) ValBool() bool {
	return mf.Value.(bool)
}

func (mf MalForm) ValString() string {
	return mf.Value.(string)
}

func (mf MalForm) ValInt() int {
	return mf.Value.(int)
}

func (mf MalForm) ValList() []MalForm {
	return mf.Value.([]MalForm)
}

func (mf MalForm) ValMalFunc() MalFunc {
	return mf.Value.(MalFunc)
}

func (mf MalForm) ValMalClosure() MalClosure {
	return mf.Value.(MalClosure)
}

func (mf MalForm) IsSpecialForm() bool {
	if mf.Type != MalTypeSymbol {
		return false
	}
	return IsSpecialForm(mf.ValString())
}
