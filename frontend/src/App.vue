<script setup>
  import Logo from "./components/Logo.vue";
  import TabsWrapper from "./components/TabsWrapper.vue";
  import Tab from "./components/Tab.vue";
  // import Prolabore from "./components/Prolabore.vue";
</script>

<template class="container">
  <!-- <Prolabore/> -->
  <TabsWrapper>
    <Tab title="Menu 1">Hello from Tab 1</Tab>
    <Tab title="Menu 2">Hello from Tab 2</Tab>
    <Tab title="Menu 3">Hello from Tab 3</Tab>
  </TabsWrapper>
  <Logo/>
</template>

<style scoped>
  .container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
  }
</style>