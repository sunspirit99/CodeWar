package kata

func FindOutlier(integers []int) int {
  a := []int{}
  b := []int{}
  for i:= range integers{
     if integers[i] % 2 == 0{
       a = append(a,integers[i])
     }else{
       b = append(b,integers[i])
     }  
  }  
  if(len(a) > len(b)){
    return b[0]
  }
  return a[0]
}