import {inject, ref, toValue} from 'vue';

// url must be a ref
export default function useBackendDelete(url) {
    const data = ref(null);
    const status = ref("");
    const loading = ref(false);
    const $cookies = inject("$cookies");

    const deleteFn = async () => {
        data.value = null;
        status.value = "";
        loading.value = true;

        const res = await fetch(import.meta.env.VITE_BACKEND_URL + toValue(url), {
            method: "DELETE",
            cors: "cors",
            cache: "no-cache",
            headers: {
                "Authorization": $cookies.get("Authorization")
            }
        });

        if (res.status >= 400) {
            status.value = JSON.parse(await res.text()).error;
            loading.value = false;
            return;
        }

        if (res.headers.get("Authorization")) {
            $cookies.set("Authorization", res.headers.get("Authorization").split(' ')[1]);
        }

        if (res.headers.get("Refresh-Token")) {
            $cookies.set("Refresh-Token", res.headers.get("Refresh-Token"));
        }

        const text = await res.text();
        if (text) {
            data.value = JSON.parse(text);
        }
        status.value = "Success";
        loading.value = false;
    };

    return { deleteFn, data, status, loading };
}