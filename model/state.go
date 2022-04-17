package model

type State struct {
	Id string
}

func Genesis() State {
	return State{Id: ""}
}

func NewState(id string) State {
	return State{Id: id}
}

func (s *State) IsGenesis() bool {
	return s.Id == ""
}
