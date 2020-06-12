package main

import (
	"log"
	"os"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
)

func main() {
	e := os.Args[1]
	log.Printf("evaluating %q", e)

	env, err := cel.NewEnv(cel.Declarations())
	if err != nil {
		log.Fatalln(err)
	}
	ast, iss := env.Compile(e)
	if iss.Err() != nil {
		// Syntax error.
		log.Fatalln(iss.Err())
	}

	prg, err := env.Program(ast)
	out, _, err := prg.Eval(map[string]interface{}{})
	if err != nil {
		// Error evaluating.
		log.Fatalln(err)
	}

	if !types.IsBool(out) {
		log.Fatalf("result (%v) is not boolean (type: %s)", out.Value(), out.Type().TypeName())
	}
	if out != types.Bool(true) {
		log.Fatal("result is false")
	}
	log.Println("result is true")
}
