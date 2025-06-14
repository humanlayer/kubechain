//go:build !ignore_autogenerated

/*
Copyright 2025 the Agent Control Plane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeySource) DeepCopyInto(out *APIKeySource) {
	*out = *in
	out.SecretKeyRef = in.SecretKeyRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeySource.
func (in *APIKeySource) DeepCopy() *APIKeySource {
	if in == nil {
		return nil
	}
	out := new(APIKeySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Agent) DeepCopyInto(out *Agent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Agent.
func (in *Agent) DeepCopy() *Agent {
	if in == nil {
		return nil
	}
	out := new(Agent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Agent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentList) DeepCopyInto(out *AgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Agent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentList.
func (in *AgentList) DeepCopy() *AgentList {
	if in == nil {
		return nil
	}
	out := new(AgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpec) DeepCopyInto(out *AgentSpec) {
	*out = *in
	out.LLMRef = in.LLMRef
	if in.MCPServers != nil {
		in, out := &in.MCPServers, &out.MCPServers
		*out = make([]LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.HumanContactChannels != nil {
		in, out := &in.HumanContactChannels, &out.HumanContactChannels
		*out = make([]LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.SubAgents != nil {
		in, out := &in.SubAgents, &out.SubAgents
		*out = make([]LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpec.
func (in *AgentSpec) DeepCopy() *AgentSpec {
	if in == nil {
		return nil
	}
	out := new(AgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentStatus) DeepCopyInto(out *AgentStatus) {
	*out = *in
	if in.ValidMCPServers != nil {
		in, out := &in.ValidMCPServers, &out.ValidMCPServers
		*out = make([]ResolvedMCPServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ValidHumanContactChannels != nil {
		in, out := &in.ValidHumanContactChannels, &out.ValidHumanContactChannels
		*out = make([]ResolvedContactChannel, len(*in))
		copy(*out, *in)
	}
	if in.ValidSubAgents != nil {
		in, out := &in.ValidSubAgents, &out.ValidSubAgents
		*out = make([]ResolvedSubAgent, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentStatus.
func (in *AgentStatus) DeepCopy() *AgentStatus {
	if in == nil {
		return nil
	}
	out := new(AgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnthropicConfig) DeepCopyInto(out *AnthropicConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnthropicConfig.
func (in *AnthropicConfig) DeepCopy() *AnthropicConfig {
	if in == nil {
		return nil
	}
	out := new(AnthropicConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseConfig) DeepCopyInto(out *BaseConfig) {
	*out = *in
	if in.MaxTokens != nil {
		in, out := &in.MaxTokens, &out.MaxTokens
		*out = new(int)
		**out = **in
	}
	if in.TopK != nil {
		in, out := &in.TopK, &out.TopK
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseConfig.
func (in *BaseConfig) DeepCopy() *BaseConfig {
	if in == nil {
		return nil
	}
	out := new(BaseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactChannel) DeepCopyInto(out *ContactChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactChannel.
func (in *ContactChannel) DeepCopy() *ContactChannel {
	if in == nil {
		return nil
	}
	out := new(ContactChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContactChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactChannelList) DeepCopyInto(out *ContactChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContactChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactChannelList.
func (in *ContactChannelList) DeepCopy() *ContactChannelList {
	if in == nil {
		return nil
	}
	out := new(ContactChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContactChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactChannelSpec) DeepCopyInto(out *ContactChannelSpec) {
	*out = *in
	if in.APIKeyFrom != nil {
		in, out := &in.APIKeyFrom, &out.APIKeyFrom
		*out = new(APIKeySource)
		**out = **in
	}
	if in.ChannelAPIKeyFrom != nil {
		in, out := &in.ChannelAPIKeyFrom, &out.ChannelAPIKeyFrom
		*out = new(APIKeySource)
		**out = **in
	}
	if in.Slack != nil {
		in, out := &in.Slack, &out.Slack
		*out = new(SlackChannelConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(EmailChannelConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactChannelSpec.
func (in *ContactChannelSpec) DeepCopy() *ContactChannelSpec {
	if in == nil {
		return nil
	}
	out := new(ContactChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactChannelStatus) DeepCopyInto(out *ContactChannelStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactChannelStatus.
func (in *ContactChannelStatus) DeepCopy() *ContactChannelStatus {
	if in == nil {
		return nil
	}
	out := new(ContactChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailChannelConfig) DeepCopyInto(out *EmailChannelConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailChannelConfig.
func (in *EmailChannelConfig) DeepCopy() *EmailChannelConfig {
	if in == nil {
		return nil
	}
	out := new(EmailChannelConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(EnvVarSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVarSource) DeepCopyInto(out *EnvVarSource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(SecretKeyRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVarSource.
func (in *EnvVarSource) DeepCopy() *EnvVarSource {
	if in == nil {
		return nil
	}
	out := new(EnvVarSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleConfig) DeepCopyInto(out *GoogleConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleConfig.
func (in *GoogleConfig) DeepCopy() *GoogleConfig {
	if in == nil {
		return nil
	}
	out := new(GoogleConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLM) DeepCopyInto(out *LLM) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLM.
func (in *LLM) DeepCopy() *LLM {
	if in == nil {
		return nil
	}
	out := new(LLM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LLM) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMList) DeepCopyInto(out *LLMList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LLM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMList.
func (in *LLMList) DeepCopy() *LLMList {
	if in == nil {
		return nil
	}
	out := new(LLMList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LLMList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMSpec) DeepCopyInto(out *LLMSpec) {
	*out = *in
	if in.APIKeyFrom != nil {
		in, out := &in.APIKeyFrom, &out.APIKeyFrom
		*out = new(APIKeySource)
		**out = **in
	}
	in.Parameters.DeepCopyInto(&out.Parameters)
	if in.OpenAI != nil {
		in, out := &in.OpenAI, &out.OpenAI
		*out = new(OpenAIConfig)
		**out = **in
	}
	if in.Anthropic != nil {
		in, out := &in.Anthropic, &out.Anthropic
		*out = new(AnthropicConfig)
		**out = **in
	}
	if in.Vertex != nil {
		in, out := &in.Vertex, &out.Vertex
		*out = new(VertexConfig)
		**out = **in
	}
	if in.Mistral != nil {
		in, out := &in.Mistral, &out.Mistral
		*out = new(MistralConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Google != nil {
		in, out := &in.Google, &out.Google
		*out = new(GoogleConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMSpec.
func (in *LLMSpec) DeepCopy() *LLMSpec {
	if in == nil {
		return nil
	}
	out := new(LLMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMStatus) DeepCopyInto(out *LLMStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMStatus.
func (in *LLMStatus) DeepCopy() *LLMStatus {
	if in == nil {
		return nil
	}
	out := new(LLMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MCPServer) DeepCopyInto(out *MCPServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MCPServer.
func (in *MCPServer) DeepCopy() *MCPServer {
	if in == nil {
		return nil
	}
	out := new(MCPServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MCPServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MCPServerList) DeepCopyInto(out *MCPServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MCPServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MCPServerList.
func (in *MCPServerList) DeepCopy() *MCPServerList {
	if in == nil {
		return nil
	}
	out := new(MCPServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MCPServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MCPServerSpec) DeepCopyInto(out *MCPServerSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ApprovalContactChannel != nil {
		in, out := &in.ApprovalContactChannel, &out.ApprovalContactChannel
		*out = new(LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MCPServerSpec.
func (in *MCPServerSpec) DeepCopy() *MCPServerSpec {
	if in == nil {
		return nil
	}
	out := new(MCPServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MCPServerStatus) DeepCopyInto(out *MCPServerStatus) {
	*out = *in
	if in.Tools != nil {
		in, out := &in.Tools, &out.Tools
		*out = make([]MCPTool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MCPServerStatus.
func (in *MCPServerStatus) DeepCopy() *MCPServerStatus {
	if in == nil {
		return nil
	}
	out := new(MCPServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MCPTool) DeepCopyInto(out *MCPTool) {
	*out = *in
	in.InputSchema.DeepCopyInto(&out.InputSchema)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MCPTool.
func (in *MCPTool) DeepCopy() *MCPTool {
	if in == nil {
		return nil
	}
	out := new(MCPTool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Message) DeepCopyInto(out *Message) {
	*out = *in
	if in.ToolCalls != nil {
		in, out := &in.ToolCalls, &out.ToolCalls
		*out = make([]MessageToolCall, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Message.
func (in *Message) DeepCopy() *Message {
	if in == nil {
		return nil
	}
	out := new(Message)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessageToolCall) DeepCopyInto(out *MessageToolCall) {
	*out = *in
	out.Function = in.Function
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessageToolCall.
func (in *MessageToolCall) DeepCopy() *MessageToolCall {
	if in == nil {
		return nil
	}
	out := new(MessageToolCall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MistralConfig) DeepCopyInto(out *MistralConfig) {
	*out = *in
	if in.MaxRetries != nil {
		in, out := &in.MaxRetries, &out.MaxRetries
		*out = new(int)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.RandomSeed != nil {
		in, out := &in.RandomSeed, &out.RandomSeed
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MistralConfig.
func (in *MistralConfig) DeepCopy() *MistralConfig {
	if in == nil {
		return nil
	}
	out := new(MistralConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenAIConfig) DeepCopyInto(out *OpenAIConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenAIConfig.
func (in *OpenAIConfig) DeepCopy() *OpenAIConfig {
	if in == nil {
		return nil
	}
	out := new(OpenAIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfig) DeepCopyInto(out *ProviderConfig) {
	*out = *in
	if in.OpenAIConfig != nil {
		in, out := &in.OpenAIConfig, &out.OpenAIConfig
		*out = new(OpenAIConfig)
		**out = **in
	}
	if in.AnthropicConfig != nil {
		in, out := &in.AnthropicConfig, &out.AnthropicConfig
		*out = new(AnthropicConfig)
		**out = **in
	}
	if in.VertexConfig != nil {
		in, out := &in.VertexConfig, &out.VertexConfig
		*out = new(VertexConfig)
		**out = **in
	}
	if in.MistralConfig != nil {
		in, out := &in.MistralConfig, &out.MistralConfig
		*out = new(MistralConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.GoogleConfig != nil {
		in, out := &in.GoogleConfig, &out.GoogleConfig
		*out = new(GoogleConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfig.
func (in *ProviderConfig) DeepCopy() *ProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedContactChannel) DeepCopyInto(out *ResolvedContactChannel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedContactChannel.
func (in *ResolvedContactChannel) DeepCopy() *ResolvedContactChannel {
	if in == nil {
		return nil
	}
	out := new(ResolvedContactChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedMCPServer) DeepCopyInto(out *ResolvedMCPServer) {
	*out = *in
	if in.Tools != nil {
		in, out := &in.Tools, &out.Tools
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedMCPServer.
func (in *ResolvedMCPServer) DeepCopy() *ResolvedMCPServer {
	if in == nil {
		return nil
	}
	out := new(ResolvedMCPServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedSubAgent) DeepCopyInto(out *ResolvedSubAgent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedSubAgent.
func (in *ResolvedSubAgent) DeepCopy() *ResolvedSubAgent {
	if in == nil {
		return nil
	}
	out := new(ResolvedSubAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ResourceList) DeepCopyInto(out *ResourceList) {
	{
		in := &in
		*out = make(ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceList.
func (in ResourceList) DeepCopy() ResourceList {
	if in == nil {
		return nil
	}
	out := new(ResourceList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyRef) DeepCopyInto(out *SecretKeyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyRef.
func (in *SecretKeyRef) DeepCopy() *SecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(SecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackChannelConfig) DeepCopyInto(out *SlackChannelConfig) {
	*out = *in
	if in.AllowedResponderIDs != nil {
		in, out := &in.AllowedResponderIDs, &out.AllowedResponderIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackChannelConfig.
func (in *SlackChannelConfig) DeepCopy() *SlackChannelConfig {
	if in == nil {
		return nil
	}
	out := new(SlackChannelConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpanContext) DeepCopyInto(out *SpanContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpanContext.
func (in *SpanContext) DeepCopy() *SpanContext {
	if in == nil {
		return nil
	}
	out := new(SpanContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Task) DeepCopyInto(out *Task) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Task.
func (in *Task) DeepCopy() *Task {
	if in == nil {
		return nil
	}
	out := new(Task)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Task) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskList) DeepCopyInto(out *TaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Task, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskList.
func (in *TaskList) DeepCopy() *TaskList {
	if in == nil {
		return nil
	}
	out := new(TaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSpec) DeepCopyInto(out *TaskSpec) {
	*out = *in
	out.AgentRef = in.AgentRef
	if in.ContextWindow != nil {
		in, out := &in.ContextWindow, &out.ContextWindow
		*out = make([]Message, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ContactChannelRef != nil {
		in, out := &in.ContactChannelRef, &out.ContactChannelRef
		*out = new(LocalObjectReference)
		**out = **in
	}
	if in.ChannelTokenFrom != nil {
		in, out := &in.ChannelTokenFrom, &out.ChannelTokenFrom
		*out = new(SecretKeyRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSpec.
func (in *TaskSpec) DeepCopy() *TaskSpec {
	if in == nil {
		return nil
	}
	out := new(TaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStatus) DeepCopyInto(out *TaskStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.ContextWindow != nil {
		in, out := &in.ContextWindow, &out.ContextWindow
		*out = make([]Message, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SpanContext != nil {
		in, out := &in.SpanContext, &out.SpanContext
		*out = new(SpanContext)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStatus.
func (in *TaskStatus) DeepCopy() *TaskStatus {
	if in == nil {
		return nil
	}
	out := new(TaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolCall) DeepCopyInto(out *ToolCall) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolCall.
func (in *ToolCall) DeepCopy() *ToolCall {
	if in == nil {
		return nil
	}
	out := new(ToolCall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ToolCall) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolCallFunction) DeepCopyInto(out *ToolCallFunction) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolCallFunction.
func (in *ToolCallFunction) DeepCopy() *ToolCallFunction {
	if in == nil {
		return nil
	}
	out := new(ToolCallFunction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolCallList) DeepCopyInto(out *ToolCallList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ToolCall, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolCallList.
func (in *ToolCallList) DeepCopy() *ToolCallList {
	if in == nil {
		return nil
	}
	out := new(ToolCallList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ToolCallList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolCallSpec) DeepCopyInto(out *ToolCallSpec) {
	*out = *in
	out.TaskRef = in.TaskRef
	out.ToolRef = in.ToolRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolCallSpec.
func (in *ToolCallSpec) DeepCopy() *ToolCallSpec {
	if in == nil {
		return nil
	}
	out := new(ToolCallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolCallStatus) DeepCopyInto(out *ToolCallStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.SpanContext != nil {
		in, out := &in.SpanContext, &out.SpanContext
		*out = new(SpanContext)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolCallStatus.
func (in *ToolCallStatus) DeepCopy() *ToolCallStatus {
	if in == nil {
		return nil
	}
	out := new(ToolCallStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VertexConfig) DeepCopyInto(out *VertexConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VertexConfig.
func (in *VertexConfig) DeepCopy() *VertexConfig {
	if in == nil {
		return nil
	}
	out := new(VertexConfig)
	in.DeepCopyInto(out)
	return out
}
