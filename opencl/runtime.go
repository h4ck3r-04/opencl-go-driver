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

func CLCreateCommandQueue(context CL_CONTEXT, device CL_DEVICE_ID, properties CL_COMMAND_QUEUE_PROPERTIES, errcodeRet *CL_INT) CL_COMMAND_QUEUE {
	context_ := (C.cl_context)(context)
	device_ := (C.cl_device_id)(device)
	properties_ := (C.cl_command_queue_properties)(properties)
	errcodeRet_ := (*C.cl_int)(errcodeRet)
	return CL_COMMAND_QUEUE(C.clCreateCommandQueue(context_, device_, properties_, errcodeRet_))
}

func CLRetainCommandQueue(commandQueue CL_COMMAND_QUEUE) CL_INT {
	commandQueue_ := (C.cl_command_queue)(commandQueue)
	return CL_INT(C.clRetainCommandQueue(commandQueue_))
}

func CLReleaseCommandQueue(commandQueue CL_COMMAND_QUEUE) CL_INT {
	commandQueue_ := (C.cl_command_queue)(commandQueue)
	return CL_INT(C.clReleaseCommandQueue(commandQueue_))
}

func CLGetCommandQueueInfo(commandQueue CL_COMMAND_QUEUE, paramName CL_COMMAND_QUEUE_INFO, paramValueSize CL_SIZE_T, paramValue unsafe.Pointer, paramValueSizeRet *CL_SIZE_T) CL_INT {
	commandQueue_ := (C.cl_command_queue)(commandQueue)
	paramName_ := (C.cl_command_queue_info)(paramName)
	paramValueSize_ := (C.size_t)(paramValueSize)
	paramValueSizeRet_ := (*C.size_t)(paramValueSizeRet)
	return CL_INT(C.clGetCommandQueueInfo(commandQueue_, paramName_, paramValueSize_, paramValue, paramValueSizeRet_))
}
