package main

import (
	"fmt"
)

type MCTSNode[STATE any, ACTION comparable] struct {
	action  *ACTION //action to get to that state
	state   *STATE
	nbPlays int
	nbWins  int
	nbDraws int
	parent  *MCTSNode[STATE, ACTION]
	// possible actions from that state
	// if node is nil it is still not expanded
	children map[*ACTION]*MCTSNode[STATE, ACTION]
}

func (node *MCTSNode[STATE, ACTION]) String() string {
	if node == nil {
		return ""
	}
	return fmt.Sprintf(
		"{action: %v, state: %v, nbPlays: %d, nbWins: %d, nbDraws: %d, parentState: %v, possibleActions: %v}",
		node.action,
		node.state,
		node.nbPlays,
		node.nbWins,
		node.nbDraws,
		node.GetParentState(),
		node.GetPossibleActions(),
	)
}

func (node *MCTSNode[STATE, ACTION]) AddChild(child *MCTSNode[STATE, ACTION]) {
	if node == nil || child == nil {
		return
	}
	node.children[child.action] = child
}

func (node *MCTSNode[STATE, ACTION]) GetChild(action *ACTION) *MCTSNode[STATE, ACTION] {
	if node == nil || action == nil {
		return nil
	}
	for key, child := range node.children {
		if key == nil {
			continue
		}
		if key == action || *key == *action {
			return child
		}
	}
	return nil
}

func (node *MCTSNode[STATE, ACTION]) GetParentState() *STATE {
	if node == nil || node.parent == nil {
		return nil
	}
	return node.parent.state
}

func (node *MCTSNode[STATE, ACTION]) GetPossibleActions() []*ACTION {
	actions := make([]*ACTION, 0)
	for action := range node.children {
		actions = append(actions, action)
	}
	return actions
}

func (node *MCTSNode[STATE, ACTION]) GetUnexploredActions() []*ACTION {
	actions := make([]*ACTION, 0)
	for action, child := range node.children {
		if child == nil {
			actions = append(actions, action)
		}
	}
	return actions
}

func (node *MCTSNode[STATE, ACTION]) IsLeaf() bool {
	return len(node.children) == 0
}

func NewMCTSNode[STATE any, ACTION comparable](
	action *ACTION,
	state *STATE,
	parent *MCTSNode[STATE, ACTION],
	possibleActions []*ACTION,
) *MCTSNode[STATE, ACTION] {
	children := make(map[*ACTION]*MCTSNode[STATE, ACTION])
	for _, action := range possibleActions {
		children[action] = nil
	}
	return &MCTSNode[STATE, ACTION]{
		action:   action,
		state:    state,
		nbPlays:  0,
		nbWins:   0,
		nbDraws:  0,
		parent:   parent,
		children: children,
	}
}
