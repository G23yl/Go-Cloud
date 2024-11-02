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
          beforeEnter(to, from) {
            if (to.fullPath === "/dashboard/files") {
              to.fullPath = "/dashboard/files?path=/"
            }
          },
          props: (route) => ({ query: route.query.path }),
        },
        {
          path: "docs",
          component: () => import("@/views/Files/Docs.vue"),
          name: "docs",
          meta: {
            requireAuth: true,
          },
        },
        {
          path: "images",
          component: () => import("@/views/Files/Images.vue"),
          name: "images",
          meta: {
            requireAuth: true,
          },
        },
        {
          path: "videos",
          component: () => import("@/views/Files/Videos.vue"),
          name: "videos",
          meta: {
            requireAuth: true,
          },
        },
        {
          path: "audios",
          component: () => import("@/views/Files/Audios.vue"),
          name: "audios",
          meta: {
            requireAuth: true,
          },
        },
        {
          path: "others",
          component: () => import("@/views/Files/Others.vue"),
          name: "others",
          meta: {
            requireAuth: true,
          },
        },
        {
          path: "settings",
          component: () => import("@/views/settings/Settings.vue"),
          name: "settings",
          meta: {
            requireAuth: true,
          },
        },
        {
          path: "helps",
          component: () => import("@/views/help/Help.vue"),
          name: "helps",
          meta: {
            requireAuth: true,
          },
        },
      ],
    },
    {
      path: "/404",
      component: () => import("@/views/NotFound/404.vue"),
      name: "404",
    },
    {
      path: "/:pathMatch(.*)*",
      redirect: "/404",
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
