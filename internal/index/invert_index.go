package index

import (
	"github.com/bj-wangjia/searcher/internal/query"
)

type InvertList interface {
	HasNext()
	Next() query.DocId
	GetGE(id query.DocId) query.DocId
}

type InvertIndex interface {
	GetInvertList(fieldName string) InvertList
}
