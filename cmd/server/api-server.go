package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"resume-backend/pkg/config"
	"resume-backend/pkg/database"
	"resume-backend/dto"
	"resume-backend/pkg/middleware"
	"resume-backend/pkg/service"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Load()

	if cfg.RootKey == "" {
		log.Fatal("ROOT_KEY environment variable is required")
	}

	db, err := database.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.CORS())

	rootKeyAuth := middleware.RootKeyAuth(cfg.RootKey)

	bootcampService := service.NewBootcampService(db)
	journalService := service.NewJournalService(db)
	memeService := service.NewMemeService(db)
	storyService := service.NewStoryService(db)
	contactService := service.NewContactService(db)
	chatService := service.NewChatService(db)
	openRouterService := service.NewOpenRouterService()

	api := e.Group("/api")

	api.GET("/bootcamps", func(c echo.Context) error {
		bootcamps, err := bootcampService.GetAll()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, bootcamps)
	})

	api.POST("/bootcamps", func(c echo.Context) error {
		var req dto.CreateBootcampRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if errors := middleware.ValidateStruct(req); len(errors) > 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
		}
		bootcamp, err := bootcampService.Create(req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, bootcamp)
	}, rootKeyAuth)

	api.GET("/bootcamps/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		bootcamp, err := bootcampService.GetByID(id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, bootcamp)
	})

	api.PUT("/bootcamps/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		var req dto.UpdateBootcampRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if errors := middleware.ValidateStruct(req); len(errors) > 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
		}
		bootcamp, err := bootcampService.Update(id, req)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, bootcamp)
	}, rootKeyAuth)

	api.DELETE("/bootcamps/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		if err := bootcampService.Delete(id); err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Bootcamp deleted successfully"})
	}, rootKeyAuth)

	api.GET("/journal", func(c echo.Context) error {
		entries, err := journalService.GetAll()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, entries)
	})

	api.POST("/journal", func(c echo.Context) error {
		var req dto.CreateJournalRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if errors := middleware.ValidateStruct(req); len(errors) > 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
		}
		entry, err := journalService.Create(req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, entry)
	}, rootKeyAuth)

	api.GET("/journal/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		entry, err := journalService.GetByID(id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, entry)
	})

	api.PUT("/journal/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		var req dto.UpdateJournalRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if errors := middleware.ValidateStruct(req); len(errors) > 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
		}
		entry, err := journalService.Update(id, req)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, entry)
	}, rootKeyAuth)

	api.DELETE("/journal/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		if err := journalService.Delete(id); err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Journal entry deleted successfully"})
	}, rootKeyAuth)

	api.GET("/memes/categories", func(c echo.Context) error {
		categories, err := memeService.GetAllCategoriesWithMemes()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, categories)
	})

	api.POST("/memes/categories", func(c echo.Context) error {
		var req dto.CreateMemeCategoryRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if errors := middleware.ValidateStruct(req); len(errors) > 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
		}
		category, err := memeService.CreateCategory(req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, category)
	}, rootKeyAuth)

	api.GET("/memes/categories/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		category, err := memeService.GetCategoryWithMemes(id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, category)
	})

	api.DELETE("/memes/categories/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		if err := memeService.DeleteCategory(id); err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Category deleted successfully"})
	}, rootKeyAuth)

	api.POST("/memes", func(c echo.Context) error {
		var req dto.CreateMemeRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if errors := middleware.ValidateStruct(req); len(errors) > 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
		}
		meme, err := memeService.CreateMeme(req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, meme)
	}, rootKeyAuth)

	api.GET("/memes/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		meme, err := memeService.GetMemeByID(id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, meme)
	})

	api.PUT("/memes/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		var req dto.UpdateMemeRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if errors := middleware.ValidateStruct(req); len(errors) > 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
		}
		meme, err := memeService.UpdateMeme(id, req)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, meme)
	}, rootKeyAuth)

	api.DELETE("/memes/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		if err := memeService.DeleteMeme(id); err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Meme deleted successfully"})
	}, rootKeyAuth)

	api.GET("/stories", func(c echo.Context) error {
		stories, err := storyService.GetAll()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, stories)
	})

	api.POST("/stories", func(c echo.Context) error {
		var req dto.CreateStoryRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if errors := middleware.ValidateStruct(req); len(errors) > 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
		}
		story, err := storyService.Create(req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, story)
	}, rootKeyAuth)

	api.GET("/stories/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		story, err := storyService.GetByID(id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, story)
	})

	api.PUT("/stories/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		var req dto.UpdateStoryRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if errors := middleware.ValidateStruct(req); len(errors) > 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
		}
		story, err := storyService.Update(id, req)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, story)
	}, rootKeyAuth)

	api.DELETE("/stories/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		if err := storyService.Delete(id); err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Story deleted successfully"})
	}, rootKeyAuth)

	api.GET("/contacts", func(c echo.Context) error {
		page := 1
		pageSize := 20
		if pageStr := c.QueryParam("page"); pageStr != "" {
			if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
				page = parsedPage
			}
		}
		if pageSizeStr := c.QueryParam("page_size"); pageSizeStr != "" {
			if parsedPageSize, err := strconv.Atoi(pageSizeStr); err == nil && parsedPageSize > 0 && parsedPageSize <= 100 {
				pageSize = parsedPageSize
			}
		}
		contacts, total, err := contactService.GetAll(page, pageSize)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"contacts":    contacts,
			"total":       total,
			"page":        page,
			"page_size":   pageSize,
			"total_pages": (total + pageSize - 1) / pageSize,
		})
	}, rootKeyAuth)

	api.POST("/contacts", func(c echo.Context) error {
		var req dto.CreateContactRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if errors := middleware.ValidateStruct(req); len(errors) > 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
		}
		contact, err := contactService.Create(req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, contact)
	})

	api.GET("/contacts/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		contact, err := contactService.GetByID(id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, contact)
	}, rootKeyAuth)

	api.DELETE("/contacts/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}
		if err := contactService.Delete(id); err != nil {
			if strings.Contains(err.Error(), "not found") {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Contact deleted successfully"})
	}, rootKeyAuth)

	api.POST("/chat", func(c echo.Context) error {
		var req dto.ChatRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
		if req.Message == "" || req.SessionID == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "message and session_id are required")
		}

		session, err := chatService.GetOrCreateSession(req.SessionID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get or create session: "+err.Error())
		}

		lastMessages, err := chatService.GetLastMessages(session.ID, 5)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get last messages: "+err.Error())
		}

		_, err = chatService.SaveMessage(session.ID, "user", req.Message, nil)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save user message: "+err.Error())
		}

		model := cfg.OpenRouterModel
		response, statusCode, err := openRouterService.CreateChatCompletionWithContext(
			lastMessages,
			req.Message,
			model,
			c.Request().Header.Get("Referer"),
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if statusCode != http.StatusOK {
			return c.JSON(statusCode, response)
		}

		choices, ok := response["choices"].([]interface{})
		if !ok || len(choices) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Invalid response from OpenRouter")
		}

		choice, ok := choices[0].(map[string]interface{})
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Invalid response format")
		}

		message, ok := choice["message"].(map[string]interface{})
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Invalid message format")
		}

		content, ok := message["content"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Invalid content format")
		}

		reasoningDetails, _ := message["reasoning_details"]

		assistantMessageID, err := chatService.SaveMessage(session.ID, "assistant", content, reasoningDetails)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save assistant message: "+err.Error())
		}

		chatResponse := dto.ChatResponse{
			Answer:    content,
			SessionID: session.ID,
			MessageID: assistantMessageID,
		}

		return c.JSON(http.StatusOK, chatResponse)
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	log.Printf("Server starting on %s", addr)
	log.Printf("Environment: %s", cfg.Environment)
	log.Fatal(e.Start(addr))
}

