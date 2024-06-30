import {inject, ref} from 'vue';

export default function useBackendPost(url) {
    const data = ref(null);
    const status = ref("");
    const loading = ref(false);
    const $cookies = inject("$cookies");

    const post = async (postData) => {
        data.value = null;
        status.value = "";
        loading.value = true;

        const res = await fetch("http://localhost:8080" + url, {
            method: "POST",
            cors: "cors",
            cache: "no-cache",
            headers: {
                "Content-Type": "application/json",
                "Authorization": $cookies.get("Authorization")
            },
            body: JSON.stringify(postData)
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
    }

    return { post, data, status, loading };
}