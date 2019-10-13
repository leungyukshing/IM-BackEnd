# IM BackEnd Data Transporation Design Doc

**Date:** 2019.10.13

## Goal

This document try to clarify which data type we use to communicate with front end. That is, what data type we handle in backend.

## ProtoBuf

ProtoBuf, PB in short, is a data type used to transport data between front end and back end. It's more convenient to use PB than json.

We only need to edit `.proto` file. And use the script to generate `.go` automatically. Those generated files can be used in backend code.

## Reference

1. [PB Usage in Go](https://www.jianshu.com/p/c1723e5f6a46)