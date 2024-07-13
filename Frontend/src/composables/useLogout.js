import router from "@/router.js";
import useBackendGet from "@/composables/useBackendGet.js";
import {inject} from "vue";

export default function useLogout() {
    const { get } = useBackendGet("/user/logout", false);
    const $cookies = inject("$cookies");

    const logout = async () => {
        await get();
        $cookies.remove("Authorization");
        $cookies.remove("Refresh-Token");
        $cookies.remove("email");
        $cookies.remove("user_type");
        await router.push("/auth/login");
    };

    return { logout };
}