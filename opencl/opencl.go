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
import (
	"unsafe"
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

func CLGetPlatformInfo(platform CL_PLATFORM_ID, paramName CL_PLATFORM_INFO, paramValueSize CL_SIZE_T, paramValue unsafe.Pointer, paramValueSizeRet *CL_SIZE_T) CL_INT {
	platform_ := (C.cl_platform_id)(platform)
	paramName_ := (C.cl_platform_info)(paramName)
	paramValueSize_ := (C.size_t)(paramValueSize)
	paramValueSizeRet_ := (*C.size_t)(paramValueSizeRet)
	return CL_INT(C.clGetPlatformInfo(platform_, paramName_, paramValueSize_, paramValue, paramValueSizeRet_))
}

func CLGetDeviceIDs(platform CL_PLATFORM_ID, deviceType CL_DEVICE_TYPE, numEntries CL_UINT, devices *CL_DEVICE_ID, numDevices *CL_UINT) CL_INT {
	platform_ := (C.cl_platform_id)(platform)
	deviceType_ := (C.cl_device_type)(deviceType)
	numEntries_ := (C.cl_uint)(numEntries)
	devices_ := (*C.cl_device_id)(devices)
	numDevices_ := (*C.cl_uint)(numDevices)
	return CL_INT(C.clGetDeviceIDs(platform_, deviceType_, numEntries_, devices_, numDevices_))
}

func CLGetDeviceInfo(device CL_DEVICE_ID, paramName CL_DEVICE_INFO, paramValueSize CL_SIZE_T, paramValue unsafe.Pointer, paramValueSizeRet *CL_SIZE_T) CL_INT {
	device_ := (C.cl_device_id)(device)
	paramName_ := (C.cl_device_info)(paramName)
	paramValueSize_ := (C.size_t)(paramValueSize)
	paramValueSizeRet_ := (*C.size_t)(paramValueSizeRet)
	return CL_INT(C.clGetDeviceInfo(device_, paramName_, paramValueSize_, paramValue, paramValueSizeRet_))
}

func CLCreateSubDevices(inDevice CL_DEVICE_ID, properties *CL_DEVICE_PARTITION_PROPERTY, numDevices CL_UINT, outDevices *CL_DEVICE_ID, numDevicesRet *CL_UINT) CL_INT {
	inDevice_ := (C.cl_device_id)(inDevice)
	properties_ := (*C.cl_device_partition_property)(properties)
	numDevices_ := (C.cl_uint)(numDevices)
	outDevices_ := (*C.cl_device_id)(outDevices)
	numDevicesRet_ := (*C.cl_uint)(numDevicesRet)
	return CL_INT(C.clCreateSubDevices(inDevice_, properties_, numDevices_, outDevices_, numDevicesRet_))
}

func CLRetainDevice(device CL_DEVICE_ID) CL_INT {
	device_ := (C.cl_device_id)(device)
	return CL_INT(C.clRetainDevice(device_))
}

func CLReleaseDevice(device CL_DEVICE_ID) CL_INT {
	device_ := (C.cl_device_id)(device)
	return CL_INT(C.clReleaseDevice(device_))
}

// func clCreateContext
// func clCreateContextFromType

func CLRetainContext(context CL_CONTEXT) CL_INT {
	context_ := (C.cl_context)(context)
	return CL_INT(C.clRetainContext(context_))
}

func CLReleaseContext(context CL_CONTEXT) CL_INT {
	context_ := (C.cl_context)(context)
	return CL_INT(C.clReleaseContext(context_))
}

func CLGetContextInfo(context CL_CONTEXT, paramName CL_CONTEXT_INFO, paramValueSize CL_SIZE_T, paramValue unsafe.Pointer, paramValueSizeRet *CL_SIZE_T) CL_INT {
	context_ := (C.cl_context)(context)
	paramName_ := (C.cl_context_info)(paramName)
	paramValueSize_ := (C.size_t)(paramValueSize)
	paramValueSizeRet_ := (*C.size_t)(paramValueSizeRet)
	return CL_INT(C.clGetContextInfo(context_, paramName_, paramValueSize_, paramValue, paramValueSizeRet_))
}

func CLGetExtensionFunctionAddressForPlatform(platform CL_PLATFORM_ID, funcname *CL_CHAR) unsafe.Pointer {
	platform_ := (C.cl_platform_id)(platform)
	funcname_ := (*C.char)(funcname)
	return unsafe.Pointer(C.clGetExtensionFunctionAddressForPlatform(platform_, funcname_))
}
