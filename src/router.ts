import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import EmailsView from '@/views/EmailsView.vue'; 
import EmailContentView from '@/views/EmailContentView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: EmailsView,
  },
  {
    path: '/EmailContent/:id',
    name: 'EmailContent',
    component: EmailContentView,
    props: true
  }
  
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;


// const routes = [
//   {
//     path: "/",
//     name: "EmailsView",
//     component: EmailsView,
//     },
//     {
//         path: '/EmailContent/:id',
//         name: 'EmailContent',
//         component: EmailContentView,
//         props: true
//     }
  
// ];

// const router = createRouter({
//   history: createWebHistory(),
//   routes,
// });

// export default router;