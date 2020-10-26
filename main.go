package main

import (
	_6_merge_intervals "explore-go-leetcode/56-merge-intervals"
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
	objID := uint64(65662821525882663)
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
	//ret := _8_four_sum.FourSum([]int{-3,-2,-1,0,0,1,2,3}, 0)
	//fmt.Println(ret)
	//ret := _2_generate_parenttheses.GenerateParenthesis(1)
	//fmt.Println(ret)
 	//ret := _3_unique_path_II.UniquePathsWithObstacles([][]int{[]int{0,0,0}, []int{0,1,0}, []int{0,0,0}})
	//ret := _5_max_rectangle.MaximalRectangle([][]byte{[]byte{'1','0','1','0','0'}, []byte{'1','0','1','1','1'},
	//	[]byte{'1','1','1','1','1'}, []byte{'1','0','0','1','0'}})
	ret := _6_merge_intervals.Merge([][]int{[]int{1, 3}, []int{2, 6}, []int{15,18}, []int{8, 10}})
	fmt.Println(ret)


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
