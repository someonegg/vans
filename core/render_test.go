package core

import (
	"bytes"
	"testing"
	"text/template"
)

func Test_templateFuncMap(t *testing.T) {
	tt,_:=template.New("test").Option("missingkey=error").Funcs(templateFuncMap).Parse("{{ \"hello!\" | ToUpper | repeat 5 }}")
	wr:=&bytes.Buffer{}
	err:=tt.Execute(wr,map[string]interface{}{})
	if err!=nil{
		t.Error(err)
	}
	if wr.String()!="HELLO!HELLO!HELLO!HELLO!HELLO!"{
		t.Error(wr.String())
	}
}

func Test_templateFuncMap2(t *testing.T) {
	tt,_:=template.New("test").Option("missingkey=error").Funcs(templateFuncMap).Parse("{{- int 3.1415926 -}}")
	wr:=&bytes.Buffer{}
	err:=tt.Execute(wr,map[string]interface{}{})
	if err!=nil{
		t.Error(err)
	}
	if wr.String()!="3"{
		t.Error(wr.String())
	}
}
