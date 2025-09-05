package main

func findClosest(x int, y int, z int) int {

    difx := abs(x-z)
    dify := abs(y-z)

    if difx>dify {return 2}
    if difx<dify {return 1}
    return 0
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n

}