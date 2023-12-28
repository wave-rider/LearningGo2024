# LearningGo2024

## 28.12.2023
To speed up go on Windows without disabling the antivirus.
1. Add the working dir to the Avira exclusion list
2. powershell
`$env:GOTMPDIR = "C:\Temp\GoBuilds"`
3. `go run -work .\hello.go

## 27.12.2023
I figured out what was causing the slowness after running the same build commands on 10 years old Mac Book Pro.

The problem was caused by the antivurus on my Windows laptop, disabling it fixed the issue

## 26.12.2023
1. Install go from go.dev
2. Install vs code
3. Install go plugin
4. View/ Command Palette / Go Install/Update tools

For some reason `go run .`  takes a very long time to execute

So I use `go build .`

` .\hello.exe`
