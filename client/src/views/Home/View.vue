<template>
  <div>
    <Divider text="Home" />
    <div class="flex space-x-10 justify-center">
      <button
        @click="chatWith()"
        class="transition flex space-x-0.5 hover:bg-gray-100 px-3 py-1 rounded-lg"
      >
        <mdicon name="message-outline" />
        <span>Start Chat</span>
      </button>
      <button
        @click="$router.push({ name: 'Profile' })"
        class="transition flex space-x-0.5 hover:bg-gray-100 px-3 py-1 rounded-lg"
      >
        <mdicon name="account-outline" />
        <span>Profile</span>
      </button>
    </div>
  </div>
</template>
<script>
import Divider from "../../components/Divider.vue";
export default {
  components: { Divider },
  mounted() {
    this.generateUUID();
  },
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
    chatWith() {
      this.$router.push({
        name: "Chat",
        params: { name: this.generateUUID() },
      });
    },
    generateUUID() {
      var timestamp = new Date().getTime();
      const uuid = "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(
        /[xy]/g,
        function (c) {
          var r = (timestamp + Math.random() * 16) % 16 | 0;
          timestamp = Math.floor(timestamp / 16);
          return (c === "x" ? r : (r & 0x3) | 0x8).toString(16);
        }
      );
      return uuid;
    },
  },
};
</script>
