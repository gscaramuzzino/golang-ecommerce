import http from 'k6/http';
import { check, group } from 'k6';

export let options = {
   scenarios: {
      user_tests: {
         executor: 'constant-vus',
         vus: 20,
         duration: '5m',
         exec: 'userScenario'
      },
      store_tests: {
         executor: 'constant-vus',
         vus: 100,
         startTime: '5m',
         duration: '10m',
         exec: 'storeScenario'
      },
   },
};

export function setup() {
   const user = {
      "name": "k6user",
      "email": "k6user@email.com",
      "password": "k6user"
   };
   http.post('http://localhost:8080/users', JSON.stringify(user), {
      headers: { 'Content-Type': 'application/json' },
   });
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