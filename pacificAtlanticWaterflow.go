func pacificAtlantic(heights [][]int) [][]int {
    atl:=make(map[string][]int)
    pac:=make(map[string][]int)
    ans:=[][]int{}
    for i:=0;i<len(heights[0]);i++{
        dfs(0,i,pac,heights[0][i],heights)
        dfs(len(heights)-1,i,atl,heights[len(heights)-1][i],heights)
    }
    
    for i:=0;i<len(heights);i++{
        dfs(i,0,pac,heights[i][0],heights)
        dfs(i,len(heights[0])-1,atl,heights[i][len(heights[0])-1],heights)
    }
    for k,v:=range atl {
        _,ok:=pac[k]
      if ok {
            ans = append(ans,v)
        }
    }
    return ans
}
func dfs(r,c int, seen map[string][]int, prevHeight int, heights [][]int) {
    if r<0 || c<0 || r>=len(heights) || c>=len(heights[r]) || prevHeight>heights[r][c]{
        return 
    }
    _,ok:=seen[string(r)+","+string(c)]
    if ok {
        return
    }
    seen[string(r)+","+string(c)] = []int{r,c}
    dfs(r+1,c,seen,heights[r][c],heights)
    dfs(r-1,c,seen,heights[r][c],heights)
    dfs(r,c+1,seen,heights[r][c],heights)
    dfs(r,c-1,seen,heights[r][c],heights)
}
