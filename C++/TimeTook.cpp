#include <iostream>
#include <chrono>

int main() {
    auto start_time = std::chrono::high_resolution_clock::now();

    // Code to measure

    for (int i = 0; i < 1000000; i++) {
        std::cout << i << std::endl;
    }

    auto end_time = std::chrono::high_resolution_clock::now();
    auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(end_time - start_time);

    std::cout << "Time took: " << duration.count() << " milliseconds" << std::endl;
}