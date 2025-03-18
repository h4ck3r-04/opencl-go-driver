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

func CLEnqueueReadBuffer(commandQueue CL_COMMAND_QUEUE, buffer CL_MEM, blockingRead CL_BOOL, offset CL_SIZE_T, size CL_SIZE_T, ptr unsafe.Pointer, numEventsInWaitList CL_UINT, eventWaitList *CL_EVENT, event *CL_EVENT) CL_INT {
	commandQueue_ := (C.cl_command_queue)(commandQueue)
	buffer_ := (C.cl_mem)(buffer)
	blockingRead_ := (C.cl_bool)(blockingRead)
	offset_ := (C.size_t)(offset)
	size_ := (C.size_t)(size)
	numEventsInWaitList_ := (C.cl_uint)(numEventsInWaitList)
	eventWaitList_ := (*C.cl_event)(eventWaitList)
	event_ := (*C.cl_event)(event)
	return CL_INT(C.clEnqueueReadBuffer(commandQueue_, buffer_, blockingRead_, offset_, size_, ptr, numEventsInWaitList_, eventWaitList_, event_))
}

func CLEnqueueReadBufferRect(commandQueue CL_COMMAND_QUEUE, buffer CL_MEM, blockingRead CL_BOOL, bufferOrigin *CL_SIZE_T, hostOrigin *CL_SIZE_T, region *CL_SIZE_T, bufferRowPitch CL_SIZE_T, bufferSlicePitch CL_SIZE_T, hostRowPitch CL_SIZE_T, hostSlicePitch CL_SIZE_T, ptr unsafe.Pointer, numEventsInWaitList CL_UINT, eventWaitList *CL_EVENT, event *CL_EVENT) CL_INT {
	commandQueue_ := (C.cl_command_queue)(commandQueue)
	buffer_ := (C.cl_mem)(buffer)
	blockingRead_ := (C.cl_bool)(blockingRead)
	bufferOrigin_ := (*C.size_t)(bufferOrigin)
	hostOrigin_ := (*C.size_t)(hostOrigin)
	region_ := (*C.size_t)(region)
	bufferRowPitch_ := (C.size_t)(bufferRowPitch)
	bufferSlicePitch_ := (C.size_t)(bufferSlicePitch)
	hostRowPitch_ := (C.size_t)(hostRowPitch)
	hostSlicePitch_ := (C.size_t)(hostSlicePitch)
	numEventsInWaitList_ := (C.cl_uint)(numEventsInWaitList)
	eventWaitList_ := (*C.cl_event)(eventWaitList)
	event_ := (*C.cl_event)(event)
	return CL_INT(C.clEnqueueReadBufferRect(commandQueue_, buffer_, blockingRead_, bufferOrigin_, hostOrigin_, region_, bufferRowPitch_, bufferSlicePitch_, hostRowPitch_, hostSlicePitch_, ptr, numEventsInWaitList_, eventWaitList_, event_))
}

func CLEnqueueWriteBuffer(commandQueue CL_COMMAND_QUEUE, buffer CL_MEM, blockingWrite CL_BOOL, offset CL_SIZE_T, size CL_SIZE_T, ptr unsafe.Pointer, numEventsInWaitList CL_UINT, eventWaitList *CL_EVENT, event *CL_EVENT) CL_INT {
	commandQueue_ := (C.cl_command_queue)(commandQueue)
	buffer_ := (C.cl_mem)(buffer)
	blockingWrite_ := (C.cl_bool)(blockingWrite)
	offset_ := (C.size_t)(offset)
	size_ := (C.size_t)(size)
	numEventsInWaitList_ := (C.cl_uint)(numEventsInWaitList)
	eventWaitList_ := (*C.cl_event)(eventWaitList)
	event_ := (*C.cl_event)(event)
	return CL_INT(C.clEnqueueWriteBuffer(commandQueue_, buffer_, blockingWrite_, offset_, size_, ptr, numEventsInWaitList_, eventWaitList_, event_))
}

func CLEnqueueWriteBufferRect(commandQueue CL_COMMAND_QUEUE, buffer CL_MEM, blockingWrite CL_BOOL, bufferOrigin *CL_SIZE_T, hostOrigin *CL_SIZE_T, region *CL_SIZE_T, bufferRowPitch CL_SIZE_T, bufferSlicePitch CL_SIZE_T, hostRowPitch CL_SIZE_T, hostSlicePitch CL_SIZE_T, ptr unsafe.Pointer, numEventsInWaitList CL_UINT, eventWaitList *CL_EVENT, event *CL_EVENT) CL_INT {
	commandQueue_ := (C.cl_command_queue)(commandQueue)
	buffer_ := (C.cl_mem)(buffer)
	blockingWrite_ := (C.cl_bool)(blockingWrite)
	bufferOrigin_ := (*C.size_t)(bufferOrigin)
	hostOrigin_ := (*C.size_t)(hostOrigin)
	region_ := (*C.size_t)(region)
	bufferRowPitch_ := (C.size_t)(bufferRowPitch)
	bufferSlicePitch_ := (C.size_t)(bufferSlicePitch)
	hostRowPitch_ := (C.size_t)(hostRowPitch)
	hostSlicePitch_ := (C.size_t)(hostSlicePitch)
	numEventsInWaitList_ := (C.cl_uint)(numEventsInWaitList)
	eventWaitList_ := (*C.cl_event)(eventWaitList)
	event_ := (*C.cl_event)(event)
	return CL_INT(C.clEnqueueWriteBufferRect(commandQueue_, buffer_, blockingWrite_, bufferOrigin_, hostOrigin_, region_, bufferRowPitch_, bufferSlicePitch_, hostRowPitch_, hostSlicePitch_, ptr, numEventsInWaitList_, eventWaitList_, event_))
}

func CLEnqueueFillBuffer(commandQueue CL_COMMAND_QUEUE, buffer CL_MEM, pattern unsafe.Pointer, patternSize CL_SIZE_T, offset CL_SIZE_T, size CL_SIZE_T, numEventsInWaitList CL_UINT, eventWaitList *CL_EVENT, event *CL_EVENT) CL_INT {
	commandQueue_ := (C.cl_command_queue)(commandQueue)
	buffer_ := (C.cl_mem)(buffer)
	patternSize_ := (C.size_t)(patternSize)
	offset_ := (C.size_t)(offset)
	size_ := (C.size_t)(size)
	numEventsInWaitList_ := (C.cl_uint)(numEventsInWaitList)
	eventWaitList_ := (*C.cl_event)(eventWaitList)
	event_ := (*C.cl_event)(event)
	return CL_INT(C.clEnqueueFillBuffer(commandQueue_, buffer_, pattern, patternSize_, offset_, size_, numEventsInWaitList_, eventWaitList_, event_))
}

func CLEnqueueCopyBuffer(commandQueue CL_COMMAND_QUEUE, srcBuffer CL_MEM, dstBuffer CL_MEM, srcOffset CL_SIZE_T, dstOffset CL_SIZE_T, size CL_SIZE_T, numEventsInWaitList CL_UINT, eventWaitList *CL_EVENT, event *CL_EVENT) CL_INT {
	commandQueue_ := (C.cl_command_queue)(commandQueue)
	srcBuffer_ := (C.cl_mem)(srcBuffer)
	dstBuffer_ := (C.cl_mem)(dstBuffer)
	srcOffset_ := (C.size_t)(srcOffset)
	dstOffset_ := (C.size_t)(dstOffset)
	size_ := (C.size_t)(size)
	numEventsInWaitList_ := (C.cl_uint)(numEventsInWaitList)
	eventWaitList_ := (*C.cl_event)(eventWaitList)
	event_ := (*C.cl_event)(event)
	return CL_INT(C.clEnqueueCopyBuffer(commandQueue_, srcBuffer_, dstBuffer_, srcOffset_, dstOffset_, size_, numEventsInWaitList_, eventWaitList_, event_))
}

func CLEnqueueCopyBufferRect(commandQueue CL_COMMAND_QUEUE, srcBuffer CL_MEM, dstBuffer CL_MEM, srcOrigin *CL_SIZE_T, dstOrigin *CL_SIZE_T, region *CL_SIZE_T, srcRowPitch CL_SIZE_T, srcSlicePitch CL_SIZE_T, dstRowPitch CL_SIZE_T, dstSlicePitch CL_SIZE_T, numEventsInWaitList CL_UINT, eventWaitList *CL_EVENT, event *CL_EVENT) CL_INT {
	commandQueue_ := (C.cl_command_queue)(commandQueue)
	srcBuffer_ := (C.cl_mem)(srcBuffer)
	dstBuffer_ := (C.cl_mem)(dstBuffer)
	srcOrigin_ := (*C.size_t)(srcOrigin)
	dstOrigin_ := (*C.size_t)(dstOrigin)
	region_ := (*C.size_t)(region)
	srcRowPitch_ := (C.size_t)(srcRowPitch)
	srcSlicePitch_ := (C.size_t)(srcSlicePitch)
	dstRowPitch_ := (C.size_t)(dstRowPitch)
	dstSlicePitch_ := (C.size_t)(dstSlicePitch)
	numEventsInWaitList_ := (C.cl_uint)(numEventsInWaitList)
	eventWaitList_ := (*C.cl_event)(eventWaitList)
	event_ := (*C.cl_event)(event)
	return CL_INT(C.clEnqueueCopyBufferRect(commandQueue_, srcBuffer_, dstBuffer_, srcOrigin_, dstOrigin_, region_, srcRowPitch_, srcSlicePitch_, dstRowPitch_, dstSlicePitch_, numEventsInWaitList_, eventWaitList_, event_))
}

// func clEnqueueMapBuffer

func CLRetainMemObject(memobj CL_MEM) CL_INT {
	memobj_ := (C.cl_mem)(memobj)
	return CL_INT(C.clRetainMemObject(memobj_))
}

func CLReleaseMemObject(memobj CL_MEM) CL_INT {
	memobj_ := (C.cl_mem)(memobj)
	return CL_INT(C.clReleaseMemObject(memobj_))
}
