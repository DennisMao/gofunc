package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

// go/parser包提供语法分析功能，将这些token转换为AST（Abstract Syntax Tree, 抽象语法树）
func main() {
	goFile, err := ioutil.ReadFile("./var/test.go")
	if err != nil {
		panic(err)
	}

	// 创建AST树
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", goFile, 0)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, f)

	//结果
	//go的编译检查也是在parse为AST树过程执行,如果代码异常是不能转换为AST树。
	/*
	     0  *ast.File {
	     1  .  Package: 1:1
	     2  .  Name: *ast.Ident {
	     3  .  .  NamePos: 1:9
	     4  .  .  Name: "main"
	     5  .  }
	     6  .  Decls: []ast.Decl (len = 3) {
	     7  .  .  0: *ast.GenDecl {
	     8  .  .  .  TokPos: 3:1
	     9  .  .  .  Tok: import
	    10  .  .  .  Lparen: 3:8
	    11  .  .  .  Specs: []ast.Spec (len = 1) {
	    12  .  .  .  .  0: *ast.ImportSpec {
	    13  .  .  .  .  .  Path: *ast.BasicLit {
	    14  .  .  .  .  .  .  ValuePos: 4:2
	    15  .  .  .  .  .  .  Kind: STRING
	    16  .  .  .  .  .  .  Value: "\"fmt\""
	    17  .  .  .  .  .  }
	    18  .  .  .  .  .  EndPos: -
	    19  .  .  .  .  }
	    20  .  .  .  }
	    21  .  .  .  Rparen: 5:1
	    22  .  .  }
	    23  .  .  1: *ast.GenDecl {
	    24  .  .  .  TokPos: 7:1
	    25  .  .  .  Tok: type
	    26  .  .  .  Lparen: -
	    27  .  .  .  Specs: []ast.Spec (len = 1) {
	    28  .  .  .  .  0: *ast.TypeSpec {
	    29  .  .  .  .  .  Name: *ast.Ident {
	    30  .  .  .  .  .  .  NamePos: 7:6
	    31  .  .  .  .  .  .  Name: "Hello"
	    32  .  .  .  .  .  .  Obj: *ast.Object {
	    33  .  .  .  .  .  .  .  Kind: type
	    34  .  .  .  .  .  .  .  Name: "Hello"
	    35  .  .  .  .  .  .  .  Decl: *(obj @ 28)
	    36  .  .  .  .  .  .  }
	    37  .  .  .  .  .  }
	    38  .  .  .  .  .  Assign: -
	    39  .  .  .  .  .  Type: *ast.StructType {
	    40  .  .  .  .  .  .  Struct: 7:12
	    41  .  .  .  .  .  .  Fields: *ast.FieldList {
	    42  .  .  .  .  .  .  .  Opening: 7:19
	    43  .  .  .  .  .  .  .  List: []*ast.Field (len = 1) {
	    44  .  .  .  .  .  .  .  .  0: *ast.Field {
	    45  .  .  .  .  .  .  .  .  .  Names: []*ast.Ident (len = 1) {
	    46  .  .  .  .  .  .  .  .  .  .  0: *ast.Ident {
	    47  .  .  .  .  .  .  .  .  .  .  .  NamePos: 8:2
	    48  .  .  .  .  .  .  .  .  .  .  .  Name: "Content"
	    49  .  .  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
	    50  .  .  .  .  .  .  .  .  .  .  .  .  Kind: var
	    51  .  .  .  .  .  .  .  .  .  .  .  .  Name: "Content"
	    52  .  .  .  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 44)
	    53  .  .  .  .  .  .  .  .  .  .  .  }
	    54  .  .  .  .  .  .  .  .  .  .  }
	    55  .  .  .  .  .  .  .  .  .  }
	    56  .  .  .  .  .  .  .  .  .  Type: *ast.Ident {
	    57  .  .  .  .  .  .  .  .  .  .  NamePos: 8:10
	    58  .  .  .  .  .  .  .  .  .  .  Name: "string"
	    59  .  .  .  .  .  .  .  .  .  }
	    60  .  .  .  .  .  .  .  .  .  Tag: *ast.BasicLit {
	    61  .  .  .  .  .  .  .  .  .  .  ValuePos: 8:17
	    62  .  .  .  .  .  .  .  .  .  .  Kind: STRING
	    63  .  .  .  .  .  .  .  .  .  .  Value: "`json:\"content\"\\`"
	    64  .  .  .  .  .  .  .  .  .  }
	    65  .  .  .  .  .  .  .  .  }
	    66  .  .  .  .  .  .  .  }
	    67  .  .  .  .  .  .  .  Closing: 9:1
	    68  .  .  .  .  .  .  }
	    69  .  .  .  .  .  .  Incomplete: false
	    70  .  .  .  .  .  }
	    71  .  .  .  .  }
	    72  .  .  .  }
	    73  .  .  .  Rparen: -
	    74  .  .  }
	    75  .  .  2: *ast.FuncDecl {
	    76  .  .  .  Name: *ast.Ident {
	    77  .  .  .  .  NamePos: 12:6
	    78  .  .  .  .  Name: "main"
	    79  .  .  .  .  Obj: *ast.Object {
	    80  .  .  .  .  .  Kind: func
	    81  .  .  .  .  .  Name: "main"
	    82  .  .  .  .  .  Decl: *(obj @ 75)
	    83  .  .  .  .  }
	    84  .  .  .  }
	    85  .  .  .  Type: *ast.FuncType {
	    86  .  .  .  .  Func: 12:1
	    87  .  .  .  .  Params: *ast.FieldList {
	    88  .  .  .  .  .  Opening: 12:10
	    89  .  .  .  .  .  Closing: 12:11
	    90  .  .  .  .  }
	    91  .  .  .  }
	    92  .  .  .  Body: *ast.BlockStmt {
	    93  .  .  .  .  Lbrace: 12:13
	    94  .  .  .  .  List: []ast.Stmt (len = 1) {
	    95  .  .  .  .  .  0: *ast.AssignStmt {
	    96  .  .  .  .  .  .  Lhs: []ast.Expr (len = 1) {
	    97  .  .  .  .  .  .  .  0: *ast.Ident {
	    98  .  .  .  .  .  .  .  .  NamePos: 13:2
	    99  .  .  .  .  .  .  .  .  Name: "_"
	   100  .  .  .  .  .  .  .  }
	   101  .  .  .  .  .  .  }
	   102  .  .  .  .  .  .  TokPos: 13:4
	   103  .  .  .  .  .  .  Tok: =
	   104  .  .  .  .  .  .  Rhs: []ast.Expr (len = 1) {
	   105  .  .  .  .  .  .  .  0: *ast.CompositeLit {
	   106  .  .  .  .  .  .  .  .  Type: *ast.Ident {
	   107  .  .  .  .  .  .  .  .  .  NamePos: 13:6
	   108  .  .  .  .  .  .  .  .  .  Name: "Hello"
	   109  .  .  .  .  .  .  .  .  .  Obj: *(obj @ 32)
	   110  .  .  .  .  .  .  .  .  }
	   111  .  .  .  .  .  .  .  .  Lbrace: 13:11
	   112  .  .  .  .  .  .  .  .  Rbrace: 13:12
	   113  .  .  .  .  .  .  .  .  Incomplete: false
	   114  .  .  .  .  .  .  .  }
	   115  .  .  .  .  .  .  }
	   116  .  .  .  .  .  }
	   117  .  .  .  .  }
	   118  .  .  .  .  Rbrace: 14:1
	   119  .  .  .  }
	   120  .  .  }
	   121  .  }
	   122  .  Scope: *ast.Scope {
	   123  .  .  Objects: map[string]*ast.Object (len = 2) {
	   124  .  .  .  "Hello": *(obj @ 32)
	   125  .  .  .  "main": *(obj @ 79)
	   126  .  .  }
	   127  .  }
	   128  .  Imports: []*ast.ImportSpec (len = 1) {
	   129  .  .  0: *(obj @ 12)
	   130  .  }
	   131  .  Unresolved: []*ast.Ident (len = 1) {
	   132  .  .  0: *(obj @ 56)
	   133  .  }
	   134  }
	*/
}
