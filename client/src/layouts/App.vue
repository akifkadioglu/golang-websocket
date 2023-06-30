<template>
  <div>
    <slot />
  </div>
  <IMG_Dialog />
  <vue3-snackbar reverse bottom left :duration="4000">
    <template #message-content="{ message, title, text }">
      <div class="flex space-x-5 items-center">
        <div>
          <img
            v-if="message.picture != null"
            :src="message.picture"
            class="w-16 h-16"
            alt=""
          />
        </div>
        <div>
          <strong v-text="title"></strong>
          <p v-text="text"></p>
        </div>
      </div>
    </template>
  </vue3-snackbar>
</template>
<script>
import IMG_Dialog from "../components/IMG_Dialog.vue";
import { Vue3Snackbar } from "vue3-snackbar";

export default {
  components: { IMG_Dialog, Vue3Snackbar },
};
</script>
