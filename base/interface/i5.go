/*
go语言中的排序 sort.Sort()
*/

package main

import (
	"fmt"
	"sort"
)

type MyStrings []string

func (ms MyStrings) Len() int {
	return len(ms)
}

func (ms MyStrings) Less(i, j int) bool {
	return ms[i] < ms[j]
}

func (ms MyStrings) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

//自定义切片进行排序
func sortMyStrings() {
	slice := MyStrings{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(slice)

	for _, s := range slice {
		fmt.Println(s)
	}
}

// 使用go自带StringSlice排序
func sortSliceString() {
	slice := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(slice)

	for _, s := range slice {
		fmt.Println(s)
	}
}

// 使用go自带SliceInt排序
func sortSliceInt() {
	slice := sort.IntSlice{9, 44, 454646, 54454, 22, 45}

	sort.Sort(slice)

	for _, s := range slice {
		fmt.Println(s)
	}
}

// 对象排序
type HeroKind int

const (
	Tank HeroKind = iota + 1
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

func (h Heros) Len() int {
	return len(h)
}

func (h Heros) Less(i, j int) bool {
	if h[i].Kind != h[j].Kind {
		return h[i].Kind < h[j].Kind
	}
	return h[i].Name < h[j].Name
}

func (h Heros) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 对对象进行排序
func sortHeros() {
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	sort.Sort(heros)
	for _, v := range heros {
		fmt.Println(v)
	}
}

func main() {
	sortMyStrings()
	sortSliceString()
	sortSliceInt()
	sortHeros()
}
