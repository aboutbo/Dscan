package lib

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/interpreter/functions"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

// CustomLib 实现cel.Library interface
type CustomLib struct {
	envOptions     []cel.EnvOption
	programOptions []cel.ProgramOption
}

// CompileOptions cel.Library接口函数实现
func (c *CustomLib) CompileOptions() []cel.EnvOption {
	return c.envOptions
}

// ProgramOptions cel.Library接口函数实现
func (c *CustomLib) ProgramOptions() []cel.ProgramOption {
	return c.programOptions
}

// NewEnv 生成cel环境变量
func NewEnv(c *CustomLib) (*cel.Env, error) {
	return cel.NewEnv(cel.Lib(c))
}

// NewEnvOption 声明cel变量声明
func NewEnvOption() CustomLib {
	c := CustomLib{}
	c.envOptions = []cel.EnvOption{
		cel.Container("lib"),
		cel.Types(&Response{}),
		cel.Declarations(decls.NewIdent("response", decls.NewObjectType("lib.Response"), nil)),
		cel.Declarations(decls.NewFunction("bcontains", decls.NewInstanceOverload("bytes_bcontains_bytes", []*exprpb.Type{decls.Bytes, decls.Bytes}, decls.Bool))),
		cel.Declarations(decls.NewFunction("contains_string", decls.NewInstanceOverload("strings_contails_strings", []*exprpb.Type{decls.String, decls.String}, decls.Bool))),
	}
	c.programOptions = []cel.ProgramOption{
		cel.Functions(
			&functions.Overload{
				Operator: "bytes_bcontains_bytes",
				Binary:   bytesContainsBytes,
			},
			&functions.Overload{
				Operator: "strings_contails_strings",
				Binary:   stringsContainsStrings,
			},
		),
	}
	return c
}

func stringsContainsStrings(lhs ref.Val, rhs ref.Val) ref.Val {
	v1, ok := lhs.(types.String)
	if !ok {
		return types.ValOrErr(lhs, "unexpected type '%v' passed to bcontains", lhs.Type())
	}
	v2, ok := rhs.(types.String)
	if !ok {
		return types.ValOrErr(rhs, "unexpected type '%v' passed to bcontains", rhs.Type())
	}
	return types.Bool(strings.Contains(string(v1), string(v2)))
}

func bytesContainsBytes(lhs ref.Val, rhs ref.Val) ref.Val {
	v1, ok := lhs.(types.Bytes)
	if !ok {
		return types.ValOrErr(lhs, "unexpected type '%v' passed to bcontains", lhs.Type())
	}
	v2, ok := rhs.(types.Bytes)
	if !ok {
		return types.ValOrErr(rhs, "unexpected type '%v' passed to bcontains", rhs.Type())
	}
	return types.Bool(bytes.Contains(v1, v2))
}

// Evaluate 执行表达式
func Evaluate(env *cel.Env, expression string, params map[string]interface{}) (ref.Val, error) {
	ast, iss := env.Compile(expression)
	if iss.Err() != nil {
		fmt.Errorf("compile: ", iss.Err())
		return nil, iss.Err()
	}

	prg, err := env.Program(ast)
	if err != nil {
		fmt.Errorf("Program creation error: %v", err)
		return nil, err
	}

	out, _, err := prg.Eval(params)
	if err != nil {
		fmt.Errorf("Evaluation error: %v", err)
		return nil, err
	}
	return out, nil
}
