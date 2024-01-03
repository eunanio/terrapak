# Terrapak
Terrapak is a Terraform module registry with the goal of making module version management easier. Terrapak integrates with your github pull requests to automatically publish new versions of your terraform modules.

## How to setup
1. create config file, see Docs
2. Create a postgres & Redis instance for Terrapak to use
3. Run the Terrapak server
   ```bash
   docker run -v $(pwd)/config.yml:/app/config.yml -e CONFIG_PATH=/app/config.yml -p 5551:80 -d monoci/terrapak:v1
   ```
4. Setup [Terrapak-Action](http://api.github.com) in your project
5. Generate Credentials for Terrapak to use
   ```bash
   terraform login registry.host.io
   ```