package ac

type Unit rune    // rune or int32
type Entry []Unit // single entry

type Node struct {
	res      []int          // matching entry index in root.entries,
	children map[Unit]*Node // tail has no children
	fail     *Node          // root has no fail
}

// IsEntry is the end node of some entries
func (n *Node) IsEntry() bool {
	return len(n.res) > 0
}

type AC struct {
	texts   []string
	root    *Node
	nodeNum int
}

func NewAC(texts []string) (ac *AC) {
	ac = &AC{}
	ac.texts = texts
	ac.root = ac.NewNode()
	ac.Build()
	return ac
}

func (ac *AC) NewNode() (node *Node) {
	ac.nodeNum++
	return &Node{}
}

func (ac *AC) IsRoot(node *Node) bool {
	return ac.root == node
}

func (ac *AC) Build() {
	// first step
	ac.BuildTrieTree()
	// second step
	ac.BuildFail()
}

func (ac *AC) BuildTrieTree() {
	entries := AnalyseTextList(ac.texts)
	for i, entry := range entries {
		// every entry begin with root
		node := ac.root
		for _, unit := range entry {
			if len(node.children) == 0 {
				node.children = make(map[Unit]*Node)
			}
			n, ok := node.children[unit]
			if !ok {
				n = ac.NewNode()
				node.children[unit] = n
			}
			node = n
		}
		// node is the tail of entry here, add res
		node.res = append(node.res, i)
	}
}

func (ac *AC) BuildFail() {
	var nch = make(chan *Node, ac.nodeNum)

	// all root.child.fail = root
	for _, child := range ac.root.children {
		child.fail = ac.root
		nch <- child
	}

	// breadth first traversal
	for len(nch) > 0 {
		node := <-nch
		// current node.child match all node.fail.children and fail.fail.children ...
		for unit, child := range node.children {
			var fail *Node
			for fail = node.fail; fail != nil; fail = fail.fail {
				if match, ok := fail.children[unit]; ok {
					child.fail = match
					break
				}
			}
			// if fail is nil, node.child.fail match root
			if fail == nil {
				child.fail = ac.root
			}
			// push current node.child into chan
			nch <- child
		}
	}
}

// node.fail must not nil
func (ac *AC) matchFail(node *Node) {
	for unit, child := range node.children {
		// current node.child match all node.fail.children and fail.fail.children ...
		var fail *Node
		for fail = node.fail; fail != nil; fail = fail.fail {
			if match, ok := fail.children[unit]; ok {
				child.fail = match
				break
			}
		}
		// if fail is nil, node.child.fail match root
		if fail == nil {
			child.fail = ac.root
		}
	}
}

// Match the specified text, return all matching entries
func (ac *AC) Match(text string) (matches []string) {
	res := ac.MatchIndex(text)
	matches = make([]string, len(res))
	for i, idx := range res {
		matches[i] = ac.texts[idx]
	}
	return matches
}

// MatchIndex the specified entry, return all matching entries' index
func (ac *AC) MatchIndex(text string) (res []int) {
	entry := AnalyseText(text)
	node := ac.root
	for _, unit := range entry {
		next, ok := node.children[unit]
		// fail is nil when node is root
		for !ok && node.fail != nil {
			node = node.fail
			next, ok = node.children[unit]
		}
		if ok {
			node = next
			if node.IsEntry() {
				res = append(res, node.res...)
			}
		}
	}
	return res
}

func AnalyseText(text string) (entry Entry) {
	return Entry(text)
}

func AnalyseTextList(texts []string) (entries []Entry) {
	entries = make([]Entry, len(texts))
	for i, text := range texts {
		entries[i] = AnalyseText(text)
	}
	return entries
}
