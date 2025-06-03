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

--- Python Django
>> cd python
>> docker build -t container-python-django .
>> poetry env info --path
>> & "<path yang muncul>\Scripts\Activate.ps1"
   cnth:  & "C:\Users\clari\AppData\Local\pypoetry\Cache\virtualenvs\demo-VLs62pmx-py3.12\Scripts\Activate.ps1"
>> python manage.py runserver
>> open http://127.0.0.1:8000/
