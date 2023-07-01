import http from 'k6/http';
import { check, group } from 'k6';

export let options = {
   scenarios: {
      user_tests: {
         executor: 'constant-vus',
         vus: 100,
         duration: '3m', // user registration and login test will run for 1 minute
         exec: 'userScenario' // the function to be executed for this scenario
      },
      store_tests: {
         executor: 'constant-vus',
         vus: 40,
         startTime: '3m', // product test will start after 1 minute
         duration: '25m',  // product test will run for 1 minute
         exec: 'storeScenario' // the function to be executed for this scenario
      },
   },
};

export function setup() {
   let user = { "name": "bwGZrLeX8v", "email": "bwGZrLeX8v@outlook.com", "password": "mypassword" };
   let res = http.post('http://localhost:8080/users/login', JSON.stringify(user));
   if (res.status !== 200) throw new Error('Auth failed');
   let authToken = res.body;
   return authToken;
}

export function userScenario() {
   group('user', function () {
      const randomName = getRandomString(getRandomInt(4, 15));
      const user = {
         "name": randomName,
         "email": randomName + "@email.com",
         "password": randomName
      };
      let resRegister = http.post('http://localhost:8080/users', JSON.stringify(user), {
         headers: { 'Content-Type': 'application/json' },
      });

      check(resRegister, {
         'register status was 201': (r) => r.status == 201,
      });

      let resLogin = http.post('http://localhost:8080/users/login', JSON.stringify(user));

      check(resLogin, {
         'login status was 200': (r) => r.status == 200,
      });
   });
}

export function storeScenario(authToken) {
   
   group('products', function () {
      let resListProducts = http.get('http://localhost:8080/products');

      check(resListProducts, {
         'list products status was 200': (r) => r.status == 200,
      });

      let products = JSON.parse(resListProducts.body);
      let randomProduct = products[getRandomInt(0, products.length - 1)];

      let resProductDetails = http.get(`http://localhost:8080/products/${randomProduct.id}`);

      check(resProductDetails, {
         'get product details status was 200': (r) => r.status == 200,
      });
   });

   group('cart', function () {
      // Add a random product to the cart
      let headers = { 'Authorization': `Bearer ${authToken}` };
      let resAddToCart = http.post('http://localhost:8080/cart/add/1', {}, { headers: headers });

      check(resAddToCart, {
         'add to cart status was 204': (r) => r.status == 204,
      });

      let resViewToCart = http.get('http://localhost:8080/cart', { headers: headers });
      check(resViewToCart, {
         'view cart status was 200': (r) => r.status === 200,
      });
   });

   group('checkout', function () {
      let headers = { 'Authorization': `Bearer ${authToken}` };

      let checkoutRes = http.post('http://localhost:8080/cart/checkout', {}, { headers: headers });
      check(checkoutRes, {
          'checkout status was 200': (r) => r.status === 200,
      });
   });


}

// Helper functions
function getRandomInt(min, max) {
   min = Math.ceil(min);
   max = Math.floor(max);
   return Math.floor(Math.random() * (max - min + 1)) + min;
}

function getRandomString(length) {
   var result = '';
   var characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
   for (var i = 0; i < length; i++) {
      result += characters.charAt(Math.floor(Math.random() * characters.length));
   }
   return result;
}