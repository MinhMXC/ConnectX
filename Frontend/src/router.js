import { createRouter, createWebHistory} from "vue-router";

import Login from './pages/auth/Login.vue';
import SignUp from "@/pages/auth/SignUp.vue";
import SetupChoose from "@/pages/setup/SetupChoose.vue";
import UserSetup from "@/pages/setup/UserSetup.vue";
import TutorSetup from "@/pages/setup/TutorSetup.vue";
import TuitionCenterSetup from "@/pages/setup/TuitionCenterSetup.vue";
import NotImplemented from "@/pages/NotImplemented.vue";
import LandingPage from "@/pages/LandingPage.vue";
import TeacherDisplay from "@/components/teacher/TeacherDisplay.vue";

const routes = [
    { path: '/auth/login', component: Login },
    { path: '/auth/signup', component: SignUp },
    { path: '/setup', component: SetupChoose },
    { path: '/setup/user', component: UserSetup },
    { path: '/setup/tutor', component: TutorSetup },
    { path: '/setup/tuition_center', component: TuitionCenterSetup },
    { path: '/temp', component: NotImplemented },
    { path: '/', component: LandingPage },
    { path: '/teacher', component: TeacherDisplay }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router;