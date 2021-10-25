package parser

// Implementation of the hash based symbol table.
// For internal use only, only the interface SymbolTable will be exposed.
// Lists are used to solve hash collisions
type stImplementation struct {

	// HashTable represented as an array of arrays of Symbols.
	// the outer array contains buckets of symbols, each bucket is associated with a hash code.
	hashTable [][]Symbol 
}

//Takes as parameters the symbol to hash and the size of the hash table 
//returns a int which is the hash code of the bucket in which the symbol will be stored.
//The sum of the bytes of a symbols is used as a toy hash function. 
func naiveHash(s Symbol, tableSize int) int {
	hash := 0
	symbolBytes := []byte(s)

	for _, byte := range symbolBytes {
		hash = (hash + int(byte)) % tableSize
	}

	return hash
}

// Implementation of the SaveSymbol method from the SymbolTable interface in accordance with its specification.
// Uses lists / buckets for solving hash collisions.
func (st *stImplementation) SaveSymbol(symbol Symbol) (index STIndex) {
	hash := naiveHash(symbol, len(st.hashTable))
	index = STIndex{tableIndex: hash, bucketIndex: 0}

	tableBucket := st.hashTable[hash]

	for i, s := range tableBucket {
		if symbol == s {
			index.bucketIndex = i
			return
		}
	}

	st.hashTable[hash] = append(tableBucket, symbol)
	index.bucketIndex = len(st.hashTable[hash]) - 1
	return
}

// Implementation of the GetSymbol method from the SymbolTable interface in accordance with its specification.
func (st *stImplementation) GetSymbol(index STIndex) (symbol Symbol, err error) {
	symbol, err = Symbol(""), nil

	if index.tableIndex >= len(st.hashTable) {
		err = invalidIndex(index)
		return
	}

	bucket := st.hashTable[index.tableIndex]

	if index.bucketIndex >= len(bucket) {
		err = invalidIndex(index)
		return
	}

	symbol = bucket[index.bucketIndex]

	return

}

// Acts as a constructor for SymbolTable
// Returns a hash based symbol table with list hash collision resolution.
func NewSymbolTable() SymbolTable {
	return &stImplementation{hashTable: make([][]Symbol, 10, 20)}
}
