## Requirement
docker (if do not have yet, see [Install Docker Tutorial])
node 20 (if you wanna do faster development, skip if just running on production, [Install Node 20 via n])

## For Production
sudo docker compose up --build
open http://localhost:3000

## For Development (faster developement)
### use node 20
use node 20
cd frontend
yarn install
yarn dev
cd ../
sudo docker compose -f docker-compose-dev.yaml up
open http://localhost:3000

## API Documentation
http://localhost:8080/swagger/index.html

## Test Backend Go (ensure use go 1.23.12)
cd backend
go test ./...

## [Install Docker Tutorial]
sudo apt update
sudo apt upgrade -y
sudo apt install -y \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | \
    sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) \
  signed-by=/etc/apt/keyrings/docker.gpg] \
  https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
docker --version
docker run hello-world
sudo systemctl enable docker
sudo systemctl start docker
sudo usermod -aG docker $USER
newgrp docker  # or log out and back in
docker compose version


## [Install Node 20 via n]
sudo apt install -y nodejs npm
sudo npm install -g n
sudo n 20
node -v
npm -v
sudo mkdir -p /usr/local/n
sudo chown -R $(whoami) /usr/local/n
sudo chown -R $(whoami) /usr/local/bin /usr/local/lib /usr/local/include /usr/local/share
n 20




