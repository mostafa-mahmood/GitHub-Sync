@echo off
echo Building GitHub-Sync...

:: Get Git commit hash
for /f "delims=" %%a in ('git rev-parse --short HEAD 2^>nul') do set GIT_COMMIT=%%a
if "%GIT_COMMIT%"=="" set GIT_COMMIT=unknown

:: Get current date
for /f "tokens=1-3 delims=/" %%a in ('echo %DATE%') do set BUILD_DATE=%%c-%%a-%%b

:: Build the binary with version information
go build -o ghs.exe -ldflags "-X 'github.com/mostafa-mahmood/GitHub-Sync/cmd.Version=1.0.0' -X 'github.com/mostafa-mahmood/GitHub-Sync/cmd.BuildDate=%BUILD_DATE%' -X 'github.com/mostafa-mahmood/GitHub-Sync/cmd.GitCommit=%GIT_COMMIT%'"
if errorlevel 1 (
    echo ❌ Failed to build GitHub-Sync. Please ensure Go is installed and your project is set up correctly.
    pause
    exit /b 1
)

echo Creating bin directory in your user profile...
if not exist "%USERPROFILE%\bin" mkdir "%USERPROFILE%\bin"

echo Testing write permissions...
echo. > "%USERPROFILE%\bin\testfile"
if errorlevel 1 (
    echo ❌ You do not have write permissions to %USERPROFILE%\bin.
    pause
    exit /b 1
)
del "%USERPROFILE%\bin\testfile"

echo Installing GitHub-Sync...
copy ghs.exe "%USERPROFILE%\bin\"
if errorlevel 1 (
    echo ❌ Failed to copy ghs.exe to %USERPROFILE%\bin.
    pause
    exit /b 1
)

:: Check if bin directory is already in PATH
echo Checking PATH configuration...
set "binPath=%USERPROFILE%\bin"
call :CheckPath
if %ERRORLEVEL% EQU 0 (
    echo Path already exists in your PATH.
) else (
    echo Adding to PATH...
    setx PATH "%PATH%;%binPath%"
    set PATH=%PATH%;%binPath%
)

echo.
echo Installation complete! 
echo Please restart your command prompt and type 'ghs' to use.
pause
exit /b

:CheckPath
echo %PATH% | findstr /C:"%binPath%" > nul
exit /b %ERRORLEVEL%