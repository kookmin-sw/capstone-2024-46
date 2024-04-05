import {ApiError, BackendClient, ConversationMessage} from "./backend";

const appClient = new BackendClient({
    BASE: 'http://localhost:9090',
    WITH_CREDENTIALS: true,
    CREDENTIALS: 'include',
    HEADERS: {
        'X-Api-Key': '5b5265e7d5034c1db7337f73bda0e53b',
    },
});

async function main() {
    try {
        const resp = await appClient.playgroundChatService.playgroundChatServiceCreateChatConversation(
            {
                messages: [
                    {
                        content: "hello",
                        role: ConversationMessage.role.USER,
                    }
                ],
                systemVersion: "dev",
            }
        );
        console.log(resp);
    } catch (error) {
        if (error instanceof ApiError) {
            const errResp = error.body;
            console.log("ApiError", error.status, errResp.code, errResp.message);
            return
        }

        // Maybe network error
        throw error;
    }
}

main().catch(console.error);
