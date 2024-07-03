package statemachine

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type State struct {
	Name        string
	Transitions []Transition
	IsInitial   bool
	IsFinal     bool
}

type Transition struct {
	Event       string
	TargetState string
}

func ReadMermaidDiagram() (string, error) {
	var lines []string
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}

	return strings.Join(lines, "\n"), nil
}

func ParseMermaidDiagram(diagram string) ([]State, error) {
	var states []State
	lines := strings.Split(diagram, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "stateDiagram-v2" || line == "" {
			continue
		}
		if strings.Contains(line, "-->") {
			parts := strings.Split(line, "-->")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid transition format: %s", line)
			}
			fromState := strings.TrimSpace(parts[0])
			toStatePart := strings.TrimSpace(parts[1])

			var event, toState string
			if strings.Contains(toStatePart, ":") {
				eventAndToState := strings.SplitN(toStatePart, ":", 2)
				event = strings.TrimSpace(eventAndToState[0])
				toState = strings.TrimSpace(eventAndToState[1])
			} else {
				toState = toStatePart
			}

			fromInitial := fromState == "[*]"
			toFinal := toState == "[*]"

			if fromInitial {
				fromState = "Initial"
			}
			if toFinal {
				toState = "Final"
			}

			var state *State
			for i := range states {
				if states[i].Name == fromState {
					state = &states[i]
					break
				}
			}
			if state == nil {
				states = append(states, State{Name: fromState, IsInitial: fromInitial})
				state = &states[len(states)-1]
			}

			if !toFinal {
				state.Transitions = append(state.Transitions, Transition{
					Event:       event,
					TargetState: toState,
				})
			}

			var exists bool
			for i := range states {
				if states[i].Name == toState {
					exists = true
					states[i].IsFinal = toFinal
					break
				}
			}
			if !exists {
				states = append(states, State{Name: toState, IsFinal: toFinal})
			}
		}
	}

	if len(states) == 0 {
		return nil, fmt.Errorf("no valid states found in the diagram")
	}

	return states, nil
}

func GenerateBoostCode(states []State) (string, error) {
	if len(states) == 0 {
		return "", fmt.Errorf("no states to generate code for")
	}

	var code strings.Builder

	code.WriteString("#include <boost/statechart/state_machine.hpp>\n")
	code.WriteString("#include <boost/statechart/simple_state.hpp>\n")
	code.WriteString("#include <boost/statechart/transition.hpp>\n")
	code.WriteString("#include <boost/statechart/termination.hpp>\n\n")

	code.WriteString("// Forward declarations\n")
	for _, state := range states {
		code.WriteString(fmt.Sprintf("struct %s;\n", state.Name))
	}
	code.WriteString("\n")

	var initialState string
	for _, state := range states {
		if state.IsInitial {
			initialState = state.Name
			break
		}
	}
	if initialState == "" {
		initialState = states[0].Name
	}

	code.WriteString("// State machine\n")
	code.WriteString(fmt.Sprintf("struct StateMachine : boost::statechart::state_machine<StateMachine, %s> {};\n\n", initialState))

	code.WriteString("// States and transitions\n")
	for _, state := range states {
		code.WriteString(fmt.Sprintf("struct %s : boost::statechart::simple_state<%s, StateMachine> {\n", state.Name, state.Name))
		if len(state.Transitions) > 0 || state.IsFinal {
			code.WriteString("    typedef boost::mpl::list<\n")
			for i, transition := range state.Transitions {
				code.WriteString(fmt.Sprintf("        boost::statechart::transition<%s, %s>", transition.Event, transition.TargetState))
				if i < len(state.Transitions)-1 || state.IsFinal {
					code.WriteString(",")
				}
				code.WriteString("\n")
			}
			if state.IsFinal {
				code.WriteString("        boost::statechart::termination<>\n")
			}
			code.WriteString("    > reactions;\n")
		}
		code.WriteString("};\n\n")
	}

	code.WriteString("// Events\n")
	events := make(map[string]bool)
	for _, state := range states {
		for _, transition := range state.Transitions {
			if transition.Event != "" && !events[transition.Event] {
				events[transition.Event] = true
				code.WriteString(fmt.Sprintf("struct %s : boost::statechart::event<%s> {};\n", transition.Event, transition.Event))
			}
		}
	}

	return code.String(), nil
}

func Use() {
	mermaidDiagram, err := ReadMermaidDiagram()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading Mermaid diagram: %v\n", err)
		os.Exit(1)
	}

	if mermaidDiagram == "" {
		fmt.Fprintf(os.Stderr, "No input provided. Please provide a Mermaid diagram.\n")
		os.Exit(1)
	}

	states, err := ParseMermaidDiagram(mermaidDiagram)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing Mermaid diagram: %v\n", err)
		os.Exit(1)
	}

	boostCode, err := GenerateBoostCode(states)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating Boost code: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(boostCode)
}
