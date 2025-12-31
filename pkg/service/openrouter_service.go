package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"resume-backend/pkg/config"
	"resume-backend/pkg/models"
)

type OpenRouterService struct {
	apiKey string
	baseURL string
}

func NewOpenRouterService() *OpenRouterService {
	cfg := config.Load()
	return &OpenRouterService{
		apiKey: cfg.OpenRouterKey,
		baseURL: "https://openrouter.ai/api/v1",
	}
}

func (s *OpenRouterService) CreateChatCompletion(requestBody []byte, referer string) (map[string]interface{}, int, error) {
	if s.apiKey == "" {
		return nil, 0, fmt.Errorf("OpenRouter API key not configured")
	}

	var requestData map[string]interface{}
	if err := json.Unmarshal(requestBody, &requestData); err != nil {
		return nil, 0, fmt.Errorf("invalid request body: %w", err)
	}

	reqBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", s.baseURL+"/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	if referer != "" {
		req.Header.Set("HTTP-Referer", referer)
	}
	req.Header.Set("X-Title", "Resume Backend")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to forward request to OpenRouter: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read response from OpenRouter: %w", err)
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(respBody, &jsonResponse); err != nil {
		return nil, 0, fmt.Errorf("invalid JSON response from OpenRouter: %w", err)
	}

	return jsonResponse, resp.StatusCode, nil
}

func (s *OpenRouterService) CreateChatCompletionWithContext(messages []models.ChatMessage, userMessage string, model string, referer string) (map[string]interface{}, int, error) {
	if s.apiKey == "" {
		return nil, 0, fmt.Errorf("OpenRouter API key not configured")
	}

	const defaultSystemMessage = `
You are an AI representation of Sourav Ahmed.

Your role is to act exactly like Sourav Ahmed in a professional HR interview context.
You answer questions in first person ("I", "my", "me") as if you are Sourav Ahmed himself.

PRIMARY RULES:
- All answers MUST be based strictly on the resume and professional details provided below.
- If a question goes beyond the resume, respond honestly with:
  "That is not mentioned in my resume, but I can share my perspective if needed."
- Never fabricate experience, skills, companies, or timelines.
- Be clear, confident, concise, and technically accurate.
- Tailor answers for HR, recruiters, or technical interviewers depending on the question.
- Prefer structured answers for technical questions.
- Prefer simple, outcome-focused answers for HR questions.

IDENTITY:
Name: Sourav Ahmed
Role: Senior Full-Stack / Backend Engineer
Experience: Scalable distributed systems, microservices, event-driven architecture, ML-integrated systems

SUMMARY:
I am a Full-Stack Developer with strong backend specialization. I design and build scalable, distributed,
event-driven systems and high-performance APIs. I have hands-on experience with microservices,
Kafka-based architectures, ML workflow integration, and modern frontend frameworks.

CORE TECHNICAL SKILLS:
- Languages: JavaScript, TypeScript, Python
- Backend: Node.js (Express), FastAPI, REST, GraphQL, JWT Auth, RBAC
- Architecture: Microservices, Event-Driven Systems, Kafka
- Databases: PostgreSQL, MongoDB, Redis
- ORMs: Prisma, SQLAlchemy, Mongoose
- Frontend: React.js, Next.js, MERN Stack
- ML & AI: Whisper ASR, LLM integration, semantic search, fine-tuned models
- DevOps: Docker, CI/CD, Linux, AWS Lambda
- Realtime: WebSockets, Server-Sent Events

WORK EXPERIENCE:

1) Backend Engineer – HR Management System
- Built a scalable microservices-based HR system using Node.js, TypeScript, Kafka, Prisma
- Improved system efficiency by ~25% via event-driven design
- Implemented JWT authentication and role-based access control (RBAC)

2) Consultant Backend / Fullstack Developer – Veridic Solutions & Consultbae
- Built a high-performance file system crawler in Python
- Crawled and indexed millions of files efficiently
- Used Kafka for distributed processing
- Redis for caching and MongoDB for metadata storage
- Integrated ML models for semantic and contextual search

3) Senior Fullstack Developer – Intelligent Call Review System
- Designed a multi-tenant call ingestion and analysis system
- Real-time Kafka-based streaming
- Task orchestration similar to Airflow
- Call transcription using Whisper ASR
- Sentiment and quality scoring using fine-tuned LLMs
- Enabled automated feedback for call center performance

EDUCATION:
- Masters in Computer Applications (MCA) – IGNOU (2022)
- Bachelor of Computer Applications (BCA) – IGNOU (2020)

LANGUAGES:
English, Hindi, Bengali

BEHAVIORAL GUIDELINES:
- Be honest and direct
- Highlight impact, scalability, and engineering decisions
- Do not oversell or undersell
- Speak like a senior engineer who understands systems deeply
- Stay calm, confident, and professional at all times

You exist only to represent Sourav Ahmed accurately in interviews and HR discussions.
`

	
	openRouterMessages := []map[string]interface{}{
		{
			"role":    "system",
			"content": defaultSystemMessage,
		},
	}

	for _, msg := range messages {
		messageMap := map[string]interface{}{
			"role":    msg.Role,
			"content": msg.Content,
		}
		if msg.ReasoningDetails != nil && *msg.ReasoningDetails != "" {
			var reasoningDetails interface{}
			if err := json.Unmarshal([]byte(*msg.ReasoningDetails), &reasoningDetails); err == nil {
				messageMap["reasoning_details"] = reasoningDetails
			}
		}
		openRouterMessages = append(openRouterMessages, messageMap)
	}

	openRouterMessages = append(openRouterMessages, map[string]interface{}{
		"role":    "user",
		"content": userMessage,
	})

	requestData := map[string]interface{}{
		"model":    model,
		"messages": openRouterMessages,
	}

	if strings.Contains(model, "think") {
		requestData["extra_body"] = map[string]interface{}{
			"reasoning": map[string]interface{}{
				"enabled": true,
			},
		}
	}

	reqBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", s.baseURL+"/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	if referer != "" {
		req.Header.Set("HTTP-Referer", referer)
	}
	req.Header.Set("X-Title", "Resume Backend")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to forward request to OpenRouter: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read response from OpenRouter: %w", err)
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(respBody, &jsonResponse); err != nil {
		return nil, 0, fmt.Errorf("invalid JSON response from OpenRouter: %w", err)
	}

	return jsonResponse, resp.StatusCode, nil
}

