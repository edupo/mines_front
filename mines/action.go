package mines

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Action int

type Command struct {
	Id     int
	Action Action
}

const (
	ActionNone = iota
	ActionFlag
	ActionUncover
)

var (
	ActionName = map[int]string{
		0: "nothing",
		1: "flag",
		2: "uncover",
	}
	ActionValue = map[string]int{
		"nothing": 0,
		"flag":    1,
		"uncover": 2,
	}
)

func (a Action) String() string {
	return ActionName[int(a)]
}

func ParseAction(s string) (Action, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := ActionValue[s]
	if !ok {
		return Action(0), fmt.Errorf("%q is not a valid action", s)
	}
	return Action(value), nil
}

func (a Action) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a *Action) UnmarshalJSON(data []byte) (err error) {
	var action string
	if err := json.Unmarshal(data, &action); err != nil {
		return err
	}
	if *a, err = ParseAction(action); err != nil {
		return err
	}
	return nil
}
