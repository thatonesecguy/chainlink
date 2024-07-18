// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/smartcontractkit/chainlink-common/pkg/sqlutil/models"
	mock "github.com/stretchr/testify/mock"

	pipeline "github.com/smartcontractkit/chainlink/v2/core/services/pipeline"

	sqlutil "github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	time "time"

	uuid "github.com/google/uuid"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

type ORM_Expecter struct {
	mock *mock.Mock
}

func (_m *ORM) EXPECT() *ORM_Expecter {
	return &ORM_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *ORM) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type ORM_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *ORM_Expecter) Close() *ORM_Close_Call {
	return &ORM_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *ORM_Close_Call) Run(run func()) *ORM_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ORM_Close_Call) Return(_a0 error) *ORM_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_Close_Call) RunAndReturn(run func() error) *ORM_Close_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRun provides a mock function with given fields: ctx, run
func (_m *ORM) CreateRun(ctx context.Context, run *pipeline.Run) error {
	ret := _m.Called(ctx, run)

	if len(ret) == 0 {
		panic("no return value specified for CreateRun")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipeline.Run) error); ok {
		r0 = rf(ctx, run)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_CreateRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRun'
type ORM_CreateRun_Call struct {
	*mock.Call
}

// CreateRun is a helper method to define mock.On call
//   - ctx context.Context
//   - run *pipeline.Run
func (_e *ORM_Expecter) CreateRun(ctx interface{}, run interface{}) *ORM_CreateRun_Call {
	return &ORM_CreateRun_Call{Call: _e.mock.On("CreateRun", ctx, run)}
}

func (_c *ORM_CreateRun_Call) Run(run func(ctx context.Context, run *pipeline.Run)) *ORM_CreateRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pipeline.Run))
	})
	return _c
}

func (_c *ORM_CreateRun_Call) Return(err error) *ORM_CreateRun_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ORM_CreateRun_Call) RunAndReturn(run func(context.Context, *pipeline.Run) error) *ORM_CreateRun_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSpec provides a mock function with given fields: ctx, _a1, maxTaskTimeout
func (_m *ORM) CreateSpec(ctx context.Context, _a1 pipeline.Pipeline, maxTaskTimeout models.Interval) (int32, error) {
	ret := _m.Called(ctx, _a1, maxTaskTimeout)

	if len(ret) == 0 {
		panic("no return value specified for CreateSpec")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pipeline.Pipeline, models.Interval) (int32, error)); ok {
		return rf(ctx, _a1, maxTaskTimeout)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pipeline.Pipeline, models.Interval) int32); ok {
		r0 = rf(ctx, _a1, maxTaskTimeout)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pipeline.Pipeline, models.Interval) error); ok {
		r1 = rf(ctx, _a1, maxTaskTimeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_CreateSpec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSpec'
type ORM_CreateSpec_Call struct {
	*mock.Call
}

// CreateSpec is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 pipeline.Pipeline
//   - maxTaskTimeout models.Interval
func (_e *ORM_Expecter) CreateSpec(ctx interface{}, _a1 interface{}, maxTaskTimeout interface{}) *ORM_CreateSpec_Call {
	return &ORM_CreateSpec_Call{Call: _e.mock.On("CreateSpec", ctx, _a1, maxTaskTimeout)}
}

func (_c *ORM_CreateSpec_Call) Run(run func(ctx context.Context, _a1 pipeline.Pipeline, maxTaskTimeout models.Interval)) *ORM_CreateSpec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pipeline.Pipeline), args[2].(models.Interval))
	})
	return _c
}

func (_c *ORM_CreateSpec_Call) Return(_a0 int32, _a1 error) *ORM_CreateSpec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_CreateSpec_Call) RunAndReturn(run func(context.Context, pipeline.Pipeline, models.Interval) (int32, error)) *ORM_CreateSpec_Call {
	_c.Call.Return(run)
	return _c
}

// DataSource provides a mock function with given fields:
func (_m *ORM) DataSource() sqlutil.DataSource {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DataSource")
	}

	var r0 sqlutil.DataSource
	if rf, ok := ret.Get(0).(func() sqlutil.DataSource); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sqlutil.DataSource)
		}
	}

	return r0
}

// ORM_DataSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DataSource'
type ORM_DataSource_Call struct {
	*mock.Call
}

// DataSource is a helper method to define mock.On call
func (_e *ORM_Expecter) DataSource() *ORM_DataSource_Call {
	return &ORM_DataSource_Call{Call: _e.mock.On("DataSource")}
}

func (_c *ORM_DataSource_Call) Run(run func()) *ORM_DataSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ORM_DataSource_Call) Return(_a0 sqlutil.DataSource) *ORM_DataSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_DataSource_Call) RunAndReturn(run func() sqlutil.DataSource) *ORM_DataSource_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRun provides a mock function with given fields: ctx, id
func (_m *ORM) DeleteRun(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRun")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_DeleteRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRun'
type ORM_DeleteRun_Call struct {
	*mock.Call
}

// DeleteRun is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *ORM_Expecter) DeleteRun(ctx interface{}, id interface{}) *ORM_DeleteRun_Call {
	return &ORM_DeleteRun_Call{Call: _e.mock.On("DeleteRun", ctx, id)}
}

func (_c *ORM_DeleteRun_Call) Run(run func(ctx context.Context, id int64)) *ORM_DeleteRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *ORM_DeleteRun_Call) Return(_a0 error) *ORM_DeleteRun_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_DeleteRun_Call) RunAndReturn(run func(context.Context, int64) error) *ORM_DeleteRun_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRunsOlderThan provides a mock function with given fields: _a0, _a1
func (_m *ORM) DeleteRunsOlderThan(_a0 context.Context, _a1 time.Duration) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRunsOlderThan")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_DeleteRunsOlderThan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRunsOlderThan'
type ORM_DeleteRunsOlderThan_Call struct {
	*mock.Call
}

// DeleteRunsOlderThan is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 time.Duration
func (_e *ORM_Expecter) DeleteRunsOlderThan(_a0 interface{}, _a1 interface{}) *ORM_DeleteRunsOlderThan_Call {
	return &ORM_DeleteRunsOlderThan_Call{Call: _e.mock.On("DeleteRunsOlderThan", _a0, _a1)}
}

func (_c *ORM_DeleteRunsOlderThan_Call) Run(run func(_a0 context.Context, _a1 time.Duration)) *ORM_DeleteRunsOlderThan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(time.Duration))
	})
	return _c
}

func (_c *ORM_DeleteRunsOlderThan_Call) Return(_a0 error) *ORM_DeleteRunsOlderThan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_DeleteRunsOlderThan_Call) RunAndReturn(run func(context.Context, time.Duration) error) *ORM_DeleteRunsOlderThan_Call {
	_c.Call.Return(run)
	return _c
}

// FindRun provides a mock function with given fields: ctx, id
func (_m *ORM) FindRun(ctx context.Context, id int64) (pipeline.Run, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindRun")
	}

	var r0 pipeline.Run
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (pipeline.Run, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) pipeline.Run); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pipeline.Run)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_FindRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindRun'
type ORM_FindRun_Call struct {
	*mock.Call
}

// FindRun is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *ORM_Expecter) FindRun(ctx interface{}, id interface{}) *ORM_FindRun_Call {
	return &ORM_FindRun_Call{Call: _e.mock.On("FindRun", ctx, id)}
}

func (_c *ORM_FindRun_Call) Run(run func(ctx context.Context, id int64)) *ORM_FindRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *ORM_FindRun_Call) Return(_a0 pipeline.Run, _a1 error) *ORM_FindRun_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_FindRun_Call) RunAndReturn(run func(context.Context, int64) (pipeline.Run, error)) *ORM_FindRun_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllRuns provides a mock function with given fields: ctx
func (_m *ORM) GetAllRuns(ctx context.Context) ([]pipeline.Run, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllRuns")
	}

	var r0 []pipeline.Run
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]pipeline.Run, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []pipeline.Run); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipeline.Run)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_GetAllRuns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllRuns'
type ORM_GetAllRuns_Call struct {
	*mock.Call
}

// GetAllRuns is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ORM_Expecter) GetAllRuns(ctx interface{}) *ORM_GetAllRuns_Call {
	return &ORM_GetAllRuns_Call{Call: _e.mock.On("GetAllRuns", ctx)}
}

func (_c *ORM_GetAllRuns_Call) Run(run func(ctx context.Context)) *ORM_GetAllRuns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ORM_GetAllRuns_Call) Return(_a0 []pipeline.Run, _a1 error) *ORM_GetAllRuns_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_GetAllRuns_Call) RunAndReturn(run func(context.Context) ([]pipeline.Run, error)) *ORM_GetAllRuns_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnfinishedRuns provides a mock function with given fields: _a0, _a1, _a2
func (_m *ORM) GetUnfinishedRuns(_a0 context.Context, _a1 time.Time, _a2 func(pipeline.Run) error) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetUnfinishedRuns")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, func(pipeline.Run) error) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_GetUnfinishedRuns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnfinishedRuns'
type ORM_GetUnfinishedRuns_Call struct {
	*mock.Call
}

// GetUnfinishedRuns is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 time.Time
//   - _a2 func(pipeline.Run) error
func (_e *ORM_Expecter) GetUnfinishedRuns(_a0 interface{}, _a1 interface{}, _a2 interface{}) *ORM_GetUnfinishedRuns_Call {
	return &ORM_GetUnfinishedRuns_Call{Call: _e.mock.On("GetUnfinishedRuns", _a0, _a1, _a2)}
}

func (_c *ORM_GetUnfinishedRuns_Call) Run(run func(_a0 context.Context, _a1 time.Time, _a2 func(pipeline.Run) error)) *ORM_GetUnfinishedRuns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(time.Time), args[2].(func(pipeline.Run) error))
	})
	return _c
}

func (_c *ORM_GetUnfinishedRuns_Call) Return(_a0 error) *ORM_GetUnfinishedRuns_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_GetUnfinishedRuns_Call) RunAndReturn(run func(context.Context, time.Time, func(pipeline.Run) error) error) *ORM_GetUnfinishedRuns_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with given fields:
func (_m *ORM) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// ORM_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type ORM_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *ORM_Expecter) HealthReport() *ORM_HealthReport_Call {
	return &ORM_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *ORM_HealthReport_Call) Run(run func()) *ORM_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ORM_HealthReport_Call) Return(_a0 map[string]error) *ORM_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_HealthReport_Call) RunAndReturn(run func() map[string]error) *ORM_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// InsertFinishedRun provides a mock function with given fields: ctx, run, saveSuccessfulTaskRuns
func (_m *ORM) InsertFinishedRun(ctx context.Context, run *pipeline.Run, saveSuccessfulTaskRuns bool) error {
	ret := _m.Called(ctx, run, saveSuccessfulTaskRuns)

	if len(ret) == 0 {
		panic("no return value specified for InsertFinishedRun")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipeline.Run, bool) error); ok {
		r0 = rf(ctx, run, saveSuccessfulTaskRuns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_InsertFinishedRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertFinishedRun'
type ORM_InsertFinishedRun_Call struct {
	*mock.Call
}

// InsertFinishedRun is a helper method to define mock.On call
//   - ctx context.Context
//   - run *pipeline.Run
//   - saveSuccessfulTaskRuns bool
func (_e *ORM_Expecter) InsertFinishedRun(ctx interface{}, run interface{}, saveSuccessfulTaskRuns interface{}) *ORM_InsertFinishedRun_Call {
	return &ORM_InsertFinishedRun_Call{Call: _e.mock.On("InsertFinishedRun", ctx, run, saveSuccessfulTaskRuns)}
}

func (_c *ORM_InsertFinishedRun_Call) Run(run func(ctx context.Context, run *pipeline.Run, saveSuccessfulTaskRuns bool)) *ORM_InsertFinishedRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pipeline.Run), args[2].(bool))
	})
	return _c
}

func (_c *ORM_InsertFinishedRun_Call) Return(err error) *ORM_InsertFinishedRun_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ORM_InsertFinishedRun_Call) RunAndReturn(run func(context.Context, *pipeline.Run, bool) error) *ORM_InsertFinishedRun_Call {
	_c.Call.Return(run)
	return _c
}

// InsertFinishedRunWithSpec provides a mock function with given fields: ctx, run, saveSuccessfulTaskRuns
func (_m *ORM) InsertFinishedRunWithSpec(ctx context.Context, run *pipeline.Run, saveSuccessfulTaskRuns bool) error {
	ret := _m.Called(ctx, run, saveSuccessfulTaskRuns)

	if len(ret) == 0 {
		panic("no return value specified for InsertFinishedRunWithSpec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipeline.Run, bool) error); ok {
		r0 = rf(ctx, run, saveSuccessfulTaskRuns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_InsertFinishedRunWithSpec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertFinishedRunWithSpec'
type ORM_InsertFinishedRunWithSpec_Call struct {
	*mock.Call
}

// InsertFinishedRunWithSpec is a helper method to define mock.On call
//   - ctx context.Context
//   - run *pipeline.Run
//   - saveSuccessfulTaskRuns bool
func (_e *ORM_Expecter) InsertFinishedRunWithSpec(ctx interface{}, run interface{}, saveSuccessfulTaskRuns interface{}) *ORM_InsertFinishedRunWithSpec_Call {
	return &ORM_InsertFinishedRunWithSpec_Call{Call: _e.mock.On("InsertFinishedRunWithSpec", ctx, run, saveSuccessfulTaskRuns)}
}

func (_c *ORM_InsertFinishedRunWithSpec_Call) Run(run func(ctx context.Context, run *pipeline.Run, saveSuccessfulTaskRuns bool)) *ORM_InsertFinishedRunWithSpec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pipeline.Run), args[2].(bool))
	})
	return _c
}

func (_c *ORM_InsertFinishedRunWithSpec_Call) Return(err error) *ORM_InsertFinishedRunWithSpec_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ORM_InsertFinishedRunWithSpec_Call) RunAndReturn(run func(context.Context, *pipeline.Run, bool) error) *ORM_InsertFinishedRunWithSpec_Call {
	_c.Call.Return(run)
	return _c
}

// InsertFinishedRuns provides a mock function with given fields: ctx, run, saveSuccessfulTaskRuns
func (_m *ORM) InsertFinishedRuns(ctx context.Context, run []*pipeline.Run, saveSuccessfulTaskRuns bool) error {
	ret := _m.Called(ctx, run, saveSuccessfulTaskRuns)

	if len(ret) == 0 {
		panic("no return value specified for InsertFinishedRuns")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*pipeline.Run, bool) error); ok {
		r0 = rf(ctx, run, saveSuccessfulTaskRuns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_InsertFinishedRuns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertFinishedRuns'
type ORM_InsertFinishedRuns_Call struct {
	*mock.Call
}

// InsertFinishedRuns is a helper method to define mock.On call
//   - ctx context.Context
//   - run []*pipeline.Run
//   - saveSuccessfulTaskRuns bool
func (_e *ORM_Expecter) InsertFinishedRuns(ctx interface{}, run interface{}, saveSuccessfulTaskRuns interface{}) *ORM_InsertFinishedRuns_Call {
	return &ORM_InsertFinishedRuns_Call{Call: _e.mock.On("InsertFinishedRuns", ctx, run, saveSuccessfulTaskRuns)}
}

func (_c *ORM_InsertFinishedRuns_Call) Run(run func(ctx context.Context, run []*pipeline.Run, saveSuccessfulTaskRuns bool)) *ORM_InsertFinishedRuns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*pipeline.Run), args[2].(bool))
	})
	return _c
}

func (_c *ORM_InsertFinishedRuns_Call) Return(err error) *ORM_InsertFinishedRuns_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ORM_InsertFinishedRuns_Call) RunAndReturn(run func(context.Context, []*pipeline.Run, bool) error) *ORM_InsertFinishedRuns_Call {
	_c.Call.Return(run)
	return _c
}

// InsertRun provides a mock function with given fields: ctx, run
func (_m *ORM) InsertRun(ctx context.Context, run *pipeline.Run) error {
	ret := _m.Called(ctx, run)

	if len(ret) == 0 {
		panic("no return value specified for InsertRun")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipeline.Run) error); ok {
		r0 = rf(ctx, run)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_InsertRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertRun'
type ORM_InsertRun_Call struct {
	*mock.Call
}

// InsertRun is a helper method to define mock.On call
//   - ctx context.Context
//   - run *pipeline.Run
func (_e *ORM_Expecter) InsertRun(ctx interface{}, run interface{}) *ORM_InsertRun_Call {
	return &ORM_InsertRun_Call{Call: _e.mock.On("InsertRun", ctx, run)}
}

func (_c *ORM_InsertRun_Call) Run(run func(ctx context.Context, run *pipeline.Run)) *ORM_InsertRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pipeline.Run))
	})
	return _c
}

func (_c *ORM_InsertRun_Call) Return(_a0 error) *ORM_InsertRun_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_InsertRun_Call) RunAndReturn(run func(context.Context, *pipeline.Run) error) *ORM_InsertRun_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *ORM) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ORM_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type ORM_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *ORM_Expecter) Name() *ORM_Name_Call {
	return &ORM_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *ORM_Name_Call) Run(run func()) *ORM_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ORM_Name_Call) Return(_a0 string) *ORM_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_Name_Call) RunAndReturn(run func() string) *ORM_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with given fields:
func (_m *ORM) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type ORM_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *ORM_Expecter) Ready() *ORM_Ready_Call {
	return &ORM_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *ORM_Ready_Call) Run(run func()) *ORM_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ORM_Ready_Call) Return(_a0 error) *ORM_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_Ready_Call) RunAndReturn(run func() error) *ORM_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *ORM) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type ORM_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *ORM_Expecter) Start(_a0 interface{}) *ORM_Start_Call {
	return &ORM_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *ORM_Start_Call) Run(run func(_a0 context.Context)) *ORM_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ORM_Start_Call) Return(_a0 error) *ORM_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_Start_Call) RunAndReturn(run func(context.Context) error) *ORM_Start_Call {
	_c.Call.Return(run)
	return _c
}

// StoreRun provides a mock function with given fields: ctx, run
func (_m *ORM) StoreRun(ctx context.Context, run *pipeline.Run) (bool, error) {
	ret := _m.Called(ctx, run)

	if len(ret) == 0 {
		panic("no return value specified for StoreRun")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pipeline.Run) (bool, error)); ok {
		return rf(ctx, run)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pipeline.Run) bool); ok {
		r0 = rf(ctx, run)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pipeline.Run) error); ok {
		r1 = rf(ctx, run)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_StoreRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreRun'
type ORM_StoreRun_Call struct {
	*mock.Call
}

// StoreRun is a helper method to define mock.On call
//   - ctx context.Context
//   - run *pipeline.Run
func (_e *ORM_Expecter) StoreRun(ctx interface{}, run interface{}) *ORM_StoreRun_Call {
	return &ORM_StoreRun_Call{Call: _e.mock.On("StoreRun", ctx, run)}
}

func (_c *ORM_StoreRun_Call) Run(run func(ctx context.Context, run *pipeline.Run)) *ORM_StoreRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pipeline.Run))
	})
	return _c
}

func (_c *ORM_StoreRun_Call) Return(restart bool, err error) *ORM_StoreRun_Call {
	_c.Call.Return(restart, err)
	return _c
}

func (_c *ORM_StoreRun_Call) RunAndReturn(run func(context.Context, *pipeline.Run) (bool, error)) *ORM_StoreRun_Call {
	_c.Call.Return(run)
	return _c
}

// Transact provides a mock function with given fields: _a0, _a1
func (_m *ORM) Transact(_a0 context.Context, _a1 func(pipeline.ORM) error) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Transact")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(pipeline.ORM) error) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_Transact_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transact'
type ORM_Transact_Call struct {
	*mock.Call
}

// Transact is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 func(pipeline.ORM) error
func (_e *ORM_Expecter) Transact(_a0 interface{}, _a1 interface{}) *ORM_Transact_Call {
	return &ORM_Transact_Call{Call: _e.mock.On("Transact", _a0, _a1)}
}

func (_c *ORM_Transact_Call) Run(run func(_a0 context.Context, _a1 func(pipeline.ORM) error)) *ORM_Transact_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(pipeline.ORM) error))
	})
	return _c
}

func (_c *ORM_Transact_Call) Return(_a0 error) *ORM_Transact_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_Transact_Call) RunAndReturn(run func(context.Context, func(pipeline.ORM) error) error) *ORM_Transact_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTaskRunResult provides a mock function with given fields: ctx, taskID, result
func (_m *ORM) UpdateTaskRunResult(ctx context.Context, taskID uuid.UUID, result pipeline.Result) (pipeline.Run, bool, error) {
	ret := _m.Called(ctx, taskID, result)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTaskRunResult")
	}

	var r0 pipeline.Run
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, pipeline.Result) (pipeline.Run, bool, error)); ok {
		return rf(ctx, taskID, result)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, pipeline.Result) pipeline.Run); ok {
		r0 = rf(ctx, taskID, result)
	} else {
		r0 = ret.Get(0).(pipeline.Run)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, pipeline.Result) bool); ok {
		r1 = rf(ctx, taskID, result)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uuid.UUID, pipeline.Result) error); ok {
		r2 = rf(ctx, taskID, result)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ORM_UpdateTaskRunResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTaskRunResult'
type ORM_UpdateTaskRunResult_Call struct {
	*mock.Call
}

// UpdateTaskRunResult is a helper method to define mock.On call
//   - ctx context.Context
//   - taskID uuid.UUID
//   - result pipeline.Result
func (_e *ORM_Expecter) UpdateTaskRunResult(ctx interface{}, taskID interface{}, result interface{}) *ORM_UpdateTaskRunResult_Call {
	return &ORM_UpdateTaskRunResult_Call{Call: _e.mock.On("UpdateTaskRunResult", ctx, taskID, result)}
}

func (_c *ORM_UpdateTaskRunResult_Call) Run(run func(ctx context.Context, taskID uuid.UUID, result pipeline.Result)) *ORM_UpdateTaskRunResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(pipeline.Result))
	})
	return _c
}

func (_c *ORM_UpdateTaskRunResult_Call) Return(run pipeline.Run, start bool, err error) *ORM_UpdateTaskRunResult_Call {
	_c.Call.Return(run, start, err)
	return _c
}

func (_c *ORM_UpdateTaskRunResult_Call) RunAndReturn(run func(context.Context, uuid.UUID, pipeline.Result) (pipeline.Run, bool, error)) *ORM_UpdateTaskRunResult_Call {
	_c.Call.Return(run)
	return _c
}

// WithDataSource provides a mock function with given fields: _a0
func (_m *ORM) WithDataSource(_a0 sqlutil.DataSource) pipeline.ORM {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for WithDataSource")
	}

	var r0 pipeline.ORM
	if rf, ok := ret.Get(0).(func(sqlutil.DataSource) pipeline.ORM); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pipeline.ORM)
		}
	}

	return r0
}

// ORM_WithDataSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithDataSource'
type ORM_WithDataSource_Call struct {
	*mock.Call
}

// WithDataSource is a helper method to define mock.On call
//   - _a0 sqlutil.DataSource
func (_e *ORM_Expecter) WithDataSource(_a0 interface{}) *ORM_WithDataSource_Call {
	return &ORM_WithDataSource_Call{Call: _e.mock.On("WithDataSource", _a0)}
}

func (_c *ORM_WithDataSource_Call) Run(run func(_a0 sqlutil.DataSource)) *ORM_WithDataSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sqlutil.DataSource))
	})
	return _c
}

func (_c *ORM_WithDataSource_Call) Return(_a0 pipeline.ORM) *ORM_WithDataSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_WithDataSource_Call) RunAndReturn(run func(sqlutil.DataSource) pipeline.ORM) *ORM_WithDataSource_Call {
	_c.Call.Return(run)
	return _c
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewORM(t interface {
	mock.TestingT
	Cleanup(func())
}) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
