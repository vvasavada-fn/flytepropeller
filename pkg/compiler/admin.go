package compiler

import (
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
)

// Should we call this something else? It can be used more generally if ever necessary.
// This object is meant to satisfy github.com/lyft/flytepropeller/pkg/compiler/common.InterfaceProvider
// This file is pretty much copied from Admin,
// https://github.com/lyft/flyteadmin/blob/1acce744b8c7839ab77a0eb1ed922905af15baa5/pkg/workflowengine/impl/interface_provider.go
// but that implementation relies on the internal Admin Gorm model. We should consider deprecating that one in favor
// of this one as Admin already has a dependency on the Propeller compiler
type LaunchPlanInterfaceProvider struct {
	expectedInputs  core.ParameterMap
	expectedOutputs core.VariableMap
	identifier      *core.Identifier
}

func (p *LaunchPlanInterfaceProvider) GetID() *core.Identifier {
	return p.identifier
}
func (p *LaunchPlanInterfaceProvider) GetExpectedInputs() *core.ParameterMap {
	return &p.expectedInputs

}
func (p *LaunchPlanInterfaceProvider) GetExpectedOutputs() *core.VariableMap {
	return &p.expectedOutputs
}

func NewLaunchPlanInterfaceProvider(launchPlan admin.LaunchPlan) *LaunchPlanInterfaceProvider {
	return &LaunchPlanInterfaceProvider{
		expectedInputs:  *launchPlan.Closure.ExpectedInputs,
		expectedOutputs: *launchPlan.Closure.ExpectedOutputs,
		identifier:      launchPlan.Id,
	}
}
