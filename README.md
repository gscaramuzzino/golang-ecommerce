# MyAmazon

MyAmazon is a robust e-commerce API, akin to a minimalistic version of the Amazon marketplace. The application is written in Golang, using libraries such as Gorilla Mux for HTTP routing, JWT-Go for JSON Web Token authentication, and Crypto for secure password hashing. It utilizes PostgreSQL as the backend database for storing information about users, products, and shopping carts.

## Endpoints

### User Routes

- `POST /users`: Registers a new user. The request body should include `username`, `password`, and `email` fields.

- `POST /users/login`: Authenticates a user. The request body should include `username` and `password` fields.

### Product Routes

- `GET /products`: Retrieves a list of all available products.

- `GET /products/{productID}`: Retrieves detailed information about a specific product. Replace `{productID}` with the ID of the desired product.

### Cart Routes

- `GET /cart`: View the current user's shopping cart.

- `POST /cart/add/{productID}`: Adds a specific product to the user's shopping cart. Replace `{productID}` with the ID of the desired product.

- `POST /cart/checkout`: Checks out the user's shopping cart, finalizing the purchase.

## Setup

### Requirements

- Golang (latest version recommended)
- PostgreSQL

### Installation

Follow these steps to set up the project locally:

1. Clone the repository:
```
git clone https://github.com/gscaramuzzino/golang-ecommerce
```

2. Navigate to the project directory:
```
cd golang-ecommerce
```

3. Install the necessary Go dependencies:
```
go mod download
```

4. Run the server:
```
docker compose up -d
```

## Running Load Tests

We use k6 for load testing the application. This tool allows us to simulate traffic to our application and measure its performance.

### Prerequisites

Install k6 on your machine. You can follow the [official installation guide](https://k6.io/docs/getting-started/installation/) to install it on your operating system.

### Load Testing

Once k6 is installed, navigate to the directory containing your loadtesting.js file. You can run your k6 tests using the following command:

```bash
k6 run loadtesting.js
```

This will start the load test based on the parameters specified in the `loadtesting.js` file.

### Configuring Load Testing

The configuration of your load testing, such as the number of virtual users, test duration, and the API endpoints to be tested, can be adjusted in the `loadtesting.js` script.

## Analyzing Test Results

k6 provides detailed test results in the terminal after the test run. You can analyze these results to understand the system's performance under load.

For more complex analysis and visualization, consider using Grafana with k6's backend, which allows real-time result visualization and historical data analysis.

## Integrating InfluxDB

InfluxDB is a time series database that can be used to store the results of the load tests for further analysis. In conjunction with Grafana, it allows for easy visualization of the k6 test results.

### Prerequisites

Ensure InfluxDB is installed and properly configured on your machine. You can follow the [official InfluxDB installation guide](https://docs.influxdata.com/influxdb/v2.0/install/) to install it on your operating system.

### Running k6 with InfluxDB

After installing and setting up InfluxDB, you can run your load tests and send the results to InfluxDB using the following command:

```bash
k6 run --out influxdb=http://localhost:8086/myk6db loadtesting.js
```

In the command above, replace `http://localhost:8086/myk6db` with your InfluxDB URL and database name.

### Viewing Test Results with Grafana

After running the tests and storing the results in InfluxDB, you can use Grafana to visualize the data.

1. Set up a data source in Grafana that points to your InfluxDB instance.
2. Create a new dashboard in Grafana.
3. Add panels and select the k6 data from InfluxDB.

This way, you can create a powerful dashboard that visualizes the load testing results and helps you understand the system's performance under various conditions.

Remember to keep your load testing environment separate from your production environment to prevent accidental overloads.
