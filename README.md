Docker Challenge (Multi-staged Docker Build)

git clone https://github.com/edosyhptra/container-assignment.git
--- C++
>> cd cpp
>> docker build -t container-cpp .
>> docker run --rm container-cpp

--- Go
>> cd go
>> docker build -t container-go .
>> docker run -p 8080:8080 container-go
