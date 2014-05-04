package main

import (
    "fmt"
    "github.com/guanjie/reddit"
    "log"
)

func (i Item) String() string {
    com := ""
    switch i.Comments {
    case 0:
        // nothing
    case 1:
        com = " (1 comment)"
    default:
        com = fmt.Sprintf(" (%d comments)", i.Comments)
    }
    return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}

func main() {
    items, err := reddit.Get("golang")
    if err != nil {
        log.Fatal(err)
    }
    for _, item := range items {
        fmt.Println(item.Title)
    }
}
