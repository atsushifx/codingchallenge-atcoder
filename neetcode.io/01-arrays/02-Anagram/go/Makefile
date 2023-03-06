## Makefile for atcoder, go

## commands
RM		:= rm

## lang , compiler
GO		:= go

## directries, flags
APPDIR	:= dist
TARGET	:= $(APPDIR)/main
SRCS	:= main.go solver.go


## Make
.PHONY: clean build

## tasks
all: clean build

clean:
	$(GO) clean
	-$(RM) -fr $(APPDIR)

build: $(TARGET) ;

$(TARGET): $(SRCS)
	$(GO) build -o $(TARGET)
