// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package asset

import (
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1beta1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/api/option"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockAssetServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	assetpb.AssetServiceServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockAssetServer) ExportAssets(ctx context.Context, req *assetpb.ExportAssetsRequest) (*longrunningpb.Operation, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*longrunningpb.Operation), nil
}

func (s *mockAssetServer) BatchGetAssetsHistory(ctx context.Context, req *assetpb.BatchGetAssetsHistoryRequest) (*assetpb.BatchGetAssetsHistoryResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*assetpb.BatchGetAssetsHistoryResponse), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockAsset mockAssetServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	assetpb.RegisterAssetServiceServer(serv, &mockAsset)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestAssetServiceExportAssets(t *testing.T) {
	var expectedResponse *assetpb.ExportAssetsResponse = &assetpb.ExportAssetsResponse{}

	mockAsset.err = nil
	mockAsset.reqs = nil

	any, err := ptypes.MarshalAny(expectedResponse)
	if err != nil {
		t.Fatal(err)
	}
	mockAsset.resps = append(mockAsset.resps[:0], &longrunningpb.Operation{
		Name:   "longrunning-test",
		Done:   true,
		Result: &longrunningpb.Operation_Response{Response: any},
	})

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var outputConfig *assetpb.OutputConfig = &assetpb.OutputConfig{}
	var request = &assetpb.ExportAssetsRequest{
		Parent:       formattedParent,
		OutputConfig: outputConfig,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	respLRO, err := c.ExportAssets(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := respLRO.Wait(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockAsset.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestAssetServiceExportAssetsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockAsset.err = nil
	mockAsset.resps = append(mockAsset.resps[:0], &longrunningpb.Operation{
		Name: "longrunning-test",
		Done: true,
		Result: &longrunningpb.Operation_Error{
			Error: &status.Status{
				Code:    int32(errCode),
				Message: "test error",
			},
		},
	})

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var outputConfig *assetpb.OutputConfig = &assetpb.OutputConfig{}
	var request = &assetpb.ExportAssetsRequest{
		Parent:       formattedParent,
		OutputConfig: outputConfig,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	respLRO, err := c.ExportAssets(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := respLRO.Wait(context.Background())

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestAssetServiceBatchGetAssetsHistory(t *testing.T) {
	var expectedResponse *assetpb.BatchGetAssetsHistoryResponse = &assetpb.BatchGetAssetsHistoryResponse{}

	mockAsset.err = nil
	mockAsset.reqs = nil

	mockAsset.resps = append(mockAsset.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var contentType assetpb.ContentType = assetpb.ContentType_CONTENT_TYPE_UNSPECIFIED
	var readTimeWindow *assetpb.TimeWindow = &assetpb.TimeWindow{}
	var request = &assetpb.BatchGetAssetsHistoryRequest{
		Parent:         formattedParent,
		ContentType:    contentType,
		ReadTimeWindow: readTimeWindow,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.BatchGetAssetsHistory(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockAsset.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestAssetServiceBatchGetAssetsHistoryError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockAsset.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var contentType assetpb.ContentType = assetpb.ContentType_CONTENT_TYPE_UNSPECIFIED
	var readTimeWindow *assetpb.TimeWindow = &assetpb.TimeWindow{}
	var request = &assetpb.BatchGetAssetsHistoryRequest{
		Parent:         formattedParent,
		ContentType:    contentType,
		ReadTimeWindow: readTimeWindow,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.BatchGetAssetsHistory(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
