// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: proto/job.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobStatus int32

const (
	JobStatus_JOB_STATUS_UNKNOWN   JobStatus = 0
	JobStatus_JOB_STATUS_PENDING   JobStatus = 1
	JobStatus_JOB_STATUS_SCHEDULED JobStatus = 2
	JobStatus_JOB_STATUS_RUNNING   JobStatus = 3
	JobStatus_JOB_STATUS_SUCCEEDED JobStatus = 4
	JobStatus_JOB_STATUS_FAILED    JobStatus = 5
)

// Enum value maps for JobStatus.
var (
	JobStatus_name = map[int32]string{
		0: "JOB_STATUS_UNKNOWN",
		1: "JOB_STATUS_PENDING",
		2: "JOB_STATUS_SCHEDULED",
		3: "JOB_STATUS_RUNNING",
		4: "JOB_STATUS_SUCCEEDED",
		5: "JOB_STATUS_FAILED",
	}
	JobStatus_value = map[string]int32{
		"JOB_STATUS_UNKNOWN":   0,
		"JOB_STATUS_PENDING":   1,
		"JOB_STATUS_SCHEDULED": 2,
		"JOB_STATUS_RUNNING":   3,
		"JOB_STATUS_SUCCEEDED": 4,
		"JOB_STATUS_FAILED":    5,
	}
)

func (x JobStatus) Enum() *JobStatus {
	p := new(JobStatus)
	*p = x
	return p
}

func (x JobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_job_proto_enumTypes[0].Descriptor()
}

func (JobStatus) Type() protoreflect.EnumType {
	return &file_proto_job_proto_enumTypes[0]
}

func (x JobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobStatus.Descriptor instead.
func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{0}
}

type JobType int32

const (
	JobType_JOB_TYPE_EMAIL JobType = 0
	JobType_JOB_TYPE_SMS   JobType = 1
)

// Enum value maps for JobType.
var (
	JobType_name = map[int32]string{
		0: "JOB_TYPE_EMAIL",
		1: "JOB_TYPE_SMS",
	}
	JobType_value = map[string]int32{
		"JOB_TYPE_EMAIL": 0,
		"JOB_TYPE_SMS":   1,
	}
)

func (x JobType) Enum() *JobType {
	p := new(JobType)
	*p = x
	return p
}

func (x JobType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_job_proto_enumTypes[1].Descriptor()
}

func (JobType) Type() protoreflect.EnumType {
	return &file_proto_job_proto_enumTypes[1]
}

func (x JobType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobType.Descriptor instead.
func (JobType) EnumDescriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{1}
}

type CreateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ScheduleTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	JobType      JobType                `protobuf:"varint,4,opt,name=job_type,json=jobType,proto3,enum=proto.JobType" json:"job_type,omitempty"`
	JobData      string                 `protobuf:"bytes,5,opt,name=job_data,json=jobData,proto3" json:"job_data,omitempty"`
}

func (x *CreateJobRequest) Reset() {
	*x = CreateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobRequest) ProtoMessage() {}

func (x *CreateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobRequest.ProtoReflect.Descriptor instead.
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateJobRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateJobRequest) GetScheduleTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduleTime
	}
	return nil
}

func (x *CreateJobRequest) GetJobType() JobType {
	if x != nil {
		return x.JobType
	}
	return JobType_JOB_TYPE_EMAIL
}

func (x *CreateJobRequest) GetJobData() string {
	if x != nil {
		return x.JobData
	}
	return ""
}

type CreateJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateJobResponse) Reset() {
	*x = CreateJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobResponse) ProtoMessage() {}

func (x *CreateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobResponse.ProtoReflect.Descriptor instead.
func (*CreateJobResponse) Descriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{2}
}

func (x *GetJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListJobsRequest) Reset() {
	*x = ListJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsRequest.ProtoReflect.Descriptor instead.
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{3}
}

func (x *ListJobsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListJobsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64  `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	TotalPages int64  `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	Page       int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Size       int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	HasMore    bool   `protobuf:"varint,5,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	Jobs       []*Job `protobuf:"bytes,6,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *ListJobsResponse) Reset() {
	*x = ListJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsResponse.ProtoReflect.Descriptor instead.
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{4}
}

func (x *ListJobsResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListJobsResponse) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *ListJobsResponse) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListJobsResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListJobsResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *ListJobsResponse) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type UpdateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ScheduleTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	JobType      JobType                `protobuf:"varint,5,opt,name=job_type,json=jobType,proto3,enum=proto.JobType" json:"job_type,omitempty"`
	JobData      string                 `protobuf:"bytes,6,opt,name=job_data,json=jobData,proto3" json:"job_data,omitempty"`
}

func (x *UpdateJobRequest) Reset() {
	*x = UpdateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobRequest) ProtoMessage() {}

func (x *UpdateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateJobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateJobRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateJobRequest) GetScheduleTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduleTime
	}
	return nil
}

func (x *UpdateJobRequest) GetJobType() JobType {
	if x != nil {
		return x.JobType
	}
	return JobType_JOB_TYPE_EMAIL
}

func (x *UpdateJobRequest) GetJobData() string {
	if x != nil {
		return x.JobData
	}
	return ""
}

type UpdateJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateJobResponse) Reset() {
	*x = UpdateJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobResponse) ProtoMessage() {}

func (x *UpdateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobResponse.ProtoReflect.Descriptor instead.
func (*UpdateJobResponse) Descriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateJobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteJobRequest) Reset() {
	*x = DeleteJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobRequest) ProtoMessage() {}

func (x *DeleteJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteJobResponse) Reset() {
	*x = DeleteJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobResponse) ProtoMessage() {}

func (x *DeleteJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobResponse.ProtoReflect.Descriptor instead.
func (*DeleteJobResponse) Descriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteJobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ScheduleTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	CreatedTime  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdatedTime  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	Status       JobStatus              `protobuf:"varint,7,opt,name=status,proto3,enum=proto.JobStatus" json:"status,omitempty"`
	JobType      JobType                `protobuf:"varint,8,opt,name=job_type,json=jobType,proto3,enum=proto.JobType" json:"job_type,omitempty"`
	JobData      string                 `protobuf:"bytes,9,opt,name=job_data,json=jobData,proto3" json:"job_data,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_proto_job_proto_rawDescGZIP(), []int{9}
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Job) GetScheduleTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduleTime
	}
	return nil
}

func (x *Job) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Job) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Job) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_JOB_STATUS_UNKNOWN
}

func (x *Job) GetJobType() JobType {
	if x != nil {
		return x.JobType
	}
	return JobType_JOB_TYPE_EMAIL
}

func (x *Job) GetJobData() string {
	if x != nil {
		return x.JobData
	}
	return ""
}

var File_proto_job_proto protoreflect.FileDescriptor

var file_proto_job_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x39, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xb7, 0x01, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x68,
	0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68,
	0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62,
	0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xfa, 0x02, 0x0a, 0x03, 0x4a,
	0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x29, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6a, 0x6f, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x2a, 0x9e, 0x01, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x16, 0x0a, 0x12, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x4a, 0x4f, 0x42, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10,
	0x04, 0x12, 0x15, 0x0a, 0x11, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x2a, 0x2f, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x4a, 0x4f, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4a, 0x4f, 0x42, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x4d, 0x53, 0x10, 0x01, 0x32, 0xc0, 0x02, 0x0a, 0x0b, 0x4a, 0x6f,
	0x62, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_job_proto_rawDescOnce sync.Once
	file_proto_job_proto_rawDescData = file_proto_job_proto_rawDesc
)

func file_proto_job_proto_rawDescGZIP() []byte {
	file_proto_job_proto_rawDescOnce.Do(func() {
		file_proto_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_job_proto_rawDescData)
	})
	return file_proto_job_proto_rawDescData
}

var file_proto_job_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_job_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_job_proto_goTypes = []interface{}{
	(JobStatus)(0),                // 0: proto.JobStatus
	(JobType)(0),                  // 1: proto.JobType
	(*CreateJobRequest)(nil),      // 2: proto.CreateJobRequest
	(*CreateJobResponse)(nil),     // 3: proto.CreateJobResponse
	(*GetJobRequest)(nil),         // 4: proto.GetJobRequest
	(*ListJobsRequest)(nil),       // 5: proto.ListJobsRequest
	(*ListJobsResponse)(nil),      // 6: proto.ListJobsResponse
	(*UpdateJobRequest)(nil),      // 7: proto.UpdateJobRequest
	(*UpdateJobResponse)(nil),     // 8: proto.UpdateJobResponse
	(*DeleteJobRequest)(nil),      // 9: proto.DeleteJobRequest
	(*DeleteJobResponse)(nil),     // 10: proto.DeleteJobResponse
	(*Job)(nil),                   // 11: proto.Job
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_proto_job_proto_depIdxs = []int32{
	12, // 0: proto.CreateJobRequest.schedule_time:type_name -> google.protobuf.Timestamp
	1,  // 1: proto.CreateJobRequest.job_type:type_name -> proto.JobType
	11, // 2: proto.ListJobsResponse.jobs:type_name -> proto.Job
	12, // 3: proto.UpdateJobRequest.schedule_time:type_name -> google.protobuf.Timestamp
	1,  // 4: proto.UpdateJobRequest.job_type:type_name -> proto.JobType
	12, // 5: proto.Job.schedule_time:type_name -> google.protobuf.Timestamp
	12, // 6: proto.Job.created_time:type_name -> google.protobuf.Timestamp
	12, // 7: proto.Job.updated_time:type_name -> google.protobuf.Timestamp
	0,  // 8: proto.Job.status:type_name -> proto.JobStatus
	1,  // 9: proto.Job.job_type:type_name -> proto.JobType
	2,  // 10: proto.JobsService.CreateJob:input_type -> proto.CreateJobRequest
	4,  // 11: proto.JobsService.GetJob:input_type -> proto.GetJobRequest
	5,  // 12: proto.JobsService.ListJobs:input_type -> proto.ListJobsRequest
	7,  // 13: proto.JobsService.UpdateJob:input_type -> proto.UpdateJobRequest
	9,  // 14: proto.JobsService.DeleteJob:input_type -> proto.DeleteJobRequest
	3,  // 15: proto.JobsService.CreateJob:output_type -> proto.CreateJobResponse
	11, // 16: proto.JobsService.GetJob:output_type -> proto.Job
	6,  // 17: proto.JobsService.ListJobs:output_type -> proto.ListJobsResponse
	8,  // 18: proto.JobsService.UpdateJob:output_type -> proto.UpdateJobResponse
	10, // 19: proto.JobsService.DeleteJob:output_type -> proto.DeleteJobResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_job_proto_init() }
func file_proto_job_proto_init() {
	if File_proto_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_job_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_job_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_job_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_job_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_job_proto_goTypes,
		DependencyIndexes: file_proto_job_proto_depIdxs,
		EnumInfos:         file_proto_job_proto_enumTypes,
		MessageInfos:      file_proto_job_proto_msgTypes,
	}.Build()
	File_proto_job_proto = out.File
	file_proto_job_proto_rawDesc = nil
	file_proto_job_proto_goTypes = nil
	file_proto_job_proto_depIdxs = nil
}