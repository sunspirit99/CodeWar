package kata

func BreakChocolate(n, m int) int {
  // math goes here
   if n == 0 || m == 0 {
     return 0
   }

   return m * n - 1
}