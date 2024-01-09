# Terrapak
Terrapak is a version manager for Terraform modules. Terrapak integrates with your GitHub pull requests to automaticlly publish new versions of your Terraform modules to a private Terraform registry. 

## Feature Overview
- Automatic versioning of Terraform modules
- Monorepo friendly CI/CD workflow
- Supports S3 as storage backend
- Automatic cleanup of draft modules when pull request is closed unsucessfully
- Support Gtihub Organisations for Authorization
- Support for `terraform login` with github


### MVP for v1
- [x] Github-driven automatic versioning of Terraform modules
- [x] Support S3 as storage backend
- [x] Support for future oauth2 providers
- [ ] rule based assignment for RBAC users
- [ ] Known issues

*HTTPS is required to use this application, we recommend using a reverse proxy such as ngrok for local development.*


> [!NOTE]  
> This project is currently in development and not ready for production use.