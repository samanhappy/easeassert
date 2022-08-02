# EaseEval
EaseEval is a simple text eval tool, built on DSL, written in Go.
## 1. Overview
EaseEval is designed to extract values from text and evaluate expressions using a DSL, which contains Function, Typecast and Compute.

### 1.1 Function
Function is very useful in many scenarios, we plan to supply different kind of functions.
### 1.1.1 Extract Function
 Extract functions extract target values from source text in many types: plain text, json, html, xml, etc.
 - **jq**. extract values from json format text using [jq](https://github.com/savaki/jq)
 - **regex**. extract values from text using regex (to be implemented)
### 1.1.2 Other Function (to be implemented)

### 1.2 Typecast
Typecast auto converts values on demand to specific types like string, number, date, etc. We support types below.
- **bool**. bool type can be used in assertions
- **int64**. int64 is the only int type supported, other int type like int8,int32 and int strings will auto be converted to int64.
- **datetime**. (to be implemented)
- **other types**. (to be implemented)
### 1.3 Compute
Compute does the actual eval work for different operators. 
- **logical operators**. &&, ||, !
- **arithmetic operators**. +, -, *, /
- **relational operators**. >, <, ==, !=, >=, <=
- **other operators**. (to be implemented)

## 2 Usage

### 2.1 Installation
```
go get github.com/samanhappy/easeeval
```

### 2.2 Example
```
package main

import (
	"fmt"

	"github.com/samanhappy/easeeval"
)

func main() {
	// pure eval
	v, _ := easeeval.Eval(`1 + 2 * 3`, "")
	fmt.Printf("value %v, type %T\n", v, v)

	// use jq to extract value for eval
	v, _ = easeeval.Eval(`1 + jq(".key") * 3`, `{"key":"2"}`)
	fmt.Printf("value %v, type %T\n", v, v)
}
```