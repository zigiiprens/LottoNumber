package main

import(
    "testing"
)

func Testlottonumber(t *testing.T){
    x := lottonumber(1)
    for i, v := range x{
            if (v[i] > 49 == true) {
            t.Error("expected", "[0-49]", "got", "[50--]")
            }         
        
    }
    
}
