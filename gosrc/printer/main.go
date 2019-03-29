package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
)

var GFset *token.FileSet

// go/printer包提供把AST树重新生成为go文件功能
//
// 我们尝试把test.go文件中的结构体
// ```
// 	 type Hello struct {
//		  	Content string `json:"content"`
//	 }
//  ```
// 名称从Hello改为World,字段Content增加一个标签xml:"content"
// 然后保存到var/test2.go
func main() {
	goFile, err := ioutil.ReadFile("./var/test.go")
	if err != nil {
		panic(err)
	}

	// 创建AST树
	fset := token.NewFileSet() // positions are relative to fset
	GFset = fset
	f, err := parser.ParseFile(fset, "", goFile, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	fmt.Println("========== Before modify ============ ")
	//ast.Print(fset, f) // 打印修改前的

	ast.Inspect(f, WalkModify) // 遍历ast树并且找到我们需要修改的地方

	fmt.Println("========== After modify ============ ")
	//ast.Print(fset, f) // 打印修改后的AST树

	fGo, err := os.OpenFile("./var/test2.go", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = printer.Fprint(fGo, fset, f)
	if err != nil {
		panic(err)
	}

	//修改前的AST树 结果
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
	    63  .  .  .  .  .  .  .  .  .  .  Value: "`json:\"content\"\`"
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

// 设置处理函数
func WalkModify(n ast.Node) bool {
	s, ok := n.(*ast.TypeSpec)
	if !ok {
		return true
	}

	fmt.Printf("[DEBUG] CurName:%s \n", s.Name.Name)

	// 判断结构体名称是否为Hello
	if s.Name.Name != "Hello" {
		return true
	}

	fmt.Printf("[DEBUG] Find struct 'Hello' pos:%s \n", GFset.Position(s.Pos()))

	// 修改结构体名称
	s.Name.Name = "World"

	if s.Type == nil {
		fmt.Printf("[Warning] Type类型缺失,位置:%s\n", GFset.Position(s.Pos()))
		return true
	}

	t, ok := (s.Type).(*ast.StructType)
	if !ok {
		return true
	}

	// 因为 Hello结构体肯定是有属性的,如果fields数组都为空的那肯定不是,我们要找的结构体,忽略即可
	if t.Fields.NumFields() == 0 {
		return false
	}

	// 遍历所有属性
	for i, _ := range t.Fields.List {
		// 寻找字段名称为Content的属性
		if len(t.Fields.List[i].Names) == 0 {
			continue
		}

		for j, _ := range t.Fields.List[i].Names {
			if t.Fields.List[i].Names[j].Name != "Content" {
				continue
			}

			fmt.Printf("[DEBUG] Find field 'Content' pos:%s \n", GFset.Position(t.Fields.List[i].Pos()))

			// 修改标签值
			t.Fields.List[i].Tag.Value = "`json:\"content\" xml:\"content\"`"
		}

	}
	return true
}
