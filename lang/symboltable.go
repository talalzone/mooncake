package lang

type SymbolTable struct {
	table map[string]interface{}
}

func NewSymbolTable() SymbolTable {
	return SymbolTable{table: map[string]interface{}{}}
}

func (st *SymbolTable) Get(key string) interface{} {
	return st.table[key]
}

func (st *SymbolTable) Put(key string, value interface{}) {
	st.table[key] = value
}

func (st *SymbolTable) Delete(key string) {
	delete(st.table, key)
}
