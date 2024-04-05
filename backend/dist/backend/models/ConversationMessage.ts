/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type ConversationMessage = {
    content: string;
    role: ConversationMessage.role;
};
export namespace ConversationMessage {
    export enum role {
        UNSPECIFIED = 'unspecified',
        USER = 'user',
        ASSISTANT = 'assistant',
    }
}

