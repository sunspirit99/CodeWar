package kata

func CreateSpiral(n int) [][]int {
  // your code here
  if n < 1 {
    return [][]int{}
  }
  a := make([][]int, n)
  for i := range a {
    a[i] = make([]int, n)
  }


  minC := 0
  minR := 0
  maxC := n - 1
  maxR := n - 1
  s := 1
  for s <= n * n  {
  if maxC >= minC && maxR >= minR {
    for i:= minC; i <= maxC; i++{
      a[minR][i] = s
      s += 1
    }
    minR += 1           // minC = 0, minR = 1, maxC = 0, maxR = 0
  }
  if maxC >= minC && maxR >= minR {
    for i:= minR; i <= maxR; i++{
      a[i][maxC] = s
      s += 1
    }
    maxC -= 1                   // minC = 0, minR = 1, maxC = 0, maxR = 1
  }               
  if maxC >= minC && maxR >= minR {
    for i:= maxC; i >= minC; i--{
      a[maxR][i] = s
      s += 1
    }
    maxR -= 1                   // minC = 0, minR = 1, maxC = 0, maxR = 0
  }
  if maxC >= minC && maxR >= minR {
    for i:= maxR; i >= minR; i--{
      a[i][minC] = s
      s += 1
    }
    minC += 1
  }
 
  }
  return a
}