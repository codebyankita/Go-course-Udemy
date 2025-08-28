Go-course-Udemy/
├── cmd/
│   └── api/
│       ├── server.go            # Main entry point for the API server
│       ├── .env                 # Environment variables
│       ├── cert.pem             # TLS certificate
│       └── key.pem              # TLS private key
│
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── execs.go         # Handlers for exec routes
│   │   │   ├── helpers.go       # Utility/helper functions
│   │   │   ├── students.go      # Handlers for student routes
│   │   │   └── teachers.go      # Handlers for teacher routes
│   │   │
│   │   ├── middlewares/
│   │   │   ├── compression.go   # Response compression
│   │   │   ├── cors.go          # CORS middleware
│   │   │   ├── exclude_routes.go # Excluded routes handling
│   │   │   ├── hpp.go           # HTTP parameter pollution prevention
│   │   │   ├── jwt_middleware.go # JWT authentication
│   │   │   ├── rate_limiter.go  # Rate limiting
│   │   │   ├── response_time.go # Response time tracking
│   │   │   ├── sanitize.go      # XSS/data sanitization
│   │   │   └── security_header.go # Security headers
│   │   │
│   │   └── router/
│   │       ├── router.go        # Root router setup
│   │       ├── exec_router.go   # Exec routes
│   │       ├── students_router.go # Student routes
│   │       └── teachers_router.go # Teacher routes
│   │
│   ├── models/
│   │   ├── exec.go              # Exec model
│   │   ├── student.go           # Student model
│   │   └── teacher.go           # Teacher model
│   │
│   └── repository/
│       └── sqlconnect/
│           ├── sqlconfig.go     # DB connection config
│           ├── exec_crud.go     # Exec DB operations
│           ├── students_crud.go # Student DB operations
│           └── teachers_crud.go # Teacher DB operations
│
├── pkg/
│   └── utils/
│       ├── authorize_user.go    # Authorization utils
│       ├── database_utils.go    # DB query utils
│       ├── error_handler.go     # Error handling utils
│       ├── jwt.go               # JWT utilities
│       ├── middlewaresutil.go   # Common middleware utils
│       └── password.go          # Password hashing utils
│
├── data/
│   ├── execsdata.json           # JSON seed data: execs
│   ├── studentsdata.json        # JSON seed data: students
│   └── teachersdata.json        # JSON seed data: teachers
│
├── .gitignore                   # Git ignore rules
├── go.mod                       # Go module definition
├── go.sum                       # Go dependencies lockfile
│
├── docs/
│   ├── helpful_commands.txt     # Reference commands
│   ├── ROLEFORROUTES.txt        # Role-to-route mapping
│   └── MailHog.txt              # MailHog usage notes
│
└── openssl.cnf                  # OpenSSL config for certs
