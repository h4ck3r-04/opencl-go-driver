package opencl

/*
#cgo darwin  LDFLAGS: -framework OpenCL
#cgo linux   LDFLAGS: -lOpenCL
#cgo windows LDFLAGS: -lOpenCL

#ifdef __APPLE__
    #include <OpenCL/opencl.h>
#elif defined(_WIN32)
    #include <CL/opencl.h>
#else
    #include <CL/cl.h>
#endif
*/
import "C"

type (
	CL_UINT                      C.cl_uint
	CL_PLATFORM_ID               C.cl_platform_id
	CL_INT                       C.cl_int
	CL_PLATFORM_INFO             C.cl_platform_info
	CL_SIZE_T                    C.size_t
	CL_DEVICE_TYPE               C.cl_device_type
	CL_DEVICE_ID                 C.cl_device_id
	CL_DEVICE_INFO               C.cl_device_info
	CL_DEVICE_PARTITION_PROPERTY C.cl_device_partition_property
	CL_CONTEXT                   C.cl_context
	CL_CONTEXT_PROPERTIES        C.cl_context_properties
	CL_CONTEXT_INFO              C.cl_context_info
	CL_CHAR                      C.char
	CL_COMMAND_QUEUE             C.cl_command_queue
	CL_COMMAND_QUEUE_PROPERTIES  C.cl_command_queue_properties
	CL_COMMAND_QUEUE_INFO        C.cl_command_queue_info
	CL_MEM                       C.cl_mem
	CL_MEM_FLAGS                 C.cl_mem_flags
	CL_BUFFER_CREATE_TYPE        C.cl_buffer_create_type
	CL_BOOL                      C.cl_bool
	CL_EVENT                     C.cl_event
)
