cd ..
go run . left standard --align=right
go run . right standard --align=left
go run . hello shadow --align=center
go run . "1 Two 4" shadow --align=justify
go run . 23/32 standard --align=right
go run . ABCabc123 thinkertoy --align=right
go run . "#$%&" thinkertoy --align=center
go run . "23Hello World!" standard --align=left
go run . "HELLO there HOW are YOU?!" thinkertoy --align=justify
go run . "a -> A b -> B c -> C" shadow --align=right
go run . abcd shadow --align=right
go run . ola standard --align=center