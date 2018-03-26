package text

import (
	"errors"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger("Test Logger", true)
	logger.AddOutput(func(message []byte) {
		os.Stdout.Write(message)
	})
	logger.WriteString("Raw message")
	logger.Info("Logger working.")
	logger.Debug("Debug message.", "another debug")
	var err error
	logger.LogError(err) // err is nil, does not print anything.
	err = errors.New("error")
	logger.LogError(err)
	logger.LogStack()

	logger.Wait()
}

func TestDefault(t *testing.T) {
	DefaultLogger.Debug("Debug message")
	DefaultLogger.Notice("Notice message")
	DefaultLogger.LogStack()
	DefaultLogger.Wait()
}

func TestMultipleWait(t *testing.T) {
	logger := NewLogger("Test Logger", true)
	logger.AddOutput(func(message []byte) {
		os.Stdout.Write(message)
	})
	logger.LogStack()
	logger.Wait()
	logger.LogStack()
	logger.Wait()
}
