namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

struct LoginRequest {
    1:  string user_name (api.raw = "user_name" api.vd = "len($) > 0 && len($) < 33>")
    2:  string password (api.raw = "password" api.vd = "len($) > 0 && len($) < 33>")
}

struct LogoutRequest {
    1: i32 uid (api.raw = "uid" api.vd = "len($) > 0 && len($) < 33>")
}

struct RegisterRequest {
    1: string user_name (api.raw = "user_name" api.vd = "len($) > 0 && len($) < 33>")
    2: string password (api.raw = "password" api.vd = "len($) > 0 && len($) < 33>")
    3: string phone_number (api.raw = "phone_number")
    4: string avatar_url (api.raw = "avatar_url")
    5: string email (api.raw = "email")
    6: string open_id (api.raw = "open_id")
    7: string personal_signature (api.raw = "personal_signature")
}

struct LoginResponse {
    1: i16 code,   // Status code, 0-success, other values-failure
    2: string message, // Return status description
    3: i32 data, // uid
}

service UserService {
    // 登陆
    LoginResponse Login(1: LoginRequest req) (api.post = "/login")
    // 注册
    LoginResponse Register(1: RegisterRequest req) (api.post = "/register")
    // 登出
    common.BaseResponse Logout(1: LogoutRequest req) (api.post = "/logout")
}