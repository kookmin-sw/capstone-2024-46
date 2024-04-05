/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ConversationMessage } from './ConversationMessage';
export type CreateChatConversationRequest = {
    /**
     * system version for testing
     */
    systemVersion?: string;
    /**
     * array of messages
     */
    messages: Array<ConversationMessage>;
    /**
     * thread id
     */
    threadId: string;
};

