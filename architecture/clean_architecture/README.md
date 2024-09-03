/go-todo-app
│
├── /entities
│   └── task.go
│
├── /usecases
│   ├── create_task.go
│   └── get_tasks.go
│
├── /interfaces
│   ├── /repositories
│   │   └── task_repository.go
│   └── /controllers
│       └── task_controller.go
│
├── /infrastructure
│   ├── /db
│   │   └── task_repository_impl.go
│   └── /web
│       └── router.go
│
├── /constants
│   └── config.go
│
└── main.go
