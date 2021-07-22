package kata


func Tickets(peopleInLine []int) string {
    // Do the magic
  balance := map[int] int{ 25 : 0, 50 : 0, 100 : 0 }
    
  for i:= 0; i < len(peopleInLine) ; i++{
    b := peopleInLine[i]
      switch b{
        case 100 :
          if balance[25] == 0{
            return "NO"
          }
          if balance[50] == 0{
            if balance[25] < 3{
              return "NO"
            }
            balance[25] -= 3
  
          }else{
            balance[100] += 1
            balance[50] -= 1
            balance[25] -= 1
          }
        case 50 :
          if balance[25] == 0{
            return "NO"
          }
            balance[50] += 1
            balance[25] -= 1
        case 25 :
            balance[25] += 1     
      }
  }
  return "YES"
}