package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
)

// go/scanner包提供词法分析功能，将源代码转换为一系列的token，以供go/parser使用
func main() {
	goFile, err := ioutil.ReadFile("./var/test.go")
	if err != nil {
		panic(err)
	}

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(goFile))
	s.Init(
		file,                 // token的结构体
		goFile,               // go文件内容
		nil,                  // 错误处理函数
		scanner.ScanComments, // 扫描模式,是否扫描注释
	)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
	/*
		结果
		位置(行:列)  token     内容
			1:1   package "package"
			1:9   IDENT   "main"
			1:13  ;       "\n"
			2:2   import  "import"
			2:9   (       ""
			3:3   STRING  "\"fmt\""
			3:8   ;       "\n"
			4:2   )       ""
			4:3   ;       "\n"
			6:2   type    "type"
			6:7   IDENT   "Hello"
			6:13  struct  "struct"
			6:19  {       ""
			6:20  }       ""
			6:21  ;       "\n"
			8:2   func    "func"
			8:7   (       ""
			8:8   IDENT   "this"
			8:13  *       ""
			8:14  )       ""
			8:16  IDENT   "Print"
			8:21  (       ""
			8:22  )       ""
			8:24  {       ""
			9:3   IDENT   "fmt"
			9:6   .       ""
			9:7   IDENT   "Println"
			9:14  (       ""
			9:15  STRING  "\"hello\""
			9:22  )       ""
			9:23  ;       "\n"
			10:2  }       ""
			10:3  ;       "\n"
			12:2  COMMENT "// Main"
			13:2  func    "func"
			13:7  IDENT   "main"
			13:11 (       ""
			13:12 )       ""
			13:13 {       ""
			14:3  IDENT   "h"
			14:5  :=      ""
			14:8  &       ""
			14:9  IDENT   "Hello"
			14:14 {       ""
			14:15 }       ""
			14:16 ;       "\n"
			15:3  IDENT   "h"
			15:4  .       ""
			15:5  IDENT   "Print"
			15:10 (       ""
			15:11 )       ""
			15:12 ;       "\n"
			16:2  }       ""
			16:3  ;       "\n"
			17:2  EOF     ""
	*/
}
