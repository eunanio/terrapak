## Configuration Values

| Env | yaml | Description | Required |
|-----|-----|-------------|----------|
| `TP_DB_HOST`| `database:host`| Database host | true |
| `TP_DB_USER` | `database:username` | Database name | true |
| `TP_DB_PASS` | `database:password` | Database password | true |
| `TP_HOSTNAME` | `hostname` | Hostname of the server | true |
| `TP_ORGANIZATION` | `organization` | Organization name | true |
| `TP_AUTH_TYPE` | `auth:type` | Authentication type - Only support value is github currently | true |
| `TP_AUTH_CLIENT_ID` | `auth:client_id` | OAuth Client ID | true |
| `TP_AUTH_SECRET` | `auth:client_secret` | OAuth Client Secret | true |
| `TP_SECRET` | `secret` | Secret used for signing tokens | true |
| `TP_GITHUB_ORG` | `auth:organization` | Github organization name | false |
| `TP_CONFIG_FILE` | N/A | Path to config file | false |
| `TP_STORAGE` | `storage` | Storage type - Only support value is `s3://bucket-name` currently | true |
| `TP_USER` | N/A | Default root user for the server | false |
| `TP_PASSWORD` | N/A | Default root password for the server | false |
| `ENV_REDIS_HOST` | `redis:host` | Redis host | true |
| `TP_REDIS_PASSWORD` | `redis:password` | Redis password | true |

