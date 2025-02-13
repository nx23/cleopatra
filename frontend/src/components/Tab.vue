<script>
  import { inject } from "vue";
  import Logo from "./Logo.vue";
  export default {
    props: ["title"],
    setup() {
      const selectedTab = inject("selectedTab");
      const home = "Home";

      return {
        selectedTab,
      };
    },
  };
</script>

<template>
  <div class="tab-content" v-show="selectedTab === home">
    <Logo/>
  </div>
  <div class="tab-content" v-show="selectedTab === title">
    <slot />
  </div>
</template>