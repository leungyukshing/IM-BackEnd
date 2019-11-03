# IM Back End Gorm Issues

This doc lists some issues when using gorm in Go.

## Capitals

when we use gorm, we need to pay attention to the captical spelling in codes. gorm is case-sensitive. Carmel spelling may be transformed. e.g, `UserName` may be transformed to `user_name` in gorm.

## TimeOut

we need to specify connection timeout when we are using connection pool. A long timeout may increase DB stress, while a short timeout may cause may bad connection, resulting in many handling errors.

## Find in SQL

we use `var u []User` in `Find(&u)`, instead of `var u User`. Because the returning result could be an array, containing all results that matches the query condition.