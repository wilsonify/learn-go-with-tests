module test-S01-fundamentals

go 1.18

require (
	github.com/approvals/go-approval-tests v0.0.0-20220429035603-e60d9ca44d3e
	github.com/wilsonify/learn-go-with-tests/S01-fundamentals v0.0.0
)

require github.com/gomarkdown/markdown v0.0.0-20220510115730-2372b9aa33e5 // indirect

replace github.com/wilsonify/learn-go-with-tests/S01-fundamentals => ../../pkg/S01-fundamentals
