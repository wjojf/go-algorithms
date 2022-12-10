package main

func canJump(nums []int) bool {
    var reachable int 

    for i := 0; i < len(nums); i++{
        if i > reachable {
            return false
        }
        r := i + nums[i]
        if reachable < r {
            reachable = r
        }
    }  
    return true 
}
