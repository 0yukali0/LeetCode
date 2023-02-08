import (
    "container/heap"
    // "fmt"
)

type Item struct {
	value int
    sliceIndex int
	index int 
}

type MaxQueue []*Item
func (pq MaxQueue) Len() int { return len(pq) }
func (pq MaxQueue) Less(i, j int) bool { return pq[i].value > pq[j].value }
func (pq MaxQueue) Swap(i, j int) { 
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
	pq[j].index = j
 }
func (pq *MaxQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MaxQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func maxProfit(prices []int) int {
    max := 0
    if len(prices) < 2 {
        return max
    }

    maxs := make(MaxQueue, 0)
    for index, value := range prices {
        item := &Item {
            value: value,
            sliceIndex: index,
        }
        // fmt.Printf("Parse %v\n", item)
        heap.Push(&maxs, item)
    }

    for index, value := range prices {
        if index == len(prices) { continue }
        // fmt.Printf("Buy %d %d\n", index, value)
        bk := make([]*Item, 0)
        for maxs.Len() > 0 {
            tmp := heap.Pop(&maxs).(*Item)
            // fmt.Printf("Try %d %d\n", tmp.sliceIndex, tmp.value)
            if profit := tmp.value - value ;tmp.sliceIndex <= index {
                continue
            } else if profit < 0 || profit <= max {
                // fmt.Print("Invalid\n")
                bk = append(bk, tmp)
                break
            } else {
                // fmt.Printf("Buy and profit is %d\n", profit)
                max = profit
                bk = append(bk, tmp)
            }
        }

        for _, item := range bk {
            heap.Push(&maxs, item)
        }
    }
    return max
}
