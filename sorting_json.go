package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Extra struct {
	Balance float64 `json:"balance"`
}

type Account struct {
	Balance       float64 `json:"balance"`
	AccountNumber int64   `json:"account_no"`
	Extra         *Extra  `json:"extra,omitempty"`
}

type SortedAccount struct {
	Name    string
	Balance float64
}

var sortedArr []SortedAccount

func jsonSorting(input string) {
	var data map[string]Account
	newInput := convertJson(input)
	err := json.Unmarshal([]byte(newInput), &data)
	if err != nil {
		fmt.Println(err)
	}
	for name, acct := range data {
		balance := acct.Balance
		if acct.Extra != nil {
			balance = acct.Extra.Balance
		}
		sortedArr = append(sortedArr, SortedAccount{Name: name, Balance: balance})
	}
}
func convertJson(input string) string {
	var raw map[string]json.RawMessage
	err := json.Unmarshal([]byte(input), &raw)
	if err != nil {
		fmt.Println(err)
	}
	extraRaw, hasExtra := raw["extra"]
	if !hasExtra {
		return input
	}
	delete(raw, "extra")
	for name := range raw {
		var acc map[string]json.RawMessage
		json.Unmarshal(raw[name], &acc)
		acc["extra"] = extraRaw
		fixed, _ := json.Marshal(acc)
		raw[name] = fixed
	}
	result, _ := json.Marshal(raw)

	return string(result)
}
func PrintSorted() {
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i].Balance < sortedArr[j].Balance
	})
	for _, e := range sortedArr {
		fmt.Printf("%s: %s\n", e.Name, formatBalance(e.Balance))
	}
}
func formatBalance(n float64) string {
	s := fmt.Sprintf("%d", int64(n))
	if len(s) <= 3 {
		return s
	}

	var result []byte
	for i, c := range s {
		remaining := len(s) - i
		if i > 0 && remaining%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, byte(c))
	}
	return string(result)
}
