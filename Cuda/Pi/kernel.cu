
#include "cuda_runtime.h"
#include "device_launch_parameters.h"

#include <iostream>
#include <limits>

using ullong = unsigned long long;

constexpr ullong N_BLOCKS =     10u;
constexpr ullong N_THREADS =    1'000u;
constexpr ullong ITERATIONS =   10'000'000ul;

#define CUDA_CALL(x) \
    do { if((x)!=cudaSuccess) { \
    printf("Error at %s:%d\n",__FILE__,__LINE__);\
    return EXIT_FAILURE;}} while(0)


__device__ unsigned RandomUInt(unsigned mod) {
    static unsigned seed = 676767676767676;

    seed = seed * 1103515245 + 123456789 * mod;

    return seed;
}

__global__ void MonteCarlo(unsigned long long* counter) {

    for (unsigned long long i = 0; i < ITERATIONS; i++) {

        double x = (double)RandomUInt(blockIdx.x * blockDim.x + threadIdx.x) / UINT_MAX;
        double y = (double)RandomUInt(blockIdx.x * blockDim.x + threadIdx.x) / UINT_MAX;

        if (x * x + y * y <= 1) {
            atomicAdd(counter, 1);
        }

    }
}

int main() {

    //host memory
    unsigned long long output;

    //device memory
    unsigned long long* deviceCounter = nullptr;

    //setting device memory
    CUDA_CALL(cudaMalloc((void**)&deviceCounter, sizeof(unsigned long long)));
    cudaMemset(deviceCounter, 0, sizeof(unsigned long long));

    //calling kernel
    MonteCarlo <<<BLOCKS, THREADS>>> (deviceCounter);

    cudaDeviceSynchronize();

    //retrieving info from device
    cudaMemcpy(&output, deviceCounter, sizeof(unsigned long long), cudaMemcpyDeviceToHost);

    std::cout << "total points: " << N_BLOCKS * N_THREADS * ITERATIONS << "\t circle points: " 
        << output << std::endl << std::endl;

    //outputing results
    printf("pi: %Lf", ((long double)output / N_BLOCKS * N_THREADS * ITERATIONS) * 4);

 

    //cleanup
    cudaFree(deviceCounter);

    return 0;
}


