package main


import (
"os"
"fmt"
"bufio"
"strconv"
)

func getIP(array string) string{

   fo,err := strconv.ParseInt(array[6:8], 16, 64)
   if err != nil {
       panic(err)
   } 
   var i = 0
   var ipr=strconv.Itoa(int(fo))
   for i=5 ; i> 0; i-=2 {
      fotemp,err := strconv.ParseInt(array[i-1:i+1], 16, 64)
      if err != nil {
          panic(err)
      }
      ipr=ipr+"."+strconv.Itoa(int(fotemp))

   }
   foport,err := strconv.ParseInt(array[9:], 16, 64)
   if err != nil {
     panic(err)
   }
   ipr=ipr+":"+strconv.Itoa(int(foport))
   return ipr
}


func main() {

  states := map[interface{}]interface{} {
     "01": "TCP_ESTABLISHED",
     "02": "TCP_SYN_SENT",
     "03": "TCP_SYN_SENT",
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

  opntcp,err := os.Open("/proc/net/tcp")
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
      
      local_ip=getIP(linearray[i][6:19])
      remote_ip=getIP(linearray[i][20:33])
      fmt.Println(local_ip+"\t\t"+remote_ip+"\t\t"+states[linearray[i][34:36]].(string))

  } 


}
