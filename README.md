Создан worker-pool для обработки строк, который позволяет динамически управлять количеством активных воркеров. Воркеры организованы в стек, что позволяет добавлять новые воркеры в конец или удалять последний добавленный. Учтены ошибки, возникающие из-за отсутствия воркеров. В функции main представлен пример использования, который можно адаптировать под собственные сценарии.

Как запустить: 

cd cmd

go run main.go