# gRPC-go simple and handy examples

# Installing dependencies ->

# Step1-
####################################################################################

Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip

Unzip
unzip protoc-3.3.0-linux-x86_64.zip -d protoc3

Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/

Optional: change owner
sudo chown $USER /usr/local/bin/protoc
sudo chown -R $USER /usr/local/include/google


######################################################################################

# Step2-
go get -u github.com/golang/protobuf/proto

# Step3-
go get -u github.com/golang/protobuf/protoc-gen-go
