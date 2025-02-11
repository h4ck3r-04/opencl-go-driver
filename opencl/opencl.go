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
import "unsafe"

type (
	CL_UINT          C.cl_uint
	CL_PLATFORM_ID   C.cl_platform_id
	CL_INT           C.cl_int
	CL_PLATFORM_INFO C.cl_platform_info
	CL_SIZE_T        C.size_t
)

// OpenCL Platform Layer

const (
	CL_PLATFORM_PROFILE    CL_PLATFORM_INFO = C.CL_PLATFORM_PROFILE
	CL_PLATFORM_VERSION    CL_PLATFORM_INFO = C.CL_PLATFORM_VERSION
	CL_PLATFORM_NAME       CL_PLATFORM_INFO = C.CL_PLATFORM_NAME
	CL_PLATFORM_VENDOR     CL_PLATFORM_INFO = C.CL_PLATFORM_VENDOR
	CL_PLATFORM_EXTENSIONS CL_PLATFORM_INFO = C.CL_PLATFORM_EXTENSIONS
)

func CLGetPlatformIDs(numEntries CL_UINT, platforms []CL_PLATFORM_ID, numPlatforms *CL_UINT) CL_INT {
	numEntries_ := (C.cl_uint)(numEntries)
	var platforms_ *C.cl_platform_id
	if len(platforms) > 0 {
		platforms_ = (*C.cl_platform_id)(unsafe.Pointer(&platforms[0]))
	}
	numPlatforms_ := (*C.cl_uint)(numPlatforms)
	return CL_INT(C.clGetPlatformIDs(numEntries_, platforms_, numPlatforms_))
}

func CLGetPlatformInfo(platform CL_PLATFORM_ID, paramName CL_PLATFORM_INFO, paramValueSize CL_SIZE_T, paramValue unsafe.Pointer, paramValueSizeRet *CL_SIZE_T) CL_INT {
	platform_ := (C.cl_platform_id)(platform)
	paramName_ := (C.cl_platform_info)(paramName)
	paramValueSize_ := (C.size_t)(paramValueSize)
	paramValueSizeRet_ := (*C.size_t)(paramValueSizeRet)
	return CL_INT(C.clGetPlatformInfo(platform_, paramName_, paramValueSize_, paramValue, paramValueSizeRet_))
}
