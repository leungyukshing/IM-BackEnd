# IM Back End Database Design

This document shows the blueprint of back-end database design of IM.

## Intro

I store most information in backend database out of simplicity. In the basic version, I don't want to care much about data synchronization or multi-device synchronization. My goal is to make it work properly, care less about the transfer time. So I only store information in a remote server, wihout backup.

---

## Techs

In development phrase, I'd like to use a local server to act as a remote server, on which I install a MySQL database. When this project is first released, I may deploy the database on remote server like Tencent Cloud. [How to deploy MySQL on Tencent Cloud Server](https://blog.csdn.net/runner1920/article/details/79495368)

---

## Content

I store most of the application's data on backend database. There are tables, including **user**, **friend**, **chat**.

### user

This table is used to store user information

| Name       | Type    | Key         | Note                      |
| ---------- | ------- | ----------- | ------------------------- |
| id         | int     | primary key | user id(auto increment)   |
| username   | varchar |             | username                  |
| password   | varchar |             | password                  |
| avatar     | varchar |             | user's avatar             |
| email      | varchar |             | user's email              |
| time       | varchar |             | when this user is created |
| encryptkey | varchar |             | key used to encrypt info  |

### friend

This table is used to store friends relationship among users. We use bi-relation, that is, only one record of a friend ship is stored in the database.[Store Friends Relationship](https://blog.csdn.net/cienit/article/details/45158149)

| Name     | Type    | Key         | Note                                  |
| -------- | ------- | ----------- | ------------------------------------- |
| id       | int     | primary key | auto increment                        |
| user_id1 | int     |             |                                       |
| user_id2 | int     |             |                                       |
| time     | varchar |             | when this relationship is established |

### chat

This table is used to store chat list.

| Name             | Type    | Key         | Note                                                         |
| ---------------- | ------- | ----------- | ------------------------------------------------------------ |
| id               | int     | primary key | auto increment, used to locate all records belong to this chat |
| create_time      | varchar |             | When this chat is created                                    |
| last_update_time | varchar |             | when this chat is last updated(can be used to sort chat list) |

---

## Reference

1. [How to store picture in database](https://blog.csdn.net/Cs_hnu_scw/article/details/74011674)