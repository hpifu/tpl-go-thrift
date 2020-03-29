namespace go echoservice

struct EchoReq {
    1: string rid;
    2: string message;
}

struct EchoRes {
    1: string message;
    2: string err;
}

service EchoService {
    EchoRes Echo(1: EchoReq req);
}
