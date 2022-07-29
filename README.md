# EaseEval
EaseEval is a DSL tool that can do various eval jobs, written in Go.
## 1. Overview
EaseEval is designed to do evals using DSL which consists of Functions, Converts and Operates.

### 1.1 Function
Function is very useful in many scenarios, we plan to supply different kind of functions.
### 1.1.1 Extract Function
 Extract functions extract the target data from source text of many formats like text, json, html, xml, and so on.
 - **jq**. extract data from json format text using [jq](https://github.com/savaki/jq)
 - **regex**. extract data from text using regex (to be implemented)
### 1.1.2 Other Function (to be implemented)

### 1.2 Convert
Converts auto convert the extracted data to specific types like string, number, date and so on on demand. We support types below.
- **bool**. bool type can be used in assertions
- **int64**. int64 is the only int type supported, other int type like int8,int32 and int strings will auto be converted to int64.
- **datetime**. (to be implemented)
- **other types**. (to be implemented)
### 1.3 Operate
Operates do a various operates for different operators. 
- **logic operators**. &&, ||, !
- **basic calculate operators**. +, -, *, /
- **compare operators**. >, <, ==, !=, >=, <=
- **other operators**. (to be implemented)