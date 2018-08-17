package types

import (
	"sync"
	"fmt"
)

type BlockPool struct {
	MaxSizeAllowed    int //允许的区块最大数量
	AffiremBlockHeigh int //用以确认区块，离最新区块的标准高度

	startIdx int          //区块池中的起始区块号
	endIdx   int          //区块池中的最新区块号
	size     int          //区块池中的区块数量
	pool     map[int]Hash //管理最近区块的池子
	lock     *sync.RWMutex
}

/// 创建 "BlockPool"实例
func InitBlockPool() *BlockPool {
	p := &BlockPool{
		startIdx:          -1,
		endIdx:            -1,
		size:              0,
		MaxSizeAllowed:    100,
		AffiremBlockHeigh: 6,
		pool:              make(map[int]Hash),
		lock:              new(sync.RWMutex),
	}
	return p

}

func (b *BlockPool) IsEmpty() bool {
	return b.size == 0
}

func (b *BlockPool) LatestNumber() int {
	return b.endIdx
}

/*
func (b *BlockPool) InsertNodes(nodes []BlockNode) []BlockNode {
	if nil == nodes || 0 == len(nodes) {
		return nil
	}
	var _nodes []BlockNode = make([]BlockNode, len(nodes))
	for _, n := range nodes {
		_, bn := b.Insert(n.BlockNumber, n.BlockHash)
		if nil != bn {
			_nodes = append(_nodes, (*bn))
		}
	}
	return _nodes
}
*/
/*
func (b *BlockPool) ReceiveBlocks(bs <-chan *BlockNode, latestNumber int) {
	for n := range bs {
		rn := b.Insert2Main(n, latestNumber)
		if nil != rn {
			b.insert2ReTrans(rn)
		}
	}
}
*/

func (b *BlockPool) ReceiveBlocks(bs chan *BlockNode) {
	for n := range bs {
		fmt.Println(n.BlockNumber)
		fmt.Println(n.BlockHash)
		//b.insert(n)
	}
}

func (b *BlockPool) insert(node *BlockNode) *BlockNode {
	if nil == node {
		return nil
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	if b.size+1 == b.MaxSizeAllowed {
		b.removeElementAtStart()
	}

	k, v := (*node).BlockNumber, (*node).BlockHash
	if k < 0 || 0 == len(v) {
		return nil
	}

	// 更新startIdx 与 endIdx
	if k < b.startIdx {
		b.startIdx = k
	} else if k > b.endIdx {
		b.endIdx = k
	}

	var n *BlockNode = nil
	if oldHash, ok := b.pool[k]; ok { //存在旧值
		if oldHash != v {
			n = &BlockNode{k, oldHash}
		}
	} else {
		b.size++
	}

	// 更新或增加元素
	b.pool[k] = v
	return n
}

/*
func (b *BlockPool) insert2ReTrans(node *BlockNode) {
	if nil == node || nil == b.needReTransactioPool {
		return
	}
	b.lock4ReTransaction.Lock()
	defer b.lock4ReTransaction.Unlock()

	k, v := node.BlockNumber, node.BlockHash
	b.needReTransactioPool[k] = v
}
*/

/// 对区块进行校验，以处理已经被确认的交易
func (b *BlockPool) LookBocks4AffirmTrans(latestNumber int) (bool, *BlockNode) {
	if latestNumber < 0 {
		return false, nil
	}

	b.lock.RLock()
	defer b.lock.RUnlock()

	for k, v := range b.pool {
		if latestNumber-k+1 >= b.AffiremBlockHeigh {
			return true, &BlockNode{k, v}
		}
	}
	return false, nil
}

func (b *BlockPool) removeElementAtStart() {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.size == 0 {
		return
	}
	if b.startIdx < 0 {
		return
	}

	k := b.startIdx
	if _, ok := b.pool[k]; ok {
		delete(b.pool, k)
	}

	if 0 == b.size {
		b.startIdx = -1
		b.endIdx = -1
	} else {
		b.startIdx++
	}
	b.size--
}
