<template>
  <div>
    <slot />
  </div>
  <IMG_Dialog />
  <vue3-snackbar reverse bottom left :duration="4000" />
</template>
<script>
import IMG_Dialog from "../components/IMG_Dialog.vue";
import { Vue3Snackbar } from "vue3-snackbar";

export default {
  components: { IMG_Dialog, Vue3Snackbar },
};
</script>
