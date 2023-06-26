<template>
  <div>
    <Divider text="Groups" />
    <ul class="list-none space-y-2">
      <li
        v-for="(item, index) in groups"
        :key="index"
        class="transition flex items-center space-x-3 hover:bg-gray-100 px-3 py-1 rounded-lg"
      >
        <button
          @click="showUserImage(item.title, 'https://placehold.co/400x400')"
        >
          <img
            class="w-10 h-10 rounded"
            src="https://placehold.co/400"
            alt="Default avatar"
          />
        </button>
        <div class="cursor-pointer w-full text-start" @click="chatWith(item.slug)">
          <dt class="font-bold text-lg">{{ item.title }}</dt>
          <dd class="font-light text-sm">{{ item.description }}</dd>
        </div>
      </li>
    </ul>
  </div>
</template>
<script>
import Divider from "../../components/Divider.vue";

export default {
  components: { Divider },
  data() {
    return {
      groups: [
        {
          title: "1st Group",
          slug: "1",
          description: "1st Description",
        },
        {
          title: "2nd Group",
          slug: "2",
          description: "2nd Description",
        },
      ],
    };
  },
  methods: {
    showUserImage(title, src) {
      this.$store.state.imgDialog.title = title;
      this.$store.state.imgDialog.src = src;
      this.$store.state.imgDialog.isopen = true;
    },
    chatWith(name) {
      this.$router.push({ name: "Chat", params: { name: name } });
    },
  },
};
</script>
