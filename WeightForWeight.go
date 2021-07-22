package kata
import ("strings")
import ("strconv")
func newWeight(str string) int{
  a,_ := strconv.Atoi(str)
  s := 0
  d := 0
  for a > 0 {
    d = a % 10
    s += d
    a = a / 10
  }
  return s
}

func OrderWeight(strng string) string {
    // your code
  if len(strng) == 0 {
    return ""
  }
  s := strings.Split(strng, " ")

  for i := 0; i < len(s); i++ {
    for j := 0; j < len(s) - 1; j++ {
      if newWeight(s[j]) > newWeight(s[j + 1]) {
        temp := s[j]
        s[j] = s[j+1]
        s[j+1] = temp
      }
      if newWeight(s[j]) == newWeight(s[j + 1]){
        if strings.Compare(s[j],s[j+1]) == 1{
          temp := s[j]
          s[j] = s[j+1]
          s[j+1] = temp
        }
      }
    }
  }
  return strings.Join(s," ")
}