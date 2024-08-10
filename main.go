package main

type Course struct {
  Name string `json:"coursename"`
}

func check(e error)  {
  if e != nil {
    panic(e)
  }
}

func main()  {
  
}
