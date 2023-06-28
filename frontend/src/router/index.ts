// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/Users',
    component: () => import('@/layouts/default/Default.vue'),
    children: [
      {
        path: '',
        name: 'Home',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "home" */ '@/views/Home/Home.vue'),
      },
      {
        path: '/contact',
        component: () => import('@/views/Contact/Contact.vue'),
      },
      {
        path: '/exercise',
        component: () => import('@/views/Exercise/Exercise.vue'),
      }
      ,
      {
        path: '/users',
        component: () => import('@/views/Users/Users.vue'),
      },
      {
        path: '/admin',
        component: () => import('@/views/Admin/Admin.vue'),
      },
      {
        path: '/schoolList',
        component: ()=> import('@/views/SchoolList/School.vue')
      }
    ],
  },
  {
    path: '/',
    component: () => import('@/views/LoginScreen/Login.vue'),
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
