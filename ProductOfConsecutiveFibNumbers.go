package kata

func ProductFib(prod uint64) [3]uint64 {
  // your code
  a := [3]uint64 {}
  var f1 uint64 = 0
  var f2 uint64 = 1

  for {
    fn := f1 + f2
    if fn * f2 >= prod {
      if prod / fn == f2 {
        a[0] = f2
        a[1] = fn
        a[2] = 1
        break
      }else{
        a[0] = f2
        a[1] = fn
        a[2] = 0
        break
      }  
    }
    f1 = f2
    f2 = fn
  }
  return a
}