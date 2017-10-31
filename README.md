# static-file-server

## Usage

Just execute `run.exe` (in Windows) or `run` (in macOS) to start serving static files present in **./public** folder on the **8081** port.

If you find any conflicts running this you can configure some application aspects via running it with some parameters:

* `-port`: You can pass by a new port different from the default **8081**
* `-folder`: You can pass by a new folder different from **./public**

## Usage Requisites

For Windows and macOS:
Nothing, since there's a compiled version of the application available in the repo

For Linux:
Golang 1.8+, since you'll need to compile the application (sorry, this will get fixed soon).

## Build

Run `build.bat` in Windows or `build.sh` in macOS / Linux

## Build requisites

Golang 1.8+