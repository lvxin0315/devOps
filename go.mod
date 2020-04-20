module github.com/lvxin0315/devOps

go 1.14

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200414173820-0848c9571904
	golang.org/x/net => github.com/golang/net v0.0.0-20200324143707-d3edc9973b7e
)

require (
	github.com/bndr/gojenkins v1.0.1
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect

)
