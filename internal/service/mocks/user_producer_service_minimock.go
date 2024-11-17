// Code generated by http://github.com/gojuno/minimock (v3.4.2). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/marinaaaniram/go-auth/internal/service.UserProducerService -o user_producer_service_minimock.go -n UserProducerServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/marinaaaniram/go-auth/internal/model"
)

// UserProducerServiceMock implements service.UserProducerService
type UserProducerServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcSendUser          func(ctx context.Context, user *model.User) (err error)
	inspectFuncSendUser   func(ctx context.Context, user *model.User)
	afterSendUserCounter  uint64
	beforeSendUserCounter uint64
	SendUserMock          mUserProducerServiceMockSendUser
}

// NewUserProducerServiceMock returns a mock for service.UserProducerService
func NewUserProducerServiceMock(t minimock.Tester) *UserProducerServiceMock {
	m := &UserProducerServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendUserMock = mUserProducerServiceMockSendUser{mock: m}
	m.SendUserMock.callArgs = []*UserProducerServiceMockSendUserParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserProducerServiceMockSendUser struct {
	optional           bool
	mock               *UserProducerServiceMock
	defaultExpectation *UserProducerServiceMockSendUserExpectation
	expectations       []*UserProducerServiceMockSendUserExpectation

	callArgs []*UserProducerServiceMockSendUserParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UserProducerServiceMockSendUserExpectation specifies expectation struct of the UserProducerService.SendUser
type UserProducerServiceMockSendUserExpectation struct {
	mock      *UserProducerServiceMock
	params    *UserProducerServiceMockSendUserParams
	paramPtrs *UserProducerServiceMockSendUserParamPtrs
	results   *UserProducerServiceMockSendUserResults
	Counter   uint64
}

// UserProducerServiceMockSendUserParams contains parameters of the UserProducerService.SendUser
type UserProducerServiceMockSendUserParams struct {
	ctx  context.Context
	user *model.User
}

// UserProducerServiceMockSendUserParamPtrs contains pointers to parameters of the UserProducerService.SendUser
type UserProducerServiceMockSendUserParamPtrs struct {
	ctx  *context.Context
	user **model.User
}

// UserProducerServiceMockSendUserResults contains results of the UserProducerService.SendUser
type UserProducerServiceMockSendUserResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSendUser *mUserProducerServiceMockSendUser) Optional() *mUserProducerServiceMockSendUser {
	mmSendUser.optional = true
	return mmSendUser
}

// Expect sets up expected params for UserProducerService.SendUser
func (mmSendUser *mUserProducerServiceMockSendUser) Expect(ctx context.Context, user *model.User) *mUserProducerServiceMockSendUser {
	if mmSendUser.mock.funcSendUser != nil {
		mmSendUser.mock.t.Fatalf("UserProducerServiceMock.SendUser mock is already set by Set")
	}

	if mmSendUser.defaultExpectation == nil {
		mmSendUser.defaultExpectation = &UserProducerServiceMockSendUserExpectation{}
	}

	if mmSendUser.defaultExpectation.paramPtrs != nil {
		mmSendUser.mock.t.Fatalf("UserProducerServiceMock.SendUser mock is already set by ExpectParams functions")
	}

	mmSendUser.defaultExpectation.params = &UserProducerServiceMockSendUserParams{ctx, user}
	for _, e := range mmSendUser.expectations {
		if minimock.Equal(e.params, mmSendUser.defaultExpectation.params) {
			mmSendUser.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendUser.defaultExpectation.params)
		}
	}

	return mmSendUser
}

// ExpectCtxParam1 sets up expected param ctx for UserProducerService.SendUser
func (mmSendUser *mUserProducerServiceMockSendUser) ExpectCtxParam1(ctx context.Context) *mUserProducerServiceMockSendUser {
	if mmSendUser.mock.funcSendUser != nil {
		mmSendUser.mock.t.Fatalf("UserProducerServiceMock.SendUser mock is already set by Set")
	}

	if mmSendUser.defaultExpectation == nil {
		mmSendUser.defaultExpectation = &UserProducerServiceMockSendUserExpectation{}
	}

	if mmSendUser.defaultExpectation.params != nil {
		mmSendUser.mock.t.Fatalf("UserProducerServiceMock.SendUser mock is already set by Expect")
	}

	if mmSendUser.defaultExpectation.paramPtrs == nil {
		mmSendUser.defaultExpectation.paramPtrs = &UserProducerServiceMockSendUserParamPtrs{}
	}
	mmSendUser.defaultExpectation.paramPtrs.ctx = &ctx

	return mmSendUser
}

// ExpectUserParam2 sets up expected param user for UserProducerService.SendUser
func (mmSendUser *mUserProducerServiceMockSendUser) ExpectUserParam2(user *model.User) *mUserProducerServiceMockSendUser {
	if mmSendUser.mock.funcSendUser != nil {
		mmSendUser.mock.t.Fatalf("UserProducerServiceMock.SendUser mock is already set by Set")
	}

	if mmSendUser.defaultExpectation == nil {
		mmSendUser.defaultExpectation = &UserProducerServiceMockSendUserExpectation{}
	}

	if mmSendUser.defaultExpectation.params != nil {
		mmSendUser.mock.t.Fatalf("UserProducerServiceMock.SendUser mock is already set by Expect")
	}

	if mmSendUser.defaultExpectation.paramPtrs == nil {
		mmSendUser.defaultExpectation.paramPtrs = &UserProducerServiceMockSendUserParamPtrs{}
	}
	mmSendUser.defaultExpectation.paramPtrs.user = &user

	return mmSendUser
}

// Inspect accepts an inspector function that has same arguments as the UserProducerService.SendUser
func (mmSendUser *mUserProducerServiceMockSendUser) Inspect(f func(ctx context.Context, user *model.User)) *mUserProducerServiceMockSendUser {
	if mmSendUser.mock.inspectFuncSendUser != nil {
		mmSendUser.mock.t.Fatalf("Inspect function is already set for UserProducerServiceMock.SendUser")
	}

	mmSendUser.mock.inspectFuncSendUser = f

	return mmSendUser
}

// Return sets up results that will be returned by UserProducerService.SendUser
func (mmSendUser *mUserProducerServiceMockSendUser) Return(err error) *UserProducerServiceMock {
	if mmSendUser.mock.funcSendUser != nil {
		mmSendUser.mock.t.Fatalf("UserProducerServiceMock.SendUser mock is already set by Set")
	}

	if mmSendUser.defaultExpectation == nil {
		mmSendUser.defaultExpectation = &UserProducerServiceMockSendUserExpectation{mock: mmSendUser.mock}
	}
	mmSendUser.defaultExpectation.results = &UserProducerServiceMockSendUserResults{err}
	return mmSendUser.mock
}

// Set uses given function f to mock the UserProducerService.SendUser method
func (mmSendUser *mUserProducerServiceMockSendUser) Set(f func(ctx context.Context, user *model.User) (err error)) *UserProducerServiceMock {
	if mmSendUser.defaultExpectation != nil {
		mmSendUser.mock.t.Fatalf("Default expectation is already set for the UserProducerService.SendUser method")
	}

	if len(mmSendUser.expectations) > 0 {
		mmSendUser.mock.t.Fatalf("Some expectations are already set for the UserProducerService.SendUser method")
	}

	mmSendUser.mock.funcSendUser = f
	return mmSendUser.mock
}

// When sets expectation for the UserProducerService.SendUser which will trigger the result defined by the following
// Then helper
func (mmSendUser *mUserProducerServiceMockSendUser) When(ctx context.Context, user *model.User) *UserProducerServiceMockSendUserExpectation {
	if mmSendUser.mock.funcSendUser != nil {
		mmSendUser.mock.t.Fatalf("UserProducerServiceMock.SendUser mock is already set by Set")
	}

	expectation := &UserProducerServiceMockSendUserExpectation{
		mock:   mmSendUser.mock,
		params: &UserProducerServiceMockSendUserParams{ctx, user},
	}
	mmSendUser.expectations = append(mmSendUser.expectations, expectation)
	return expectation
}

// Then sets up UserProducerService.SendUser return parameters for the expectation previously defined by the When method
func (e *UserProducerServiceMockSendUserExpectation) Then(err error) *UserProducerServiceMock {
	e.results = &UserProducerServiceMockSendUserResults{err}
	return e.mock
}

// Times sets number of times UserProducerService.SendUser should be invoked
func (mmSendUser *mUserProducerServiceMockSendUser) Times(n uint64) *mUserProducerServiceMockSendUser {
	if n == 0 {
		mmSendUser.mock.t.Fatalf("Times of UserProducerServiceMock.SendUser mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSendUser.expectedInvocations, n)
	return mmSendUser
}

func (mmSendUser *mUserProducerServiceMockSendUser) invocationsDone() bool {
	if len(mmSendUser.expectations) == 0 && mmSendUser.defaultExpectation == nil && mmSendUser.mock.funcSendUser == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSendUser.mock.afterSendUserCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSendUser.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// SendUser implements service.UserProducerService
func (mmSendUser *UserProducerServiceMock) SendUser(ctx context.Context, user *model.User) (err error) {
	mm_atomic.AddUint64(&mmSendUser.beforeSendUserCounter, 1)
	defer mm_atomic.AddUint64(&mmSendUser.afterSendUserCounter, 1)

	if mmSendUser.inspectFuncSendUser != nil {
		mmSendUser.inspectFuncSendUser(ctx, user)
	}

	mm_params := UserProducerServiceMockSendUserParams{ctx, user}

	// Record call args
	mmSendUser.SendUserMock.mutex.Lock()
	mmSendUser.SendUserMock.callArgs = append(mmSendUser.SendUserMock.callArgs, &mm_params)
	mmSendUser.SendUserMock.mutex.Unlock()

	for _, e := range mmSendUser.SendUserMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSendUser.SendUserMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendUser.SendUserMock.defaultExpectation.Counter, 1)
		mm_want := mmSendUser.SendUserMock.defaultExpectation.params
		mm_want_ptrs := mmSendUser.SendUserMock.defaultExpectation.paramPtrs

		mm_got := UserProducerServiceMockSendUserParams{ctx, user}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmSendUser.t.Errorf("UserProducerServiceMock.SendUser got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.user != nil && !minimock.Equal(*mm_want_ptrs.user, mm_got.user) {
				mmSendUser.t.Errorf("UserProducerServiceMock.SendUser got unexpected parameter user, want: %#v, got: %#v%s\n", *mm_want_ptrs.user, mm_got.user, minimock.Diff(*mm_want_ptrs.user, mm_got.user))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendUser.t.Errorf("UserProducerServiceMock.SendUser got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendUser.SendUserMock.defaultExpectation.results
		if mm_results == nil {
			mmSendUser.t.Fatal("No results are set for the UserProducerServiceMock.SendUser")
		}
		return (*mm_results).err
	}
	if mmSendUser.funcSendUser != nil {
		return mmSendUser.funcSendUser(ctx, user)
	}
	mmSendUser.t.Fatalf("Unexpected call to UserProducerServiceMock.SendUser. %v %v", ctx, user)
	return
}

// SendUserAfterCounter returns a count of finished UserProducerServiceMock.SendUser invocations
func (mmSendUser *UserProducerServiceMock) SendUserAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendUser.afterSendUserCounter)
}

// SendUserBeforeCounter returns a count of UserProducerServiceMock.SendUser invocations
func (mmSendUser *UserProducerServiceMock) SendUserBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendUser.beforeSendUserCounter)
}

// Calls returns a list of arguments used in each call to UserProducerServiceMock.SendUser.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendUser *mUserProducerServiceMockSendUser) Calls() []*UserProducerServiceMockSendUserParams {
	mmSendUser.mutex.RLock()

	argCopy := make([]*UserProducerServiceMockSendUserParams, len(mmSendUser.callArgs))
	copy(argCopy, mmSendUser.callArgs)

	mmSendUser.mutex.RUnlock()

	return argCopy
}

// MinimockSendUserDone returns true if the count of the SendUser invocations corresponds
// the number of defined expectations
func (m *UserProducerServiceMock) MinimockSendUserDone() bool {
	if m.SendUserMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SendUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SendUserMock.invocationsDone()
}

// MinimockSendUserInspect logs each unmet expectation
func (m *UserProducerServiceMock) MinimockSendUserInspect() {
	for _, e := range m.SendUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserProducerServiceMock.SendUser with params: %#v", *e.params)
		}
	}

	afterSendUserCounter := mm_atomic.LoadUint64(&m.afterSendUserCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SendUserMock.defaultExpectation != nil && afterSendUserCounter < 1 {
		if m.SendUserMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserProducerServiceMock.SendUser")
		} else {
			m.t.Errorf("Expected call to UserProducerServiceMock.SendUser with params: %#v", *m.SendUserMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendUser != nil && afterSendUserCounter < 1 {
		m.t.Error("Expected call to UserProducerServiceMock.SendUser")
	}

	if !m.SendUserMock.invocationsDone() && afterSendUserCounter > 0 {
		m.t.Errorf("Expected %d calls to UserProducerServiceMock.SendUser but found %d calls",
			mm_atomic.LoadUint64(&m.SendUserMock.expectedInvocations), afterSendUserCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserProducerServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockSendUserInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserProducerServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UserProducerServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendUserDone()
}
