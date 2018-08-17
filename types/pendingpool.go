package types

import "sync"

type PendingPool struct {
	startIdx int          //区块池中的起始区块号
	endIdx   int          //区块池中的最新区块号
	size     int          //区块池中的区块数量
	pool     map[int]Hash //管理需要重要打包交易相关区块的池子
	lock     *sync.RWMutex
}

func InitPendingPool() *PendingPool {
	p := &PendingPool{
		startIdx: -1,
		endIdx:   -1,
		size:     0,
		pool:     make(map[int]Hash),
		lock:     new(sync.RWMutex),
	}
	return p
}

func (p *PendingPool) IsEmpty() bool {
	return p.size == 0
}

func (p *PendingPool) Insert(node *BlockNode) {
	if nil == node {
		return
	}
	if nil == p.pool {
		return
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	k, v := node.BlockNumber, node.BlockHash
	if k < 0 || 0 == len(v) {
		return
	}

	p.pool[k] = v
}
