// enumerated types (enums) are a special case of sum types
// an enum is a type that has a fixed number of possible values, each with a distinct name
// go doesn't have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms

package main

import "fmt"

// our enum type ServerState has an underlying int type
type ServerState int

// the possible values for ServerState are defined as constants
// the special keyword iota generates successive constant values automatically; in this case 0, 1, 2 and so on
const (
	StateIdle ServerState = iota // 0
	StateConnected
	StateError
	StateRetrying
)

// by implementing the fmt.Stringer interface, values of ServerState can be printed out or converted to strings
var stateName = map[ServerState]string{
	StateIdle: "idle",
	StateConnected: "connected",
	// this can get cumbersome if there are many possible values
	// in such cases the stringer tool can be used in conjunction with go:generate to automate the process
	StateError: "error",
	StateRetrying: "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

// if we have a value of type int, we cannot pass it to transition - the compiler will complain about type mismatch
// this provides some degree of compile-time type safety for enums
func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition emulates a state transition for a server; it takes the existing state and returns a new state
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// suppose we check some predicates here to determine the next state...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
