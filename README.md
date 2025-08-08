## For Production
sudo docker compose up --build

## For Development
### use node 20
use node 20
cd frontend
yarn install
yarn dev
cd ../
sudo docker compose -f docker-compose-dev.yaml up