package index

import "github.com/bj-wangjia/searcher/internal/query"

type StorageIndex interface {
	Get(filedName string, id query.DocId) interface{}
}
