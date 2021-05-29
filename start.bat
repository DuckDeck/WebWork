@echo off
set AppName=main.exe
set AppPath=D:\Project\WebWork\
title 进程监控
cls
echo.
echo 进程监控开始……
echo.
:startjc
qprocess %AppName% > nul
if %errorlevel% equ 0 (
rem echo ^>%date:~0,10% %time:~0,8% running……
)else (
start %AppPath%%AppName% 2>nul && echo ^>%date:~0,10% %time:~0,8% start,success!
)
for /l %%i in (1,1,1) do ping -n 2 -w 100 127.0.0.1>nul
goto startjc
echo on
————————————————
版权声明：本文为CSDN博主「Asciphx」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/qq_33359662/article/details/86321027