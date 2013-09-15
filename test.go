package test

import "fmt"
import "djonwrapper"

func main() {
	fmt.Println("Hello");
   con := djonwrapper.DjondbConnectionManagerGetConnection("localhost");
   con.Open();
   con.Insert("a", "go", "{a:'1'}");
   res := con.Find("a", "go");
   for res.Next() {
      val := res.Current();
      fmt.Println(val.ToChar());
   }
   fmt.Println("Bye");
}
