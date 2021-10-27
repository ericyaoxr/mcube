package etc

// EnvExampleTemplate todo
const EnvExampleTemplate = `export APP_NAME="{{.Name}}"
export APP_KEY="key"
export HTTP_HOST="127.0.0.1"
export HTTP_PORT=8848
export GRPC_HOST="127.0.0.1"
export GRPC_PORT=18848
export KEYAUTH_HOST="127.0.0.1"
export KEYAUTH_PORT=18080
export KEYAUTH_CLIENT_ID=""
export KEYAUTH_CLIENT_SECRET=""
export LOG_LEVEL="info"
export LOG_PATH="logs"
export LOG_TO="stdout"
export MYSQL_HOST="127.0.0.1"
export MYSQL_PORT=3306
export MYSQL_USERNAME="{{.Name}}"
export MYSQL_PASSWORD=""
export MYSQL_DATABASE="{{.Name}}"`
