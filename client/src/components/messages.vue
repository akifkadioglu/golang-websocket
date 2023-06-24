<template>
  <div class="d-grid justify-content-center">
    <div
      class="card my-3"
      style="width: 18rem"
      v-for="(item, index) in messages"
      :key="index"
    >
      <div class="card-body">
        <h5 class="card-title text-center">{{ item.username }}</h5>
        <p class="card-text">
          {{ item.text }}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    messages: {
      type: Array,
      default() {
        return [];
      },
    },
  },
};
</script>
