package product

import "sync"

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var _nextProductIdMtx sync.Mutex
var _nextProductId int = 0

func nextProductId() int {
	_nextProductIdMtx.Lock()
	defer _nextProductIdMtx.Unlock()

	_nextProductId = _nextProductId + 1
	return _nextProductId
}

var ProductList = []Product{
	{
		ID:   nextProductId(),
		Name: "Ayam Goreng",
	},
	{
		ID:   nextProductId(),
		Name: "Sambal Paru Balado",
	},
}
