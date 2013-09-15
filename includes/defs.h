#ifndef DEFS_H_INCLUDED
#define DEFS_H_INCLUDED

#include <stdlib.h>

#if (defined _WIN32 | _WIN64)
#define WINDOWS
#endif
#ifdef __linux
#define LINUX
#endif

#ifndef WINDOWS
#include "config.h"
#else
#define VERSION "0.30.0"
#define PACKAGE_VERSION "0.30.0"
#endif

#if (defined __MACH__ || __APPLE__ || TARGET_OS_MAC)
#define MAC
#endif


// Define architecture
#if (defined __amd64__ || __amd64 || __x86_64__ || __x86_64)
#define _64BITS
#else
#define _32BITS
#endif

// typedefs
//
#ifndef WINDOWS
typedef int8_t            __int8;
typedef int16_t           __int16;
typedef int32_t           __int32;                
#ifndef SWIGJAVA
typedef int64_t           __int64;
#endif
#endif
typedef __int64           __LONG64;



// #define TESTING

#endif // DEFS_H_INCLUDED
