package kata
import ("strings")
func DuplicateEncode(word string) string {
  word = strings.ToLower(word)
  a := []string {}
  k := strings.Split(word,"")
  for _,v := range k{
    if (strings.Count(word,v) == 1){
      a = append(a,"(")
    }else{
      a = append(a,")") 
    }  
  } 
  return strings.Join(a,"")
}