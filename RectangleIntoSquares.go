package kata

func SquaresInRect(lng int, wdth int) []int {
    // your code
  if lng == wdth {
    return nil
  }
  a := []int {}
  for wdth > 0 && lng > 0 {
    if lng > wdth {
      a = append(a,wdth)
      lng -= wdth
    }else{
      a = append(a,lng)
      wdth -= lng
    }
  }
  return a
}