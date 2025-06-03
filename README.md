Docker Challenge (Multi-staged Docker Build)
> git clone https://github.com/edosyhptra/container-assignment.git

--- C++ ---
1. cd cpp
2. docker build -t container-cpp .
3. docker run --rm container-cpp

--- Go ---
1. cd go
2. docker build -t container-go .
3. docker run -p 8080:8080 container-go

--- Python Django ---
1. cd python
2. docker build -t container-python-django .
3. poetry env info --path
4. & "<path yang muncul>\Scripts\Activate.ps1"
> cnth:  & "C:\Users\clari\AppData\Local\pypoetry\Cache\virtualenvs\demo-VLs62pmx-py3.12\Scripts\Activate.ps1"
5. python manage.py runserver
6. open http://127.0.0.1:8000/

--- PHP ---
1. cd php
2. docker compose -f docker-compose.yml up --build -d
3. open http://localhost:8000
