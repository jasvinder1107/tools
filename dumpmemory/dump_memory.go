package main

import (

"flag"
"fmt"
"os"
"bufio"
"regexp"
"strings"
"strconv"
)

func get_mem_address(mem string) string  {

  re := regexp.MustCompile(".*r..p.*")
  var rm = ""
  readable_memory := re.FindAllString(mem, -1)
  if readable_memory != nil {
   rm := strings.Join(readable_memory," ")
   return(strings.Split(rm," ")[0])

  }
  return rm

}


func dump_memory(a string,b string,openmem *os.File,processPtr string) {

   Outputfile := fmt.Sprintf("dump-%s-%s-%s.bin",a,b,processPtr)
   Output,err := os.Create(Outputfile)
   if err != nil {
     fmt.Println("not able to create Ouput file")
   }
   defer Output.Close()
   start, err1 := strconv.ParseInt(a,16,64)
   if err1 != nil {
     fmt.Println(err1)

   }
   end, err2 := strconv.ParseInt(b,16,64)
   if err2 != nil {
    fmt.Println(err2)
   } 
   _,err4 := openmem.Seek(start,0)
   if err4 != nil{
     fmt.Println(err4)
   }
   data := make([]byte,(end-start))
   memchunk,err3 := openmem.Read(data)
   if err3 != nil{
     fmt.Println(err3)
   }
   Output.Write(data[:memchunk]) 
 
}

func main() {
 processPtr := flag.String("pid", "0","Process id to dump memory")
 flag.Parse()
 if *processPtr != "0" {
  file2open := fmt.Sprintf("/proc/%s/maps",*processPtr)
  memfile := fmt.Sprintf("/proc/%s/mem", *processPtr)
  openmaps,err := os.Open(file2open)
  openmem,err1 := os.Open(memfile)
  if err != nil {
   fmt.Println("couldn't be able to open the maps file")
  }
  if err1 != nil {
   fmt.Println("Couldn't be able to open the mem file")
  }
  defer openmaps.Close()
  defer openmem.Close()
  fileScanner := bufio.NewScanner(openmaps)
  fileScanner.Split(bufio.ScanLines)
  var linearray []string
  for fileScanner.Scan() {
     linearray = append(linearray,fileScanner.Text())
  }
  if err = fileScanner.Err(); err != nil {
     fmt.Println("Could not close the file due to this error %s error \n", err)
  }
  for i := 0 ; i < len(linearray); i++ {
     if strings.Contains(linearray[i], "vsyscall") || strings.Contains(linearray[i], "vvar") {
       continue
     }
     rm := get_mem_address(linearray[i])
     if rm != "" {
       ab := strings.Split(rm,"-")
       dump_memory(strings.TrimSpace(ab[0]),strings.TrimSpace(ab[1]),openmem,*processPtr)
     }

  } 
 }else{
  fmt.Println("This function takes a mandatory argument of --pid")
 }
 
}
