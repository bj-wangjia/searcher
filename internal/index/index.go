package index

type Index struct {
	InvertIndex  InvertIndex
	StorageIndex StorageIndex
}

func NewIndex() *Index {
	return &Index{}
}
