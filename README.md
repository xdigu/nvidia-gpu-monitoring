# nvidia-gpu-monitoring

### Api to monitore Nvidia GPU in go lang

To run this api you need configure nvidia-smi.exe in your environment path, you can find nvidia-smi in `%programfiles%\NVIDIA Corporation\NVSMI`

Main exemple:

```Go
package main

import
"nvidia-gpu-monitoring/router"


func main() {
	router.RunApi("8080")

}
```

you can check stats of your gpu in JSON:

```
localhost:8080/v1/gpuinfo
```

or in XML:

```
localhost:8080/v1/gpuinfo?format=xml
```

### Data of gpu:
```
Gpu Name
Bios Version
Fan Speed
Gpu Usage
Gpu Memory Usage
Gpu Clock
Gpu Memory Clock
```