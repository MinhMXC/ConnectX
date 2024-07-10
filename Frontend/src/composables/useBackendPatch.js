import {inject, ref} from 'vue';

export default function useBackendPatch(url) {
    const data = ref(null);
    const status = ref("");
    const loading = ref(false);
    const $cookies = inject("$cookies");

    const put = async (putData) => {
        data.value = null;
        status.value = "";
        loading.value = true;

        const res = await fetch(import.meta.env.VITE_BACKEND_URL + url, {
            method: "PATCH",
            cors: "cors",
            cache: "no-cache",
            headers: {
                "Content-Type": "application/json",
                "Authorization": $cookies.get("Authorization")
            },
            body: JSON.stringify(putData)
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

    return { put, data, status, loading };
}