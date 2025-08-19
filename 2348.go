package main

func zeroFilledSubarray(nums []int) int64 {
    n := len(nums)
    var cnt, run int64 = 0,0
    for i:=0; i<n; i++{
        if nums[i] == 0 {
            run++
        } else {
            run = 0
        }
        cnt += run
    }
    return cnt
}