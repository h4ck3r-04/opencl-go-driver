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

// Platform Info Constants
const (
	CL_PLATFORM_PROFILE    CL_PLATFORM_INFO = C.CL_PLATFORM_PROFILE
	CL_PLATFORM_VERSION    CL_PLATFORM_INFO = C.CL_PLATFORM_VERSION
	CL_PLATFORM_NAME       CL_PLATFORM_INFO = C.CL_PLATFORM_NAME
	CL_PLATFORM_VENDOR     CL_PLATFORM_INFO = C.CL_PLATFORM_VENDOR
	CL_PLATFORM_EXTENSIONS CL_PLATFORM_INFO = C.CL_PLATFORM_EXTENSIONS
)

// Device Info Constants
const (
	CL_DEVICE_NAME                          CL_DEVICE_INFO = C.CL_DEVICE_NAME
	CL_DEVICE_VENDOR                        CL_DEVICE_INFO = C.CL_DEVICE_VENDOR
	CL_DEVICE_PROFILE                       CL_DEVICE_INFO = C.CL_DEVICE_PROFILE
	CL_DEVICE_TYPE_                         CL_DEVICE_INFO = C.CL_DEVICE_TYPE
	CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR      CL_DEVICE_INFO = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR
	CL_DEVICE_NATIVE_VECTOR_WIDTH_INT       CL_DEVICE_INFO = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_INT
	CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG      CL_DEVICE_INFO = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG
	CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT     CL_DEVICE_INFO = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT
	CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE    CL_DEVICE_INFO = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE
	CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF      CL_DEVICE_INFO = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF
	CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT     CL_DEVICE_INFO = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR   CL_DEVICE_INFO = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT    CL_DEVICE_INFO = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG   CL_DEVICE_INFO = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT  CL_DEVICE_INFO = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE CL_DEVICE_INFO = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF   CL_DEVICE_INFO = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF
	CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT  CL_DEVICE_INFO = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT
	CL_DEVICE_PREFERRED_INTEROP_USER_SYNC   CL_DEVICE_INFO = C.CL_DEVICE_PREFERRED_INTEROP_USER_SYNC
	CL_DEVICE_ADDRESS_BITS                  CL_DEVICE_INFO = C.CL_DEVICE_ADDRESS_BITS
	CL_DEVICE_AVAILABLE                     CL_DEVICE_INFO = C.CL_DEVICE_AVAILABLE
	CL_DEVICE_BUILT_IN_KERNELS              CL_DEVICE_INFO = C.CL_DEVICE_BUILT_IN_KERNELS
	CL_DEVICE_COMPILER_AVAILABLE            CL_DEVICE_INFO = C.CL_DEVICE_COMPILER_AVAILABLE
	CL_DEVICE_DOUBLE_FP_CONFIG              CL_DEVICE_INFO = C.CL_DEVICE_DOUBLE_FP_CONFIG
	CL_DEVICE_HALF_FP_CONFIG                CL_DEVICE_INFO = C.CL_DEVICE_HALF_FP_CONFIG
	CL_DEVICE_SINGLE_FP_CONFIG              CL_DEVICE_INFO = C.CL_DEVICE_SINGLE_FP_CONFIG
	CL_DEVICE_ENDIAN_LITTLE                 CL_DEVICE_INFO = C.CL_DEVICE_ENDIAN_LITTLE
	CL_DEVICE_EXTENSIONS                    CL_DEVICE_INFO = C.CL_DEVICE_EXTENSIONS
	CL_DEVICE_ERROR_CORRECTION_SUPPORT      CL_DEVICE_INFO = C.CL_DEVICE_ERROR_CORRECTION_SUPPORT
	CL_DEVICE_EXECUTION_CAPABILITIES        CL_DEVICE_INFO = C.CL_DEVICE_EXECUTION_CAPABILITIES
	CL_DEVICE_GLOBAL_MEM_CACHE_SIZE         CL_DEVICE_INFO = C.CL_DEVICE_GLOBAL_MEM_CACHE_SIZE
	CL_DEVICE_GLOBAL_MEM_CACHE_TYPE         CL_DEVICE_INFO = C.CL_DEVICE_GLOBAL_MEM_CACHE_TYPE
	CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE     CL_DEVICE_INFO = C.CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE
	CL_DEVICE_GLOBAL_MEM_SIZE               CL_DEVICE_INFO = C.CL_DEVICE_GLOBAL_MEM_SIZE
	CL_DEVICE_HOST_UNIFIED_MEMORY           CL_DEVICE_INFO = C.CL_DEVICE_HOST_UNIFIED_MEMORY
	CL_DEVICE_IMAGE_MAX_ARRAY_SIZE          CL_DEVICE_INFO = C.CL_DEVICE_IMAGE_MAX_ARRAY_SIZE
	CL_DEVICE_IMAGE_MAX_BUFFER_SIZE         CL_DEVICE_INFO = C.CL_DEVICE_IMAGE_MAX_BUFFER_SIZE
	CL_DEVICE_IMAGE_SUPPORT                 CL_DEVICE_INFO = C.CL_DEVICE_IMAGE_SUPPORT
	CL_DEVICE_IMAGE2D_MAX_WIDTH             CL_DEVICE_INFO = C.CL_DEVICE_IMAGE2D_MAX_WIDTH
	CL_DEVICE_IMAGE2D_MAX_HEIGHT            CL_DEVICE_INFO = C.CL_DEVICE_IMAGE2D_MAX_HEIGHT
	CL_DEVICE_IMAGE3D_MAX_WIDTH             CL_DEVICE_INFO = C.CL_DEVICE_IMAGE3D_MAX_WIDTH
	CL_DEVICE_IMAGE3D_MAX_HEIGHT            CL_DEVICE_INFO = C.CL_DEVICE_IMAGE3D_MAX_HEIGHT
	CL_DEVICE_IMAGE3D_MAX_DEPTH             CL_DEVICE_INFO = C.CL_DEVICE_IMAGE3D_MAX_DEPTH
	CL_DEVICE_LOCAL_MEM_TYPE                CL_DEVICE_INFO = C.CL_DEVICE_LOCAL_MEM_TYPE
	CL_DEVICE_LOCAL_MEM_SIZE                CL_DEVICE_INFO = C.CL_DEVICE_LOCAL_MEM_SIZE
	CL_DEVICE_MAX_READ_IMAGE_ARGS           CL_DEVICE_INFO = C.CL_DEVICE_MAX_READ_IMAGE_ARGS
	CL_DEVICE_MAX_WRITE_IMAGE_ARGS          CL_DEVICE_INFO = C.CL_DEVICE_MAX_WRITE_IMAGE_ARGS
	CL_DEVICE_MAX_CLOCK_FREQUENCY           CL_DEVICE_INFO = C.CL_DEVICE_MAX_CLOCK_FREQUENCY
	CL_DEVICE_MAX_COMPUTE_UNITS             CL_DEVICE_INFO = C.CL_DEVICE_MAX_COMPUTE_UNITS
	CL_DEVICE_MAX_CONSTANT_ARGS             CL_DEVICE_INFO = C.CL_DEVICE_MAX_CONSTANT_ARGS
	CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE      CL_DEVICE_INFO = C.CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE
	CL_DEVICE_MAX_MEM_ALLOC_SIZE            CL_DEVICE_INFO = C.CL_DEVICE_MAX_MEM_ALLOC_SIZE
	CL_DEVICE_MAX_PARAMETER_SIZE            CL_DEVICE_INFO = C.CL_DEVICE_MAX_PARAMETER_SIZE
	CL_DEVICE_MAX_SAMPLERS                  CL_DEVICE_INFO = C.CL_DEVICE_MAX_SAMPLERS
	CL_DEVICE_MAX_WORK_GROUP_SIZE           CL_DEVICE_INFO = C.CL_DEVICE_MAX_WORK_GROUP_SIZE
	CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS      CL_DEVICE_INFO = C.CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS
	CL_DEVICE_MAX_WORK_ITEM_SIZES           CL_DEVICE_INFO = C.CL_DEVICE_MAX_WORK_ITEM_SIZES
	CL_DEVICE_MEM_BASE_ADDR_ALIGN           CL_DEVICE_INFO = C.CL_DEVICE_MEM_BASE_ADDR_ALIGN
	CL_DEVICE_OPENCL_C_VERSION              CL_DEVICE_INFO = C.CL_DEVICE_OPENCL_C_VERSION
	CL_DEVICE_PARENT_DEVICE                 CL_DEVICE_INFO = C.CL_DEVICE_PARENT_DEVICE
	CL_DEVICE_PARTITION_AFFINITY_DOMAIN     CL_DEVICE_INFO = C.CL_DEVICE_PARTITION_AFFINITY_DOMAIN
	CL_DEVICE_PARTITION_MAX_SUB_DEVICES     CL_DEVICE_INFO = C.CL_DEVICE_PARTITION_MAX_SUB_DEVICES
	CL_DEVICE_PARTITION_PROPERTIES          CL_DEVICE_INFO = C.CL_DEVICE_PARTITION_PROPERTIES
	CL_DEVICE_PARTITION_TYPE                CL_DEVICE_INFO = C.CL_DEVICE_PARTITION_TYPE
	CL_DEVICE_PLATFORM                      CL_DEVICE_INFO = C.CL_DEVICE_PLATFORM
	CL_DEVICE_PRINTF_BUFFER_SIZE            CL_DEVICE_INFO = C.CL_DEVICE_PRINTF_BUFFER_SIZE
	CL_DEVICE_PROFILING_TIMER_RESOLUTION    CL_DEVICE_INFO = C.CL_DEVICE_PROFILING_TIMER_RESOLUTION
	CL_DEVICE_QUEUE_PROPERTIES              CL_DEVICE_INFO = C.CL_DEVICE_QUEUE_PROPERTIES
	CL_DEVICE_REFERENCE_COUNT               CL_DEVICE_INFO = C.CL_DEVICE_REFERENCE_COUNT
	CL_DEVICE_VENDOR_ID                     CL_DEVICE_INFO = C.CL_DEVICE_VENDOR_ID
	CL_DEVICE_VERSION                       CL_DEVICE_INFO = C.CL_DEVICE_VERSION
	CL_DRIVER_VERSION                       CL_DEVICE_INFO = C.CL_DRIVER_VERSION
)

// Partition Properties Constants
const (
	CL_DEVICE_PARTITION_EQUALLY                  CL_DEVICE_PARTITION_PROPERTY = C.CL_DEVICE_PARTITION_EQUALLY
	CL_DEVICE_PARTITION_BY_COUNTS                CL_DEVICE_PARTITION_PROPERTY = C.CL_DEVICE_PARTITION_BY_COUNTS
	CL_DEVICE_PARTITION_BY_AFFINITY_DOMAIN       CL_DEVICE_PARTITION_PROPERTY = C.CL_DEVICE_PARTITION_BY_AFFINITY_DOMAIN
	CL_DEVICE_AFFINITY_DOMAIN_NUMA               CL_DEVICE_PARTITION_PROPERTY = C.CL_DEVICE_AFFINITY_DOMAIN_NUMA
	CL_DEVICE_AFFINITY_DOMAIN_L4_CACHE           CL_DEVICE_PARTITION_PROPERTY = C.CL_DEVICE_AFFINITY_DOMAIN_L4_CACHE
	CL_DEVICE_AFFINITY_DOMAIN_L3_CACHE           CL_DEVICE_PARTITION_PROPERTY = C.CL_DEVICE_AFFINITY_DOMAIN_L3_CACHE
	CL_DEVICE_AFFINITY_DOMAIN_L2_CACHE           CL_DEVICE_PARTITION_PROPERTY = C.CL_DEVICE_AFFINITY_DOMAIN_L2_CACHE
	CL_DEVICE_AFFINITY_DOMAIN_L1_CACHE           CL_DEVICE_PARTITION_PROPERTY = C.CL_DEVICE_AFFINITY_DOMAIN_L1_CACHE
	CL_DEVICE_AFFINITY_DOMAIN_NEXT_PARTITIONABLE CL_DEVICE_PARTITION_PROPERTY = C.CL_DEVICE_AFFINITY_DOMAIN_NEXT_PARTITIONABLE
)

// OpenCL Context Properties Constants
const (
	CL_CONTEXT_PLATFORM          CL_CONTEXT_PROPERTIES = C.CL_CONTEXT_PLATFORM
	CL_CONTEXT_INTEROP_USER_SYNC CL_CONTEXT_PROPERTIES = C.CL_CONTEXT_INTEROP_USER_SYNC
)

// OpenCL Context Info Constants
const (
	CL_CONTEXT_REFERENCE_COUNT CL_CONTEXT_INFO = C.CL_CONTEXT_REFERENCE_COUNT
	CL_CONTEXT_DEVICES         CL_CONTEXT_INFO = C.CL_CONTEXT_DEVICES
	CL_CONTEXT_NUM_DEVICES     CL_CONTEXT_INFO = C.CL_CONTEXT_NUM_DEVICES
	CL_CONTEXT_PROPERTIES_     CL_CONTEXT_INFO = C.CL_CONTEXT_PROPERTIES
)
