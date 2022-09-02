var ROWS int
var COLS int
func pacificAtlantic(heights [][]int) [][]int {
    atl:=make(map[int][]int)
    pac:=make(map[int][]int)
    ans:=[][]int{}
    ROWS = len(heights)
    COLS = len(heights[0])
    for i:=0;i<COLS;i++{
        dfs(0,i,pac,heights[0][i],heights)
        dfs(len(heights)-1,i,atl,heights[ROWS-1][i],heights)
    }
    
    for i:=0;i<len(heights);i++{
        dfs(i,0,pac,heights[i][0],heights)
        dfs(i,COLS-1,atl,heights[i][COLS-1],heights)
    }
    for k,v:=range atl {
        _,ok:=pac[k]
        if ok {
            ans = append(ans,v)
        }
    }
    return ans
}
func dfs(r,c int, seen map[int][]int, prevHeight int, heights [][]int) {
    if r<0 || c<0 || r>=ROWS || c>=COLS || prevHeight>heights[r][c]{
        return 
    }
    _,ok:=seen[c+(r*COLS)]
    if ok {
        return
    }
    seen[c+(r*COLS)] = []int{r,c}
    dfs(r+1,c,seen,heights[r][c],heights)
    dfs(r-1,c,seen,heights[r][c],heights)
    dfs(r,c+1,seen,heights[r][c],heights)
    dfs(r,c-1,seen,heights[r][c],heights)
}
