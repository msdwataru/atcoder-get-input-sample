# atcoder-get-input-sample
Atcoderの入力例を表示するコマンドラインツール

## Dependency
goquery
```go get https://github.com/msdwataru/atcoder-get-input-sample.git```

## Usage
```go run main.go -c <CONTEST_NAME> -q <WHICH_QUESTION> -i <INDEX_OF_INPUT_SAMPLE>```
for example if you'd like to get first input sample from ABC123, question A,
```go run main.go -c abc123 -q a -i 1```
