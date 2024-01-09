# Terrapak
Terrapak is a private registry for your Terraform modules. Terrapak integrates with your GitHub pull requests to automatically publish new versions of your Terraform modules. This server works in conjunction with the [Terrapack-Action](https://github.com/eunanhardy/terrapak-action) to deliver a configuration driven workflow that allows you more flexability in how you structure your Terraform project.

## Requirements
- Postgres DB
- Redis sidecar
- Github OAuth App
- S3 Bucket for modules

## Feature Overview
- Automatic versioning of Terraform modules
- Monorepo friendly CI/CD workflow
- Supports S3 as storage backend
- Automatic cleanup of draft modules when pull request is closed unsuccessfully
- Support GitHub Organisations for Authorization
- Support for `terraform login` with github


### MVP for v1
- [x] Github-driven automatic versioning of Terraform modules
- [x] Support S3 as storage backend
- [x] Support for future oauth2 providers
- [ ] Improve UX of Github Action
- [ ] Known Bug Fixes

*HTTPS is required to use this application, I recommend using a reverse proxy such as ngrok for local development.*


> [!NOTE]  
> This project is currently in development and not ready for production use.