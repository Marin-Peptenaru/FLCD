package parser

// github link: https://github.com/Marin-Peptenaru/FLCD/tree/master/app
import "fmt"

// type alias for symbols. A Symbol is just a string.
type Symbol string


// Type used for indexing symbols in the symbol table.
type STIndex struct {
	// tableIndex represents the bucket in which the symbol is stored.
	tableIndex  int
	// bucketIndex represents the index of a symbol inside its bucket. 
	bucketIndex int
}


// Interface of the Hash Table based symbol table.
type SymbolTable interface {
	// Takes a Symbol as parameter. Computes the hash of a symbol and checks if it is not already in the symbol table.
	// If it's not, it will add it to the symbol table based on it's hash code 
	// and will return the index at which the symbol was put inside the symbol table. 
	// If the symbol already is present in the table, it will just return the index at which it is stored.
	SaveSymbol(s Symbol) STIndex

	// Takes an STIndex as a parameter and uses it to retrieve a symbol from the symbol the table.
	// If the index is a valid one, returns a Symbol and nil.
	// If the index is an invalid one, returns an empty Symbol ("") and an error
	GetSymbol(i STIndex) (Symbol, error)
}


//Just a utility function for creating error out of invalid indexes
func invalidIndex(index STIndex) error {
	return fmt.Errorf("invalid symbol table index: %v", index);
}

func NullIndex() STIndex {
	return STIndex{tableIndex: -1, bucketIndex: -1}
}