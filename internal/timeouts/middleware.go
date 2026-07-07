package timeouts

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
)

const defaultTimeout = 10 * time.Second

func WithDefaultTimeoutHandler() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		goCtx, cancel := context.WithTimeout(ctx.Context(), defaultTimeout)
		defer cancel()

		ctx.SetContext(goCtx)

		chanel := make(chan error, 1)
		go func() {
			chanel <- ctx.Next()
		}()

		select {
		case err := <-chanel:
			if err != nil {
				var fiberErr *fiber.Error
				if errors.As(err, &fiberErr) {
					return ctx.Status(fiberErr.Code).JSON(fiber.Map{"error": fiberErr.Message})
				}
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			return nil
		case <-goCtx.Done():
			return ctx.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{"error": "request timeout"})
		}
	}
}
