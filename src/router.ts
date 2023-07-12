import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import EmailsListView from '@/views/EmailsListView.vue'; 
import EmailContentView from '@/views/EmailContentView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'EmailsList',
    component: EmailsListView,
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