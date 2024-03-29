// Package specs provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package specs

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/chi/v5"
)

// Defines values for BookingStatus.
const (
	BookingStatusBooked  BookingStatus = "booked"
	BookingStatusBooking BookingStatus = "booking"
	BookingStatusFree    BookingStatus = "free"
)

// Defines values for SeatPosition.
const (
	Aisle  SeatPosition = "aisle"
	Middle SeatPosition = "middle"
	Window SeatPosition = "window"
)

// Defines values for SeatType.
const (
	Buisness SeatType = "buisness"
	Economy  SeatType = "economy"
)

// Airport defines model for Airport.
type Airport struct {
	// City Город аэропорта.
	City string `json:"city"`

	// Code Код аэропорта.
	Code string `json:"code"`

	// Country Страна аэропорта.
	Country string `json:"country"`

	// Description Описание аэропорта.
	Description *string `json:"description,omitempty"`

	// Id Идентификатор аэропорта.
	Id openapi_types.UUID `json:"id"`

	// Name Название аэропорта.
	Name string `json:"name"`
}

// AirportsList defines model for AirportsList.
type AirportsList = []Airport

// Booking Бронирование
type Booking struct {
	// FlightId Идентификатор рейса.
	FlightId openapi_types.UUID `json:"flight_id"`

	// Id Идентификатор бронирования.
	Id openapi_types.UUID `json:"id"`

	// SeatId Идентификатор места.
	SeatId openapi_types.UUID `json:"seat_id"`

	// SeatOptionsIds Идентификаторы услуг для места
	SeatOptionsIds *[]openapi_types.UUID `json:"seat_options_ids,omitempty"`

	// Status Статус бронирования.
	Status BookingStatus `json:"status"`

	// TransactionId Идентификатор транзакции.
	TransactionId *openapi_types.UUID `json:"transaction_id,omitempty"`

	// UserProfileId Идентификатор пользователя.
	UserProfileId openapi_types.UUID `json:"user_profile_id"`
}

// BookingCreateRequest Тело запроса на создание бронирования
type BookingCreateRequest struct {
	// SeatId Идентификатор места.
	SeatId openapi_types.UUID `json:"seat_id"`

	// UserProfileId Идентификатор пользователя.
	UserProfileId openapi_types.UUID `json:"user_profile_id"`
}

// BookingList Список бронирований
type BookingList = []Booking

// BookingStatus Статус бронирования.
type BookingStatus string

// Flight defines model for Flight.
type Flight struct {
	// Code Код рейса.
	Code string `json:"code"`

	// From Время вылета.
	From time.Time `json:"from"`

	// FromAirportId Идентификатор аэропорта вылета.
	FromAirportId openapi_types.UUID `json:"from_airport_id"`

	// Id Идентификатор рейса.
	Id openapi_types.UUID `json:"id"`

	// Status Статус.
	Status string `json:"status"`

	// To Время прилёта.
	To time.Time `json:"to"`

	// ToAirportId Идентификатор аэропорта прилёта.
	ToAirportId openapi_types.UUID `json:"to_airport_id"`
}

// FlightsList defines model for FlightsList.
type FlightsList = []Flight

// Seat defines model for Seat.
type Seat struct {
	// BookingStatus Статус бронирования.
	BookingStatus BookingStatus `json:"booking_status"`

	// Code Код места.
	Code string `json:"code"`

	// FlightId Идентификатор рейса.
	FlightId openapi_types.UUID `json:"flight_id"`

	// Id Идентификатор места.
	Id openapi_types.UUID `json:"id"`

	// Position Место места.
	Position SeatPosition `json:"position"`

	// Price Цена места.
	Price float64 `json:"price"`

	// Type Тип места.
	Type SeatType `json:"type"`
}

// SeatPosition Место места.
type SeatPosition string

// SeatType Тип места.
type SeatType string

// SeatOption defines model for SeatOption.
type SeatOption struct {
	// Description Описание услуги для места.
	Description *string `json:"description,omitempty"`

	// Id Идентификатор услуги для места.
	Id openapi_types.UUID `json:"id"`

	// Name Название услуги для места.
	Name string `json:"name"`

	// Price Цена места.
	Price float64 `json:"price"`
}

// SeatOptionsList defines model for SeatOptionsList.
type SeatOptionsList = []SeatOption

// SeatsList defines model for SeatsList.
type SeatsList = []Seat

// Transaction Транзакция
type Transaction struct {
	// Amount Сумма.
	Amount float64 `json:"amount"`

	// Id Идентификатор транзакции.
	Id openapi_types.UUID `json:"id"`

	// Timestamp Время оплаты.
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// UserProfileId Идентификатор пользователя.
	UserProfileId openapi_types.UUID `json:"user_profile_id"`
}

// TransactionCreateRequest Тело запроса на создание транзакции
type TransactionCreateRequest struct {
	// BookingId Идентификатор бронирования.
	BookingId openapi_types.UUID `json:"booking_id"`
}

// TransactionsList Список транзакций
type TransactionsList = []Transaction

// UserProfile defines model for UserProfile.
type UserProfile struct {
	// AvatarUrl URL по которому можно получить аватар пользователя.
	AvatarUrl string `json:"avatarUrl"`

	// Id Идентификатор пользователя.
	Id openapi_types.UUID `json:"id"`

	// Login Логин пользователя в системе.
	Login string `json:"login"`
}

// CreateBookingJSONRequestBody defines body for CreateBooking for application/json ContentType.
type CreateBookingJSONRequestBody = BookingCreateRequest

// CreateTransactionJSONRequestBody defines body for CreateTransaction for application/json ContentType.
type CreateTransactionJSONRequestBody = TransactionCreateRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получить список аэропортов.
	// (GET /v1/airports)
	GetAirportsList(w http.ResponseWriter, r *http.Request)
	// Получить список заявок на бронирование.
	// (GET /v1/booking)
	GetBookingList(w http.ResponseWriter, r *http.Request)
	// Создать заявку на бронирования
	// (POST /v1/booking)
	CreateBooking(w http.ResponseWriter, r *http.Request)
	// Получить список рейсов.
	// (GET /v1/flights)
	GetFlightsList(w http.ResponseWriter, r *http.Request)
	// Получить список услуг для места.
	// (GET /v1/seat_options)
	GetSeatOptionsList(w http.ResponseWriter, r *http.Request)
	// Получить список мест.
	// (GET /v1/seats)
	GetSeatsList(w http.ResponseWriter, r *http.Request)
	// Получить место по id
	// (GET /v1/seats/{seatId})
	GetSeatById(w http.ResponseWriter, r *http.Request, seatId openapi_types.UUID)
	// Получить список транзакций.
	// (GET /v1/transactions)
	GetTransactionsList(w http.ResponseWriter, r *http.Request)
	// Создать транзакцию.
	// (POST /v1/transactions)
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	// Информация об аутентифицированном пользователе.
	// (GET /v1/user)
	GetAuthUser(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Получить список аэропортов.
// (GET /v1/airports)
func (_ Unimplemented) GetAirportsList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Получить список заявок на бронирование.
// (GET /v1/booking)
func (_ Unimplemented) GetBookingList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Создать заявку на бронирования
// (POST /v1/booking)
func (_ Unimplemented) CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Получить список рейсов.
// (GET /v1/flights)
func (_ Unimplemented) GetFlightsList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Получить список услуг для места.
// (GET /v1/seat_options)
func (_ Unimplemented) GetSeatOptionsList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Получить список мест.
// (GET /v1/seats)
func (_ Unimplemented) GetSeatsList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Получить место по id
// (GET /v1/seats/{seatId})
func (_ Unimplemented) GetSeatById(w http.ResponseWriter, r *http.Request, seatId openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Получить список транзакций.
// (GET /v1/transactions)
func (_ Unimplemented) GetTransactionsList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Создать транзакцию.
// (POST /v1/transactions)
func (_ Unimplemented) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Информация об аутентифицированном пользователе.
// (GET /v1/user)
func (_ Unimplemented) GetAuthUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetAirportsList operation middleware
func (siw *ServerInterfaceWrapper) GetAirportsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAirportsList(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetBookingList operation middleware
func (siw *ServerInterfaceWrapper) GetBookingList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetBookingList(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateBooking operation middleware
func (siw *ServerInterfaceWrapper) CreateBooking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateBooking(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetFlightsList operation middleware
func (siw *ServerInterfaceWrapper) GetFlightsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFlightsList(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSeatOptionsList operation middleware
func (siw *ServerInterfaceWrapper) GetSeatOptionsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSeatOptionsList(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSeatsList operation middleware
func (siw *ServerInterfaceWrapper) GetSeatsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSeatsList(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSeatById operation middleware
func (siw *ServerInterfaceWrapper) GetSeatById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "seatId" -------------
	var seatId openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "seatId", runtime.ParamLocationPath, chi.URLParam(r, "seatId"), &seatId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "seatId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSeatById(w, r, seatId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetTransactionsList operation middleware
func (siw *ServerInterfaceWrapper) GetTransactionsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTransactionsList(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateTransaction operation middleware
func (siw *ServerInterfaceWrapper) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTransaction(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetAuthUser operation middleware
func (siw *ServerInterfaceWrapper) GetAuthUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAuthUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/airports", wrapper.GetAirportsList)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/booking", wrapper.GetBookingList)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/booking", wrapper.CreateBooking)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/flights", wrapper.GetFlightsList)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/seat_options", wrapper.GetSeatOptionsList)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/seats", wrapper.GetSeatsList)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/seats/{seatId}", wrapper.GetSeatById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/transactions", wrapper.GetTransactionsList)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/transactions", wrapper.CreateTransaction)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/user", wrapper.GetAuthUser)
	})

	return r
}
