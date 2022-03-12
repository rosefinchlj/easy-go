package gen

// need install golang.org/x/tools/cmd/stringer
//go:generate stringer -type Code -linecomment  -output code_msg_gen.go

type Code int

func (i Code) Message() string {
    return i.String()
}

func (i Code) Code() int {
    return int(i)
}