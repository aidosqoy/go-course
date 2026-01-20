package main

import "fmt"

/*
TODO: Напиши своё объяснение defer здесь

1. Что такое defer?
defer позволяет нам отложить какую либо функцию. То есть defer сработает только после выполнения окружной функции

2. В каком порядке выполняются несколько defer?
снизу вверх, то есть какой defer написан первый, тот будет последним

3. Примеры использования defer в реальном коде:
   - Пример 1:
	pool, err := database.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("failed to create pool: %v", err)
	}

	defer pool.Close() После подключения к бд, в конце нужно закрыть подключение чтобы не было утечки

	fmt.Println("Success", pool)


   - Пример 2:
func (userRepo *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	rows, err := userRepo.Pool.Query(ctx, "SELECT id, full_name, email, department_id, created_at, updated_at, deleted_at FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}

	defer rows.Close() здесь defer нужен чтобы освободить ресурсы которые пришли с бд, и вернуть соединение обратно в pool

	for rows.Next() {
		var user models.User

		err = rows.Scan(
			&user.Id,
			&user.FullName,
			&user.Email,
			&user.DepartmentId,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan a row: %w", err)
		}

		users = append(users, user)
	}

   - Пример 3:
	ch := make(chan int)
	defer close(ch) тут мы используем defer чтобы гарантированно закрыть канал

	ch <- 1
	ch <- 2

*/

func main() {
	// Демонстрация работы defer
	fmt.Println("Начало функции main")

	defer fmt.Println("Это выполнится последним (defer 1)")
	defer fmt.Println("Это выполнится предпоследним (defer 2)")
	defer fmt.Println("Это выполнится третьим с конца (defer 3)")

	fmt.Println("Конец функции main (но до defer)")
}
