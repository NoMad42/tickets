// Package specs provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package specs

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// Defines values for BookingStatus.
const (
	BookingStatusBooked BookingStatus = "booked"

	BookingStatusBooking BookingStatus = "booking"

	BookingStatusFree BookingStatus = "free"
)

// Defines values for SeatPosition.
const (
	SeatPositionAisle SeatPosition = "aisle"

	SeatPositionMiddle SeatPosition = "middle"

	SeatPositionWindow SeatPosition = "window"
)

// Defines values for SeatType.
const (
	SeatTypeBuisness SeatType = "buisness"

	SeatTypeEconomy SeatType = "economy"
)

// Airport defines model for Airport.
type Airport struct {
	// Город аэропорта.
	City string `json:"city"`

	// Код аэропорта.
	Code string `json:"code"`

	// Страна аэропорта.
	Country string `json:"country"`

	// Описание аэропорта.
	Description *string `json:"description,omitempty"`

	// Идентификатор аэропорта.
	Id string `json:"id"`

	// Название аэропорта.
	Name string `json:"name"`
}

// AirportsList defines model for AirportsList.
type AirportsList []Airport

// Booking defines model for Booking.
type Booking struct {
	// Идентификатор рейса.
	FlightId string `json:"flight_id"`

	// Идентификатор услуги для места.
	Id string `json:"id"`

	// Идентификатор места.
	SeatId string `json:"seat_id"`

	// Идентификаторы услуг для места
	SeatOptionsIds *[]string `json:"seat_options_ids,omitempty"`

	// Статус бронирования.
	Status BookingStatus `json:"status"`

	// Идентификатор транзакции.
	TransactionId *string `json:"transaction_id,omitempty"`

	// Идентификатор пользователя.
	UserProfileId string `json:"user_profile_id"`
}

// Тело запроса на создание бронирования
type BookingCreateRequest struct {
	// Идентификатор места.
	SeatId string `json:"seat_id"`

	// Идентификатор пользователя.
	UserProfileId string `json:"user_profile_id"`
}

// BookingList defines model for BookingList.
type BookingList []Booking

// Статус бронирования.
type BookingStatus string

// Flight defines model for Flight.
type Flight struct {
	// Код рейса.
	Code string `json:"code"`

	// Время вылета.
	From time.Time `json:"from"`

	// Идентификатор аэропорта вылета.
	FromAirportId string `json:"from_airport_id"`

	// Идентификатор рейса.
	Id string `json:"id"`

	// Статус.
	Status string `json:"status"`

	// Время прилёта.
	To time.Time `json:"to"`

	// Идентификатор аэропорта прилёта.
	ToAirportId string `json:"to_airport_id"`
}

// FlightsList defines model for FlightsList.
type FlightsList []Flight

// Seat defines model for Seat.
type Seat struct {
	// Статус бронирования.
	BookingStatus BookingStatus `json:"booking_status"`

	// Код места.
	Code string `json:"code"`

	// Идентификатор рейса.
	FlightId string `json:"flight_id"`

	// Идентификатор места.
	Id string `json:"id"`

	// Место места.
	Position SeatPosition `json:"position"`

	// Цена места.
	Price float64 `json:"price"`

	// Тип места.
	Type SeatType `json:"type"`
}

// Место места.
type SeatPosition string

// Тип места.
type SeatType string

// SeatOption defines model for SeatOption.
type SeatOption struct {
	// Описание услуги для мета.
	Description *string `json:"description,omitempty"`

	// Идентификатор услуги для места.
	Id string `json:"id"`

	// Название услуги для мета.
	Name string `json:"name"`

	// Цена места.
	Price float64 `json:"price"`
}

// SeatOptionsList defines model for SeatOptionsList.
type SeatOptionsList []SeatOption

// SeatsList defines model for SeatsList.
type SeatsList []Seat

// Transaction defines model for Transaction.
type Transaction struct {
	// Сумма.
	Amount float64 `json:"amount"`

	// Идентификатор транзакции.
	Id string `json:"id"`

	// Время оплаты.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// TransactionsList defines model for TransactionsList.
type TransactionsList []Transaction

// UserProfile defines model for UserProfile.
type UserProfile struct {
	// URL по которому можно получить аватар пользователя.
	AvatarUrl string `json:"avatarUrl"`

	// Идентификатор пользователя.
	Id string `json:"id"`

	// Логин пользователя в системе.
	Login string `json:"login"`
}

// CreateBookingJSONBody defines parameters for CreateBooking.
type CreateBookingJSONBody BookingCreateRequest

// CreateBookingJSONRequestBody defines body for CreateBooking for application/json ContentType.
type CreateBookingJSONRequestBody CreateBookingJSONBody

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

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// GetAirportsList operation middleware
func (siw *ServerInterfaceWrapper) GetAirportsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAirportsList(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetBookingList operation middleware
func (siw *ServerInterfaceWrapper) GetBookingList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetBookingList(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateBooking operation middleware
func (siw *ServerInterfaceWrapper) CreateBooking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateBooking(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetFlightsList operation middleware
func (siw *ServerInterfaceWrapper) GetFlightsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFlightsList(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetSeatOptionsList operation middleware
func (siw *ServerInterfaceWrapper) GetSeatOptionsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSeatOptionsList(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetSeatsList operation middleware
func (siw *ServerInterfaceWrapper) GetSeatsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSeatsList(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetTransactionsList operation middleware
func (siw *ServerInterfaceWrapper) GetTransactionsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTransactionsList(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateTransaction operation middleware
func (siw *ServerInterfaceWrapper) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTransaction(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetAuthUser operation middleware
func (siw *ServerInterfaceWrapper) GetAuthUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAuthUser(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
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
