# EaseEval
[![GitHub license](https://img.shields.io/github/license/samanhappy/easeeval)](https://github.com/samanhappy/easeeval/blob/main/LICENSE)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/samanhappy/easeeval)](https://github.com/samanhappy/easeeval/blob/main/go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/samanhappy/easeeval)](https://goreportcard.com/report/github.com/samanhappy/easeeval)
[![Build](https://github.com/samanhappy/easeeval/actions/workflows/test.yaml/badge.svg)](https://github.com/samanhappy/easeeval/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/samanhappy/easeeval/branch/main/graph/badge.svg)](https://codecov.io/gh/samanhappy/easeeval)

EaseEval is a simple text eval tool, built on DSL, written in Go.

## Quick Start
### Install
```
go get github.com/samanhappy/easeeval
```
### Example
```
// use jq to extract value from json, v should be 7(int64)
v, err = easeeval.Eval(`1 + jq(".key") * 3`, `{"key":"2"}`)

// use time functions, v should be true(bool)
v, err = easeeval.Eval(`unixTime(jq(".time")) < now()`, `{"time":"2022-08-01T12:00:00"}`)
```

## Overview
EaseEval is designed to extract values from text and evaluate expressions using a DSL, which contains Function, Typecast and Compute.

### Function
Function is very useful in many scenarios, we plan to supply different kind of functions.
#### Extract Function
Extract functions extract target values from source text in many types: plain text, json, html, xml, etc.
|  Name   |  Description | Usage
|  ----  | ----  | ----  |
| jq | extract values from json format text using [https://github.com/savaki/jq](https://github.com/savaki/jq) | jq(".key") |
| regex | extract values from text using regex | (to be implemented) | 
#### Time Function
Time functions supply ways to get unix time from string or for now.
|  Name   |  Description | Usage
|  ----  | ----  | ----  |
| unixTime | return unix time seconds in int64 for any format string using [https://github.com/araddon/dateparse](https://github.com/araddon/dateparse) | unixTime("2022-08-01T12:00:00") |
| now | return unix time seconds in int64 for now | now() | 
### Typecast
Typecast auto converts values on demand to specific types like string, number, date, etc. We support types below.
|  Name   |  Description |
|  ----  | ----  |
| bool | bool type can be used in assertions |
| int64 | int64 is the only int type supported, other int type like int8,int32 and int strings will auto be converted to int64 |
| Time | Time type will be converted to int64 for compute |
### Compute
Compute does the actual eval work for different kind of operators. 
|  Kind   |  Operators |
|  ----  | ----  |
| logical | && \|\| ! |
| arithmetic | + - * / |
| relational | > < == != >= <= |
