package qsdiff

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	keyWidth = 10
	colWidth = 35
)

// KeyValue stores for key, left and right values
type KeyValue struct {
	key, left, right string
}

// Print diff status to stdout
func (k *KeyValue) Print() {
	var c *color.Color
	var l, r string

	if k.left == k.right {
		c = color.New(color.FgGreen)
	} else {
		c = color.New(color.FgRed)
	}

	if l = k.left; l == "" {
		l = "<nil>"
	}

	if r = k.right; r == "" {
		r = "<nil>"
	}

	fmt.Printf("%s", padRight(k.key, keyWidth))
	c.Printf("%s%s\n", padRight(l, colWidth), padRight(r, colWidth))
	color.Unset()
}

func padRight(s string, l int) string {
	if l-len(s) < 0 {
		return s
	}
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

	for _, part := range kvs {
		eq := strings.Index(part, "=")
		key := part[:eq]
		value := part[eq+1:]
		m[key] = value
	}

	return m
}

func getUniqueKeys(list []string) []*string {
	m := make(map[string]bool)

	// Makeshift Set type
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

// Diff compares two query strings and returns a map of key => KeyValues
func Diff(a, b string) map[string]*KeyValue {
	checkNotEmpty(a)
	checkNotEmpty(b)

	return buildResultList(a, b)
}
