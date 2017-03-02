package qsdiff

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// KeyValue stores for key, left and right values
type KeyValue struct {
	key, left, right string
}

// Print outputs the left and right value to stdout
func (k *KeyValue) Print() {
	var l, r string
	var c *color.Color

	if l = k.left; k.left == "" {
		l = "-"
	}
	if r = k.right; k.right == "" {
		r = "-"
	}

	if l == r {
		c = color.New(color.FgGreen)
	} else {
		c = color.New(color.FgRed)
	}

	fmt.Printf("%s\t", k.key)
	c.Printf("%s%s\n", padRight(l, 40), padRight(r, 40))
	color.Unset()
}

func padRight(s string, l int) string {
	return s + strings.Repeat(" ", l-len(s))
}

func checkNotEmpty(str string) bool {
	if len(str) == 0 {
		fmt.Println("Query string must not be empty")
		return false
	}

	return true
}

func split(str string) map[string]string {
	m := make(map[string]string)
	kvs := strings.Split(strings.Trim(str, "? "), "&")
	//res := make([]*string, len(kvs))

	for _, part := range kvs {
		eq := strings.Index(part, "=")
		key := part[:eq]
		value := part[eq+1:]
		m[key] = value
	}

	return m
}

func appendKey(keys []*string, newKey string) []*string {
	if len(keys)+1 > cap(keys) {
		// Double the capacity, if needed (+1 if cap is 0)
		new := make([]*string, len(keys), (cap(keys)*2)+1)
		for i := range keys {
			new[i] = keys[i]
		}
		keys = new
	}

	// Re-slice keys in order to add new key
	end := len(keys)
	keys = keys[:end+1]
	keys[end] = &newKey

	return keys
}

func getUniqueKeys(list []string) []*string {
	m := make(map[string]bool)

	// Amateur set
	for _, k := range list {
		m[k] = true
	}

	s := make([]*string, 0, 10)
	for k := range m {
		tmp := k
		s = append(s, &tmp)
	}

	return s
}

func buildResultList(left, right string) map[string]*KeyValue {
	m := make(map[string]*KeyValue)

	l := split(left)
	r := split(right)

	keyList := make([]string, 0, len(l)+len(r))
	for k := range l {
		keyList = append(keyList, k)
	}

	for k := range r {
		keyList = append(keyList, k)
	}

	for _, uk := range getUniqueKeys(keyList) {
		v := KeyValue{*uk, l[*uk], r[*uk]}
		m[*uk] = &v
	}

	return m
}

// Diff compares two query strings and returns a map of key to KeyValues
func Diff(a, b string) map[string]*KeyValue {
	checkNotEmpty(a)
	checkNotEmpty(b)

	return buildResultList(a, b)
}
