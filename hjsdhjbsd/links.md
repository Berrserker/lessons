+ https://github.com/qiniu/checkstyle
+ https://github.com/bradleyfalzon/apicompat
+ https://github.com/KyleBanks/depth
+ https://github.com/go-swagger/go-swagger
+ https://github.com/golangci/golangci-lint
+ https://github.com/ofabry/go-callvis
+ https://github.com/inconshreveable/gonative
+ https://godbolt.org/
+ https://dave.cheney.net/2020/06/19/how-to-dump-the-gossafunc-graph-for-a-method
+ makefile 
+ linter rules

# pprof info 
+ cpu: профиль ЦП определяет, где программа тратит свое время, активно потребляя циклы ЦП (в отличие от сна или ожидания ввода-вывода).
+ heap: профиль кучи сообщает о распределении памяти; используется для мониторинга текущего и исторического использования памяти, а также для проверки утечек памяти.
+ threadcreate: профиль создания потока сообщает о разделах программы, которые ведут создание новых потоков ОС.
+ goroutine: профиль goroutine сообщает о следах стека (stack traces) всех текущих goroutines.
+ block: профиль block показывает, где блокирующиеся программы ожидают примитивы синхронизации (включая каналы таймера). Профиль block не включен по умолчанию; используйте runtime.SetBlockProfileRate, чтобы включить его.
+ mutex: профиль mutex сообщает о блокировках. Если вы считаете, что ваш процессор не используется полностью из-за мьютекса, используйте этот профиль. Профиль mutex не включен по умолчанию, используйте Runtime.SetMutexProfileFraction, чтобы включить его.