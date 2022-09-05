/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    if root == nil {
        return [][]int{}
    }

    ans:=[][]int{}
    q:=[]*Node{root}
    for len(q)>0 {
        n:=len(q)
        level:=[]int{}
        for i:=0;i<n;i++{
            root, q = q[0],q[1:]
            level = append(level,root.Val)
            q = append(q,root.Children...)
        }
        ans = append(ans,level)
    }
    return ans
}
