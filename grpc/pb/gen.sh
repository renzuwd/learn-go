echo "生成 rpc server 代码"

OUT=../server/pb

mkdir -p ${OUT}
protoc \
--go_out=${OUT} \
--go-grpc_out=${OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
*.proto

mv ${OUT}/pb/*  ${OUT}/
rm -rf ${OUT}/pb


echo "生成 rpc client 代码"

OUT=../client/pb

mkdir -p ${OUT}
protoc \
--go_out=${OUT} \
--go-grpc_out=${OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
*.proto

mv ${OUT}/pb/*  ${OUT}/
rm -rf ${OUT}/pb
