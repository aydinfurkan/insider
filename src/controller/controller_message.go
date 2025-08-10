package controller

import (
	"insider/src/db"
	"insider/src/domain"
	"insider/src/infra/api"
	"insider/src/infra/myerror"
	"insider/src/service"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CreateMessageRequest struct {
	Content         string `json:"content"`
	RecipientNumber string `json:"recipient_number"`
}

type MessageController struct {
	db             *db.MessageDb
	messageService *service.MessageService
}

func NewMessageController(messageDb *db.MessageDb, messageService *service.MessageService) *MessageController {
	return &MessageController{
		db:             messageDb,
		messageService: messageService,
	}
}

// CreateMessage creates a new message
// @Summary Create a new message
// @Description Create a new message with content and recipient number. Phone number must be in E.164 international format: +[country code][number]
// @Tags messages
// @Accept json
// @Produce json
// @Param message body CreateMessageRequest true "Message data"
// @Success 202 {object} api.ApiResponse{data=string} "Message created successfully"
// @Failure 400 {object} api.ApiResponse{error=api.ApiResponseError} "Bad request - Invalid phone number format or validation error"
// @Failure 500 {object} api.ApiResponse{error=api.ApiResponseError} "Internal server error"
// @Router /messages [post]
func (m *MessageController) CreateMessage(ctx echo.Context) error {

	var req CreateMessageRequest

	if err := ctx.Bind(&req); err != nil {
		return myerror.NewBadRequestError(err, "Invalid request body", 4001)
	}

	if len(req.Content) > 100 {
		return myerror.NewBadRequestError(echo.ErrBadRequest, "Content length must be at most 100 characters", 4002)
	}

	matched, _ := regexp.MatchString(`^\+[1-9]\d{1,14}$`, req.RecipientNumber)
	if !matched {
		return myerror.NewBadRequestError(echo.ErrBadRequest, "Invalid phone number format. Must be in E.164 format (+[country code][number])", 4003)
	}

	message := domain.NewMessage(req.Content, req.RecipientNumber)

	err := m.db.CreateMessage(message)
	if err != nil {
		return myerror.NewInternalServerError(err, "Failed to create message", 5001)
	}

	return ctx.JSON(http.StatusAccepted, api.NewSuccessResponse(message.MessageId))
}

// GetSentMessages retrieves sent messages with pagination
// @Summary Get sent messages
// @Description Retrieve sent messages with optional pagination
// @Tags messages
// @Accept json
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {object} api.ApiResponse{data=[]domain.Message} "List of sent messages"
// @Failure 500 {object} api.ApiResponse{error=api.ApiResponseError} "Internal server error"
// @Router /sentmessages [get]
func (m *MessageController) GetSentMessages(ctx echo.Context) error {
	offset, err := strconv.Atoi(ctx.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	messages, err := m.db.GetSentMessages(offset, limit)
	if err != nil {
		return myerror.NewInternalServerError(err, "Failed to get sent messages", 5001)
	}

	return ctx.JSON(http.StatusOK, api.NewSuccessResponse(messages))
}

// TopggleMessageService toggles the message service start/stop state
// @Summary Start or stop message service
// @Description Toggles the message service between running and stopped states
// @Tags messages
// @Accept json
// @Produce json
// @Success 200 {object} api.ApiResponse{data=string} "Service state toggled successfully"
// @Failure 500 {object} api.ApiResponse{error=api.ApiResponseError} "Internal server error"
// @Router /messages/toggle [post]
func (m *MessageController) ToggleMessageService(ctx echo.Context) error {
	result := m.messageService.Toggle()
	if result {
		return ctx.JSON(http.StatusOK, api.NewSuccessResponse("Message service started successfully"))
	} else {
		return ctx.JSON(http.StatusOK, api.NewSuccessResponse("Message service stopped successfully"))
	}
}
