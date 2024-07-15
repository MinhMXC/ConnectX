<script setup>
import {computed, inject, ref, watch} from "vue";
import Button from "primevue/button";
import CXInputText from "@/components/form_input/CXInputText.vue";
import CXDatePicker from "@/components/form_input/CXDatePicker.vue";
import useBackendGet from "@/composables/useBackendGet.js";
import CXSelect from "@/components/form_input/CXSelect.vue";
import useBackendPost from "@/composables/useBackendPost.js";
import FormStatusText from "@/components/FormStatusText.vue";
import useBackendPatch from "@/composables/useBackendPatch.js";
import router from "@/router.js";

const props = defineProps(['data']);
const data = computed(() => props.data);

const name = ref("");
const description = ref("");
const time = ref(null);
const level = ref(null);

const $cookies = inject("$cookies");

const { data: levelData } = useBackendGet("/level");
const postObject = useBackendPost("/qualification");
const patchObject = computed(() => useBackendPatch(`/qualification/${props.data?.id}`));

const post = postObject.post;
const patch = computed(() => patchObject.value.patch);

const status = computed(() => props.data ? patchObject.value.status.value : postObject.status.value);
const loading = computed(() => props.data ? patchObject.value.loading.value : postObject.loading.value);

watch([data, levelData], ([newData]) => {
  if (!newData) {
    return;
  }
  name.value = newData.name;
  description.value = newData.description;
  time.value = new Date(newData.time);
  level.value = levelData.value?.find((e) => e.id === newData.level_id).name;
});

const submitOnClick = async () => {
  if (props.data) {
    await patch.value({
      name: name.value,
      description: description.value,
      time: new Date(time.value).getTime(),
      level_id: levelData.value.find((e) => e.name === level.value).id,
      tutor_id: data.value.tutor_id,
    });
  } else {
    await post({
      name: name.value,
      description: description.value,
      time: new Date(time.value).getTime(),
      level_id: levelData.value.find((e) => e.name === level.value).id,
      tutor_id: Number($cookies.get("user_id")),
    });
  }

  if (status.value === "Success") {
    await router.back();
  }
};
</script>

<template>
  <div class="vertical-form-container">
    <CXSelect label="Level of Education" :options="levelData?.map((e) => e.name)" v-model="level" />
    <CXInputText label="Name" v-model="name" />
    <CXInputText label="Description" v-model="description" />
    <CXDatePicker label="Time of Attainment" v-model="time" />
    <FormStatusText :status="status" />
    <Button :loading="loading" @click="submitOnClick">Submit</Button>
  </div>
</template>

<style scoped>

</style>