namespace go base

struct BaseResponse {
    1: i16 code,   // Status code, 0-success, other values-failure
    2: string message, // Return status description
}