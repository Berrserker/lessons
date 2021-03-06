package unsafe

import (
	"fmt"
	"unsafe"
)

func Unf() {
	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{0, 0, 0, 0}

	// Преобразуем указатель структуры в общий указатель
	p := unsafe.Pointer(&s)
	// Сохраняем адрес структуры для будущего использования (смещение 0)
	up0 := uintptr(p)
	// Преобразуем общий указатель в байтовый указатель
	pb := (*byte)(p)
	// Присваиваем значение преобразованному указателю
	*pb = 10
	// Содержимое структуры меняется соответственно
	fmt.Println(s)
	
	// Смещение ко второму полю
	up := up0 + unsafe.Offsetof(s.b)
	// Преобразуем адрес смещения в общий указатель
	p = unsafe.Pointer(up)
	// Преобразуем общий указатель в байтовый указатель
	pb = (*byte)(p)
	// Присваиваем значение преобразованному указателю
	*pb = 20
	// Содержимое структуры меняется соответственно
	fmt.Println(s)
	
	// Смещение к 3-му полю
	up = up0 + unsafe.Offsetof(s.c)
	// Преобразуем адрес смещения в общий указатель
	p = unsafe.Pointer(up)
	// Преобразуем общий указатель в байтовый указатель
	pb = (*byte)(p)
	// Присваиваем значение преобразованному указателю
	*pb = 30
	// Содержимое структуры меняется соответственно
	fmt.Println(s)
	
	// Смещение до 4-го поля
	up = up0 + unsafe.Offsetof(s.d)
	// Преобразуем адрес смещения в общий указатель
	p = unsafe.Pointer(up)
	// Преобразуем общий указатель в указатель int64
	pi := (*int64)(p)
	// Присваиваем значение преобразованному указателю
	*pi = 40
	// Содержимое структуры меняется соответственно
	fmt.Println(s)
}