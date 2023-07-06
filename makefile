docker-compose up -d

docker logs goole-ecommerce

k6 run loadtesting.js
run:
	k6 run loadtesting.js --out influxdb=http://localhost:8086/k6