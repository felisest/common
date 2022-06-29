#include <cstdint>
#include <iostream>
#include <vector>
#include <string>

class Csort {

public:

    Csort();

    void bubble(void* in_buff, int size);

};

extern "C" void bubble_uint(void* in_buff, size_t size, size_t type_size);
extern "C" void bubble_int(void* in_buff, size_t size, size_t type_size);

template<typename T>
static inline void bubble(T*, size_t size);


