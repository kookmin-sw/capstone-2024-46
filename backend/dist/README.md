# Tickets Client SDK

private-llm-backend 서버에서 제공하는 API 를 자동 생성된 SDK 를 이용하여 각 method 로 호출 할 수 있습니다.
제공하는 API 의 목록은 서버의 `/docs` 에서 제공하는 swagger-ui 를 통해 확인할 수 있습니다.

## Dependency

- TypeScript
- fetch (대부분의 browser 환경에서 built-in 으로 제공되므로 별도 dependency 가 없습니다.)

## 시작하기

### SDK 복사

`backend` 디렉토리를 자신의 프로젝트에 복사합니다. 이 작업은 private-llm-backend 서버가 업데이트 될때마다 반복되어야 합니다.

```bash
git clone .../private-llm-backend
cd private-llm-backend
cp -r dist/backend /path/to/your/project/
```

### 사용법

`example.ts` 파일은 `CreateChatConversation`의 사용 예를 제공합니다.

```typescript
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
```
