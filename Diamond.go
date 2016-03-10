//Package Diamond prints a diamond 
//made out of asteriks
package Diamond

import "fmt"

//Some global variables for the loops
var n,c,k,dummy1,dummy2,space int

//Print Diamond is a global function

func PrintDiamond() {
	n:=5
	space := n-1
	for k := 1; k <=n ; k++ {
		
		for c:=1 ; c<= space;c++{
	  		fmt.Print(" ")
		}
		
		space--
		
		for c=1;c<=2*k-1;c++{
			fmt.Print("*")
		}
		
		fmt.Println("")
	}
	
	space = 1;
 
  for k = 1; k <= n - 1; k++{
	  for c = 1; c <= space; c++{
		  fmt.Print(" ")
	  }
	  
 
    space++
 
	  for c = 1 ; c <= 2*(n-k)-1; c++{
		  fmt.Print("*")		
	  }
		
		fmt.Println("")
	}
	
	space = 1;
 
  for k = 1; k <= n - 1; k++{
	  for c = 1; c <= space; c++{
		  fmt.Print(" ")
	  }
	  
 
    space++
 
	  for c = 1 ; c <= 2*(n-k)-1; c++{
		  fmt.Print("*")
	  }
 
    fmt.Println("")
  }
	
}

