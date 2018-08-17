package types

/*
type CLWalletList struct {
	size              uint64
	latestBlockNumber int
	head              *CLWalletNode
	tail              *CLWalletNode
}

type CLWalletNode struct {
	blockNumber int
	blockHash   [32]byte
	next        *CLWalletNode
	prev        *CLWalletNode
}

/// 初始化链表
func (list *CLWalletList) Init() {
	(*list).size = 0
	(*list).head = nil
	(*list).tail = nil
}

/// 在尾部增加元素
func (list *CLWalletList) AppendAtTail(node *CLWalletNode) bool {
	if nil == node {
		return false
	}
	(*node).next = nil
	(*node).prev = nil

	if 0 == (*list).size {
		(*list).head = node
		(*list).tail = node
	} else {
		oldTail := (*list).tail
		(*node).prev = oldTail
		(*oldTail).next = node
		(*list).tail = node
	}
	(*list).size++
	return true
}

/// 在头部增加元素
func (list *CLWalletList) AppendAtHead(node *CLWalletNode) bool {
	if nil == node {
		return false
	}
	(*node).next = nil
	(*node).prev = nil

	if 0 == (*list).size {
		(*list).head = node
		(*list).tail = node
	} else {
		oldHead := (*list).head
		(*node).next = oldHead
		(*oldHead).prev = node
		(*list).head = node
	}
	(*list).size++
	(*list).latestBlockNumber = node.blockNumber
	return true
}

func (list *CLWalletList) RemoveTail() {
	if nil == (*list).tail || 0 == (*list).size {
		return
	}
	if (*list).head == (*list).tail {
		(*list).head = nil
		(*list).tail = nil
	} else {
		tail := (*list).tail
		(*list).tail = tail.prev
		(*list).tail.next = nil
	}

	(*list).size--
}

func (list *CLWalletList) Compare(latestList *CLWalletList) (bool, *CLWalletList) {
	if nil == latestList {
		return false, nil
	}

	oriHead := list.head       //原链表头
	newHead := latestList.head //新链表头

	//找到新结点与原结点BlockNumber相等的结点

	return false, nil
}
*/
