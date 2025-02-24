package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func StringToInt64(s string) int64 {

	var i int64

	if strings.ContainsAny(s, ".,") {
		s = strings.ReplaceAll(s, ".", "")
		s = strings.ReplaceAll(s, ",", "")
	}

	if strings.ContainsAny(s, "ABDCEFGHIJKLMNÑOPQRSTUVWXYZÇabcdefghijklmnñopqrstuvwxyzç!\"·$%^&*()#~€¬|@¡¿?¿¡") {
		return 0
	}

	for _, c := range s {
		i = i*10 + int64(c-'0')
	}
	return i
}

func StringToInt(s string) (int, error) {
	if strings.ContainsAny(s, "\"≠”´!@#$%^&*()_+-=[]{};':,./<>?çæ·") {
		return 0, fmt.Errorf("invalid number, contains special characters")
	}
	if strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyzñABCDEFGHIJKLMNOPQRSTUVWXYZÑ") {
		return 0, fmt.Errorf("invalid number, contains letters")
	}
	var number int

	_, err := fmt.Sscan(s, &number)

	if err != nil {
		return 0, err
	}
	return number, err
}

func StringToFloat64(s string) (float64, error) {
	// Verificar si la cadena contiene caracteres no válidos
	if strings.ContainsAny(s, "ABDCEFGHIJKLMNÑOPQRSTUVWXYZÇabcdefghijklmnñopqrstuvwxyzç!\"·$%^&*()#~€¬|@¡¿?¿¡") {
		return 0, fmt.Errorf("cadena contiene caracteres no válidos")
	}

	// Reemplazar ',' por '.' si es necesario
	s = strings.ReplaceAll(s, ",", ".")

	// Dividir la cadena en parte entera y parte decimal
	parts := strings.Split(s, ".")
	var integerPart, decimalPart string

	if len(parts) > 0 {
		integerPart = parts[0]
	}

	if len(parts) > 1 {
		decimalPart = parts[1]
	}

	// if strings.Contains(decimalPart, ".") {
	// 	decimalPart = strings.ReplaceAll(decimalPart, ".", "")
	// }

	// Convertir la parte entera a float64
	result, err := strconv.ParseFloat(integerPart, 64)
	if err != nil {
		return 0, fmt.Errorf("no se pudo convertir la parte entera a float64: %v", err)
	}

	// Procesar la parte decimal
	if decimalPart != "" {
		decimal, err := strconv.ParseFloat("0."+decimalPart, 64)
		if err != nil {
			return 0, fmt.Errorf("no se pudo convertir la parte decimal a float64: %v", err)
		}
		result += decimal
	}

	return result, nil
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
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	}
	return fmt.Sprintf("%d", n)
}

func FisrtIdentifier(id uuid.UUID) string {
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
	s := strings.Split(id.String(), "-")
	return s[index]
}

func UnixTimestampStringToDate(timestamp string) string {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return ""
	}
	i /= 1000
	return time.Unix(i, 0).Format("02/01/2006")
}
