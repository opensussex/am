package main

import (
    "fmt"
    "os"
    //"io/ioutil"
    //"encoding/json"
    "time"

)


func main() {
    //fmt.Println(len(os.Args), os.Args)
    
    if(len(os.Args) > 1){
    	arg_values := os.Args
        switch(arg_values[1]){
            case `help`,`h`:
            
            message := `
                Welcome to am help
                am is a command line time tracker
                `;
                fmt.Println(message)
            break

            case `time`,`t`:
               fmt.Println(getTime())
            break

            case `start`,`s`:
                fmt.Printf("tracking started at %v\n", getTime())
            break

            case `end`,`e`:
                fmt.Printf("tracking ended at %v\n", getTime())
            break

            default:
                message := arg_values[1] + ` :argument not recognised try :am help for how to use`
                fmt.Println(message)
            break;
        }

    }else{
        fmt.Println("No arguments provided try :am help for how to use")
    }
}

func getTime() time.Time{
    return  time.Now()
}