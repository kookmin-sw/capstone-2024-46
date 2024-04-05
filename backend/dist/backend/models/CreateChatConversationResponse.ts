/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ConversationMessage } from './ConversationMessage';
export type CreateChatConversationResponse = {
    /**
     * A unique identifier for the chat conversation.
     */
    id: string;
    /**
     * The generated chat conversation message.
     */
    message: ConversationMessage;
    /**
     * The Unix timestamp (in seconds) of when the chat completion was created.
     */
    createTime: string;
    /**
     * This fingerprint represents the backend configuration that the model runs with.
     */
    systemFingerprint?: string;
    /**
     * The object type, which is always `chat_completion`.
     */
    objectType: CreateChatConversationResponse.objectType;
};
export namespace CreateChatConversationResponse {
    /**
     * The object type, which is always `chat_completion`.
     */
    export enum objectType {
        UNSPECIFIED = 'unspecified',
        CHAT_COMPLETION = 'chat_completion',
    }
}

