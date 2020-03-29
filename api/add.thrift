namespace go addservice

struct AddReq {
    1: i64 a;
    2: i64 b;
}

struct AddRes {
    1: i64 v;
    2: string err;
}

service AddService {
    AddRes Add(1: AddReq req);
}
