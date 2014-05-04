package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Item struct {
    Title string
    URL   string
}

type Response struct {
    Data struct {
        Children []struct {
            Data Item
        }
    }
}

func Get(reddit string) ([]Item, error) {
    url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, errors.New(resp.Status)
    }
    r := new(Response)
    err = json.NewDecoder(resp.Body).Decode(r)
    if err != nil {
        return nil, err
    }
    items := make([]Item, len(r.Data.Children))
    for i, child := range r.Data.Children {
        items[i] = child.Data
    }
    return items, nil
}

func main() {
    items, err := Get("golang")
    //resp, err := http.Get("http://reddit.com/r/golang.json")
    if err != nil {
        log.Fatal(err)
    }
    for _, item := range items {
        fmt.Println(item.Title)
    }

    //if resp.StatusCode != http.StatusOK {
    //    log.Fatal(resp.Status)
    //}
    ////    _, err = io.Copy(os.Stdout, resp.Body)
    //r := new(Response)
    //err = json.NewDecoder(resp.Body).Decode(r)
    //for _, child := range r.Data.Children {
    //    fmt.Println(child.Data.Title)
    //}

    //if err != nil {
    //    log.Fatal(err)
    //}
}