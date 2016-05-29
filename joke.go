// Copyright 2016 The go-deaf Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the go LICENSE file.

// package main

// import "fmt"

// func main() {
//     fmt.Printf("hello world, joke! \n")
// }

package main

import (
    "github.com/dongjinxian/deaf/log"
)

func main() {

    name := "djx"
    log.Debug("My name is %v", name)
    log.Release("My name is %v", name)
    log.Error("My name is %v", name)
    // log.Fatal("My name is %v", name)

    // logger, err := log.New("debug", "test_djx_log")
    // if err != nil {
    //     return
    // }
    // defer logger.Close()

    log.Debug("will not print")
    log.Release("My name is %v", name)

    // log.Export(logger)

    log.Debug("will not print")
    log.Release("My name is %v", name)
}


// package main

// import (
//     "fmt"
//     "log"
//     "os"
// )

// func main() {

//     fmt.Println("begin TestLog ...")
//     file, err := os.Create("test_djx.log")
//     if err != nil {
//         log.Fatalln("fail to create test.log file!")
//     }
//     logger := log.New(file, "", log.LstdFlags|log.Llongfile)
//     log.Println("1.Println log with log.LstdFlags ...")
//     logger.Println("1.Println log with log.LstdFlags ...")

//     logger.SetFlags(log.LstdFlags)

//     log.Println("2.Println log without log.LstdFlags ...")
//     logger.Println("2.Println log without log.LstdFlags ...")

//     //log.Panicln("3.std Panicln log without log.LstdFlags ...")
//     //fmt.Println("3 Will this statement be execute ?")
//     //logger.Panicln("3.Panicln log without log.LstdFlags ...")

//     log.Println("4.Println log without log.LstdFlags ...")
//     logger.Println("4.Println log without log.LstdFlags ...")

//     log.Fatal("5.std Fatal log without log.LstdFlags ...")
//     fmt.Println("5 Will this statement be execute ?")
//     logger.Fatal("5.Fatal log without log.LstdFlags ...")
// }