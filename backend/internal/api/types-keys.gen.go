// Code generated by type_keys_generator DO NOT EDIT.
package api

func (c ConversationMessage) Keys() map[string]struct{} {
	return map[string]struct{}{"content": {}, "role": {}}
}

func (c CreateChatConversationRequest) Keys() map[string]struct{} {
	return map[string]struct{}{"messages": {}, "systemVersion": {}, "threadId": {}}
}

func (c CreateChatConversationResponse) Keys() map[string]struct{} {
	return map[string]struct{}{"createTime": {}, "id": {}, "message": {}, "objectType": {}, "systemFingerprint": {}}
}

func (c CreateThreadResponse) Keys() map[string]struct{} {
	return map[string]struct{}{"threadId": {}}
}

func (s Status) Keys() map[string]struct{} {
	return map[string]struct{}{"code": {}, "details": {}, "message": {}}
}