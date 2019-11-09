# IM Back End CI Design

This documents describes the CI design of IM backend.

## What is CI?

CI, continuous integration, is a development practice where developers integrate code into a shared repository frequently. Each integration can then be verified by **an automated build and automated tests**.

---

## Our Goal

I'd like to cover the following things in CI:

+ go test: to make sure no stupid bug in unit test.
+ go cover: I want a high coverage in our code. e.g, 80%?
+ go vet: Vet examines source code and reports suspicious contructs. It can find errors not caught by the compilers.
+ go fmt: code format check.
+ check commit message

---

## CI Platform

1. In first stage, we use **Travis**, which provides a free CI platform.
2. In second stage, we may use CI/CD provided by Github, more features. But not stable and few documents can be referenced.
3. If we have enough resources, we can build up a jenkins in a seperated machine.

## Process

1. build a docker, containing all dipendencies, including mysql, redis....
2. download code in the docker(How???)
3. run a .sh file that we have written to test our commit code.

---

## Reference

1. [Travis for CI in Github](https://maiyang.me/post/2017-06-17-github-travis-ci-coveralls/)
2. [Travis Docs](https://docs.travis-ci.com/user/languages/go/)
3. [go vet](https://golang.org/cmd/vet/)
4. [go vet(CN)](https://studygolang.com/articles/9619)
5. [go cover](https://golang.org/cmd/cover/)
6. [go test framework](https://golang.org/pkg/testing/)
7. [travis Ci for Github](https://lingxiankong.github.io/2018-06-28-travis-ci-integration.html)
8. [watching git commit message](https://www.npmjs.com/package/git-commit-msg-linter)
9. [Github Actions Guide](http://www.ruanyifeng.com/blog/2019/09/getting-started-with-github-actions.html)
10. [docker-compose usage](https://www.jianshu.com/p/46db38b94200)
