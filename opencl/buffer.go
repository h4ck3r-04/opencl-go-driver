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

func CLCreateBuffer(context CL_CONTEXT, flags CL_MEM_FLAGS, size CL_SIZE_T, hostPtr unsafe.Pointer, errCodeRet *CL_INT) CL_MEM {
	context_ := (C.cl_context)(context)
	flags_ := (C.cl_mem_flags)(flags)
	size_ := (C.size_t)(size)
	errCodeRet_ := (*C.cl_int)(errCodeRet)
	return CL_MEM(C.clCreateBuffer(context_, flags_, size_, hostPtr, errCodeRet_))
}

func CLCreateSubBuffer(buffer CL_MEM, flags CL_MEM_FLAGS, bufferCreateType CL_BUFFER_CREATE_TYPE, bufferCreateInfo unsafe.Pointer, errCodeRet *CL_INT) CL_MEM {
	buffer_ := (C.cl_mem)(buffer)
	flags_ := (C.cl_mem_flags)(flags)
	bufferCreateType_ := (C.cl_buffer_create_type)(bufferCreateType)
	errCodeRet_ := (*C.cl_int)(errCodeRet)
	return CL_MEM(C.clCreateSubBuffer(buffer_, flags_, bufferCreateType_, bufferCreateInfo, errCodeRet_))
}
