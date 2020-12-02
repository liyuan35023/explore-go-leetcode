package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)
type Person struct {
	age int
	child *Person
}

func (p *Person) modify() {
	p.age = 10
	p.child.age++
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[int(d)]
}

func TestGenBlockIdSeqId() {
	objID := uint64(122442667136123138)
	blockID, seqID := GenBlockIdSeqId(objID)
	groupID := GetGroupID(objID)
	fmt.Println(blockID, seqID, groupID)
	//assert.Equal(t, uint32(objID>>32), blockID)
	//assert.Equal(t, uint32(objID&0xffffffff), seqID)
}

func TestGenGroupId() {
	blockId := uint32(2836495)
	fmt.Println(GenGroupId(blockId))
	seqId := uint32(7355)
	objId := GenObjectId(blockId, seqId)
	fmt.Println(objId)
}

func GenBlockIdSeqId(objectId uint64) (blockId, seqId uint32) {
	blockId = uint32(objectId >> 32)
	seqId = uint32((objectId << 32) >> 32)

	return
}
func GenGroupId(storeId uint32) (groupId uint32) {
	groupId = uint32(storeId >> 10)
	return
}

func GenObjectId(blockId uint32, seqId uint32) uint64 {
	objectId := (uint64(blockId) << 32) + uint64(seqId)

	return objectId
}

func GetGroupID(objID uint64) uint32 {
	blockID, seqID := GenBlockIdSeqId(objID)
	groupID := GenGroupId(blockID)
	fmt.Println("Group ID:", groupID)
	//fmt.Println("Block ID:", blockID)
	fmt.Println("Sequence ID:", seqID)
	return groupID
}

func escapee() {
	randSize := rand.Int()
	s := make([]string, 10, 100)
	//str := "hello"
	for i:=0;i<randSize;i++ {
		s = append(s, strconv.Itoa(randSize))

	}
	_ = s
}

func reEnter(rw sync.RWMutex) {
	rw.RLock()
	defer rw.RUnlock()
	fmt.Println("keyi")
}

func finbina(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		if i < 2 {
			dp[i] = i
			continue
		}
		dp[i] = (dp[i-1] % 1000000007 + dp[i-2] % 1000000007) % 1000000007
	}
	return dp[n]
}


func main() {
	sli := make([]string, 0)
	s := "fads"
	sli = append(sli, s)
	s = "fda"
	fmt.Println(sli[0])

	//third := _4_binary_tree_inorder.TreeNode{Val: 3}
	//second := _4_binary_tree_inorder.TreeNode{Val: 2, Left: &third}
	//root := _4_binary_tree_inorder.TreeNode{Val: 1, Right: &second}
	//
	//ret := _4_binary_tree_inorder.InorderTraversalLoop(&root)

	//head := &_48_sort_list.ListNode{Val: -1}
	//two := &_48_sort_list.ListNode{Val: 5}
	//tree := &_48_sort_list.ListNode{Val: 3}
	//four := &_48_sort_list.ListNode{Val: 4}
	//five := &_48_sort_list.ListNode{Val: 0}
	//head.Next = two
	//two.Next = tree
	//tree.Next = four
	//four.Next = five
	//newHead := _48_sort_list.SortListLoop(head)
	//for newHead != nil {
	//	fmt.Println(newHead.Val)
	//	newHead = newHead.Next
	//}


	//[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]

 	//ans := finbina(45)
 	//fmt.Println(ans)

	//TestGenGroupId()
	TestGenBlockIdSeqId()
	//s := "9223372036854775808"
	//fmt.Println(__Atoi.MyAtoi(s))



	//s := __Longest_subString.LengthOfLongestSubstringBruteForce(" ")
	//s := __Longest_subString.LengthOfLongestSubstring("abcabcbb")
	//fmt.Println(s)
}
