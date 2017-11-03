# static-file-server

## Usage

Download latest release from releases

Just execute `sfs-win.exe` (in Windows) or `sfs-macos` (in macOS) to start serving static files present in the **./www** folder on the **8080** port.

If you find any conflicts running this you can configure some application aspects via running it with some parameters:

* `-port`: You can pass by a new port different from the default **8080**
* `-folder`: You can pass by a new folder different from the **./www** folder

## Usage Requisites

For Windows and macOS:
Nothing, since there's a compiled version of the application available in the repo

For Linux:
Golang 1.8+, since you'll need to compile the application (sorry, this will get fixed soon).

## Build

Run `build.bat` in Windows or `build.sh` in macOS / Linux

## Build requisites

Golang 1.8+