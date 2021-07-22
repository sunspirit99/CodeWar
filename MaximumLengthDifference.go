package kata
import ("math")
func MxDifLg(a1 []string, a2 []string) int {
    // your code
  if len(a1) == 0 || len(a2) == 0{
    return -1
  }
  min1 , min2 , max1 , max2 := len(a1[0]) , len(a2[0]) , len(a1[0]) , len(a2[0])
  
  for _,v := range a1{
    if len(v) < min1{
      min1 = len(v)
    }else if len(v) > max1{
      max1 = len(v)
    }
  }
  
  for _,v := range a2{
    if len(v) > max2{
      max2 = len(v)
    }else if len(v) < min2{
      min2 = len(v)
    }
  }
  if (max1 - min2 > max2 - min1) {
    return int(math.Abs(float64(min2 - max1)))
  }
    return int(math.Abs(float64(min1 - max2)))
  
}