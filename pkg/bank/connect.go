package bank

import "sort"

// NewBank is the function that returns an instance of a Bank. Multiple accounts
// in the same bank can thus be handled properly.
type NewBank func() Bank

// rBanks holds the registered banking instancers. Each instancer has a name,
// which is often the name of the bank.
var rBanks = map[string]NewBank{}

// Banks returns the sorted list of registered banking instancers.
func Banks() []string {
	list := make([]string, 0, len(rBanks))
	for name := range rBanks {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

// SetBank registers a banking instancer.
func SetBank(name string, b NewBank) {
	if _, ok := rBanks[name]; ok {
		panic("bank `" + name + "` is already registered")
	}
	rBanks[name] = b
}

// GetBank returns a specific bank instancer.
func GetBank(name string) (b NewBank, ok bool) {
	b, ok = rBanks[name]
	return
}
