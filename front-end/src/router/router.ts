import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/login",
      component: () => import("@/views/login/Login.vue"),
      name: "login",
    },
    {
      path: "/",
      redirect: "/login",
    },
    {
      path: "/sign-up",
      component: () => import("@/views/signUp/SignUp.vue"),
      name: "signUp",
    },
  ],
})

export default router
