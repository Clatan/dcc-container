FROM gcc:13 AS build

WORKDIR /myapp
COPY main.cpp .
RUN g++ -o myapp main.cpp

FROM scratch
COPY --from=build /myapp/myapp /myapp
CMD ["/myapp"]