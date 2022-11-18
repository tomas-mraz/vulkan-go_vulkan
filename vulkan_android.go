//go:build android
// +build android

package vulkan

/*
#cgo android LDFLAGS: -Wl,--no-warn-mismatch
#cgo android CFLAGS: -DVK_USE_PLATFORM_ANDROID_KHR -D_NDK_MATH_NO_SOFTFP=1 -mfpu=vfp

#include <android/native_window.h>

#include "vulkan/vulkan.h"
#include "vk_wrapper.h"
#include "vk_bridge.h"
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

const (
	// UsePlatformAndroid as defined in https://www.khronos.org/registry/vulkan/specs/1.0-wsi_extensions/xhtml/vkspec.html
	UsePlatformAndroid = 1
	// KhrAndroidSurface as defined in vulkan/vulkan.h:3669
	KhrAndroidSurface = 1
	// KhrAndroidSurfaceSpecVersion as defined in vulkan/vulkan.h:3672
	KhrAndroidSurfaceSpecVersion = 6
	// KhrAndroidSurfaceExtensionName as defined in vulkan/vulkan.h:3673
	KhrAndroidSurfaceExtensionName = "VK_KHR_android_surface"
)

// CreateWindowSurface creates a Vulkan surface (VK_KHR_android_surface) for ANativeWindow from Android NDK.
func CreateWindowSurface(instance Instance, nativeWindow uintptr, pAllocator *AllocationCallbacks, pSurface *Surface) Result {
	cinstance, _ := *(*C.VkInstance)(unsafe.Pointer(&instance)), cgoAllocsUnknown
	cpAllocator, _ := (*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)), cgoAllocsUnknown
	cpSurface, _ := (*C.VkSurfaceKHR)(unsafe.Pointer(pSurface)), cgoAllocsUnknown
	cpWindow, _ := (*ANativeWindow)(unsafe.Pointer(nativeWindow)), cgoAllocsUnknown
	// VkAndroidSurfaceCreateInfoKHR
	pCreateInfo := &AndroidSurfaceCreateInfo{
		SType:  StructureTypeAndroidSurfaceCreateInfo,
		PNext:  nil,
		Flags:  0,
		Window: cpWindow,
	}
	cpCreateInfo, _ := (*C.VkAndroidSurfaceCreateInfoKHR)(unsafe.Pointer(pCreateInfo)), cgoAllocsUnknown
	//cpCreateInfo := (*C.VkSurfaceKHR)(unsafe.Pointer(pSurface)), cgoAllocsUnknown
	__ret := C.callVkCreateAndroidSurfaceKHR(cinstance, cpCreateInfo, cpAllocator, cpSurface)
	__v := (Result)(__ret)
	return __v
}

// GetRequiredInstanceExtensions should be used to query instance extensions required for surface initialization.
func GetRequiredInstanceExtensions() []string {
	return []string{
		"VK_KHR_surface\x00",
		"VK_KHR_android_surface\x00",
	}
}

// allocAndroidSurfaceCreateInfoMemory allocates memory for type C.VkAndroidSurfaceCreateInfoKHR in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAndroidSurfaceCreateInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAndroidSurfaceCreateInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAndroidSurfaceCreateInfoValue = unsafe.Sizeof([1]C.VkAndroidSurfaceCreateInfoKHR{})

// ANativeWindow as declared in android/native_window.h:36
type ANativeWindow C.ANativeWindow

// AndroidSurfaceCreateFlags type as declared in https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidSurfaceCreateFlagsKHR.html
type AndroidSurfaceCreateFlags uint32

// AndroidSurfaceCreateInfo as declared in https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAndroidSurfaceKHR.html
type AndroidSurfaceCreateInfo struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Flags  AndroidSurfaceCreateFlags
	Window *ANativeWindow
}

// Ref returns a reference to C object as it is.
func (x *ANativeWindow) Ref() *C.ANativeWindow {
	if x == nil {
		return nil
	}
	return (*C.ANativeWindow)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *ANativeWindow) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewANativeWindowRef converts the C object reference into a raw struct reference without wrapping.
func NewANativeWindowRef(ref unsafe.Pointer) *ANativeWindow {
	return (*ANativeWindow)(ref)
}

// NewANativeWindow allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewANativeWindow() *ANativeWindow {
	return (*ANativeWindow)(allocANativeWindowMemory(1))
}

// allocANativeWindowMemory allocates memory for type C.ANativeWindow in C.
// The caller is responsible for freeing the this memory via C.free.
func allocANativeWindowMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfANativeWindowValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfANativeWindowValue = unsafe.Sizeof([1]C.ANativeWindow{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *ANativeWindow) PassRef() *C.ANativeWindow {
	if x == nil {
		x = (*ANativeWindow)(allocANativeWindowMemory(1))
	}
	return (*C.ANativeWindow)(unsafe.Pointer(x))
}
