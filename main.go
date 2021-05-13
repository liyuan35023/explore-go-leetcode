package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Person struct {
	age   int
	child *Person
}

func (p *Person) modify() {
	p.age = 10
	p.child.age++
}

type Person2 struct {
	age   int
	child *Person
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

func TestGenBlockIdSeqId() *Person {
	objID := uint64(341975299086352768)
	blockID, seqID := GenBlockIdSeqId(objID)
	groupID := GetGroupID(objID)
	fmt.Println(blockID, seqID, groupID)
	//assert.Equal(t, uint32(objID>>32), blockID)
	//assert.Equal(t, uint32(objID&0xffffffff), seqID)
	p1 := new(Person)
	p2 := new(Person)
	//p1 := &Person{}
	//p2 := &Person{}
	println(p1)
	println(p2)
	//fmt.Printf(" p2:%p", p2)
	return p2
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
	for i := 0; i < randSize; i++ {
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
		dp[i] = (dp[i-1]%1000000007 + dp[i-2]%1000000007) % 1000000007
	}
	return dp[n]
}

func dataRace() (ret int) {
	ret = 5
	go func() {
		i := ret
		fmt.Println(i)
	}()
	return 6
}

func goMaxProcs() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(19)
	for i := -1; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("A: %d\n", i)
		}()
	}

	for i := -1; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("B: %d\n", i)
		}(i)
	}
	wg.Wait()

}

func structCompare() {
	type com struct {
		a int
		_ string
		s [3]int
		sli []int
	}
	c1 := com{sli: []int{3,4}}
	c2 := com{sli: []int{1,2,3}}
	fmt.Println(reflect.DeepEqual(c1, c2))

	v1 := struct {
		A int
		B string
	} {
		A: 10,
		B: "bb",
	}
	v2 := struct {
		A int
		B string
	} {
		A: 10,
		B: "bbc",
	}
	fmt.Println(v1 == v2)
}

func gmp() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
		}
	}()

	time.Sleep(time.Millisecond)
	println("OK")
}

func regex() {
	bucketUpper := regexp.MustCompile("^http://(.*?)\\.bjcnc.*?/(.*)")
	//bucketLower := regexp.MustCompile("http://bjcnc.*?/(.*?)/(.*)")
	res := bucketUpper.FindStringSubmatch("http://app-content.bjcnc.scs.sohucs.com/article/6784756376276041728")
	match := bucketUpper.Match([]byte("http://app-content.bjcnc.scs.sohucs.com/article/6784756376276041728"))
	fmt.Println(match)
	fmt.Println(len(res))
	fmt.Println(res)
	//
	//res2 := bucketLower.FindAllStringSubmatch("http://bjcnc.scs-internal.sohucs.com/app-content/article/6783718609412685839", -1)
	//
	//fmt.Println(res2)

	//numberRegex := regexp.MustCompile("abc([0-9]*)(f)")
	//res3 := numberRegex.FindAllStringSubmatch("abc44fabc6f", 1)
	//fmt.Println(res3)


}

func main() {

	//regex()
	//gmp()
	//structCompare()
	//dataRace()
	//goMaxProcs()

	//ans := finbina(45)
	//fmt.Println(ans)

	//TestGenGroupId()
	p := TestGenBlockIdSeqId()
	fmt.Println(p)
	//s := "9223372036854775808"
	//fmt.Println(__Atoi.MyAtoi(s))

	var inter *Person
	fmt.Println(reflect.TypeOf(inter))
	//var yy *Person2 = nil
	fmt.Println(inter == nil)

	//s := __Longest_subString.LengthOfLongestSubstringBruteForce(" ")
	//s := __Longest_subString.LengthOfLongestSubstring("abcabcbb")
	//fmt.Println(s)
}
