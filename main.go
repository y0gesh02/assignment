The intuition is to find the linear ordering in which the tasks will be performed if it is possible to perform all the tasks otherwise, to return an empty array.

So, first, we need to identify a graph as a directed acyclic graph and if it is so we need to find out the linear ordering of the nodes as well.

Now, we have successfully reduced the problem to one of our known concepts: Detect cycle in a directed graph. We will solve this problem using the Topological Sort Algorithm or Kahn’s Algorithm.

We will return true if the length of the ordering equals the number of tasks. otherwise, we will return false.



//Main Fuction To Solve The Problem
//Recieving n and Prerequisites from main Func
func findOrder(n int, prerequisites [][]int) []int {
    
    graph := make(map[int][]int, n) //prereqs-course list of next courses
    indegree := make([]int, n)       // indegree[i] denotes number of prerequisite courses for ith course
    ans := []int{} //
   
    //forming adjacency list graph
    for _, node := range prerequisites {  
        graph[node[1]] = append(graph[node[1]], node[0])
        indegree[node[0]]++
    }
    
    // Applying Topological Sort
    q := list.New()
    for i:=0; i<n; i++ {
        if indegree[i] == 0 {
            q.PushBack(i)
        }
    }
    for q.Len() > 0 {
        u := q.Front().Value.(int)
        ans = append(ans, u)
        q.Remove(q.Front())
        for _, v := range graph[u] {
            indegree[v]--
            if indegree[v] == 0 {
                q.PushBack(v)
            }
        }
    }
    
    if len(ans) == n {  // returning ans if  all courses can be finished
        return ans
    }
    return []int{}  //return empty array
}
