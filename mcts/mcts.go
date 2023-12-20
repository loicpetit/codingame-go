package main

// Monte Carlo Tree Search algorithm

import (
	"math"
	"math/rand"
	"time"

	"github.com/loicpetit/codingame-go/runner"
)

type MCTS[STATE Hashable, ACTION comparable] struct {
	exploreParam      float64
	game              runner.Game[STATE, ACTION]
	simulationTimeout time.Duration // 0 = no timeout
	simulationDepth   int           // max nb turn by simulation, 0 = infinite
	tree              map[string]*MCTSNode[STATE, ACTION]
}

func (mcts *MCTS[STATE, ACTION]) makeNode(state *STATE) *MCTSNode[STATE, ACTION] {
	if mcts == nil || state == nil {
		return nil
	}
	hash := (*state).Hash()
	if mcts.tree[hash] != nil {
		return mcts.tree[hash]
	}
	mcts.tree[hash] = NewMCTSNode(nil, state, nil, mcts.game.GetAvailableActions(state, 1))
	return mcts.tree[hash]
}

func (mcts *MCTS[STATE, ACTION]) computeUCB1(node *MCTSNode[STATE, ACTION]) float64 {
	// UCB1 selection factor
	// (Wi / Si) + c * sqrt(ln(Sp) / Si)
	// Wi = nb of win simulations of node
	// to give value to draw games, add half of the draws in Wi value
	// Si = total nb of simulations of node
	// c = exporer parameter (usually sqrt(2))
	// Sp = total nb of simulations of node parent
	// first part is the exploitation part
	// second is the exploration part
	// ln(100) = 2
	Wi := float64(node.nbWins) + float64(node.nbDraws/2)
	Si := float64(node.nbPlays)
	Sp := float64(node.parent.nbPlays)
	return (Wi / Si) + (mcts.exploreParam * math.Sqrt(math.Log(Sp)/Si))
}

func (mcts *MCTS[STATE, ACTION]) selectNode(root *MCTSNode[STATE, ACTION]) *MCTSNode[STATE, ACTION] {
	node := root
	for mcts.isNodeFullyExpanded(node) && !mcts.isNodeLeaf(node) {
		var bestNode *MCTSNode[STATE, ACTION]
		var bestUCB1 float64
		for _, child := range node.children {
			if bestNode == nil {
				bestNode = child
				bestUCB1 = mcts.computeUCB1(bestNode)
				continue
			}
			childUCB1 := mcts.computeUCB1(child)
			if childUCB1 > bestUCB1 {
				bestNode = child
				bestUCB1 = childUCB1
			}
		}
		node = bestNode
	}
	return node
}

func (mcts *MCTS[STATE, ACTION]) expandNode(node *MCTSNode[STATE, ACTION]) *MCTSNode[STATE, ACTION] {
	unexploredActions := node.GetUnexploredActions()
	randomIndex := rand.Intn(len(unexploredActions))
	action := unexploredActions[randomIndex]
	childState := mcts.game.Play(node.state, action)
	childPossibleAction := mcts.game.GetAvailableActions(childState, mcts.game.GetNextPlayer(childState))
	childNode := NewMCTSNode(action, childState, node, childPossibleAction)
	node.AddChild(childNode)
	return childNode
}

func (mcts *MCTS[STATE, ACTION]) simulate(node *MCTSNode[STATE, ACTION]) int {
	timeout := false
	go func() {
		if mcts.simulationTimeout > 0 {
			<-time.After(mcts.simulationTimeout)
			timeout = true
		}
	}()
	state := node.state
	winner := mcts.game.Winner(state)
	player := mcts.game.GetNextPlayer(state)
	possibleActions := mcts.game.GetAvailableActions(state, player)
	count := 0
	for winner == 0 && len(possibleActions) > 0 {
		if timeout {
			panic("Simulation timeout")
		}
		if mcts.simulationDepth > 0 && count == mcts.simulationDepth {
			break
		}
		randomIndex := rand.Intn(len(possibleActions))
		action := possibleActions[randomIndex]
		state = mcts.game.Play(state, action)
		winner = mcts.game.Winner(state)
		player = mcts.game.GetNextPlayer(state)
		possibleActions = mcts.game.GetAvailableActions(state, player)
		count++
	}
	return winner
}

func (mcts *MCTS[STATE, ACTION]) backPropagateResult(node *MCTSNode[STATE, ACTION], winner int) {
	for node != nil {
		node.nbPlays += 1
		if winner == 0 {
			node.nbDraws += 1
		} else if mcts.game.GetLastPlayer(node.state) == winner {
			node.nbWins += 1
		}
		node = node.parent
	}
}

func (mcts *MCTS[STATE, ACTION]) isNodeFullyExpanded(node *MCTSNode[STATE, ACTION]) bool {
	return len(node.GetUnexploredActions()) == 0
}

func (mcts *MCTS[STATE, ACTION]) isNodeLeaf(node *MCTSNode[STATE, ACTION]) bool {
	return len(node.GetPossibleActions()) == 0
}

func (mcts *MCTS[STATE, ACTION]) Search(state *STATE, maxTime time.Time) {
	if mcts == nil {
		return
	}
	root := mcts.makeNode(state)
	for maxTime.After(time.Now()) {
		mcts.searchOnce(root)
	}
}

func (mcts *MCTS[STATE, ACTION]) SearchN(state *STATE, n int) {
	if mcts == nil {
		return
	}
	root := mcts.makeNode(state)
	i := 0
	for i < n {
		i++
		mcts.searchOnce(root)
	}
}

func (mcts *MCTS[STATE, ACTION]) searchOnce(root *MCTSNode[STATE, ACTION]) {
	if mcts == nil || root == nil {
		return
	}
	node := mcts.selectNode(root)
	winner := mcts.game.Winner(node.state)
	if !node.IsLeaf() && winner == 0 {
		node = mcts.expandNode(node)
		winner = mcts.simulate(node)
	}
	mcts.backPropagateResult(node, winner)
}

func (mcts *MCTS[STATE, ACTION]) GetBestAction(state *STATE) *ACTION {
	if mcts == nil || state == nil {
		return nil
	}
	node := mcts.makeNode(state)
	var action *ACTION
	nbPlays := -1
	for _, child := range node.children {
		if child != nil && child.nbPlays > nbPlays {
			action = child.action
			nbPlays = child.nbPlays
		}
	}
	return action
}

func NewMCTS[STATE Hashable, ACTION comparable](game runner.Game[STATE, ACTION]) *MCTS[STATE, ACTION] {
	return &MCTS[STATE, ACTION]{
		exploreParam:      2,
		game:              game,
		simulationTimeout: 5 * time.Millisecond,
		simulationDepth:   20,
		tree:              map[string]*MCTSNode[STATE, ACTION]{},
	}
}
