#!/bin/bash

# Update system
yum update -y

# Install Docker
yum install -y docker
systemctl start docker
systemctl enable docker
usermod -a -G docker ec2-user

# Install Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

# Install Git
yum install -y git

# Create application directory
mkdir -p /opt/app
cd /opt/app

# Create .env file from template variables
cat << 'EOF' > /opt/app/.env
${backend_env_vars}
EOF

# Create docker-compose.yml for both services
cat << 'EOF' > /opt/app/docker-compose.yml
version: '3.8'

services:
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    env_file:
      - .env
    ports:
      - "8080:8080"
    restart: unless-stopped
    networks:
      - app-network

  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "80:80"
    restart: unless-stopped
    depends_on:
      - backend
    networks:
      - app-network

networks:
  app-network:
    driver: bridge
EOF

# Create deploy script
cat << 'EOF' > /opt/app/deploy.sh
#!/bin/bash

# Function to clone or update repository
clone_or_update_repo() {
    local repo_url=$1
    local target_dir=$2
    local branch=$${3:-main}

    if [ -d "$target_dir" ]; then
        echo "Updating existing repository in $target_dir..."
        cd "$target_dir"
        git fetch origin
        git reset --hard origin/$branch
        cd /opt/app
    else
        echo "Cloning repository to $target_dir..."
        git clone -b $branch $repo_url $target_dir
    fi
}

# You'll need to replace these URLs with your actual repository URLs
# clone_or_update_repo "https://github.com/yourusername/your-backend-repo.git" "backend" "main"
# clone_or_update_repo "https://github.com/yourusername/your-frontend-repo.git" "frontend" "main"

echo "Repository URLs are commented out. Please update deploy.sh with your actual repository URLs."
echo "For now, you'll need to manually copy your application code to /opt/app/backend and /opt/app/frontend"

# Build and start services
echo "Building and starting services..."
docker-compose down
docker-compose build --no-cache
docker-compose up -d

# Show status
docker-compose ps

echo "Deployment complete!"
echo "Frontend: http://$(curl -s http://169.254.169.254/latest/meta-data/public-ipv4)"
echo "Backend API: http://$(curl -s http://169.254.169.254/latest/meta-data/public-ipv4):8080"
EOF

chmod +x /opt/app/deploy.sh

# Set ownership
chown -R ec2-user:ec2-user /opt/app

# Create systemd service for auto-start
cat << 'EOF' > /etc/systemd/system/fullstack-app.service
[Unit]
Description=Fullstack Application
Requires=docker.service
After=docker.service

[Service]
Type=oneshot
RemainAfterExit=yes
WorkingDirectory=/opt/app
ExecStart=/usr/bin/docker-compose up -d
ExecStop=/usr/bin/docker-compose down
User=ec2-user
Group=ec2-user

[Install]
WantedBy=multi-user.target
EOF

systemctl enable fullstack-app.service

# Log completion
echo "Bootstrap script completed at $(date)" > /var/log/bootstrap.log