go run . "Hello There Test" "standard" --output=test.txt
go run . --reverse=test.txt
go run . "One two free four five word something" "standard" --align=justify
go run . "One two free four five" "shadow" --align=right
go run . "Test" "thinkertoy" --align=center
go run . "Test" "standard" --color=red
go run . "Test" "standard" --color=yellow [0]
go run . "Test Hello" "standard" --color=blue [2]
go run . "Test Hello" "standard" --color=brown [5-9]
go run . "Test Hello" "standard" --color=brown [5-9]
go run . "Hello world"