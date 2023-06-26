<template>
    <div class="relative flex py-5 items-center">
      <div class="flex-grow border-t border-gray-400" />
      <span class="flex-shrink mx-4 font-mono">{{ text }}</span>
      <div class="flex-grow border-t border-gray-400" />
    </div>
  </template>
  <script>
  export default {
    props: {
      text: {
        type: String,
        default: "",
      },
    },
  };
  </script>