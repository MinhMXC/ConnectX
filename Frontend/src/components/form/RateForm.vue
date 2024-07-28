<script setup>
import {computed, inject, ref, watch} from "vue";
import Button from "primevue/button";
import useBackendGet from "@/composables/useBackendGet.js";
import CXSelect from "@/components/form_input/CXSelect.vue";
import useBackendPost from "@/composables/useBackendPost.js";
import FormStatusText from "@/components/FormStatusText.vue";
import useBackendPatch from "@/composables/useBackendPatch.js";
import router from "@/router.js";
import CXInputNumber from "@/components/form_input/CXInputNumber.vue";

const props = defineProps(['data']);
const data = computed(() => props.data);

const name = ref("");
const description = ref("");
const time = ref(null);

const level = ref("");
const amount = ref(0);
const isOpen = ref("false");
const subject = ref("");

const $cookies = inject("$cookies");

const { data: levelData } = useBackendGet("/level");
const { data: subjectData } = useBackendGet("/subject");
let subjectList = subjectData.value;
const postObject = useBackendPost("/rate");
const patchObject = computed(() => useBackendPatch(`/rate/${props.data?.id}`));

const post = postObject.post;
const patch = computed(() => patchObject.value.patch);

const status = computed(() => props.data ? patchObject.value.status.value : postObject.status.value);
const loading = computed(() => props.data ? patchObject.value.loading.value : postObject.loading.value);

watch([data, subjectData, levelData], ([newData]) => {
  if (!newData || !subjectData || !levelData) {
    return;
  }
  const sub = subjectData.value?.find((e) => e.id === newData.subject_id);
  subject.value = sub?.name;
  level.value = levelData.value?.find((e) => e.id === sub?.level_id)?.name;
  amount.value = newData?.amount;
  isOpen.value = newData?.is_open ? "Yes" : "No";
});

watch(level, (newData) => {
  subjectList = subjectData.value?.filter((subject) => subject.level_id === levelData.value.find((e) => e.name === newData).id);
  subject.value = null;
});

const submitOnClick = async () => {
  const uploadData = {
    amount: amount.value,
    is_open: isOpen.value === "Yes",
    subject_id: subjectData.value
        .find((sub) => sub.name === subject.value && sub.level_id === levelData.value.find((e) => e.name === level.value).id).id
  };


  if ($cookies.get("user_type") === "1") {
    uploadData.tutor_id = Number($cookies.get("user_id"));
  } else if ($cookies.get("user_type") === "2") {
    uploadData.tuition_center_id = Number($cookies.get("user_id"));
  }

  if (props.data) {
    await patch.value(uploadData);

  } else {
    await post(uploadData);
  }

  if (status.value === "Success") {
    await router.back();
  }
};
</script>

<template>
  <div class="vertical-form-container">
    <CXSelect label="Level" :options="levelData?.map((e) => e.name)" v-model="level" />
    <CXSelect label="Subject" :options="subjectList?.map((e) => e.name)" v-model="subject" />
    <CXInputNumber label="Amount / Hour" v-model="amount" prefix="$ " />
    <CXSelect label="Taking in Student" :options="['Yes', 'No']" v-model="isOpen" />

    <FormStatusText :status="status" />
    <Button :loading="loading" @click="submitOnClick">Submit</Button>
  </div>
</template>

<style scoped>

</style>