package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// StringToInt64 converts a string to an int64.
// It handles optional decimal/thousands separators ('.', ',') by removing them.
// It returns 0 if the string contains non-numeric characters after cleanup, or if parsing fails.
func StringToInt64(s string) int64 {
	// Remove common thousands/decimal separators if present.
	// Note: This assumes a specific locale interpretation (e.g., "1.000" is one thousand, not one point zero).
	// For robust internationalization, a more sophisticated locale-aware parser would be needed.
	if strings.ContainsAny(s, ".,") {
		s = strings.ReplaceAll(s, ".", "")
		s = strings.ReplaceAll(s, ",", "")
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0 // Return 0 on error, consistent with original behavior for invalid input.
	}
	return i
}

// StringToInt converts a string to an int.
// It returns an error if the string is not a valid integer.
func StringToInt(s string) (int, error) {
	// strconv.Atoi is a shorthand for ParseInt(s, 10, 0)
	return strconv.Atoi(s)
}

// StringToFloat64 converts a string to a float64.
// It handles comma as a decimal separator by replacing it with a dot.
// It returns an error if the string is not a valid float.
func StringToFloat64(s string) (float64, error) {
	// Check for non-numeric characters that are not part of a valid float format.
	// This check is still somewhat basic; strconv.ParseFloat will handle most invalid formats.
	if strings.ContainsAny(s, "ABDCEFGHIJKLMNÑOPQRSTUVWXYZÇabcdefghijklmnñopqrstuvwxyzç!\"·$%^&*()#~€¬|@¡¿?¿¡") && !strings.ContainsAny(s, ".,") {
		return 0, fmt.Errorf("string contains invalid characters")
	}

	// Reemplazar ',' por '.' si es necesario
	s = strings.ReplaceAll(s, ",", ".")

	// Dividir la cadena en parte entera y parte decimal
	parts := strings.Split(s, ".")
	// If there are multiple dots (e.g., "1.2.3"), strconv.ParseFloat will handle it as an error.
	if len(parts) > 2 {
		return 0, fmt.Errorf("string contains multiple dots")
	}

	// If the string is empty or only contains a dot, ParseFloat will return an error.
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("could not convert string to float64: %w", err)
	}
	return f, nil
}

func PointerToString(p *string) (s string) {
	if p != nil {
		s = *p
	}
	return
}

// format date: 17 Junio, 2023 16:13
func FormatDateTime(date time.Time) string {
	return fmt.Sprintf("%s %s, %d %s:%s", addZero(date.Day()), monthSpanish(date.Month()), date.Year(), addZero(date.Hour()), addZero(date.Minute()))
}

// format date: 17 de Junio, 2023
func FormatDateWithYear(date time.Time) string {
	return fmt.Sprintf("%s de %s, %d", addZero(date.Day()), monthSpanish(date.Month()), date.Year())
}

// format date: 17 de Junio
func FormatDateWithoutYear(date time.Time) string {
	return fmt.Sprintf("%s de %s", addZero(date.Day()), monthSpanish(date.Month()))
}

func monthSpanish(m time.Month) string {
	switch m {
	case time.January:
		return "Enero"
	case time.February:
		return "Febrero"
	case time.March:
		return "Marzo"
	case time.April:
		return "Abril"
	case time.May:
		return "Mayo"
	case time.June:
		return "Junio"
	case time.July:
		return "Julio"
	case time.August:
		return "Agosto"
	case time.September:
		return "Septiembre"
	case time.October:
		return "Octubre"
	case time.November:
		return "Noviembre"
	default:
		return "Diciembre"
	}
}

func addZero(n int) string {
	if n >= 0 && n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}

func FirstIdentifier(id uuid.UUID) string {
	return identifier(id, 0)
}

func SecondIdentifier(id uuid.UUID) string {
	return identifier(id, 1)
}

func ThirdIdentifier(id uuid.UUID) string {
	return identifier(id, 2)
}

func FourthIdentifier(id uuid.UUID) string {
	return identifier(id, 3)
}

func FifthIdentifier(id uuid.UUID) string {
	return identifier(id, 4)
}

func identifier(id uuid.UUID, index uint) string {
	str := id.String() // UUID format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	switch index {
	case 0:
		return str[:8]
	case 1:
		return str[9:13]
	case 2:
		return str[14:18]
	case 3:
		return str[19:23]
	case 4:
		return str[24:]
	default:
		return ""
	}
}

func UnixTimestampStringToDate(timestamp string) string {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return ""
	}
	i /= 1000
	return time.Unix(i, 0).Format("02/01/2006")
}
