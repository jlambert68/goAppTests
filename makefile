setUpDirectories:
	bash -c "ScriptsUsedByMake/SetUpBasicDirectories.sh"

# **********************************************************************
# 010_goAppOriginalExample
010_goAppOriginalExample_compile: 010_goAppOriginalExample_removeAll 010_goAppOriginalExample_compileAll

010_goAppOriginalExample_run:
	cd ./build/010_goAppOriginalExample/ &&./server

010_goAppOriginalExample_removeAll:
	touch build/010_goAppOriginalExample/server && rm build/010_goAppOriginalExample/server
	touch build/010_goAppOriginalExample/web/app.wasm && rm build/010_goAppOriginalExample/web/app.wasm

010_goAppOriginalExample_compileAll: 010_goAppOriginalExample_compileServer 010_goAppOriginalExample_compileWasm

010_goAppOriginalExample_compileServer:
	go build -o build/010_goAppOriginalExample/server cmd/010_goAppOriginalExample/server/main.go

010_goAppOriginalExample_compileWasm:
	GOARCH=wasm GOOS=js go build -o build/010_goAppOriginalExample/web/app.wasm cmd/010_goAppOriginalExample/wasm/main.go

# **********************************************************************

# **********************************************************************
# 020_ServerWasmNoGrpcViaHtml
020_ServerWasmNoGrpcViaHtml_compile: 020_ServerWasmNoGrpcViaHtml_removeAll 020_ServerWasmNoGrpcViaHtml_compileAll

020_ServerWasmNoGrpcViaHtml_run:
	cd ./build/020_ServerWasmNoGrpcViaHtml/ && ./backendServer

020_ServerWasmNoGrpcViaHtml_removeAll:
	touch build/020_ServerWasmNoGrpcViaHtml/backendServer && rm build/020_ServerWasmNoGrpcViaHtml/backendServer
	touch build/020_ServerWasmNoGrpcViaHtml/web/app.wasm && rm build/020_ServerWasmNoGrpcViaHtml/web/app.wasm

020_ServerWasmNoGrpcViaHtml_compileAll: 020_ServerWasmNoGrpcViaHtml_compileServer 020_ServerWasmNoGrpcViaHtml_compileWasm

020_ServerWasmNoGrpcViaHtml_compileServer:
	go build -o build/020_ServerWasmNoGrpcViaHtml/backendServer cmd/020_ServerWasmNoGrpcViaHtml/backendServer/backendServer.go

020_ServerWasmNoGrpcViaHtml_compileWasm:
	GOARCH=wasm GOOS=js go build -o build/020_ServerWasmNoGrpcViaHtml/web/app.wasm cmd/020_ServerWasmNoGrpcViaHtml/wasm/wasm.go

# **********************************************************************


# **********************************************************************
# 030_golangWasmProtosIsPrecompiled
030_golangWasmProtosIsPrecompiled_compile: 030_golangWasmProtosIsPrecompiled_removeAll 030_golangWasmProtosIsPrecompiled_compileAll

030_golangWasmProtosIsPrecompiled_run:
	cd ./build/030_golangWasmProtosIsPrecompiled/ && ./webServer

030_golangWasmProtosIsPrecompiled_removeAll:
	touch build/030_golangWasmProtosIsPrecompiled/webServer && rm build/030_golangWasmProtosIsPrecompiled/webServer
	touch build/030_golangWasmProtosIsPrecompiled/web/app.wasm && rm build/030_golangWasmProtosIsPrecompiled/web/app.wasm

030_golangWasmProtosIsPrecompiled_compileAll: 030_golangWasmProtosIsPrecompiled_compileWebServer 030_golangWasmProtosIsPrecompiled_compileWasm

030_golangWasmProtosIsPrecompiled_compileWebServer:
	go build -o build/030_golangWasmProtosIsPrecompiled/webServer cmd/030_golangWasmProtosIsPrecompiled/webServer/main.go

030_golangWasmProtosIsPrecompiled_compileWasm:
	GOARCH=wasm GOOS=js go build -o build/030_golangWasmProtosIsPrecompiled/web/app.wasm ./cmd/030_golangWasmProtosIsPrecompiled/wasm

# **********************************************************************

# **********************************************************************
# 040_golangWasmProtosNotPrecompiled
040_golangWasmProtosNotPrecompiled_compile: 040_golangWasmProtosNotPrecompiled_removeAll 040_golangWasmProtosNotPrecompiled_compileAll

040_golangWasmProtosNotPrecompiled_run:
	cd ./build/040_golangWasmProtosNotPrecompiled/ && ./webServer

040_golangWasmProtosNotPrecompiled_removeAll:
	touch build/040_golangWasmProtosNotPrecompiled/webServer && rm build/040_golangWasmProtosNotPrecompiled/webServer
	touch build/040_golangWasmProtosNotPrecompiled/web/app.wasm && rm build/040_golangWasmProtosNotPrecompiled/web/app.wasm

040_golangWasmProtosNotPrecompiled_compileAll: 040_golangWasmProtosNotPrecompiled_compileWebServer 040_golangWasmProtosNotPrecompiled_compileWasm

040_golangWasmProtosNotPrecompiled_compileWebServer:
	go build -o build/040_golangWasmProtosNotPrecompiled/webServer cmd/040_golangWasmProtosNotPrecompiled/webServer/main.go

040_golangWasmProtosNotPrecompiled_compileWasm:
	GOARCH=wasm GOOS=js go build -o build/040_golangWasmProtosNotPrecompiled/web/app.wasm ./cmd/040_golangWasmProtosNotPrecompiled/wasm

# GRPC
040_grpc_compile_all:040_grpc_remove_grpc_result_code 040_grpc_create_client_code 040_grpc_create_server_code 040_grpc_compile_base 040_grpc_compile_client 040_grpc_compile_server

040_grpc_remove_grpc_result_code:
	touch cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-client/protoc-gen-client && rm cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-client/protoc-gen-client
	touch cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-server/protoc-gen-server && rm cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-server/protoc-gen-server
	touch cmd/040_golangWasmProtosNotPrecompiled/protos/api/api.pb.go && rm cmd/040_golangWasmProtosNotPrecompiled/protos/api/api.pb.go
	touch cmd/040_golangWasmProtosNotPrecompiled/protos/api/api.pb.client.go && rm cmd/040_golangWasmProtosNotPrecompiled/protos/api/api.pb.client.go
	touch cmd/040_golangWasmProtosNotPrecompiled/protos/api/api.pb.server.go && rm cmd/040_golangWasmProtosNotPrecompiled/protos/api/api.pb.server.go

040_grpc_create_client_code:
	cd cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-client && go build -o protoc-gen-client main.go

040_grpc_create_server_code:
	cd cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-server && go build -o protoc-gen-server main.go

040_grpc_compile_base:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/040_golangWasmProtosNotPrecompiled/protos -I cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-client --go_out=plugins=grpc:. cmd/040_golangWasmProtosNotPrecompiled/protos/api.proto

040_grpc_compile_client:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/040_golangWasmProtosNotPrecompiled/protos -I cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-client --plugin=./cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-client/protoc-gen-client --client_out=. cmd/040_golangWasmProtosNotPrecompiled/protos/api.proto

040_grpc_compile_server:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/040_golangWasmProtosNotPrecompiled/protos -I cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-server --plugin=./cmd/040_golangWasmProtosNotPrecompiled/tools/protoc-gen-server/protoc-gen-server --server_out=. cmd/040_golangWasmProtosNotPrecompiled/protos/api.proto

# **********************************************************************

# **********************************************************************
# 050_golangWasmMultiProtoServices
050_golangWasmMultiProtoServices_compile: 050_golangWasmMultiProtoServices_removeAll 050_golangWasmMultiProtoServices_compileAll

050_golangWasmMultiProtoServices_run:
	cd ./build/050_golangWasmMultiProtoServices/ && ./webServer

050_golangWasmMultiProtoServices_removeAll:
	touch build/050_golangWasmMultiProtoServices/webServer && rm build/050_golangWasmMultiProtoServices/webServer
	touch build/050_golangWasmMultiProtoServices/web/app.wasm && rm build/050_golangWasmMultiProtoServices/web/app.wasm

050_golangWasmMultiProtoServices_compileAll: 050_golangWasmMultiProtoServices_compileWebServer 050_golangWasmMultiProtoServices_compileWasm

050_golangWasmMultiProtoServices_compileWebServer:
	go build -o build/050_golangWasmMultiProtoServices/webServer cmd/050_golangWasmMultiProtoServices/webServer/main.go

050_golangWasmMultiProtoServices_compileWasm:
	GOARCH=wasm GOOS=js go build -o build/050_golangWasmMultiProtoServices/web/app.wasm ./cmd/050_golangWasmMultiProtoServices/wasm

# GRPC
050__grpc_compile_all:050__grpc_remove_grpc_result_code 050__grpc_create_client_code 050__grpc_create_server_code 050__grpc_compile_base 050__grpc_compile_client 050__grpc_compile_server

050__grpc_remove_grpc_result_code:
	touch cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-client/protoc-gen-client && rm cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-client/protoc-gen-client
	touch cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-server/protoc-gen-server && rm cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-server/protoc-gen-server
	touch cmd/050_golangWasmMultiProtoServices/protos/api/api.pb.go && rm cmd/050_golangWasmMultiProtoServices/protos/api/api.pb.go
	touch cmd/050_golangWasmMultiProtoServices/protos/api/api.pb.client.go && rm cmd/050_golangWasmMultiProtoServices/protos/api/api.pb.client.go
	touch cmd/050_golangWasmMultiProtoServices/protos/api/api.pb.server.go && rm cmd/050_golangWasmMultiProtoServices/protos/api/api.pb.server.go

050__grpc_create_client_code:
	cd cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-client && go build -o protoc-gen-client main.go

050__grpc_create_server_code:
	cd cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-server && go build -o protoc-gen-server main.go

050__grpc_compile_base:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/050_golangWasmMultiProtoServices/protos -I cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-client --go_out=plugins=grpc:. cmd/050_golangWasmMultiProtoServices/protos/api.proto

050__grpc_compile_client:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/050_golangWasmMultiProtoServices/protos -I cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-client --plugin=./cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-client/protoc-gen-client --client_out=. cmd/050_golangWasmMultiProtoServices/protos/api.proto

050__grpc_compile_server:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/050_golangWasmMultiProtoServices/protos -I cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-server --plugin=./cmd/050_golangWasmMultiProtoServices/tools/protoc-gen-server/protoc-gen-server --server_out=. cmd/050_golangWasmMultiProtoServices/protos/api.proto

# **********************************************************************

# **********************************************************************
# 060_golangWasmMultiProto_goAppv8
060_golangWasmMultiProto_goAppv8_compile: 060_golangWasmMultiProto_goAppv8_removeAll 060_golangWasmMultiProto_goAppv8_compileAll

060_golangWasmMultiProto_goAppv8_run:
	cd ./build/060_golangWasmMultiProto_goAppv8/ && ./webServer

060_golangWasmMultiProto_goAppv8_removeAll:
	touch build/060_golangWasmMultiProto_goAppv8/webServer && rm build/060_golangWasmMultiProto_goAppv8/webServer
	touch build/060_golangWasmMultiProto_goAppv8/web/app.wasm && rm build/060_golangWasmMultiProto_goAppv8/web/app.wasm

060_golangWasmMultiProto_goAppv8_compileAll: 060_golangWasmMultiProto_goAppv8_compileWebServer 060_golangWasmMultiProto_goAppv8_compileWasm

060_golangWasmMultiProto_goAppv8_compileWebServer: 060_createServerCompiledTimeStampFile
	go build -o build/060_golangWasmMultiProto_goAppv8/webServer cmd/060_golangWasmMultiProto_goAppv8/webServer/main.go

060_golangWasmMultiProto_goAppv8_compileWasm: 060_createClientCompiledTimeStampFile
	GOARCH=wasm GOOS=js go build -o build/060_golangWasmMultiProto_goAppv8/web/app.wasm ./cmd/060_golangWasmMultiProto_goAppv8/wasm

# GRPC
060_grpc_compile_all:060_grpc_remove_grpc_result_code 060_grpc_create_client_code 060_grpc_create_server_code 060_grpc_compile_base 060_grpc_compile_client 060_grpc_compile_server

060_grpc_remove_grpc_result_code:
	touch cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-client/protoc-gen-client && rm cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-client/protoc-gen-client
	touch cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-server/protoc-gen-server && rm cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-server/protoc-gen-server
	touch cmd/060_golangWasmMultiProto_goAppv8/protos/api/api.pb.go && rm cmd/060_golangWasmMultiProto_goAppv8/protos/api/api.pb.go
	touch cmd/060_golangWasmMultiProto_goAppv8/protos/api/api.pb.client.go && rm cmd/060_golangWasmMultiProto_goAppv8/protos/api/api.pb.client.go
	touch cmd/060_golangWasmMultiProto_goAppv8/protos/api/api.pb.server.go && rm cmd/060_golangWasmMultiProto_goAppv8/protos/api/api.pb.server.go

060_grpc_create_client_code:
	cd cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-client && go build -o protoc-gen-client main.go

060_grpc_create_server_code:
	cd cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-server && go build -o protoc-gen-server main.go

060_grpc_compile_base:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/060_golangWasmMultiProto_goAppv8/protos -I cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-client --go_out=plugins=grpc:. cmd/060_golangWasmMultiProto_goAppv8/protos/api.proto

060_grpc_compile_client:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/060_golangWasmMultiProto_goAppv8/protos -I cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-client --plugin=./cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-client/protoc-gen-client --client_out=. cmd/060_golangWasmMultiProto_goAppv8/protos/api.proto

060_grpc_compile_server:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/060_golangWasmMultiProto_goAppv8/protos -I cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-server --plugin=./cmd/060_golangWasmMultiProto_goAppv8/tools/protoc-gen-server/protoc-gen-server --server_out=. cmd/060_golangWasmMultiProto_goAppv8/protos/api.proto

060_createClientCompiledTimeStampFile:
	bash -c "cmd/060_golangWasmMultiProto_goAppv8/ScriptsUsedByMake/CreateClientCompiledTimeStamp.sh"


060_createServerCompiledTimeStampFile:
	bash -c "cmd/060_golangWasmMultiProto_goAppv8/ScriptsUsedByMake/CreateServerCompiledTimeStamp.sh"

# **********************************************************************

# **********************************************************************
# 070_AddListDelete_TestDomains
070_AddListDelete_TestDomains_compile: 070_AddListDelete_TestDomains_removeAll 070_AddListDelete_TestDomains_compileAll

070_AddListDelete_TestDomains_run:
	cd ./build/070_AddListDelete_TestDomains/ && ./webServer

070_AddListDelete_TestDomains_removeAll:
	touch build/070_AddListDelete_TestDomains/webServer && rm build/070_AddListDelete_TestDomains/webServer
	touch build/070_AddListDelete_TestDomains/web/app.wasm && rm build/070_AddListDelete_TestDomains/web/app.wasm

070_AddListDelete_TestDomains_compileAll: 070_AddListDelete_TestDomains_compileWebServer 070_AddListDelete_TestDomains_compileWasm

070_AddListDelete_TestDomains_compileWebServer: 070_createServerCompiledTimeStampFile
	go build -o build/070_AddListDelete_TestDomains/webServer cmd/070_AddListDelete_TestDomains/webServer/main.go

070_AddListDelete_TestDomains_compileWasm: 070_createClientCompiledTimeStampFile
	GOARCH=wasm GOOS=js go build -o build/070_AddListDelete_TestDomains/web/app.wasm ./cmd/070_AddListDelete_TestDomains/wasm

# GRPC
070_grpc_compile_all:070_grpc_remove_grpc_result_code 070_grpc_create_client_code 070_grpc_create_server_code 070_grpc_compile_base 070_grpc_compile_client 070_grpc_compile_server

070_grpc_remove_grpc_result_code:
	touch cmd/070_AddListDelete_TestDomains/tools/protoc-gen-client/protoc-gen-client && rm cmd/070_AddListDelete_TestDomains/tools/protoc-gen-client/protoc-gen-client
	touch cmd/070_AddListDelete_TestDomains/tools/protoc-gen-server/protoc-gen-server && rm cmd/070_AddListDelete_TestDomains/tools/protoc-gen-server/protoc-gen-server
	touch cmd/070_AddListDelete_TestDomains/protos/api/api.pb.go && rm cmd/070_AddListDelete_TestDomains/protos/api/api.pb.go
	touch cmd/070_AddListDelete_TestDomains/protos/api/api.pb.client.go && rm cmd/070_AddListDelete_TestDomains/protos/api/api.pb.client.go
	touch cmd/070_AddListDelete_TestDomains/protos/api/api.pb.server.go && rm cmd/070_AddListDelete_TestDomains/protos/api/api.pb.server.go

070_grpc_create_client_code:
	cd cmd/070_AddListDelete_TestDomains/tools/protoc-gen-client && go build -o protoc-gen-client main.go

070_grpc_create_server_code:
	cd cmd/070_AddListDelete_TestDomains/tools/protoc-gen-server && go build -o protoc-gen-server main.go

070_grpc_compile_base:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/070_AddListDelete_TestDomains/protos -I cmd/070_AddListDelete_TestDomains/tools/protoc-gen-client --go_out=plugins=grpc:. cmd/070_AddListDelete_TestDomains/protos/api.proto

070_grpc_compile_client:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/070_AddListDelete_TestDomains/protos -I cmd/070_AddListDelete_TestDomains/tools/protoc-gen-client --plugin=./cmd/070_AddListDelete_TestDomains/tools/protoc-gen-client/protoc-gen-client --client_out=. cmd/070_AddListDelete_TestDomains/protos/api.proto

070_grpc_compile_server:
	~/Programs/protoc-3.13.0-linux-x86_64/bin/protoc  -I cmd/070_AddListDelete_TestDomains/protos -I cmd/070_AddListDelete_TestDomains/tools/protoc-gen-server --plugin=./cmd/070_AddListDelete_TestDomains/tools/protoc-gen-server/protoc-gen-server --server_out=. cmd/070_AddListDelete_TestDomains/protos/api.proto

070_createClientCompiledTimeStampFile:
	bash -c "cmd/070_AddListDelete_TestDomains/ScriptsUsedByMake/CreateClientCompiledTimeStamp.sh"


070_createServerCompiledTimeStampFile:
	bash -c "cmd/070_AddListDelete_TestDomains/ScriptsUsedByMake/CreateServerCompiledTimeStamp.sh"

# **********************************************************************




gobuild:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

run:
	./testGoApp_v2

clean:
	@go clean
	@-rm app.wasm


clean_bazel:
	bazelisk clean


000_CreateWORKSPACEandBUILD: 005_CopyWorkspaceFilesToUse 010_reset_WORKSPACEandBUILD 020_copyBUILD 030_CancatRowsIntoOneROw 040_CreateNewWORKSPACE

# Workspace to use
WorkspaceId=v1.1

WORKSPACEFILE=originals/workspace_files_to_use_$(WorkspaceId)
WorkspaceFileToUse=`cat $(WORKSPACEFILE)`
005_CopyWorkspaceFilesToUse:
	#echo $(WORKSPACEFILE)
	cp $(WORKSPACEFILE) originals/workspace_files_to_use

010_reset_WORKSPACEandBUILD:$(FOO)
	find . -name \BUILD.bazel -type f -delete
	touch WORKSPACE && rm WORKSPACE
	touch BUILD && rm BUILD


020_copyBUILD:
	cp originals/BUILD .

030_CancatRowsIntoOneROw:
	bash -c "ScriptsUsedByMake/ConcatRowsInFile.sh"

FILE=originals/workspace_files_to_use_temp
VARIABLE=`cat $(FILE)`
040_CreateNewWORKSPACE:
	echo $(VARIABLE)
	sed ':a; N; $!ba; s/\n/ /g' $(VARIABLE) > WORKSPACE

100_GazellUpdateRepos: 110_addModLibs 120_executeGazelle 130_copy_special_BUILD.bazel_files

110_addModLibs:
	bazelisk run //:gazelle -- update-repos -from_file=go.mod

120_executeGazelle:
	bazelisk run //:gazelle

130_copy_special_BUILD.bazel_files:
	cp originals/goAppTest/wasm/BUILD.bazel.original packages/goAppTest1/wasm/BUILD.bazel

200_BuildAndRun_Main: 210_build_main 220_run_main
201_BuildAndRun_Second: 211_build_second 221_run_second
202_BuildAndRun_goAppTest1: 212_build_goAppTest1 222_run_goAppTest1

# Main App
210_build_main:
	bazelisk build //packages/main_app:main_app

220_run_main:
	./bazel-bin/packages/main_app/linux_amd64_stripped/main_app

# Second App
211_build_second:
	bazelisk build //packages/second_app:second_app

221_run_second:
	./bazel-bin/packages/second_app/linux_amd64_stripped/second_app

# goAppTest1
212_build_goAppTest1_server:
	bazelisk build //packages/goAppTest1/server:server
	touch deploy/app.server && rm deploy/app.server
	cp bazel-bin/packages/goAppTest1/server/linux_amd64_stripped/server deploy/app.server

213_build_goAppTest1_wasm:
	bazelisk build //packages/goAppTest1/wasm:wasm #--sandbox_debug --verbose_failures
	touch deploy/web/app.wasm && rm deploy/web/app.wasm
	cp bazel-out/k8-fastbuild-ST-5df9d8eb9a43/bin/packages/goAppTest1/wasm/wasm_/wasm deploy/web/app.wasm


222_run_goAppTest1:
	./bazel-bin/packages/goAppTest1/linux_amd64_stripped/goAppTest1



reset_original: reset_remove
	cp originals/WORKSPACE .
	cp originals/BUILD .




#combine_filenames:
#	sed ':a; N; $!ba; s/\n/ /g' originals/workspace_files_to_use > originals/workspace_files_to_use_temp


#combined_workspace:
#	sed ':a; N; $!ba; s/\n/ /g' $q > WORKSPACE


#bashfix:
#	bash -c "(readarray -t ARRAY < originals/workspace_files_to_use; IFS=' ' && echo "${ARRAY[*]}")"

#cattest:
#	(readarray -t ARRAY < originals/workspace_files_to_use; IFS=' '; echo "${ARRAY[*]}")






