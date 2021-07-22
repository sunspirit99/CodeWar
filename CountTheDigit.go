package kata
import ("math")
import ("strings")
import("strconv")
func NbDig(n int, d int) int {
    // your code
  k := 0.0
  sum := 0
  for i:= 0; i <= n; i++{
    k = math.Pow(float64(i),2)
    s := strconv.Itoa(int(k))
    j := strconv.Itoa(d)
    sum += strings.Count(s,j)
  }
  return sum
}