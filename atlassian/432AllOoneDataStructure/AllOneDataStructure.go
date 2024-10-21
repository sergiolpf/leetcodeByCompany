package alloonedatastructure

type AllOne struct {
	headDummy *Node
	tailDummy *Node
	stringMap map[string]*Node
}

type Node struct {
	keys       map[string]bool
	prev, next *Node
	frequency  int
}

func Constructor() AllOne {
	newNode := &Node{
		prev: nil,
		next: nil,
	}
	headNode := newNode
	tailNode := newNode
	headNode.next = tailNode
	tailNode.prev = headNode

	return AllOne{
		headDummy: headNode,
		tailDummy: tailNode,
		stringMap: make(map[string]*Node),
	}
}

func (this *AllOne) Inc(key string) {
	if currNode, ok := this.stringMap[key]; ok {
		if len(this.stringMap) == 1 {
			currNode.frequency++
			return
		}

		nextNode := currNode.next
		currentFrequency := currNode.frequency

		if nextNode.frequency == currentFrequency+1 {
			nextNode.keys[key] = true
			this.stringMap[key] = nextNode
			if len(currNode.keys) == 1 {
				currNode.DeleteThisNode()
			}
			delete(currNode.keys, key)
			return
		}

		if len(currNode.keys) == 1 {
			currNode.frequency++
			return
		}

		delete(currNode.keys, key)

		currNode.CreateNewNodeAfterCurrent(key, currentFrequency+1)
		this.stringMap[key] = currNode.next

	} else {
		firstElement := this.headDummy.next
		if firstElement.frequency == 1 {
			firstElement.keys[key] = true
			this.stringMap[key] = firstElement
			return
		}

		this.headDummy.CreateNewNodeAfterCurrent(key, 1)
		this.stringMap[key] = this.headDummy.next

	}
}

func (this *AllOne) Dec(key string) {
	currNode := this.stringMap[key]
	previousNode := currNode.prev
	nextNode := currNode.next
	currentFrequency := currNode.frequency

	if currentFrequency == 1 {
		delete(this.stringMap, key)
		if len(currNode.keys) == 1 {
			previousNode.next = nextNode
			nextNode.prev = previousNode
			currNode.next = nil
			currNode.prev = nil
		} else {
			delete(currNode.keys, key)
		}
		return
	}

	if previousNode.frequency == currentFrequency-1 {

		previousNode.keys[key] = true
		this.stringMap[key] = previousNode

		if len(currNode.keys) == 1 {
			previousNode.next = nextNode
			nextNode.prev = previousNode
			currNode.next = nil
			currNode.prev = nil
			return
		}

		delete(currNode.keys, key)
	} else {
		if len(currNode.keys) == 1 {
			currNode.frequency--
			return
		}

		currNode.CreateNewNodeBeforeCurrent(key, currentFrequency-1)
		delete(currNode.keys, key)
		this.stringMap[key] = currNode.prev
	}
}

func (this *AllOne) GetMaxKey() string {

	actualTail := this.tailDummy.prev
	for key := range actualTail.keys {
		return key
	}

	return ""
}

func (this *AllOne) GetMinKey() string {
	actualHead := this.headDummy.next

	for key := range actualHead.keys {
		return key
	}
	return ""
}

func (n *Node) CreateNewNodeAfterCurrent(str string, freq int) {
	nextNode := n.next
	newNode := Node{
		keys:      map[string]bool{str: true},
		frequency: freq,
		prev:      n,
		next:      nextNode,
	}

	nextNode.prev = &newNode
	n.next = &newNode

}

func (n *Node) CreateNewNodeBeforeCurrent(str string, freq int) {
	previousNode := n.prev
	newNode := Node{
		keys:      map[string]bool{str: true},
		frequency: freq,
		prev:      previousNode,
		next:      n,
	}

	previousNode.next = &newNode
	n.prev = &newNode

}

func (n *Node) DeleteThisNode() {
	previousNode := n.prev
	nextNode := n.next
	previousNode.next = nextNode
	nextNode.prev = previousNode
	n.next = nil
	n.prev = nil
}
