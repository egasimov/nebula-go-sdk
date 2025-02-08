// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	storage "github.com/egasimov/nebula-go-sdk/nebula/storage"
	mock "github.com/stretchr/testify/mock"
)

// GraphStorageService is an autogenerated mock type for the GraphStorageService type
type GraphStorageService struct {
	mock.Mock
}

// AddEdges provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) AddEdges(ctx context.Context, req *storage.AddEdgesRequest) (*storage.ExecResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for AddEdges")
	}

	var r0 *storage.ExecResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.AddEdgesRequest) (*storage.ExecResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.AddEdgesRequest) *storage.ExecResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ExecResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.AddEdgesRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddVertices provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) AddVertices(ctx context.Context, req *storage.AddVerticesRequest) (*storage.ExecResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for AddVertices")
	}

	var r0 *storage.ExecResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.AddVerticesRequest) (*storage.ExecResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.AddVerticesRequest) *storage.ExecResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ExecResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.AddVerticesRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainAddEdges provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) ChainAddEdges(ctx context.Context, req *storage.AddEdgesRequest) (*storage.ExecResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for ChainAddEdges")
	}

	var r0 *storage.ExecResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.AddEdgesRequest) (*storage.ExecResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.AddEdgesRequest) *storage.ExecResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ExecResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.AddEdgesRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainDeleteEdges provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) ChainDeleteEdges(ctx context.Context, req *storage.DeleteEdgesRequest) (*storage.ExecResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for ChainDeleteEdges")
	}

	var r0 *storage.ExecResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.DeleteEdgesRequest) (*storage.ExecResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.DeleteEdgesRequest) *storage.ExecResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ExecResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.DeleteEdgesRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainUpdateEdge provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) ChainUpdateEdge(ctx context.Context, req *storage.UpdateEdgeRequest) (*storage.UpdateResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for ChainUpdateEdge")
	}

	var r0 *storage.UpdateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.UpdateEdgeRequest) (*storage.UpdateResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.UpdateEdgeRequest) *storage.UpdateResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.UpdateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.UpdateEdgeRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEdges provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) DeleteEdges(ctx context.Context, req *storage.DeleteEdgesRequest) (*storage.ExecResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEdges")
	}

	var r0 *storage.ExecResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.DeleteEdgesRequest) (*storage.ExecResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.DeleteEdgesRequest) *storage.ExecResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ExecResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.DeleteEdgesRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTags provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) DeleteTags(ctx context.Context, req *storage.DeleteTagsRequest) (*storage.ExecResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTags")
	}

	var r0 *storage.ExecResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.DeleteTagsRequest) (*storage.ExecResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.DeleteTagsRequest) *storage.ExecResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ExecResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.DeleteTagsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVertices provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) DeleteVertices(ctx context.Context, req *storage.DeleteVerticesRequest) (*storage.ExecResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for DeleteVertices")
	}

	var r0 *storage.ExecResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.DeleteVerticesRequest) (*storage.ExecResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.DeleteVerticesRequest) *storage.ExecResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ExecResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.DeleteVerticesRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) Get(ctx context.Context, req *storage.KVGetRequest) (*storage.KVGetResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *storage.KVGetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.KVGetRequest) (*storage.KVGetResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.KVGetRequest) *storage.KVGetResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.KVGetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.KVGetRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDstBySrc provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) GetDstBySrc(ctx context.Context, req *storage.GetDstBySrcRequest) (*storage.GetDstBySrcResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetDstBySrc")
	}

	var r0 *storage.GetDstBySrcResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.GetDstBySrcRequest) (*storage.GetDstBySrcResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.GetDstBySrcRequest) *storage.GetDstBySrcResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.GetDstBySrcResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.GetDstBySrcRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNeighbors provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) GetNeighbors(ctx context.Context, req *storage.GetNeighborsRequest) (*storage.GetNeighborsResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetNeighbors")
	}

	var r0 *storage.GetNeighborsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.GetNeighborsRequest) (*storage.GetNeighborsResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.GetNeighborsRequest) *storage.GetNeighborsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.GetNeighborsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.GetNeighborsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProps provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) GetProps(ctx context.Context, req *storage.GetPropRequest) (*storage.GetPropResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetProps")
	}

	var r0 *storage.GetPropResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.GetPropRequest) (*storage.GetPropResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.GetPropRequest) *storage.GetPropResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.GetPropResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.GetPropRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUUID provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) GetUUID(ctx context.Context, req *storage.GetUUIDReq) (*storage.GetUUIDResp, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetUUID")
	}

	var r0 *storage.GetUUIDResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.GetUUIDReq) (*storage.GetUUIDResp, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.GetUUIDReq) *storage.GetUUIDResp); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.GetUUIDResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.GetUUIDReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupAndTraverse provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) LookupAndTraverse(ctx context.Context, req *storage.LookupAndTraverseRequest) (*storage.GetNeighborsResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for LookupAndTraverse")
	}

	var r0 *storage.GetNeighborsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.LookupAndTraverseRequest) (*storage.GetNeighborsResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.LookupAndTraverseRequest) *storage.GetNeighborsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.GetNeighborsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.LookupAndTraverseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupIndex provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) LookupIndex(ctx context.Context, req *storage.LookupIndexRequest) (*storage.LookupIndexResp, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for LookupIndex")
	}

	var r0 *storage.LookupIndexResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.LookupIndexRequest) (*storage.LookupIndexResp, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.LookupIndexRequest) *storage.LookupIndexResp); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.LookupIndexResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.LookupIndexRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) Put(ctx context.Context, req *storage.KVPutRequest) (*storage.ExecResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 *storage.ExecResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.KVPutRequest) (*storage.ExecResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.KVPutRequest) *storage.ExecResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ExecResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.KVPutRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) Remove(ctx context.Context, req *storage.KVRemoveRequest) (*storage.ExecResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 *storage.ExecResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.KVRemoveRequest) (*storage.ExecResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.KVRemoveRequest) *storage.ExecResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ExecResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.KVRemoveRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanEdge provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) ScanEdge(ctx context.Context, req *storage.ScanEdgeRequest) (*storage.ScanResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for ScanEdge")
	}

	var r0 *storage.ScanResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.ScanEdgeRequest) (*storage.ScanResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.ScanEdgeRequest) *storage.ScanResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ScanResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.ScanEdgeRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanVertex provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) ScanVertex(ctx context.Context, req *storage.ScanVertexRequest) (*storage.ScanResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for ScanVertex")
	}

	var r0 *storage.ScanResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.ScanVertexRequest) (*storage.ScanResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.ScanVertexRequest) *storage.ScanResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ScanResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.ScanVertexRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEdge provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) UpdateEdge(ctx context.Context, req *storage.UpdateEdgeRequest) (*storage.UpdateResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEdge")
	}

	var r0 *storage.UpdateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.UpdateEdgeRequest) (*storage.UpdateResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.UpdateEdgeRequest) *storage.UpdateResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.UpdateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.UpdateEdgeRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateVertex provides a mock function with given fields: ctx, req
func (_m *GraphStorageService) UpdateVertex(ctx context.Context, req *storage.UpdateVertexRequest) (*storage.UpdateResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateVertex")
	}

	var r0 *storage.UpdateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage.UpdateVertexRequest) (*storage.UpdateResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage.UpdateVertexRequest) *storage.UpdateResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.UpdateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage.UpdateVertexRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGraphStorageService creates a new instance of GraphStorageService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGraphStorageService(t interface {
	mock.TestingT
	Cleanup(func())
}) *GraphStorageService {
	mock := &GraphStorageService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
