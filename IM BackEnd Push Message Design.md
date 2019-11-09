# IM BackEnd Push Message Design Doc

**Date:** 2019.11.09

## Goal

We use **websocket** to push new messages to users.

## Intro

1. we store user connection in server. Using a map.
2. we upgrade user connection from http to websocket, which can maintain longer time and allow the server to push messages to a specific client.

## Reference

1. [Go websocket Example]([https://99mycql.github.io/application/%E7%94%A8beego%E5%AE%9E%E7%8E%B0%E8%81%8A%E5%A4%A9%E5%AE%A4-WebSocket%E6%96%B9%E5%BC%8F.html](https://99mycql.github.io/application/用beego实现聊天室-WebSocket方式.html))