# service.localstore.meta
Microservice responsible for local storage metadata

It contains:

* a gRPC server
* a gRPC client

The protobuf spec is this

  
  syntax = "proto3";
  
  package localstore;
  
  service Local {
      rpc Home(HomeReq) returns (Void) {}
      rpc Mkdir(MkdirReq) returns (Void) {}
      rpc Stat(StatReq) returns (Metadata) {}
      rpc Cp(CpReq) returns (Void) {}
      rpc Mv(MvReq) returns (Void) {}
      rpc Rm(RmReq) returns (Void) {}
  }
  
  message Void {
  }
  
  message RmReq {
      Identity idt = 1;
      string path = 2;
  }
  
  message MvReq {
      Identity idt = 1;
      string src = 2;
      string dst = 3;
  }
  
  message HomeReq {
      Identity idt = 1;    
  }
  
  message CpReq {
      Identity idt = 1;
      string src = 2;
      string dst = 3;
  }
  
  message MkdirReq {
      Identity idt = 1;
      string path = 2;
  }
  
  message StatReq {
      Identity idt = 1;
      string path = 2;
      bool children = 3;
  }
  
  message Metadata {
      string id = 1;
      string path = 2;
      uint32 size = 3;
      bool is_container = 4;
      string mime_type = 5;
      string checksum = 6;
      uint32 modified = 7;
      string etag = 8; 
      uint32 permissions = 9;
      repeated Metadata children = 10;
  } 
  
  message Identity {
      string pid = 1;
      string idp = 2;
      string email = 3;
      string display_name = 4;
  }


