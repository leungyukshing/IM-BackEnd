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

## Reference

1. [Travis for CI in Github](https://maiyang.me/post/2017-06-17-github-travis-ci-coveralls/)
2. [Travis Docs](https://docs.travis-ci.com/user/languages/go/)
3. [go vet](https://golang.org/cmd/vet/)
4. [go vet(CN)](https://studygolang.com/articles/9619)
5. [go cover](https://golang.org/cmd/cover/)
6. [go test framework](https://golang.org/pkg/testing/)
7. [travis Ci for Github](https://lingxiankong.github.io/2018-06-28-travis-ci-integration.html)
8. [watching git commit message](https://www.npmjs.com/package/git-commit-msg-linter)
9. [commit message in travis](https://segmentfault.com/a/1190000021007124?utm_source=tag-newest)
