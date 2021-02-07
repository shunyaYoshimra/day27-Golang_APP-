import Vue from "vue";
import Router from "vue-router";
import axios from "axios";
import Users from "../components/Users.vue";
import User from "../components/User.vue";
import Signup from "../components/Signup.vue";
import Login from "../components/Login.vue";
import NewProfile from "../components/NewProfile.vue";
import Profile from "../components/Profile.vue";
import EditProfile from "../components/EditProfile.vue";
import NewContact from "../components/NewContact.vue";
import Start from "../components/Start.vue";

Vue.use(Router);

const routes = [
  { path: "/", component: Users },
  { path: "/user/:id", component: User },
  { path: "/signup", component: Signup },
  { path: "/login", component: Login },
  { path: "/profiles/new", component: NewProfile },
  { path: "/profile/:id", name: "profile", component: Profile },
  { path: "/profiles/edit", component: EditProfile },
  { path: "/contacts/new", component: NewContact },
  { path: "/start", component: Start },
]

const router = new Router({
  // mode: "history",
  routes
})

router.beforeEach((to, from, next) => {
  axios.get("/api/session_check").then((res) => {
    if (to.path == "/signup" || to.path == "/login") {
      if (res.data) {
        next("/");
      } else {
        next();
      }
    } else {
      if (res.data) {
        next();
      } else {
        next("/login");
      }
    }
  })
})

export default router;