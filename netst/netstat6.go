package main


import (
"os"
"fmt"
"bufio"
"strconv"
"strings"
)

func readbytes(j int,k int,array string) string{

   var ipr6 = ""
   var i = 0
   var fortemp = ""
   for i=k; i > j; i=i-2 {
     fortemp = array[i-1:i+1]
     if i == 3 || (i == 7 &&  j != 0) || i == 11 || i == 15 || i == 19 || i == 23 || i == 27 || i == 31 {
        ipr6=ipr6+":"+fortemp
     }else{
        ipr6=ipr6+fortemp
    }
   }
  return ipr6

}


func getIPV6(array string) string{
   var ipr6 = ""
   ipr6=readbytes(0,7,array)
   ipr6= ipr6+readbytes(8,15,array)
   ipr6= ipr6+readbytes(16,23,array)
   ipr6= ipr6+readbytes(24,31,array)
   foport,err := strconv.ParseInt(strings.TrimSpace(array[33:]), 16, 64)
   if err != nil {
     panic(err)
   }
   ipr6=ipr6+":"+strconv.Itoa(int(foport))
   return ipr6
}


func main() {

  states := map[interface{}]interface{} {
     "01": "TCP_ESTABLISHED",
     "02": "TCP_SYN_SENT",
     "03": "TCP_SYN_RECV",
     "04": "TCP_FIN_WAIT1",
     "05": "TCP_FIN_WAIT2",
     "06": "TCP_TIME_WAIT",
     "07": "TCP_CLOSE",
     "08": "TCP_CLOSE_WAIT",
     "09": "TCP_LAST_ACK",
     "0A": "TCP_LISTEN",
     "0B": "TCP_CLOSING",
     "0C": "TCP_NEW_SYN_RECV",

 
  }

  opntcp,err := os.Open("/proc/net/tcp6")
  if err != nil {
    fmt.Println("Cloud not open the file")
  }
  defer opntcp.Close()
  fileScanner := bufio.NewScanner(opntcp) 
  fileScanner.Split(bufio.ScanLines)
  var linearray []string
  for fileScanner.Scan() {
        linearray = append(linearray,fileScanner.Text())
  }
  if err = fileScanner.Err(); err != nil {
       fmt.Println("cloud not close the file due to this error %s error \n", err)
  }
  var local_ip=""
  var remote_ip=""
  fmt.Println("local_address\t\tRemote_address\t\tstates")
  for i :=1; i< len(linearray); i++ {
      
      local_ip=getIPV6(linearray[i][6:43])
      remote_ip=getIPV6(linearray[i][44:82])
      fmt.Println(local_ip+"\t\t"+remote_ip+"\t\t"+states[linearray[i][82:84]].(string))

  } 


}
