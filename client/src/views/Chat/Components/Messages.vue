<template>
  <ul class="list-none">
    <li
      v-for="(item, index) in messages"
      :key="index"
      class="flex items-center space-x-3 px-3 py-1 rounded-lg"
    >
      <button>
        <img
          class="w-10 h-10 rounded"
          :src="item.picture"
          alt="Default avatar"
        />
      </button>
      <div class="cursor-pointer w-full text-start">
        <dt class="font-bold text-sm">{{ item.name }}</dt>
        <dd class="font-thin">{{ item.text }}</dd>
      </div>
    </li>
  </ul>
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
