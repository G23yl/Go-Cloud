import { getUser } from "@/utils/user"
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
    {
      path: "/dashboard",
      component: () => import("@/views/DashBoard/Dashboard.vue"),
      redirect: { name: "index" },
      children: [
        {
          path: "index",
          component: () => import("@/views/Index/Index.vue"),
          name: "index",
          meta: {
            requireAuth: true,
          },
        },
        {
          path: "files",
          component: () => import("@/views/Files/Files.vue"),
          name: "files",
          meta: {
            requireAuth: true,
          },
        },
      ],
    },
  ],
})

router.beforeEach((to, from) => {
  if (to.meta.requireAuth && getUser() == undefined) {
    return {
      name: "login",
    }
  }
})

export default router
