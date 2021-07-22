package kata
import ("strings")
func DNAStrand(dna string) string {
  // your code here
  d := strings.Split(dna,"")
  for i,v := range d{
    if v == "T"{
      d[i] = "A"
    }else if v == "A"{ 
      d[i] = "T"
    }else if v == "C"{
      d[i] = "G"
    }else if v == "G"{
      d[i] = "C"
    }
  }
  return strings.Join(d,"")
}