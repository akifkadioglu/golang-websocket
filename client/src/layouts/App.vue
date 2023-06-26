<template>
  <div>
    <slot />
  </div>
  <IMG_Dialog />
</template>
<script>
import IMG_Dialog from "../components/IMG_Dialog.vue";

export default {
  components: { IMG_Dialog },
};
</script>
