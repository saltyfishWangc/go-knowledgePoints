package design_pattern

import (
	"fmt"
	"testing"
)

type MyFuncOptions struct {
	optionStr1 string
	optionStr2 string
	optionInt1 int
	optionInt2 int
}

var defaultMyFuncOptions = MyFuncOptions{
	optionStr1: "defaultStr1",
	optionStr2: "defaultStr2",
	optionInt1: 1,
	optionInt2: 2,
}

type MyFuncOption func(options *MyFuncOptions)

func WithOptionStr1(str1 string) MyFuncOption {
	return func(options *MyFuncOptions) {
		options.optionStr1 = str1
	}
}

func WithOptionInt1(int1 int) MyFuncOption {
	return func(options *MyFuncOptions) {
		options.optionInt1 = int1
	}
}

func WithOptionStr2AndInt2(str2 string, int2 int) MyFuncOption {
	return func(options *MyFuncOptions) {
		options.optionStr2 = str2
		options.optionInt2 = int2
	}
}

func MyFunc2(requiredStr string, opts ...MyFuncOption) {
	options := defaultMyFuncOptions
	for _, o := range opts {
		o(&options)
	}

	fmt.Println(requiredStr, options.optionStr1, options.optionStr2, options.optionInt1, options.optionInt2)
}

// TestOptionPattern 选项模式
func TestOptionPattern(t *testing.T) {
	MyFunc2("requiredStr")
	MyFunc2("requiredStr", WithOptionStr1("mystr1"))
	MyFunc2("requiredStr", WithOptionStr2AndInt2("mystr2", 22), WithOptionInt1(11))
}

var defaultStuffClientOptions = StuffClientOptions{
	Retries: 3,
	Timeout: 2,
}

type StuffClientOption func(*StuffClientOptions)

type StuffClientOptions struct {
	Retries int //number of times to retry the request before giving up
	Timeout int //connection timeout in seconds
}

func WithRetries(r int) StuffClientOption {
	return func(o *StuffClientOptions) {
		o.Retries = r
	}
}
func WithTimeout(t int) StuffClientOption {
	return func(o *StuffClientOptions) {
		o.Timeout = t
	}
}

type StuffClient interface {
	DoStuff() error
}
type stuffClient struct {
	conn    Connection
	timeout int
	retries int
}
type Connection struct{}

func NewStuffClient(conn Connection, opts ...StuffClientOption) StuffClient {
	options := defaultStuffClientOptions
	for _, o := range opts {
		o(&options)
	}
	return &stuffClient{
		conn:    conn,
		timeout: options.Timeout,
		retries: options.Retries,
	}
}
func (c stuffClient) DoStuff() error {
	return nil
}

func TestOptionPattern1(t *testing.T) {
	stuClient := NewStuffClient(Connection{}, WithRetries(2), WithTimeout(10))
	stuClient.DoStuff()
}
