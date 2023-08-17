namespace go api

struct SendMessageRequest {
    1: string token,
    2: string to_user_id,
    3: string action_type,
    4: string content,
}

struct SendMessageResponse {
    1: i64 status_code,
    2: string status_msg,
}

struct GetMessageListRequest {
    1: string token,
    2: string to_user_id,
}

struct Message {
    1: i64 id,
    2: i64 to_user_id,
    3: i64 from_user_id,
    4: string content,
    5: i64 create_time
}

struct GetMessageListResponse {
    1: i64 status_code,
    2: string status_msg,
    3: list<Message> message_list,
}

service MessageService {
    SendMessageResponse sendMessage(1: SendMessageRequest req) 
    GetMessageListResponse getMessageList(1: GetMessageListRequest req)
}