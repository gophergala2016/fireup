package main

func Dispatcher() {
        for {
                key := <-chData
                for _, ch := range clients {
                        ch <- key
                }
        }
}

