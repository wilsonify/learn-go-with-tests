module test.learn.go

go 1.18

require (
	github.com/approvals/go-approval-tests v0.0.0-20220530063708-32d5677069bd
	github.com/gorilla/websocket v1.5.0
	learn.go v0.0.0
)

require github.com/gomarkdown/markdown v0.0.0-20220603122033-8f3b341fef32 // indirect

replace learn.go => ../pkg
