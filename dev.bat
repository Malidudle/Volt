@echo off
echo Running Backend Framework Development Server

echo Generating routes...
go run cmd/generate/main.go

echo Looking for Air...
where air >nul 2>&1
if %ERRORLEVEL% EQU 0 (
    echo Found Air in PATH
    set AIR_PATH=air
    goto config
)

if exist "%USERPROFILE%\go\bin\air.exe" (
    echo Found Air in %USERPROFILE%\go\bin
    set AIR_PATH=%USERPROFILE%\go\bin\air.exe
    goto config
)

if exist "%GOPATH%\bin\air.exe" (
    echo Found Air in %GOPATH%\bin
    set AIR_PATH=%GOPATH%\bin\air.exe
    goto config
)

echo Air not found. Running without hot reload...
go run main.go
goto end

:config
if not exist ".air.toml" (
    echo Creating temporary Air config...
    echo root = "." > .air.toml.tmp
    echo tmp_dir = "tmp" >> .air.toml.tmp
    echo. >> .air.toml.tmp
    echo [build] >> .air.toml.tmp
    echo cmd = "go run cmd/generate/main.go ^&^& go build -o ./tmp/main ." >> .air.toml.tmp
    echo bin = "./tmp/main" >> .air.toml.tmp
    echo include_ext = ["go"] >> .air.toml.tmp
    echo exclude_dir = ["tmp", "vendor", ".git"] >> .air.toml.tmp
    echo exclude_file = ["routes.go", "pkg/router/routes.go"] >> .air.toml.tmp
    echo exclude_regex = [".*routes\\.go$"] >> .air.toml.tmp
    echo delay = 2000 >> .air.toml.tmp
    echo stop_on_error = true >> .air.toml.tmp
    echo. >> .air.toml.tmp
    echo [env] >> .air.toml.tmp
    echo PORT = "8080" >> .air.toml.tmp
    set AIR_CONFIG=.air.toml.tmp
) else (
    set AIR_CONFIG=.air.toml
)

:start
echo Starting development server with Air...
set PORT=8080
%AIR_PATH% -c %AIR_CONFIG%

:end 