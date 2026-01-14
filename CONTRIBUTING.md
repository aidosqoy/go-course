# Инструкция для студентов

## Перед началом: Git

Для работы с курсом нужен Git — система контроля версий.

### Установка Git

**macOS:**
```bash
# Через Homebrew
brew install git

# Или скачай с сайта
# https://git-scm.com/download/mac
```

**Windows:**
```
Скачай установщик: https://git-scm.com/download/win
```

**Linux:**
```bash
# Ubuntu/Debian
sudo apt install git

# Fedora
sudo dnf install git
```

### Проверка установки

```bash
git --version
# Ожидаемый вывод: git version 2.x.x
```

### Настройка Git (один раз)

```bash
git config --global user.name "Твоё Имя"
git config --global user.email "твой@email.com"
```

### Изучение Git

Если ты новичок в Git, пройди один из курсов:

- [Learn Git Branching](https://learngitbranching.js.org/?locale=ru_RU) — интерактивный тренажёр (рекомендуем!)
- [Git How To](https://githowto.com/ru) — пошаговый туториал
- [Pro Git Book](https://git-scm.com/book/ru/v2) — полная книга (для углублённого изучения)

**Минимум что нужно знать:**
- `git clone` — скачать репозиторий
- `git checkout` — переключиться на ветку
- `git add` — добавить файлы в коммит
- `git commit` — создать коммит
- `git push` — отправить изменения на сервер
- `git pull` — получить изменения с сервера

---

## Как устроен репозиторий

```
go-course/
├── main                        # Шаблоны заданий (не трогаем!)
├── students/ivan               # Принятые решения Ивана (мержит ментор)
│   └── ivan/week-01-homework   # Рабочая ветка Ивана для PR
├── students/maria              # Принятые решения Марии
└── ...
```

- **main** — шаблоны заданий, без решений
- **students/имя** — принятые решения студента (пушит только ментор после ревью)
- **имя/week-XX-...** — рабочие ветки студента для создания PR

## Начало работы

### Шаг 1: Клонируй репозиторий

Выбери папку для проектов и клонируй туда репозиторий.

**macOS / Linux:**
```bash
# Создай папку для проектов (если ещё нет)
mkdir -p ~/projects
cd ~/projects

# Клонируй
git clone https://github.com/qazcode-mentoring/go-course.git
cd go-course
```

**Windows (PowerShell или Git Bash):**
```bash
# Создай папку для проектов
mkdir C:\projects
cd C:\projects

# Клонируй
git clone https://github.com/qazcode-mentoring/go-course.git
cd go-course
```

### Шаг 2: Создай рабочую ветку

Ментор заранее создаст для тебя базовую ветку `students/твоё-имя`. Создай от неё рабочую ветку:

```bash
# Переключись на свою базовую ветку
git checkout students/твоё-имя

# Создай рабочую ветку для домашки
git checkout -b твоё-имя/week-01-homework
```

Например:
```bash
git checkout students/ivan
git checkout -b ivan/week-01-homework
```

### Шаг 3: Проверь что Go установлен

```bash
go version
# Ожидаемый вывод: go version go1.21.0 (или новее)
```

Если Go не установлен — смотри задание 1.1 (setup).

## Структура недели

Каждая неделя имеет такую структуру:

```
week-01/
├── README.md              # Ссылка на уроки Go Tour
└── homework/
    ├── 1.1-setup/
    │   ├── README.md      # Описание задания
    │   └── main.go        # Сюда пишешь решение
    ├── 1.2-swap/
    │   ├── README.md
    │   └── main.go
    └── ...
```

## Workflow выполнения заданий

### 1. Изучи теорию

Открой `week-XX/README.md` — там указаны номера уроков Go Tour для изучения.

Пройди уроки на https://www.gocat.dev/tour/

### 2. Прочитай задание

Открой `homework/X.X-name/README.md` — там описано что нужно сделать.

### 3. Напиши решение

Открой `main.go` в папке задания и напиши свой код.

### 4. Проверь локально

```bash
cd week-01/homework/1.1-setup
go run main.go
```

### 5. Закоммить и запуш

```bash
git add .
git commit -m "week-01: выполнено задание 1.1-setup"
git push origin твоё-имя/week-01-homework
```

Например:
```bash
git push origin ivan/week-01-homework
```

### 6. Создай Pull Request для ревью

1. Зайди на GitHub в репозиторий
2. Нажми **"Pull requests"** → **"New pull request"**
3. **ВАЖНО**: выбери правильные ветки:
   - base: `students/твоё-имя` (куда мержим — твоя основная ветка)
   - compare: `твоё-имя/week-01-homework` (откуда берём — твоя рабочая ветка)

4. Напиши что сделано в описании
5. Дождись ревью от ментора
6. **Ментор замержит PR после одобрения** — ты не можешь мержить сам

## Формат коммитов

```
week-XX: краткое описание

Примеры:
- week-01: выполнено задание 1.1 setup
- week-01: выполнены все задания недели
- week-02: исправлены замечания по ревью
```

## Проверка кода перед коммитом

### Код компилируется

```bash
go build ./...
```

### Код отформатирован

```bash
go fmt ./...
```

### Нет очевидных ошибок

```bash
go vet ./...
```

## Следующая неделя

После того как ментор замержил PR:

```bash
# Обнови свою базовую ветку
git checkout students/твоё-имя
git pull origin students/твоё-имя

# Создай новую рабочую ветку для следующей недели
git checkout -b твоё-имя/week-02-homework
```

## Синхронизация с новыми заданиями

Когда ментор добавит новые задания в main, он подтянет их в твою ветку `students/имя`. Тебе нужно обновить локальную копию:

```bash
git checkout students/твоё-имя
git pull origin students/твоё-имя
```

## Частые вопросы

### Я случайно закоммитил в main, что делать?

Напиши ментору — он поможет откатить.

### Как посмотреть свои изменения?

```bash
git status          # какие файлы изменены
git diff            # что именно изменилось
git log --oneline   # история коммитов
```

### Как откатить незакоммиченные изменения?

```bash
git checkout -- файл.go    # откатить один файл
git checkout -- .          # откатить все изменения
```

## Вопросы?

Если что-то непонятно — спрашивай на встрече или пиши в чат!

---

# Инструкция для ментора

## Создание ветки для нового студента

```bash
# Создай ветку от main
git checkout main
git pull origin main
git checkout -b students/имя-студента

# Запуш ветку
git push -u origin students/имя-студента
```

## Ревью домашних заданий

1. Студент создаёт PR из `имя/week-XX-homework` в `students/имя`
2. Смотришь код на GitHub, оставляешь комментарии
3. Если нужны исправления — студент пушит в ту же ветку
4. После одобрения — мержишь PR

## Добавление новых заданий

1. Добавь задания в `main`
2. Подтяни изменения в ветки студентов:

```bash
git checkout students/имя-студента
git merge main
git push origin students/имя-студента
```

3. Уведоми студентов что нужно сделать `git pull`

## Защита веток (рекомендуется)

В настройках репозитория на GitHub:
- **main** — защитить от прямых пушей
- **students/*** — разрешить пуш только maintainers