// Code generated by Thrift Compiler (0.21.0). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/nebula-contrib/nebula-sirius/nebula"
	"github.com/nebula-contrib/nebula-sirius/nebula/graph"
)

var _ = nebula.GoUnusedProtection__
var _ = graph.GoUnusedProtection__

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  AuthResponse authenticate(string username, string password)")
	fmt.Fprintln(os.Stderr, "  void signout(i64 sessionId)")
	fmt.Fprintln(os.Stderr, "  ExecutionResponse execute(i64 sessionId, string stmt)")
	fmt.Fprintln(os.Stderr, "  ExecutionResponse executeWithParameter(i64 sessionId, string stmt,  parameterMap)")
	fmt.Fprintln(os.Stderr, "  string executeJson(i64 sessionId, string stmt)")
	fmt.Fprintln(os.Stderr, "  string executeJsonWithParameter(i64 sessionId, string stmt,  parameterMap)")
	fmt.Fprintln(os.Stderr, "  VerifyClientVersionResp verifyClientVersion(VerifyClientVersionReq req)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
	var m map[string]string = h
	return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
	parts := strings.Split(value, ": ")
	if len(parts) != 2 {
		return fmt.Errorf("header should be of format 'Key: Value'")
	}
	h[parts[0]] = parts[1]
	return nil
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	headers := make(httpHeaders)
	var parsedUrl *url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
	flag.Parse()
	
	if len(urlString) > 0 {
		var err error
		parsedUrl, err = url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}
	
	cmd := flag.Arg(0)
	var err error
	var cfg *thrift.TConfiguration = nil
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
		if len(headers) > 0 {
			httptrans := trans.(*thrift.THttpClient)
			for key, value := range headers {
				httptrans.SetHeader(key, value)
			}
		}
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans = thrift.NewTSocketConf(net.JoinHostPort(host, portStr), cfg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransportConf(trans, cfg)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := graph.NewGraphServiceClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}
	
	switch cmd {
	case "authenticate":
		if flag.NArg() - 1 != 2 {
			fmt.Fprintln(os.Stderr, "Authenticate requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		fmt.Print(client.Authenticate(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "signout":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "Signout requires 1 args")
			flag.Usage()
		}
		argvalue0, err57 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err57 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Signout(context.Background(), value0))
		fmt.Print("\n")
		break
	case "execute":
		if flag.NArg() - 1 != 2 {
			fmt.Fprintln(os.Stderr, "Execute requires 2 args")
			flag.Usage()
		}
		argvalue0, err58 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err58 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		fmt.Print(client.Execute(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "executeWithParameter":
		if flag.NArg() - 1 != 3 {
			fmt.Fprintln(os.Stderr, "ExecuteWithParameter requires 3 args")
			flag.Usage()
		}
		argvalue0, err60 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err60 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		arg62 := flag.Arg(3)
		mbTrans63 := thrift.NewTMemoryBufferLen(len(arg62))
		defer mbTrans63.Close()
		_, err64 := mbTrans63.WriteString(arg62)
		if err64 != nil {
			Usage()
			return
		}
		factory65 := thrift.NewTJSONProtocolFactory()
		jsProt66 := factory65.GetProtocol(mbTrans63)
		containerStruct2 := graph.NewGraphServiceExecuteWithParameterArgs()
		err67 := containerStruct2.ReadField3(context.Background(), jsProt66)
		if err67 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.ParameterMap
		value2 := argvalue2
		fmt.Print(client.ExecuteWithParameter(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "executeJson":
		if flag.NArg() - 1 != 2 {
			fmt.Fprintln(os.Stderr, "ExecuteJson requires 2 args")
			flag.Usage()
		}
		argvalue0, err68 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err68 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		fmt.Print(client.ExecuteJson(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "executeJsonWithParameter":
		if flag.NArg() - 1 != 3 {
			fmt.Fprintln(os.Stderr, "ExecuteJsonWithParameter requires 3 args")
			flag.Usage()
		}
		argvalue0, err70 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err70 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		arg72 := flag.Arg(3)
		mbTrans73 := thrift.NewTMemoryBufferLen(len(arg72))
		defer mbTrans73.Close()
		_, err74 := mbTrans73.WriteString(arg72)
		if err74 != nil {
			Usage()
			return
		}
		factory75 := thrift.NewTJSONProtocolFactory()
		jsProt76 := factory75.GetProtocol(mbTrans73)
		containerStruct2 := graph.NewGraphServiceExecuteJsonWithParameterArgs()
		err77 := containerStruct2.ReadField3(context.Background(), jsProt76)
		if err77 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.ParameterMap
		value2 := argvalue2
		fmt.Print(client.ExecuteJsonWithParameter(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "verifyClientVersion":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "VerifyClientVersion requires 1 args")
			flag.Usage()
		}
		arg78 := flag.Arg(1)
		mbTrans79 := thrift.NewTMemoryBufferLen(len(arg78))
		defer mbTrans79.Close()
		_, err80 := mbTrans79.WriteString(arg78)
		if err80 != nil {
			Usage()
			return
		}
		factory81 := thrift.NewTJSONProtocolFactory()
		jsProt82 := factory81.GetProtocol(mbTrans79)
		argvalue0 := graph.NewVerifyClientVersionReq()
		err83 := argvalue0.Read(context.Background(), jsProt82)
		if err83 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.VerifyClientVersion(context.Background(), value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
