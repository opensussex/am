package main

import (
    "fmt"
    "os"
    "os/user"
    "time"
    "encoding/csv"

)


func main() {
    //fmt.Println(len(os.Args), os.Args)
   
    usr, err := user.Current() 
    if err != nil {
        fmt.Println( err )
    }

    if(!file_exists(usr.HomeDir + `/.am-store`)){ // we want to create the .am-store file
        am_store, err := os.Create(usr.HomeDir + `/.am-store`)
        defer am_store.Close() // lets the file.
        if err != nil { panic(err) }
    }

    if(len(os.Args) > 1){

    	arg_values := os.Args

        csvFile, err := os.OpenFile(usr.HomeDir + `/.am-store`,os.O_RDWR , 0777)
        defer csvFile.Close()

        if err != nil {
            panic(err)
        }

        csvReader := csv.NewReader(csvFile)
        csvWriter := csv.NewWriter(csvFile)
        tasks,err := csvReader.ReadAll()
        if err != nil{
            panic(err)
        }

        current_task:= tasks[len(tasks)-1]

        switch(arg_values[1]){
            case `help`,`h`:
            
            message := `
                Welcome to am help
                am is a command line time tracker
                arguments:
                
                start (s)
                end (e)
                time (t)
                help (h)
                list (l)
                now (n)

                usage : 
                am s <task name>
                starts tracking <task name>

                am e
                ends the tracking of a current task
                
                am t
                gives current time
                `;
                fmt.Println(message)
            break

            case `time`,`t`:
               fmt.Println(getTime())
            break

            case `start`,`s`:
                if(len(os.Args) >=3){
                    fmt.Printf("tracking started at %v on task %v\n", getTime(),arg_values[2])
                    new_task := []string{arg_values[2],getTime(),`now`}
                    tasks = append(tasks,new_task)
                    csvWriter.WriteAll(tasks)
                }else{
                    fmt.Println("task to track required try : am help for how to use")
                }
            break

            case `end`,`e`:
                if(current_task[2] == `now`){
                    fmt.Printf("tracking %s ended at %v\n", current_task[0],getTime())
                    tasks[len(tasks)-1][2] = getTime()
                    csvWriter.WriteAll(tasks)
                }else{
                    fmt.Println("You're not tracking any task!")
                }
                
            break

            case `now`,`n`:
              
                if(current_task[2] == `now`){
                    fmt.Printf("You're working on %s since %s \n", current_task[0],current_task[1])
                }else{
                    fmt.Println(`You're not tracking any task`)
                }
            
                
            break

            case `list`,`l`:

                for _,task := range tasks {
                    fmt.Println(task)
                }
                
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

func getTime() string{
    return  time.Now().Format("Mon Jan _2 15:04:05 2006")
}

func file_exists(file string) bool{
    if _, err := os.Stat(file); os.IsNotExist(err) {
        //fmt.Printf("no such file : %s \n", file)
        return false
    }
    return true

}