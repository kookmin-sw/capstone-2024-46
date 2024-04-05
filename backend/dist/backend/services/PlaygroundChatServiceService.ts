/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CreateChatConversationRequest } from '../models/CreateChatConversationRequest';
import type { CreateChatConversationResponse } from '../models/CreateChatConversationResponse';
import type { CreateThreadResponse } from '../models/CreateThreadResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class PlaygroundChatServiceService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * Generate a chat conversation message.
     * @param requestBody
     * @returns CreateChatConversationResponse OK
     * @throws ApiError
     */
    public playgroundChatServiceCreateChatConversation(
        requestBody: CreateChatConversationRequest,
    ): CancelablePromise<CreateChatConversationResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/v1/playground/chat/conversation',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Creates a new thread.
     * @returns CreateThreadResponse OK
     * @throws ApiError
     */
    public playgroundChatServiceCreateThread(): CancelablePromise<CreateThreadResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/v1/threads',
        });
    }
}
