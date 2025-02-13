<template class="container">
  <TabsWrapper>
    <Tab title="Menu 1"><Prolabore/></Tab>
    <Tab title="Menu 2">Hello from Tab 2</Tab>
    <Tab title="Menu 3">Hello from Tab 3</Tab>
  </TabsWrapper>
</template>

<script setup>
  import TabsWrapper from "./components/TabsWrapper.vue";
  import Tab from "./components/Tab.vue";
  import Prolabore from "./components/Prolabore.vue";
</script>

<style scoped>
  .container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
  }
</style>