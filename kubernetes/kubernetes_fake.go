package kubernetes

import "sync"

// JobsSpecFake is a mock for the Kubernetes struct defined on the
// kubernetes.go file, JobsSpecFake has been created to mock Kubernetes
// cluster calls for testing purposes.
type JobsSpecFake struct {
	CreateFunc  func(string) error
	CreateCalls int

	DeleteFunc  func(string) error
	DeleteCalls int

	sync.Mutex
}

// Create is the mocked implementation for the Kubernetes Jobs Create func.
func (s *JobsSpecFake) Create(name string) error {
	s.Lock()
	defer s.Unlock()
	s.CreateCalls++
	return s.CreateFunc(name)
}

// DeleteJob is the mocked implementation for the Kubernetes Jobs Delete func.
func (s *JobsSpecFake) Delete(name string) error {
	s.Lock()
	defer s.Unlock()
	s.DeleteCalls++
	return s.DeleteFunc(name)
}
