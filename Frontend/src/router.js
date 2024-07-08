import { createRouter, createWebHistory} from "vue-router";

import Login from './pages/auth/Login.vue';
import SignUp from "@/pages/auth/SignUp.vue";
import SetupChoose from "@/pages/SetupChoose.vue";
import UserSetup from "@/pages/user/UserSetup.vue";
import TutorSetup from "@/pages/tutor/TutorSetup.vue";
import TuitionCenterSetup from "@/pages/tuition_center/TuitionCenterSetup.vue";
import NotImplemented from "@/pages/NotImplemented.vue";
import LandingPage from "@/pages/LandingPage.vue";
import TeacherDisplay from "@/components/teacher/TeacherDisplay.vue";
import UserProfile from "@/pages/user/UserProfile.vue";
import TutorProfile from "@/pages/tutor/TutorProfile.vue";
import TuitionCenterProfile from "@/pages/tuition_center/TuitionCenterProfile.vue";

const routes = [
    { path: '/auth/login', component: Login },
    { path: '/auth/signup', component: SignUp },
    { path: '/setup', component: SetupChoose },
    { path: '/setup/user', component: UserSetup },
    { path: '/setup/tutor', component: TutorSetup },
    { path: '/setup/tuition_center', component: TuitionCenterSetup },
    { path: '/temp', component: NotImplemented },
    { path: '/', component: LandingPage },
    { path: '/teacher', component: TeacherDisplay },

    { path: '/user/:id', component: UserProfile },
    { path: '/tutor/:id', component: TutorProfile },
    { path: '/tuition_center/:id', component: TuitionCenterProfile },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router;