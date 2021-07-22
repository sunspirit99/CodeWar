package kata
import ("strings")
import("reflect")
func count(word string) map[string]int{
  words := strings.Split(word,"")
  wc := make(map[string]int)
  for i:= range words{
    wc[words[i]]++ 
  }  
  return wc
}

func Anagrams(word string, words []string) []string {
    // your code

  m := count(word)

  a := []string {}
  for _,w := range words{
      n := count(w)
      if reflect.DeepEqual(m, n) {
        a = append(a, w)
      }
  }  
  if len(a) == 0{
    return nil
  }
  return a
}