@echo off
SET yyyymmdd=%DATE:~0,4%%DATE:~5,2%%DATE:~8,2%
REM SET yyyymmdd=%DATE:~10,4%%DATE:~4,2%%DATE:~7,2%

call cmd /c "mysqldump -r zombi -pmanager ppoint > C:\Users\zombi\ppoint_%yyyymmdd%.sql"