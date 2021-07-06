# Benchmark
Trivia: we have two funcs with same purpose, but different implementation.  
Using bench tool find the fastest one  
Criteria of acceptance: 
* Output of benchmark run saved to file `original_bench`

# Profiling
Using pprof tool find the main problem in the slowest func  
Fix the code to run it at least 2x faster  
Criteria of acceptance:
* Output svg file containing CPU usage graph of slowest func `original_cpu.svg` 
* Output svg file containing CPU usage graph of __updated__ slowest func `updated_cpu.svg`
* Output of updated code benchmark run saved to file `updated_bench`