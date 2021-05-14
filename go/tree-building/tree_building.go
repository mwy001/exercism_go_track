package tree

import (
	"sort"
	"errors"
)

type Record struct {
	ID int
	Parent int
}

type Node struct {
	ID int
	Children []*Node
}

func insertNode(n *Node, r Record) {
	if n.ID == r.Parent {
		if n.Children == nil {
			n.Children = []*Node{&Node{ID:r.ID}}
		} else {
			n.Children = append(n.Children, &Node{ID:r.ID})
		}
	} else {
		for _, n := range n.Children {
			insertNode(n, r)
		}
	}
}

func Build(input []Record) (*Node, error) {

	if len(input) == 0 {
		return nil, nil
	}

	sort.Slice(input, func(i, j int) bool { return input[i].ID < input[j].ID })

	first := input[0]

	if first.ID != 0 || first.Parent != 0 {
		return nil, errors.New("no head id 0 or head id has parent id")
	}

	for i := 1; i < len(input); i++ {
		if input[i].Parent >= input[i].ID {
			return nil, errors.New("parent id larger than or equal self id")
		}

		if input[i].ID == input[i-1].ID {
			return nil, errors.New("duplicate id")
		}

		if input[i].ID != input[i-1].ID + 1 {
			return nil, errors.New("non continuous id")
		}
	}

	var root *Node

	for _, r := range input {
		if r.ID == r.Parent {
			root = &Node{ ID:r.ID }
		} else {
			insertNode(root, r)
		}
	}

	return root, nil
}
