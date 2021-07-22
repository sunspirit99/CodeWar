package kata
import ("strings")
func CamelCase(s string) string {
    // your code here
  if len(s) == 0{
    return ""
  }
  res := []string {}
  str := strings.Fields(s)
  for _,v := range str {
    ss := strings.Split(v,"")
    ss[0] = strings.ToUpper(ss[0])
    res = append(res,strings.Join(ss,""))
  }
  return strings.Join(res,"")
}