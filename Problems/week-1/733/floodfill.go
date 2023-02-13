import (
    "fmt"
)

type pixel struct {
    x int
    y int
}

func (p pixel) Add(x, y int) pixel {
    return pixel{p.x + x, p.y + y}
}

func (p pixel) IsValid(xbound, ybound int) bool {
    if p.x >= xbound || p.x < 0 || p.y >= ybound || p.y < 0 {
        return false
    }
    return true
}

func fourways(image [][]int, input pixel) []pixel {
    result := make([]pixel, 0)
    xbound, ybound := len(image), len(image[0])
    directions := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
    for _, d := range directions {
        if e := input.Add(d[0], d[1]); e.IsValid(xbound, ybound) {
            result = append(result, e)
        }
    }
    return result
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    record := map[string]bool{fmt.Sprintf("%v,%v", sr, sc): true}
    waitToVisit := []pixel{pixel{sr, sc}}
    for index := 0; index < len(waitToVisit); index++ {
        candicates := fourways(image, waitToVisit[index])
        for _, c := range candicates {
            key := fmt.Sprintf("%v,%v", c.x, c.y)
            if _, ok := record[key]; !ok && image[c.x][c.y] == image[sr][sc] {
                waitToVisit = append(waitToVisit, c)
                record[key] = true
            }
        }
    }

    for _, v := range waitToVisit{
        image[v.x][v.y] = color
    }
    return image
}
