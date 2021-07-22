package kata
// import ("sort")
// import ("math")

func step (x int, y int) int {
    for  x != y  {
        if x > y {
            if  x % y == 0 {
              return y
            } 
            x %= y
        }else {
            if y % x == 0 {
              return x
            } 
              y %= x
        }
    }
    return x
}


func Solution(arr []int) int{
    if len(arr) == 1 {
      return arr[0]
    }

    var x int
    y := arr[0]
    for i := 1; i < len(arr); i++ {
        x = arr[i]
        x = step(x, y)
        if x == 1 {
          return len(arr)
        } 
        y = x
    }
    return x * len(arr)
}