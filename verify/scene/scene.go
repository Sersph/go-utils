package verify_scene

import (
	"strings"
)

type Scene interface {
	Match(string) bool
}

func match(name, scene string) bool {
	return strings.EqualFold(strings.ToUpper(strings.Trim(name, " ")), strings.ToUpper(scene))
}

type DEFAULTScene struct{}

func (DEFAULTScene) Match(name string) bool {
	return match(name, "")
}

type ADDScene struct{}

func (ADDScene) Match(name string) bool {
	return match(name, "ADD")
}

type UPDATEScene struct{}

func (UPDATEScene) Match(name string) bool {
	return match(name, "UPDATE")

}

type DELETEScene struct{}

func (DELETEScene) Match(name string) bool {
	return match(name, "DELETE")
}

var (
	DEFAULT = DEFAULTScene{}
	ADD     = ADDScene{}
	UPDATE  = UPDATEScene{}
	DELETE  = DELETEScene{}
)
