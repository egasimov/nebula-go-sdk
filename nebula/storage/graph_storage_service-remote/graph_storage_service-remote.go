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
	"github.com/egasimov/nebula-go-sdk/nebula"
	"github.com/egasimov/nebula-go-sdk/nebula/meta"
	"github.com/egasimov/nebula-go-sdk/nebula/storage"
)

var _ = nebula.GoUnusedProtection__
var _ = meta.GoUnusedProtection__
var _ = storage.GoUnusedProtection__

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  GetNeighborsResponse getNeighbors(GetNeighborsRequest req)")
	fmt.Fprintln(os.Stderr, "  GetDstBySrcResponse getDstBySrc(GetDstBySrcRequest req)")
	fmt.Fprintln(os.Stderr, "  GetPropResponse getProps(GetPropRequest req)")
	fmt.Fprintln(os.Stderr, "  ExecResponse addVertices(AddVerticesRequest req)")
	fmt.Fprintln(os.Stderr, "  ExecResponse addEdges(AddEdgesRequest req)")
	fmt.Fprintln(os.Stderr, "  ExecResponse deleteEdges(DeleteEdgesRequest req)")
	fmt.Fprintln(os.Stderr, "  ExecResponse deleteVertices(DeleteVerticesRequest req)")
	fmt.Fprintln(os.Stderr, "  ExecResponse deleteTags(DeleteTagsRequest req)")
	fmt.Fprintln(os.Stderr, "  UpdateResponse updateVertex(UpdateVertexRequest req)")
	fmt.Fprintln(os.Stderr, "  UpdateResponse updateEdge(UpdateEdgeRequest req)")
	fmt.Fprintln(os.Stderr, "  ScanResponse scanVertex(ScanVertexRequest req)")
	fmt.Fprintln(os.Stderr, "  ScanResponse scanEdge(ScanEdgeRequest req)")
	fmt.Fprintln(os.Stderr, "  GetUUIDResp getUUID(GetUUIDReq req)")
	fmt.Fprintln(os.Stderr, "  LookupIndexResp lookupIndex(LookupIndexRequest req)")
	fmt.Fprintln(os.Stderr, "  GetNeighborsResponse lookupAndTraverse(LookupAndTraverseRequest req)")
	fmt.Fprintln(os.Stderr, "  UpdateResponse chainUpdateEdge(UpdateEdgeRequest req)")
	fmt.Fprintln(os.Stderr, "  ExecResponse chainAddEdges(AddEdgesRequest req)")
	fmt.Fprintln(os.Stderr, "  ExecResponse chainDeleteEdges(DeleteEdgesRequest req)")
	fmt.Fprintln(os.Stderr, "  KVGetResponse get(KVGetRequest req)")
	fmt.Fprintln(os.Stderr, "  ExecResponse put(KVPutRequest req)")
	fmt.Fprintln(os.Stderr, "  ExecResponse remove(KVRemoveRequest req)")
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
	client := storage.NewGraphStorageServiceClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}
	
	switch cmd {
	case "getNeighbors":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "GetNeighbors requires 1 args")
			flag.Usage()
		}
		arg306 := flag.Arg(1)
		mbTrans307 := thrift.NewTMemoryBufferLen(len(arg306))
		defer mbTrans307.Close()
		_, err308 := mbTrans307.WriteString(arg306)
		if err308 != nil {
			Usage()
			return
		}
		factory309 := thrift.NewTJSONProtocolFactory()
		jsProt310 := factory309.GetProtocol(mbTrans307)
		argvalue0 := storage.NewGetNeighborsRequest()
		err311 := argvalue0.Read(context.Background(), jsProt310)
		if err311 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetNeighbors(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getDstBySrc":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "GetDstBySrc requires 1 args")
			flag.Usage()
		}
		arg312 := flag.Arg(1)
		mbTrans313 := thrift.NewTMemoryBufferLen(len(arg312))
		defer mbTrans313.Close()
		_, err314 := mbTrans313.WriteString(arg312)
		if err314 != nil {
			Usage()
			return
		}
		factory315 := thrift.NewTJSONProtocolFactory()
		jsProt316 := factory315.GetProtocol(mbTrans313)
		argvalue0 := storage.NewGetDstBySrcRequest()
		err317 := argvalue0.Read(context.Background(), jsProt316)
		if err317 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetDstBySrc(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getProps":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "GetProps requires 1 args")
			flag.Usage()
		}
		arg318 := flag.Arg(1)
		mbTrans319 := thrift.NewTMemoryBufferLen(len(arg318))
		defer mbTrans319.Close()
		_, err320 := mbTrans319.WriteString(arg318)
		if err320 != nil {
			Usage()
			return
		}
		factory321 := thrift.NewTJSONProtocolFactory()
		jsProt322 := factory321.GetProtocol(mbTrans319)
		argvalue0 := storage.NewGetPropRequest()
		err323 := argvalue0.Read(context.Background(), jsProt322)
		if err323 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetProps(context.Background(), value0))
		fmt.Print("\n")
		break
	case "addVertices":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "AddVertices requires 1 args")
			flag.Usage()
		}
		arg324 := flag.Arg(1)
		mbTrans325 := thrift.NewTMemoryBufferLen(len(arg324))
		defer mbTrans325.Close()
		_, err326 := mbTrans325.WriteString(arg324)
		if err326 != nil {
			Usage()
			return
		}
		factory327 := thrift.NewTJSONProtocolFactory()
		jsProt328 := factory327.GetProtocol(mbTrans325)
		argvalue0 := storage.NewAddVerticesRequest()
		err329 := argvalue0.Read(context.Background(), jsProt328)
		if err329 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.AddVertices(context.Background(), value0))
		fmt.Print("\n")
		break
	case "addEdges":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "AddEdges requires 1 args")
			flag.Usage()
		}
		arg330 := flag.Arg(1)
		mbTrans331 := thrift.NewTMemoryBufferLen(len(arg330))
		defer mbTrans331.Close()
		_, err332 := mbTrans331.WriteString(arg330)
		if err332 != nil {
			Usage()
			return
		}
		factory333 := thrift.NewTJSONProtocolFactory()
		jsProt334 := factory333.GetProtocol(mbTrans331)
		argvalue0 := storage.NewAddEdgesRequest()
		err335 := argvalue0.Read(context.Background(), jsProt334)
		if err335 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.AddEdges(context.Background(), value0))
		fmt.Print("\n")
		break
	case "deleteEdges":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteEdges requires 1 args")
			flag.Usage()
		}
		arg336 := flag.Arg(1)
		mbTrans337 := thrift.NewTMemoryBufferLen(len(arg336))
		defer mbTrans337.Close()
		_, err338 := mbTrans337.WriteString(arg336)
		if err338 != nil {
			Usage()
			return
		}
		factory339 := thrift.NewTJSONProtocolFactory()
		jsProt340 := factory339.GetProtocol(mbTrans337)
		argvalue0 := storage.NewDeleteEdgesRequest()
		err341 := argvalue0.Read(context.Background(), jsProt340)
		if err341 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DeleteEdges(context.Background(), value0))
		fmt.Print("\n")
		break
	case "deleteVertices":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteVertices requires 1 args")
			flag.Usage()
		}
		arg342 := flag.Arg(1)
		mbTrans343 := thrift.NewTMemoryBufferLen(len(arg342))
		defer mbTrans343.Close()
		_, err344 := mbTrans343.WriteString(arg342)
		if err344 != nil {
			Usage()
			return
		}
		factory345 := thrift.NewTJSONProtocolFactory()
		jsProt346 := factory345.GetProtocol(mbTrans343)
		argvalue0 := storage.NewDeleteVerticesRequest()
		err347 := argvalue0.Read(context.Background(), jsProt346)
		if err347 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DeleteVertices(context.Background(), value0))
		fmt.Print("\n")
		break
	case "deleteTags":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteTags requires 1 args")
			flag.Usage()
		}
		arg348 := flag.Arg(1)
		mbTrans349 := thrift.NewTMemoryBufferLen(len(arg348))
		defer mbTrans349.Close()
		_, err350 := mbTrans349.WriteString(arg348)
		if err350 != nil {
			Usage()
			return
		}
		factory351 := thrift.NewTJSONProtocolFactory()
		jsProt352 := factory351.GetProtocol(mbTrans349)
		argvalue0 := storage.NewDeleteTagsRequest()
		err353 := argvalue0.Read(context.Background(), jsProt352)
		if err353 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DeleteTags(context.Background(), value0))
		fmt.Print("\n")
		break
	case "updateVertex":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateVertex requires 1 args")
			flag.Usage()
		}
		arg354 := flag.Arg(1)
		mbTrans355 := thrift.NewTMemoryBufferLen(len(arg354))
		defer mbTrans355.Close()
		_, err356 := mbTrans355.WriteString(arg354)
		if err356 != nil {
			Usage()
			return
		}
		factory357 := thrift.NewTJSONProtocolFactory()
		jsProt358 := factory357.GetProtocol(mbTrans355)
		argvalue0 := storage.NewUpdateVertexRequest()
		err359 := argvalue0.Read(context.Background(), jsProt358)
		if err359 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.UpdateVertex(context.Background(), value0))
		fmt.Print("\n")
		break
	case "updateEdge":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateEdge requires 1 args")
			flag.Usage()
		}
		arg360 := flag.Arg(1)
		mbTrans361 := thrift.NewTMemoryBufferLen(len(arg360))
		defer mbTrans361.Close()
		_, err362 := mbTrans361.WriteString(arg360)
		if err362 != nil {
			Usage()
			return
		}
		factory363 := thrift.NewTJSONProtocolFactory()
		jsProt364 := factory363.GetProtocol(mbTrans361)
		argvalue0 := storage.NewUpdateEdgeRequest()
		err365 := argvalue0.Read(context.Background(), jsProt364)
		if err365 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.UpdateEdge(context.Background(), value0))
		fmt.Print("\n")
		break
	case "scanVertex":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "ScanVertex requires 1 args")
			flag.Usage()
		}
		arg366 := flag.Arg(1)
		mbTrans367 := thrift.NewTMemoryBufferLen(len(arg366))
		defer mbTrans367.Close()
		_, err368 := mbTrans367.WriteString(arg366)
		if err368 != nil {
			Usage()
			return
		}
		factory369 := thrift.NewTJSONProtocolFactory()
		jsProt370 := factory369.GetProtocol(mbTrans367)
		argvalue0 := storage.NewScanVertexRequest()
		err371 := argvalue0.Read(context.Background(), jsProt370)
		if err371 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ScanVertex(context.Background(), value0))
		fmt.Print("\n")
		break
	case "scanEdge":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "ScanEdge requires 1 args")
			flag.Usage()
		}
		arg372 := flag.Arg(1)
		mbTrans373 := thrift.NewTMemoryBufferLen(len(arg372))
		defer mbTrans373.Close()
		_, err374 := mbTrans373.WriteString(arg372)
		if err374 != nil {
			Usage()
			return
		}
		factory375 := thrift.NewTJSONProtocolFactory()
		jsProt376 := factory375.GetProtocol(mbTrans373)
		argvalue0 := storage.NewScanEdgeRequest()
		err377 := argvalue0.Read(context.Background(), jsProt376)
		if err377 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ScanEdge(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getUUID":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "GetUUID requires 1 args")
			flag.Usage()
		}
		arg378 := flag.Arg(1)
		mbTrans379 := thrift.NewTMemoryBufferLen(len(arg378))
		defer mbTrans379.Close()
		_, err380 := mbTrans379.WriteString(arg378)
		if err380 != nil {
			Usage()
			return
		}
		factory381 := thrift.NewTJSONProtocolFactory()
		jsProt382 := factory381.GetProtocol(mbTrans379)
		argvalue0 := storage.NewGetUUIDReq()
		err383 := argvalue0.Read(context.Background(), jsProt382)
		if err383 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetUUID(context.Background(), value0))
		fmt.Print("\n")
		break
	case "lookupIndex":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "LookupIndex requires 1 args")
			flag.Usage()
		}
		arg384 := flag.Arg(1)
		mbTrans385 := thrift.NewTMemoryBufferLen(len(arg384))
		defer mbTrans385.Close()
		_, err386 := mbTrans385.WriteString(arg384)
		if err386 != nil {
			Usage()
			return
		}
		factory387 := thrift.NewTJSONProtocolFactory()
		jsProt388 := factory387.GetProtocol(mbTrans385)
		argvalue0 := storage.NewLookupIndexRequest()
		err389 := argvalue0.Read(context.Background(), jsProt388)
		if err389 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.LookupIndex(context.Background(), value0))
		fmt.Print("\n")
		break
	case "lookupAndTraverse":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "LookupAndTraverse requires 1 args")
			flag.Usage()
		}
		arg390 := flag.Arg(1)
		mbTrans391 := thrift.NewTMemoryBufferLen(len(arg390))
		defer mbTrans391.Close()
		_, err392 := mbTrans391.WriteString(arg390)
		if err392 != nil {
			Usage()
			return
		}
		factory393 := thrift.NewTJSONProtocolFactory()
		jsProt394 := factory393.GetProtocol(mbTrans391)
		argvalue0 := storage.NewLookupAndTraverseRequest()
		err395 := argvalue0.Read(context.Background(), jsProt394)
		if err395 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.LookupAndTraverse(context.Background(), value0))
		fmt.Print("\n")
		break
	case "chainUpdateEdge":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "ChainUpdateEdge requires 1 args")
			flag.Usage()
		}
		arg396 := flag.Arg(1)
		mbTrans397 := thrift.NewTMemoryBufferLen(len(arg396))
		defer mbTrans397.Close()
		_, err398 := mbTrans397.WriteString(arg396)
		if err398 != nil {
			Usage()
			return
		}
		factory399 := thrift.NewTJSONProtocolFactory()
		jsProt400 := factory399.GetProtocol(mbTrans397)
		argvalue0 := storage.NewUpdateEdgeRequest()
		err401 := argvalue0.Read(context.Background(), jsProt400)
		if err401 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ChainUpdateEdge(context.Background(), value0))
		fmt.Print("\n")
		break
	case "chainAddEdges":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "ChainAddEdges requires 1 args")
			flag.Usage()
		}
		arg402 := flag.Arg(1)
		mbTrans403 := thrift.NewTMemoryBufferLen(len(arg402))
		defer mbTrans403.Close()
		_, err404 := mbTrans403.WriteString(arg402)
		if err404 != nil {
			Usage()
			return
		}
		factory405 := thrift.NewTJSONProtocolFactory()
		jsProt406 := factory405.GetProtocol(mbTrans403)
		argvalue0 := storage.NewAddEdgesRequest()
		err407 := argvalue0.Read(context.Background(), jsProt406)
		if err407 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ChainAddEdges(context.Background(), value0))
		fmt.Print("\n")
		break
	case "chainDeleteEdges":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "ChainDeleteEdges requires 1 args")
			flag.Usage()
		}
		arg408 := flag.Arg(1)
		mbTrans409 := thrift.NewTMemoryBufferLen(len(arg408))
		defer mbTrans409.Close()
		_, err410 := mbTrans409.WriteString(arg408)
		if err410 != nil {
			Usage()
			return
		}
		factory411 := thrift.NewTJSONProtocolFactory()
		jsProt412 := factory411.GetProtocol(mbTrans409)
		argvalue0 := storage.NewDeleteEdgesRequest()
		err413 := argvalue0.Read(context.Background(), jsProt412)
		if err413 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ChainDeleteEdges(context.Background(), value0))
		fmt.Print("\n")
		break
	case "get":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "Get requires 1 args")
			flag.Usage()
		}
		arg414 := flag.Arg(1)
		mbTrans415 := thrift.NewTMemoryBufferLen(len(arg414))
		defer mbTrans415.Close()
		_, err416 := mbTrans415.WriteString(arg414)
		if err416 != nil {
			Usage()
			return
		}
		factory417 := thrift.NewTJSONProtocolFactory()
		jsProt418 := factory417.GetProtocol(mbTrans415)
		argvalue0 := storage.NewKVGetRequest()
		err419 := argvalue0.Read(context.Background(), jsProt418)
		if err419 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Get(context.Background(), value0))
		fmt.Print("\n")
		break
	case "put":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "Put requires 1 args")
			flag.Usage()
		}
		arg420 := flag.Arg(1)
		mbTrans421 := thrift.NewTMemoryBufferLen(len(arg420))
		defer mbTrans421.Close()
		_, err422 := mbTrans421.WriteString(arg420)
		if err422 != nil {
			Usage()
			return
		}
		factory423 := thrift.NewTJSONProtocolFactory()
		jsProt424 := factory423.GetProtocol(mbTrans421)
		argvalue0 := storage.NewKVPutRequest()
		err425 := argvalue0.Read(context.Background(), jsProt424)
		if err425 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Put(context.Background(), value0))
		fmt.Print("\n")
		break
	case "remove":
		if flag.NArg() - 1 != 1 {
			fmt.Fprintln(os.Stderr, "Remove requires 1 args")
			flag.Usage()
		}
		arg426 := flag.Arg(1)
		mbTrans427 := thrift.NewTMemoryBufferLen(len(arg426))
		defer mbTrans427.Close()
		_, err428 := mbTrans427.WriteString(arg426)
		if err428 != nil {
			Usage()
			return
		}
		factory429 := thrift.NewTJSONProtocolFactory()
		jsProt430 := factory429.GetProtocol(mbTrans427)
		argvalue0 := storage.NewKVRemoveRequest()
		err431 := argvalue0.Read(context.Background(), jsProt430)
		if err431 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Remove(context.Background(), value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
