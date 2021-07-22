package kata
import ("math")
func FindNb(m int) int {
  // your code
  k := float64(m)
  s := 0.0
  var n = 0.0
  for s < k {
    n += 1
    s += math.Pow(n,3)
  }
  if s == k {
    return int(n)
  }
  return -1
  

}