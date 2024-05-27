# RAG 기능을 갖춘 Private LLM 서비스 솔루션 POC

<!--![Frame 2](https://github.com/kookmin-sw/capstone-2024-46/assets/13215190/0349c04f-7b10-4754-87e7-9ba5184557be)-->

# 1. 프로젝트 소개 
기업들이 LLM(Large Language Model)을 활용하여 사내 데이터를 기반으로 정확하고 관련성 높은 응답을 생성할 수 있도록 지원하는 솔루션입니다.
비전문가도 쉽게 이용할 수 있는 UI 를 제공하고, RAG(Retrieval Augmented Generation) 기술을 활용하여 최신 정보를 반영하고 환각 현상을 최소화하는 것이 주요 목표입니다.

![image](https://github.com/kookmin-sw/capstone-2024-46/assets/55116920/34916c62-5afe-4168-855d-216833fd74ee)

## 주요 기능

1. **데이터 가공 및 정제**: 고객이 보유한 다양한 형식(Excel, PDF 등)의 데이터를 효과적으로 가공하고 정제합니다.
2. **사용자 친화적인 데이터 업로드 환경**: 비 개발자도 쉽게 데이터를 업로드할 수 있는 직관적인 인터페이스를 제공합니다.
3. **데이터 업데이트 용이성**: 고객이 새로운 데이터를 추가하거나 기존 데이터를 수정할 수 있는 편리한 방법을 제공합니다.
4. **LLM 통합**: 업로드된 데이터를 효과적으로 LLM 과 통합하며, 클라우드 또는 자체 호스팅 모델 중 선택할 수 있습니다.
5. **챗봇 플러그인**: 사내 메신저 또는 웹에 통합할 수 있는 챗봇 형태의 플러그인을 제공합니다.

# 2. 시스템 구조
![ ](https://github.com/kookmin-sw/capstone-2024-46/assets/13215190/d505c9d5-455b-4dd6-8a94-57323d913ed9)

- Echo (Golang) - User Endpoint, Account, History, Slack Integration
- Flask (Python) - Retrieval, LLM Integration 
- Celery (Python) - Distribute jobs for Chunking, Embedding


# 3. 데이터 / LLM 흐름
![Screenshot 2024-05-24 at 11 18 49 AM](https://github.com/kookmin-sw/capstone-2024-46/assets/13215190/fe8b7649-0d6d-43bc-aac7-93bb3a12a3e2)

### 데이터 품질 최적화
- Semantic Chunk
- Structured Data
- Metadata

### 답변 품질 향상
- Time Weight
- Fact Verification

### 쿼리 최적화
- Hypothetical Document Embedding (HyDE)
- Multi Query

# 4. 평가
아래 평가 기준에 따라 질문과 답변 쌍의 데이터 세트를 사람이 작성하여 LLM 성능을 평가했습니다.

## 평가기준
- 적절한 답변을 생성하는가 (Relevance)
- 진실된 답변을 생성하는가 (Truthfulness)
- 부적절한 질문에 답변을 잘 거부하는가 (Adversarial Test)

## 평가 결과
#### 평균
```
🟦 RAG: 8.9
🟥 GPT-4: 3.83
```
![Screenshot 2024-05-24 at 11 15 43 AM](https://github.com/kookmin-sw/capstone-2024-46/assets/13215190/a3f5177d-ba9e-41c9-b447-26d617a3f86a)

# 5. 팀 소개

- **이민철 (20163137)** - AI / Infra

- **이민재 (20191638)** - AI / Backend

- **조준형 (20163161)** - Frontend

# 6. 레포지토리
- [Backend](backend)
- [Frontend](frontend)
- [AI Gateway](gateway)
