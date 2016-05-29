// Copyright 2016 The go-deaf Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the go LICENSE file.

// 目前仅仅是一个临时的实现，以后会大幅删改
// // 日志管道
// log_ch := make([]chan string, 1024)

package log

import (
    // "errors"
    "fmt"
    "log"
    "os"
    "path"
    "strings"
    "time"
)

// levels
const (
    kDebug = 0
    kInfo  = 1
    kWarn  = 2
    kError = 3
)

const (
    kDebugName = "[Deg] "
    kInfoName  = "[Inf] "
    kWarnName  = "[War] "
    kErrorName = "[Err] "
)

type Logger struct {
    level      int
    baseLogger *log.Logger
    baseFile   *os.File
}

// 获取日志等级：不正确的日志等级，设默认信息级别
func get_log_level(levelname string) (level int) {

    switch strings.ToLower(levelname) {
    case "debug":
        level = kDebug
    case "info":
        level = kInfo
    case "warn":
        level = kWarn
    case "error":
        level = kError
    default:
        level = kInfo
    }

    // 返回
    return level
}

func New(levelname string, pathname string) (*Logger, error) {

    // 日志等级
    level := get_log_level(levelname)

    // logger
    var baseLogger *log.Logger
    var baseFile *os.File

    now := time.Now()
    filename := fmt.Sprintf("%d%02d%02d_%02d_%02d_%02d.log",
     now.Year(), now.Month(), now.Day(), now.Hour(),
      now.Minute(), now.Second())

    file, err := os.Create(path.Join(pathname, filename))
    if err != nil {
        panic(err)
    }

    baseFile = file
    baseLogger = log.New(file, "", log.LstdFlags)
    // baseLogger = log.New(file, "", log.LstdFlags|log.Llongfile)

    // 校验
    if baseLogger == nil {
        panic("baseLogger is nil")
    }
    if baseFile == nil {
        panic("baseFile is nil")
    }

    // new
    logger := new(Logger)
    logger.level = level
    logger.baseLogger = baseLogger
    logger.baseFile = baseFile

    return logger, nil
}

// It's dangerous to call the method on logging
func (logger *Logger) Close() {

    if logger.baseFile != nil {
        logger.baseFile.Close()
    }

    logger.baseLogger = nil
    logger.baseFile = nil
}

func (logger *Logger) doPrintf(level int, levelname string, format string, a ...interface{}) {

    if level < logger.level {
        return
    }
    if logger.baseLogger == nil {
        panic("logger closed")
    }

    format = levelname + format
    logger.baseLogger.Printf(format, a...)

    if level > kInfo {
       log.Printf(format, a...) 
    }

    if level == kError {
        os.Exit(1)
    }
}

func (logger *Logger) Debug(format string, a ...interface{}) {
    logger.doPrintf(kDebug, kDebugName, format, a...)
}

func (logger *Logger) Info(format string, a ...interface{}) {
    logger.doPrintf(kInfo, kInfoName, format, a...)
}

func (logger *Logger) Warn(format string, a ...interface{}) {
    logger.doPrintf(kWarn, kWarnName, format, a...)
}

func (logger *Logger) Error(format string, a ...interface{}) {
    logger.doPrintf(kError, kErrorName, format, a...)
}

var gLogger, _ = New("debug", "")

// It's dangerous to call the method on logging
func Export(logger *Logger) {
    if logger != nil {
        gLogger = logger
    }
}

func Deg(format string, a ...interface{}) {
    gLogger.Debug(format, a...)
}

func Inf(format string, a ...interface{}) {
    gLogger.Info(format, a...)
}

func War(format string, a ...interface{}) {
    gLogger.Warn(format, a...)
}

func Err(format string, a ...interface{}) {
    gLogger.Error(format, a...)
}

func Close() {
    gLogger.Close()
}
