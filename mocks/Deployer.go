// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import service "github.com/Optum/dce-cli/pkg/service"

// Deployer is an autogenerated mock type for the Deployer type
type Deployer struct {
	mock.Mock
}

// Deploy provides a mock function with given fields: deployLocal, overrides
func (_m *Deployer) Deploy(deployLocal string, overrides *service.DeployOverrides) {
	_m.Called(deployLocal, overrides)
}