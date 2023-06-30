<template>
  <div class="container m-5 mx-auto">
    <div class="md:flex mb-6">
      <div class="flex w-full space-x-3 items-center">
        <div class="relative w-full">
          <input
            v-model="text"
            @keyup.enter="sendText"
            class="bg-gray-200 appearance-none border-2 border-gray-200 rounded-full w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500"
            id="inline-full-name"
            type="text"
            placeholder="Text here"
          />
          <button
            type="button"
            :class="isEmojiPickerOpen ? 'bg-blue-800 text-white' : ''"
            class="absolute right-0 font-medium rounded-full text-sm px-1.5 py-1.5 mt-1.5 mr-2"
            @click="isEmojiPickerOpen = !isEmojiPickerOpen"
          >
            <mdicon name="emoticon-happy-outline" />
          </button>
        </div>

        <button
          v-if="!isLoading"
          type="button"
          class="text-blue-700 border border-blue-700 hover:bg-blue-700 hover:text-white focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full text-sm p-2.5 text-center inline-flex items-center"
          @click="sendText"
        >
          <mdicon name="arrow-right" />
        </button>
        <div v-else role="status">
          <svg
            aria-hidden="true"
            class="inline w-9 h-9 mr-2 text-gray-200 animate-spin fill-gray-600"
            viewBox="0 0 100 101"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
              fill="currentColor"
            />
            <path
              d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
              fill="currentFill"
            />
          </svg>
        </div>
      </div>
    </div>
  </div>
  <div v-show="isEmojiPickerOpen" class="container pr-8 fixed bottom-0 mx-auto">
    <EmojiPicker
      style="width: 100%"
      :offset="9"
      display-recent
      :native="true"
      @select="onSelectEmoji"
    />
  </div>
</template>

<script>
/* imports */
import EmojiPicker from "vue3-emoji-picker";
import "vue3-emoji-picker/css";

export default {
  components: { EmojiPicker },
  props: {
    isLoading: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      text: "",
      isEmojiPickerOpen: false,
    };
  },

  methods: {
    sendText() {
      this.$emit("sendText", this.text);
      this.text = "";
    },

    onSelectEmoji(emoji) {
      this.text = this.text + emoji.i;
    },
  },
  emits: ["sendText"],
};
</script>
