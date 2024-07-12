import {inject, ref, toValue, watchEffect} from 'vue';

// url must be a ref
export default function useBackendGet(url, autoRun = true) {
    const first = ref(false);
    const data = ref(null);
    const status = ref("");
    const loading = ref(false);
    const $cookies = inject("$cookies");

    const get = async () => {
        first.value = true;
        data.value = null;
        status.value = "";
        loading.value = true;

        const res = await fetch(import.meta.env.VITE_BACKEND_URL + toValue(url), {
            method: "GET",
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

    watchEffect(() => {
        if (autoRun === false && first.value === false) {
            return;
        }
        get();
    });

    return { get, data, status, loading };
}