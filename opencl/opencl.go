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
	CL_UINT        C.cl_uint
	CL_PLATFORM_ID C.cl_platform_id
	CL_INT         C.cl_int
)

// OpenCL Platform Layer

func CLGetPlatformIDs(numEntries CL_UINT, platforms []CL_PLATFORM_ID, numPlatforms *CL_UINT) CL_INT {
	numEntries_ := (C.cl_uint)(numEntries)
	var platforms_ *C.cl_platform_id
	if len(platforms) > 0 {
		platforms_ = (*C.cl_platform_id)(unsafe.Pointer(&platforms[0]))
	}
	numPlatforms_ := (*C.cl_uint)(numPlatforms)
	return CL_INT(C.clGetPlatformIDs(numEntries_, platforms_, numPlatforms_))
}
