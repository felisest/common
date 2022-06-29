cmake_minimum_required(VERSION 3.14)

project(csort)

include(CheckCXXCompilerFlag)

CHECK_CXX_COMPILER_FLAG("-std=c++17" COMPILER_SUPPORTS_CXX17)

if(COMPILER_SUPPORTS_CXX17)
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++17")
else()
    message(FATAL_ERROR "The compiler ${CMAKE_CXX_COMPILER} doesn't have C++17 support.")
endif()

if(NOT CMAKE_BUILD_TYPE)
  set(CMAKE_BUILD_TYPE Release)
endif()

set(CMAKE_CXX_FLAGS "-Wall -Wextra")
set(CMAKE_CXX_FLAGS_DEBUG "-g")
set(CMAKE_CXX_FLAGS_RELEASE "-O3")

set(SRC ${SRC} csort.cpp)
set(HEADERS ${HEADERS} "csort.h")

set(CMAKE_POSITION_INDEPENDENT_CODE ON)

add_library(${PROJECT_NAME}_static STATIC ${SRC} ${HEADERS})
add_library(${PROJECT_NAME}_shared SHARED ${SRC} ${HEADERS})

