#include "csort.h"

Csort::Csort() {
}

extern "C" void Csort::bubble(void* in_buff, int size) {
}

extern "C" void bubble_uint(void* in_buff, size_t size, size_t type_size) {

    switch (type_size) {
        case 1:
            bubble(static_cast<uint8_t*>(in_buff), size);
            break;
        case 2:
            bubble(static_cast<uint16_t*>(in_buff), size);
            break;
        case 4:
            bubble(static_cast<uint32_t*>(in_buff), size);
            break;
        case 8:
            bubble(static_cast<uint64_t*>(in_buff), size);
            break;
        default: throw std::invalid_argument("Wrong type_size.");
    }
}

extern "C" void bubble_int(void* in_buff, size_t size, size_t type_size) {

    switch (type_size) {
        case 1:
            bubble(static_cast<int8_t*>(in_buff), size);
            break;
        case 2:
            bubble(static_cast<int16_t*>(in_buff), size);
            break;
        case 4:
            bubble(static_cast<int32_t*>(in_buff), size);
            break;
        case 8:
            bubble(static_cast<int64_t*>(in_buff), size);
            break;
        default: throw std::invalid_argument("Wrong type_size.");
    }
}

template<typename T>
static inline void bubble(T* buff, size_t size) {

    for (auto j = size-1; j > 0; j--) {
        for (auto i = 0; i < j; i++) {
            if (buff[i] > buff[i+1]) {
                std::swap(buff[i], buff[i+1]);
            }
        }
    }
}


