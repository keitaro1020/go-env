package env_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/keitaro1020/go-env"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	type singleStringField struct {
		String string
	}

	type singleIntField struct {
		Int int
	}
	type singleInt8Field struct {
		Int8 int8
	}
	type singleInt16Field struct {
		Int16 int16
	}
	type singleInt32Field struct {
		Int32 int32
	}
	type singleInt64Field struct {
		Int64 int64
	}

	type singleUintField struct {
		Uint uint
	}
	type singleUint8Field struct {
		Uint8 uint8
	}
	type singleUint16Field struct {
		Uint16 uint16
	}
	type singleUint32Field struct {
		Uint32 uint32
	}
	type singleUint64Field struct {
		Uint64 uint64
	}

	type singleFloat32Field struct {
		Float32 float32
	}
	type singleFloat64Field struct {
		Float64 float64
	}

	type singleBoolField struct {
		Bool bool
	}

	type notSupportType struct {
		NotSupportType interface{}
	}

	type multipleField struct {
		String string
		Int    int
	}

	type nestStructLevel3 struct {
		Level3String string
	}
	type nestStructLevel2 struct {
		Level2String string
		Level2       *nestStructLevel3
	}
	type nestPtrStructLevel1 struct {
		Level1 *nestStructLevel2
	}
	type nestStructLevel1 struct {
		Level1 nestStructLevel2
	}

	type envKey struct {
		String string `env_key:"ENV_KEY"`
	}

	i := 10
	type args struct {
		v interface{}
	}
	tests := []struct {
		name        string
		args        args
		preparation func() error
		want        interface{}
		wantErr     bool
	}{
		{
			name:    "error:not_pointer_struct",
			args:    args{v: singleStringField{}},
			want:    singleStringField{},
			wantErr: true,
		},
		{
			name:    "error:primitive_type",
			args:    args{v: i},
			want:    i,
			wantErr: true,
		},
		{
			name:    "error:primitive_type_pointer",
			args:    args{v: &i},
			want:    &i,
			wantErr: true,
		},
		{
			name: "error:not_support_type",
			args: args{v: &notSupportType{}},
			preparation: func() error {
				return os.Setenv("NOT_SUPPORT_TYPE", "aaa")
			},
			want:    &notSupportType{},
			wantErr: true,
		},
		{
			name: "error:convert_fail",
			args: args{v: &singleIntField{}},
			preparation: func() error {
				return os.Setenv("INT", "aaa")
			},
			want:    &singleIntField{},
			wantErr: true,
		},
		{
			name: "ok:string_value_not_exist",
			args: args{v: &singleStringField{}},
			preparation: func() error {
				return os.Unsetenv("STRING")
			},
			want:    &singleStringField{},
			wantErr: false,
		},
		{
			name: "ok:string_value_exist",
			args: args{v: &singleStringField{String: ""}},
			preparation: func() error {
				return os.Setenv("STRING", "hoge")
			},
			want:    &singleStringField{String: "hoge"},
			wantErr: false,
		},
		{
			name: "ok:int_value_not_exist",
			args: args{v: &singleIntField{}},
			preparation: func() error {
				return os.Unsetenv("INT")
			},
			want:    &singleIntField{},
			wantErr: false,
		},
		{
			name: "ok:int_value_exist",
			args: args{v: &singleIntField{}},
			preparation: func() error {
				return os.Setenv("INT", "2147483647")
			},
			want:    &singleIntField{Int: 2147483647},
			wantErr: false,
		},
		{
			name: "ok:int8_value_not_exist",
			args: args{v: &singleInt8Field{}},
			preparation: func() error {
				return os.Unsetenv("INT8")
			},
			want:    &singleInt8Field{},
			wantErr: false,
		},
		{
			name: "ok:int8_value_exist",
			args: args{v: &singleInt8Field{}},
			preparation: func() error {
				return os.Setenv("INT8", "127")
			},
			want:    &singleInt8Field{Int8: 127},
			wantErr: false,
		},
		{
			name: "ok:int16_value_not_exist",
			args: args{v: &singleInt16Field{}},
			preparation: func() error {
				return os.Unsetenv("INT16")
			},
			want:    &singleInt16Field{},
			wantErr: false,
		},
		{
			name: "ok:int16_value_exist",
			args: args{v: &singleInt16Field{}},
			preparation: func() error {
				return os.Setenv("INT16", "32767")
			},
			want:    &singleInt16Field{Int16: 32767},
			wantErr: false,
		},
		{
			name: "ok:int32_value_not_exist",
			args: args{v: &singleInt32Field{}},
			preparation: func() error {
				return os.Unsetenv("INT32")
			},
			want:    &singleInt32Field{},
			wantErr: false,
		},
		{
			name: "ok:int32_value_exist",
			args: args{v: &singleInt32Field{}},
			preparation: func() error {
				return os.Setenv("INT32", "2147483647")
			},
			want:    &singleInt32Field{Int32: 2147483647},
			wantErr: false,
		},
		{
			name: "ok:int64_value_not_exist",
			args: args{v: &singleInt64Field{}},
			preparation: func() error {
				return os.Unsetenv("INT64")
			},
			want:    &singleInt64Field{},
			wantErr: false,
		},
		{
			name: "ok:int64_value_exist",
			args: args{v: &singleInt64Field{}},
			preparation: func() error {
				return os.Setenv("INT64", "9223372036854775807")
			},
			want:    &singleInt64Field{Int64: 9223372036854775807},
			wantErr: false,
		},
		{
			name: "ok:uint_value_not_exist",
			args: args{v: &singleUintField{}},
			preparation: func() error {
				return os.Unsetenv("UINT")
			},
			want:    &singleUintField{},
			wantErr: false,
		},
		{
			name: "ok:uint_value_exist",
			args: args{v: &singleUintField{}},
			preparation: func() error {
				return os.Setenv("UINT", "4294967295")
			},
			want:    &singleUintField{Uint: 4294967295},
			wantErr: false,
		},
		{
			name: "ok:uint8_value_not_exist",
			args: args{v: &singleUint8Field{}},
			preparation: func() error {
				return os.Unsetenv("UINT8")
			},
			want:    &singleUint8Field{},
			wantErr: false,
		},
		{
			name: "ok:uint8_value_exist",
			args: args{v: &singleUint8Field{}},
			preparation: func() error {
				return os.Setenv("UINT8", "255")
			},
			want:    &singleUint8Field{Uint8: 255},
			wantErr: false,
		},
		{
			name: "ok:uint16_value_not_exist",
			args: args{v: &singleUint16Field{}},
			preparation: func() error {
				return os.Unsetenv("UINT16")
			},
			want:    &singleUint16Field{},
			wantErr: false,
		},
		{
			name: "ok:uint16_value_exist",
			args: args{v: &singleUint16Field{}},
			preparation: func() error {
				return os.Setenv("UINT16", "65535")
			},
			want:    &singleUint16Field{Uint16: 65535},
			wantErr: false,
		},
		{
			name: "ok:uint32_value_not_exist",
			args: args{v: &singleUint32Field{}},
			preparation: func() error {
				return os.Unsetenv("UINT32")
			},
			want:    &singleUint32Field{},
			wantErr: false,
		},
		{
			name: "ok:uint32_value_exist",
			args: args{v: &singleUint32Field{}},
			preparation: func() error {
				return os.Setenv("UINT32", "4294967295")
			},
			want:    &singleUint32Field{Uint32: 4294967295},
			wantErr: false,
		},
		{
			name: "ok:uint64_value_not_exist",
			args: args{v: &singleUint64Field{}},
			preparation: func() error {
				return os.Unsetenv("UINT64")
			},
			want:    &singleUint64Field{},
			wantErr: false,
		},
		{
			name: "ok:uint64_value_exist",
			args: args{v: &singleUint64Field{}},
			preparation: func() error {
				return os.Setenv("UINT64", "18446744073709551615")
			},
			want:    &singleUint64Field{Uint64: 18446744073709551615},
			wantErr: false,
		},
		{
			name: "ok:float32_value_not_exist",
			args: args{v: &singleFloat32Field{}},
			preparation: func() error {
				return os.Unsetenv("FLOAT32")
			},
			want:    &singleFloat32Field{},
			wantErr: false,
		},
		{
			name: "ok:float32_value_exist",
			args: args{v: &singleFloat32Field{}},
			preparation: func() error {
				return os.Setenv("FLOAT32", "42949.67295")
			},
			want:    &singleFloat32Field{Float32: 42949.67295},
			wantErr: false,
		},
		{
			name: "ok:float64_value_not_exist",
			args: args{v: &singleFloat64Field{}},
			preparation: func() error {
				return os.Unsetenv("FLOAT64")
			},
			want:    &singleFloat64Field{},
			wantErr: false,
		},
		{
			name: "ok:float64_value_exist",
			args: args{v: &singleFloat64Field{}},
			preparation: func() error {
				return os.Setenv("FLOAT64", "1844674407370.9551615")
			},
			want:    &singleFloat64Field{Float64: 1844674407370.9551615},
			wantErr: false,
		},
		{
			name: "ok:multiple_field",
			args: args{v: &multipleField{}},
			preparation: func() error {
				if err := os.Setenv("STRING", "Stringggg"); err != nil {
					return err
				}
				return os.Setenv("INT", "123456")
			},
			want: &multipleField{
				String: "Stringggg",
				Int:    123456,
			},
			wantErr: false,
		},
		{
			name: "ok:nest_ptr_struct",
			args: args{v: &nestPtrStructLevel1{Level1: &nestStructLevel2{}}},
			preparation: func() error {
				return os.Setenv("LEVEL1_LEVEL2_STRING", "vaaaaaalue")
			},
			want: &nestPtrStructLevel1{
				Level1: &nestStructLevel2{Level2String: "vaaaaaalue", Level2: &nestStructLevel3{}},
			},
			wantErr: false,
		},
		{
			name: "ok:nest_ptr_struct_nil_field",
			args: args{v: &nestPtrStructLevel1{Level1: nil}},
			preparation: func() error {
				return os.Setenv("LEVEL1_LEVEL2_STRING", "vaaaaaalue")
			},
			want: &nestPtrStructLevel1{
				Level1: &nestStructLevel2{Level2String: "vaaaaaalue", Level2: &nestStructLevel3{}},
			},
			wantErr: false,
		},
		{
			name: "ok:nest_struct",
			args: args{v: &nestStructLevel1{Level1: nestStructLevel2{}}},
			preparation: func() error {
				return os.Setenv("LEVEL1_LEVEL2_STRING", "vaaaaaalue")
			},
			want: &nestStructLevel1{
				Level1: nestStructLevel2{Level2String: "vaaaaaalue", Level2: &nestStructLevel3{}},
			},
			wantErr: false,
		},
		{
			name: "ok:nest_struct2",
			args: args{v: &nestStructLevel1{Level1: nestStructLevel2{}}},
			preparation: func() error {
				os.Unsetenv("LEVEL1_LEVEL2_STRING")
				return os.Setenv("LEVEL1_LEVEL2_LEVEL3_STRING", "vaaaaaalue")
			},
			want: &nestStructLevel1{
				Level1: nestStructLevel2{Level2: &nestStructLevel3{Level3String: "vaaaaaalue"}},
			},
			wantErr: false,
		},
		{
			name: "ok:env_key",
			args: args{v: &envKey{}},
			preparation: func() error {
				return os.Setenv("ENV_KEY", "vaaaaaalue")
			},
			want:    &envKey{String: "vaaaaaalue"},
			wantErr: false,
		},
		{
			name: "ok:bool_value",
			args: args{v: &singleBoolField{}},
			preparation: func() error {
				return os.Setenv("BOOL", "true")
			},
			want:    &singleBoolField{Bool: true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.preparation != nil {
				if err := tt.preparation(); err != nil {
					t.Errorf("preparation() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			if err := env.Parse(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.args.v, tt.want); diff != "" {
				t.Errorf("Parse() diff = %v", diff)
			}
		})
	}
}

func BenchmarkParse1(b *testing.B) {
	type bench struct {
		String string
		Int    int
	}

	os.Setenv("STRING", "Stringggg")
	os.Setenv("INT", "123456")
	for i := 0; i < b.N; i++ {
		env.Parse(&bench{})
	}
}
func BenchmarkParse2(b *testing.B) {
	type bench struct {
		Level1String string
		Level1       struct {
			Level2String string
			Level2       struct {
				Level3String string
			}
		}
	}

	os.Setenv("LEVEL1_STRING", "Stringggg")
	os.Setenv("LEVEL1_LEVEL2_STRING", "Stringggg")
	os.Setenv("LEVEL1_LEVEL2_LEVEL3_STRING", "Stringggg")
	for i := 0; i < b.N; i++ {
		env.Parse(&bench{})
	}
}

func Test_toSnakeCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		input string
		want  string
	}{
		{input: "A", want: "A"},
		{input: "a", want: "A"},
		{input: "AB", want: "A_B"},
		{input: "aB", want: "A_B"},
		{input: "Ab", want: "AB"},
		{input: "AbC", want: "AB_C"},
		{input: "AbCdE", want: "AB_CD_E"},
		{input: "Ab1Cde", want: "AB1_CDE"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v->%v", tt.input, tt.want), func(t *testing.T) {
			if got := env.ToSnakeCase(tt.input); got != tt.want {
				t.Errorf("ToSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
