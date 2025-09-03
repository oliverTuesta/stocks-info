# Terraform Deployment for Stocks Info Application

This directory contains Terraform configuration files to deploy your Go backend and Vue.js frontend on a single EC2 instance in AWS.

## Prerequisites

1. **AWS CLI configured** with appropriate permissions
2. **Terraform installed** (version >= 1.0)
3. **AWS Key Pair** created for SSH access
4. **CockroachDB** already deployed and accessible

## Architecture

- **Single EC2 instance** (t3.medium by default)
- **Nginx** serves the Vue.js frontend on port 80
- **Go backend** runs on port 8080
- **Nginx reverse proxy** forwards `/api/*` requests to the backend
- **Supervisor** manages the backend process

## Quick Start

### 1. Configure Variables

Copy the example variables file and fill in your values:

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit `terraform.tfvars` with your actual values:
- AWS region and key pair name
- CockroachDB connection details
- Any other customizations

### 2. Initialize Terraform

```bash
terraform init
```

### 3. Review the Plan

```bash
terraform plan
```

### 4. Apply the Configuration

```bash
terraform apply
```

### 5. Access Your Application

After successful deployment, you'll see outputs including:
- Public IP address
- SSH command to connect

Access your application at: `http://<public-ip>`

## Configuration Details

### Security Groups

The following ports are open:
- **22**: SSH access
- **80**: HTTP (frontend)
- **443**: HTTPS (frontend)
- **8080**: Backend API

### Instance Specifications

- **Type**: t3.medium (2 vCPU, 4GB RAM)
- **Storage**: 20GB GP3 EBS volume
- **OS**: Ubuntu 22.04 LTS

### Services

- **Nginx**: Web server and reverse proxy
- **Supervisor**: Process manager for the Go backend
- **PM2**: Node.js process manager (installed but not used by default)

## Deployment Process

The `user_data.sh` script automatically:

1. Updates the system
2. Installs Go, Node.js, Nginx, and Supervisor
3. Creates application directories
4. Configures Nginx and Supervisor
5. Sets up environment variables
6. Starts all services

## Manual Deployment Steps

If you prefer to deploy manually instead of using the user data script:

1. SSH into the instance
2. Clone your repository
3. Build the Go backend: `go build -o bin/server cmd/server/main.go`
4. Build the Vue.js frontend: `npm run build`
5. Configure environment variables
6. Start services manually

## Troubleshooting

### Check Service Status

```bash
# Check backend status
sudo supervisorctl status stocks-info-backend

# Check nginx status
sudo systemctl status nginx

# Check logs
sudo tail -f /var/log/stocks-info-backend.out.log
sudo tail -f /var/log/nginx/error.log
```

### Common Issues

1. **Backend not starting**: Check database connection and environment variables
2. **Frontend not loading**: Verify Nginx configuration and frontend build
3. **API calls failing**: Check Nginx reverse proxy configuration

### Health Check

Run the health check script:
```bash
/opt/stocks-info/health-check.sh
```

## Cost Optimization

- **Instance Type**: Consider using t3.micro for development/testing
- **Storage**: Reduce EBS volume size if not needed
- **Region**: Choose the region closest to your users
- **Reserved Instances**: For production, consider reserved instances

## Security Considerations

- **Key Pair**: Keep your private key secure
- **Security Groups**: Review and restrict access as needed
- **Database**: Ensure CockroachDB is properly secured
- **HTTPS**: Consider adding SSL/TLS certificates

## Cleanup

To destroy all resources:

```bash
terraform destroy
```

## Support

For issues or questions:
1. Check the logs on the EC2 instance
2. Verify your configuration values
3. Ensure all prerequisites are met
4. Check AWS service status
