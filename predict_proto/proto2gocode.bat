@cd /d %~dp0
@set PROTO_DIR=.
@protoc  -I  %PROTO_DIR%  --go_out=plugins=grpc:.  %PROTO_DIR%\*.proto
@pause
