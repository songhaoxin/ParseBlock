package types

//type Hash [32]byte
type Hash = string

type BlockNode struct {
	BlockNumber int
	BlockHash   Hash
}
type BlockNodeFecthed struct {
	BlockNumber   int
	BlockHash     Hash
	HighestNumber int
}

type AffiremTransInterface interface {
	AffiremTrans(blockNumber int, blockHash Hash)
}

type ReSendTransInterface interface {
	ResendTrans(pool *PendingPool)
}

type FecthBlocksInterface interface {
	FetchBlocks(localLastBlockNumber int, blocks chan *BlockNode) (isFethed bool, chainLatestBlockNumber int)
}
